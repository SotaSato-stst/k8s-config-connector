// Copyright 2024 Google LLC
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

package direct

import (
	"testing"

	"google.golang.org/protobuf/types/known/durationpb"
)

func TestStringDuration_FromProto(t *testing.T) {
	mapctx := &MapContext{}
	d := &durationpb.Duration{Seconds: 34312, Nanos: 20}
	krm := StringDuration_FromProto(mapctx, d)
	if *krm != "9h31m52.00000002s" {
		t.Fatalf("google.protobuf.Duration -> string, expect \"9h31m52.00000002s\", got %s", *krm)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> string error: %s", mapctx.Err())
	}
}

func TestStringDuration_ToProto(t *testing.T) {
	mapctx := &MapContext{}
	s := "1h1m"
	d := StringDuration_ToProto(mapctx, &s)
	if d.Seconds != 3660 || d.Nanos != 0 {
		t.Fatalf("string -> google.protobuf.Duration, expect \"seconds:3660 nanos:00\", got %s", d)
	}
	if mapctx.Err() != nil {
		t.Fatalf("google.protobuf.Duration -> String error: %s", mapctx.Err())
	}
}

// Test message types for Map functions
type testProtoMessage struct {
	Value string
}

type testKRMMessage struct {
	Value string
}

func TestMap_FromProto(t *testing.T) {
	mapCtx := &MapContext{}

	// Test mapper function
	mapper := func(mapCtx *MapContext, in *testProtoMessage) *testKRMMessage {
		if in == nil {
			return nil
		}
		return &testKRMMessage{Value: in.Value}
	}

	// Test data
	input := map[string]*testProtoMessage{
		"key1": &testProtoMessage{Value: "value1"},
		"key2": &testProtoMessage{Value: "value2"},
	}

	// Call Map_FromProto
	result := Map_FromProto(mapCtx, input, mapper)

	// Verify results
	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	if result["key1"].Value != "value1" {
		t.Errorf("Expected value1, got %s", result["key1"].Value)
	}

	if result["key2"].Value != "value2" {
		t.Errorf("Expected value2, got %s", result["key2"].Value)
	}

	// Test nil input
	nilResult := Map_FromProto(mapCtx, nil, mapper)
	if nilResult != nil {
		t.Errorf("Expected nil result for nil input")
	}

	// Test empty map
	emptyResult := Map_FromProto(mapCtx, map[string]*testProtoMessage{}, mapper)
	if len(emptyResult) != 0 {
		t.Errorf("Expected empty map, got %d items", len(emptyResult))
	}
}

func TestMap_ToProto(t *testing.T) {
	mapCtx := &MapContext{}

	// Test mapper function
	mapper := func(mapCtx *MapContext, in *testKRMMessage) *testProtoMessage {
		if in == nil {
			return nil
		}
		return &testProtoMessage{Value: in.Value}
	}

	// Test data
	input := map[string]*testKRMMessage{
		"key1": &testKRMMessage{Value: "value1"},
		"key2": &testKRMMessage{Value: "value2"},
	}

	// Call Map_ToProto
	result := Map_ToProto(mapCtx, input, mapper)

	// Verify results
	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	if result["key1"].Value != "value1" {
		t.Errorf("Expected value1, got %s", result["key1"].Value)
	}

	if result["key2"].Value != "value2" {
		t.Errorf("Expected value2, got %s", result["key2"].Value)
	}

	// Test nil input
	nilResult := Map_ToProto(mapCtx, nil, mapper)
	if nilResult != nil {
		t.Errorf("Expected nil result for nil input")
	}

	// Test empty map
	emptyResult := Map_ToProto(mapCtx, map[string]*testKRMMessage{}, mapper)
	if len(emptyResult) != 0 {
		t.Errorf("Expected empty map, got %d items", len(emptyResult))
	}
}
