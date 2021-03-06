package dictionarygroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// SetBase64URL ...
type SetBase64URL struct {
	autorest.Response `json:"-"`
	Value             *map[string]*string `json:"value,omitempty"`
}

// SetBool ...
type SetBool struct {
	autorest.Response `json:"-"`
	Value             *map[string]*bool `json:"value,omitempty"`
}

// SetByteArray ...
type SetByteArray struct {
	autorest.Response `json:"-"`
	Value             *map[string][]byte `json:"value,omitempty"`
}

// SetDate ...
type SetDate struct {
	autorest.Response `json:"-"`
	Value             *map[string]*date.Date `json:"value,omitempty"`
}

// SetDateTime ...
type SetDateTime struct {
	autorest.Response `json:"-"`
	Value             *map[string]*date.Time `json:"value,omitempty"`
}

// SetDateTimeRfc1123 ...
type SetDateTimeRfc1123 struct {
	autorest.Response `json:"-"`
	Value             *map[string]*date.TimeRFC1123 `json:"value,omitempty"`
}

// SetFloat64 ...
type SetFloat64 struct {
	autorest.Response `json:"-"`
	Value             *map[string]*float64 `json:"value,omitempty"`
}

// SetInt32 ...
type SetInt32 struct {
	autorest.Response `json:"-"`
	Value             *map[string]*int32 `json:"value,omitempty"`
}

// SetInt64 ...
type SetInt64 struct {
	autorest.Response `json:"-"`
	Value             *map[string]*int64 `json:"value,omitempty"`
}

// SetListString ...
type SetListString struct {
	autorest.Response `json:"-"`
	Value             *map[string][]string `json:"value,omitempty"`
}

// SetSetString ...
type SetSetString struct {
	autorest.Response `json:"-"`
	Value             *map[string]map[string]*string `json:"value,omitempty"`
}

// SetString ...
type SetString struct {
	autorest.Response `json:"-"`
	Value             *map[string]*string `json:"value,omitempty"`
}

// SetTimeSpan ...
type SetTimeSpan struct {
	autorest.Response `json:"-"`
	Value             *map[string]*string `json:"value,omitempty"`
}

// SetWidget ...
type SetWidget struct {
	autorest.Response `json:"-"`
	Value             *map[string]*Widget `json:"value,omitempty"`
}

// Widget ...
type Widget struct {
	Integer *int32  `json:"integer,omitempty"`
	String  *string `json:"string,omitempty"`
}
