// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bigtable

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigtable"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableAuthorizedViewGVK, NewBigtableAuthorizedViewModel)
}

func NewBigtableAuthorizedViewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBigtableAuthorizedView{config: *config}, nil
}

var _ directbase.Model = &modelBigtableAuthorizedView{}

type modelBigtableAuthorizedView struct {
	config config.ControllerConfig
}

func (m *modelBigtableAuthorizedView) client(ctx context.Context, parentProject, instanceID string) (*gcp.AdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAdminClient(ctx, parentProject, instanceID, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableAuthorizedView client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelBigtableAuthorizedView) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableAuthorizedView{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAuthorizedViewIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Extract project ID and instance ID from identity
	identityStr := id.String()
	tokens := strings.Split(identityStr, "/")
	if len(tokens) < 4 || tokens[0] != "projects" || tokens[2] != "instances" {
		return nil, fmt.Errorf("unexpected format for AuthorizedView identity")
	}
	project := tokens[1]
	instanceID := tokens[3]

	adminClient, err := m.client(ctx, project, instanceID)
	if err != nil {
		return nil, fmt.Errorf("error creating admin client: %w", err)
	}
	return &authorizedViewAdapter{
		id:        id,
		gcpClient: adminClient,
		desired:   obj,
	}, nil
}

func (m *modelBigtableAuthorizedView) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type authorizedViewAdapter struct {
	id        *krm.AuthorizedViewIdentity
	gcpClient *gcp.AdminClient
	desired   *krm.BigtableAuthorizedView
	actual    *gcp.AuthorizedViewInfo
}

var _ directbase.Adapter = &authorizedViewAdapter{}

func (a *authorizedViewAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableAuthorizedView", "name", a.id)

	// Extract table and authorized view IDs from the identity path
	path := a.id.String()
	tokens := strings.Split(path, "/")
	if len(tokens) != 8 {
		return false, fmt.Errorf("unexpected format for AuthorizedView path: %s", path)
	}
	tableID := tokens[5]
	authorizedViewID := tokens[7]

	bigtableauthorizedviewInfo, err := a.gcpClient.AuthorizedViewInfo(ctx, tableID, authorizedViewID)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableAuthorizedView %q: %w", a.id, err)
	}

	a.actual = bigtableauthorizedviewInfo
	return true, nil
}

func (a *authorizedViewAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableAuthorizedView", "name", a.id)

	// Extract table ID from the identity path
	path := a.id.String()
	tokens := strings.Split(path, "/")
	if len(tokens) != 8 {
		return fmt.Errorf("unexpected format for AuthorizedView path: %s", path)
	}
	tableID := tokens[5]
	authorizedViewID := a.id.ID()

	conf := gcp.AuthorizedViewConf{
		TableID:          tableID,
		AuthorizedViewID: authorizedViewID,
	}

	// Configure the authorized view type
	if subsetView := a.desired.Spec.SubsetView; subsetView != nil {
		subsetViewConf := &gcp.SubsetViewConf{
			RowPrefixes:   subsetView.RowPrefixes,
			FamilySubsets: make(map[string]gcp.FamilySubset),
		}
		for fam, subset := range subsetView.FamilySubsets {
			subsetViewConf.FamilySubsets[fam] = gcp.FamilySubset{
				Qualifiers:        subset.Qualifiers,
				QualifierPrefixes: subset.QualifierPrefixes,
			}
		}
		conf.AuthorizedView = subsetViewConf
	}

	// Configure deletion protection
	if a.desired.Spec.DeletionProtection != nil {
		if *a.desired.Spec.DeletionProtection {
			conf.DeletionProtection = gcp.Protected
		} else {
			conf.DeletionProtection = gcp.Unprotected
		}
	}

	err := a.gcpClient.CreateAuthorizedView(ctx, &conf)
	if err != nil {
		return fmt.Errorf("creating BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableAuthorizedView", "name", a.id)

	status := &krm.BigtableAuthorizedViewStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *authorizedViewAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableAuthorizedView", "name", a.id)

	// Extract table ID from the identity path
	path := a.id.String()
	tokens := strings.Split(path, "/")
	if len(tokens) != 8 {
		return fmt.Errorf("unexpected format for AuthorizedView path: %s", path)
	}
	tableID := tokens[5]
	authorizedViewID := a.id.ID()

	updateConf := gcp.UpdateAuthorizedViewConf{
		AuthorizedViewConf: gcp.AuthorizedViewConf{
			TableID:          tableID,
			AuthorizedViewID: authorizedViewID,
		},
	}

	hasChanges := false

	// Check for subset view changes
	if subsetView := a.desired.Spec.SubsetView; subsetView != nil {
		needsUpdate := false

		// Compare with actual if we have it
		if a.actual != nil && a.actual.AuthorizedView != nil {
			if actualSubsetView, ok := a.actual.AuthorizedView.(*gcp.SubsetViewInfo); ok {
				// Check if row prefixes changed
				if !byteSlicesEqual(subsetView.RowPrefixes, actualSubsetView.RowPrefixes) {
					needsUpdate = true
				}
				// Check if family subsets changed
				if !familySubsetsEqual(subsetView.FamilySubsets, actualSubsetView.FamilySubsets) {
					needsUpdate = true
				}
			} else {
				// Different view type, needs update
				needsUpdate = true
			}
		} else {
			// No actual state, assume update needed
			needsUpdate = true
		}

		if needsUpdate {
			subsetViewConf := &gcp.SubsetViewConf{
				RowPrefixes:   subsetView.RowPrefixes,
				FamilySubsets: make(map[string]gcp.FamilySubset),
			}
			for fam, subset := range subsetView.FamilySubsets {
				subsetViewConf.FamilySubsets[fam] = gcp.FamilySubset{
					Qualifiers:        subset.Qualifiers,
					QualifierPrefixes: subset.QualifierPrefixes,
				}
			}
			updateConf.AuthorizedViewConf.AuthorizedView = subsetViewConf
			hasChanges = true
		}
	}

	// Check for deletion protection changes
	if a.desired.Spec.DeletionProtection != nil {
		currentProtection := false
		if a.actual != nil {
			currentProtection = a.actual.DeletionProtection == gcp.Protected
		}
		desiredProtection := *a.desired.Spec.DeletionProtection
		if currentProtection != desiredProtection {
			if desiredProtection {
				updateConf.AuthorizedViewConf.DeletionProtection = gcp.Protected
			} else {
				updateConf.AuthorizedViewConf.DeletionProtection = gcp.Unprotected
			}
			hasChanges = true
		}
	}

	if !hasChanges {
		log.V(2).Info("no changes to update", "name", a.id)
	} else {
		err := a.gcpClient.UpdateAuthorizedView(ctx, updateConf)
		if err != nil {
			return fmt.Errorf("updating BigtableAuthorizedView %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated BigtableAuthorizedView", "name", a.id)
	}

	status := &krm.BigtableAuthorizedViewStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *authorizedViewAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableAuthorizedView{}

	// Set the spec based on actual state
	if a.actual.AuthorizedView != nil {
		if subsetViewInfo, ok := a.actual.AuthorizedView.(*gcp.SubsetViewInfo); ok {
			obj.Spec.SubsetView = &krm.AuthorizedView_SubsetView{
				RowPrefixes:   subsetViewInfo.RowPrefixes,
				FamilySubsets: make(map[string]*krm.AuthorizedView_FamilySubsets),
			}
			for fam, subset := range subsetViewInfo.FamilySubsets {
				obj.Spec.SubsetView.FamilySubsets[fam] = &krm.AuthorizedView_FamilySubsets{
					Qualifiers:        subset.Qualifiers,
					QualifierPrefixes: subset.QualifierPrefixes,
				}
			}
		}
	}

	// Set deletion protection
	if a.actual.DeletionProtection == gcp.Protected {
		obj.Spec.DeletionProtection = direct.LazyPtr(true)
	} else if a.actual.DeletionProtection == gcp.Unprotected {
		obj.Spec.DeletionProtection = direct.LazyPtr(false)
	}

	// Extract table reference from the identity path
	path := a.id.String()
	tokens := strings.Split(path, "/")
	if len(tokens) >= 6 {
		tableRef := fmt.Sprintf("projects/%s/instances/%s/tables/%s", tokens[1], tokens[3], tokens[5])
		obj.Spec.TableRef.External = tableRef
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BigtableAuthorizedViewGVK)

	u.Object = uObj
	return u, nil
}

func (a *authorizedViewAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableAuthorizedView", "name", a.id)

	// Extract instance and table IDs from the identity path
	path := a.id.String()
	tokens := strings.Split(path, "/")
	if len(tokens) != 8 {
		return false, fmt.Errorf("unexpected format for AuthorizedView path: %s", path)
	}
	tableID := tokens[5]
	authorizedViewID := a.id.ID()

	err := a.gcpClient.DeleteAuthorizedView(ctx, tableID, authorizedViewID)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BigtableAuthorizedView, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableAuthorizedView", "name", a.id)

	return true, nil
}

// Helper functions for comparison
func byteSlicesEqual(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func familySubsetsEqual(a map[string]*krm.AuthorizedView_FamilySubsets, b map[string]gcp.FamilySubset) bool {
	if len(a) != len(b) {
		return false
	}
	for fam, aSubset := range a {
		bSubset, ok := b[fam]
		if !ok {
			return false
		}
		if !byteSlicesEqual(aSubset.Qualifiers, bSubset.Qualifiers) {
			return false
		}
		if !byteSlicesEqual(aSubset.QualifierPrefixes, bSubset.QualifierPrefixes) {
			return false
		}
	}
	return true
}
