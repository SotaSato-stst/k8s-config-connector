//go:build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Contact) DeepCopyInto(out *Contact) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.NotificationCategorySubscriptions != nil {
		in, out := &in.NotificationCategorySubscriptions, &out.NotificationCategorySubscriptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LanguageTag != nil {
		in, out := &in.LanguageTag, &out.LanguageTag
		*out = new(string)
		**out = **in
	}
	if in.ValidateTime != nil {
		in, out := &in.ValidateTime, &out.ValidateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Contact.
func (in *Contact) DeepCopy() *Contact {
	if in == nil {
		return nil
	}
	out := new(Contact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactIdentity) DeepCopyInto(out *ContactIdentity) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(ContactParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactIdentity.
func (in *ContactIdentity) DeepCopy() *ContactIdentity {
	if in == nil {
		return nil
	}
	out := new(ContactIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactObservedState) DeepCopyInto(out *ContactObservedState) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ValidationState != nil {
		in, out := &in.ValidationState, &out.ValidationState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactObservedState.
func (in *ContactObservedState) DeepCopy() *ContactObservedState {
	if in == nil {
		return nil
	}
	out := new(ContactObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactParent) DeepCopyInto(out *ContactParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactParent.
func (in *ContactParent) DeepCopy() *ContactParent {
	if in == nil {
		return nil
	}
	out := new(ContactParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactRef) DeepCopyInto(out *ContactRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactRef.
func (in *ContactRef) DeepCopy() *ContactRef {
	if in == nil {
		return nil
	}
	out := new(ContactRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EssentialContactsContact) DeepCopyInto(out *EssentialContactsContact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EssentialContactsContact.
func (in *EssentialContactsContact) DeepCopy() *EssentialContactsContact {
	if in == nil {
		return nil
	}
	out := new(EssentialContactsContact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EssentialContactsContact) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EssentialContactsContactList) DeepCopyInto(out *EssentialContactsContactList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EssentialContactsContact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EssentialContactsContactList.
func (in *EssentialContactsContactList) DeepCopy() *EssentialContactsContactList {
	if in == nil {
		return nil
	}
	out := new(EssentialContactsContactList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EssentialContactsContactList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EssentialContactsContactObservedState) DeepCopyInto(out *EssentialContactsContactObservedState) {
	*out = *in
	if in.ValidationState != nil {
		in, out := &in.ValidationState, &out.ValidationState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EssentialContactsContactObservedState.
func (in *EssentialContactsContactObservedState) DeepCopy() *EssentialContactsContactObservedState {
	if in == nil {
		return nil
	}
	out := new(EssentialContactsContactObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EssentialContactsContactSpec) DeepCopyInto(out *EssentialContactsContactSpec) {
	*out = *in
	in.Parent.DeepCopyInto(&out.Parent)
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.NotificationCategorySubscriptions != nil {
		in, out := &in.NotificationCategorySubscriptions, &out.NotificationCategorySubscriptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LanguageTag != nil {
		in, out := &in.LanguageTag, &out.LanguageTag
		*out = new(string)
		**out = **in
	}
	if in.ValidateTime != nil {
		in, out := &in.ValidateTime, &out.ValidateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EssentialContactsContactSpec.
func (in *EssentialContactsContactSpec) DeepCopy() *EssentialContactsContactSpec {
	if in == nil {
		return nil
	}
	out := new(EssentialContactsContactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EssentialContactsContactStatus) DeepCopyInto(out *EssentialContactsContactStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(EssentialContactsContactObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EssentialContactsContactStatus.
func (in *EssentialContactsContactStatus) DeepCopy() *EssentialContactsContactStatus {
	if in == nil {
		return nil
	}
	out := new(EssentialContactsContactStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parent) DeepCopyInto(out *Parent) {
	*out = *in
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
	if in.FolderRef != nil {
		in, out := &in.FolderRef, &out.FolderRef
		*out = new(v1beta1.FolderRef)
		**out = **in
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1beta1.OrganizationRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parent.
func (in *Parent) DeepCopy() *Parent {
	if in == nil {
		return nil
	}
	out := new(Parent)
	in.DeepCopyInto(out)
	return out
}
