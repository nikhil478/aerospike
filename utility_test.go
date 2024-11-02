package main

import (
	"reflect"
	"testing"
	"time"

	as "github.com/aerospike/aerospike-client-go/v7"
)

// Define a struct that matches the expected bins
type TestStruct struct {
	Name      string    `as:"name"`
	Age       int       `as:"age"`
	Height    float64   `as:"height"`
	UpdatedAt time.Time `as:"updated_at"`
}

func TestBinsToStruct(t *testing.T) {
	// Test cases
	tests := []struct {
		name      string
		record    *as.Record
		result    *TestStruct
		expectErr bool
		expected  TestStruct
	}{
		{
			name: "valid record",
			record: &as.Record{
				Bins: as.BinMap{
					"name":        "John Doe",
					"age":         30,
					"height":      5.9,
					"updated_at":  time.Now(),
				},
			},
			result: &TestStruct{},
			expectErr: false,
		},
		{
			name: "nil record",
			record: nil,
			result: &TestStruct{},
			expectErr: true,
		},
		{
			name: "result is nil pointer",
			record: &as.Record{
				Bins: as.BinMap{
					"name": "Jane Doe",
				},
			},
			result: nil,
			expectErr: true,
		},
		{
			name: "non-pointer result",
			record: &as.Record{
				Bins: as.BinMap{
					"name": "Jane Doe",
				},
			},
			result: &TestStruct{},
			expectErr: true,
		},
		{
			name: "missing bins",
			record: &as.Record{
				Bins: as.BinMap{},
			},
			result: &TestStruct{},
			expectErr: false,
		},
		{
			name: "type mismatch",
			record: &as.Record{
				Bins: as.BinMap{
					"name":   "John",
					"age":    "notAnInt", // This should cause a mismatch
					"height": 5.5,
				},
			},
			result: &TestStruct{},
			expectErr: false, // No error on type mismatch, just won't set the field
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectErr {
				err := BinsToStruct(tt.record, tt.result)
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				err := BinsToStruct(tt.record, tt.result)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else if !reflect.DeepEqual(tt.result, &tt.expected) {
					t.Errorf("expected %+v, got %+v", tt.expected, tt.result)
				}
			}
		})
	}
}
