// Copyright The OpenTelemetry Authors
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

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package internal

import (
	"sort"

	otlplogs "go.opentelemetry.io/collector/pdata/internal/data/protogen/logs/v1"
)

// ResourceLogsSlice logically represents a slice of ResourceLogs.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewResourceLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogsSlice struct {
	// orig points to the slice otlplogs.ResourceLogs field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like EnsureCapacity.
	orig *[]*otlplogs.ResourceLogs
}

func newResourceLogsSlice(orig *[]*otlplogs.ResourceLogs) ResourceLogsSlice {
	return ResourceLogsSlice{orig}
}

// NewResourceLogsSlice creates a ResourceLogsSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewResourceLogsSlice() ResourceLogsSlice {
	orig := []*otlplogs.ResourceLogs(nil)
	return ResourceLogsSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewResourceLogsSlice()".
func (es ResourceLogsSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//   for i := 0; i < es.Len(); i++ {
//       e := es.At(i)
//       ... // Do something with the element
//   }
func (es ResourceLogsSlice) At(ix int) ResourceLogs {
	return newResourceLogs((*es.orig)[ix])
}

// CopyTo copies all elements from the current slice to the dest.
func (es ResourceLogsSlice) CopyTo(dest ResourceLogsSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newResourceLogs((*es.orig)[i]).CopyTo(newResourceLogs((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlplogs.ResourceLogs, srcLen)
	wrappers := make([]*otlplogs.ResourceLogs, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newResourceLogs((*es.orig)[i]).CopyTo(newResourceLogs(wrappers[i]))
	}
	*dest.orig = wrappers
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ResourceLogsSlice can be initialized:
//   es := NewResourceLogsSlice()
//   es.EnsureCapacity(4)
//   for i := 0; i < 4; i++ {
//       e := es.AppendEmpty()
//       // Here should set all the values for e.
//   }
func (es ResourceLogsSlice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlplogs.ResourceLogs, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty ResourceLogs.
// It returns the newly added ResourceLogs.
func (es ResourceLogsSlice) AppendEmpty() ResourceLogs {
	*es.orig = append(*es.orig, &otlplogs.ResourceLogs{})
	return es.At(es.Len() - 1)
}

// Sort sorts the ResourceLogs elements within ResourceLogsSlice given the
// provided less function so that two instances of ResourceLogsSlice
// can be compared.
//
// Returns the same instance to allow nicer code like:
//   lessFunc := func(a, b ResourceLogs) bool {
//     return a.Name() < b.Name() // choose any comparison here
//   }
//   assert.EqualValues(t, expected.Sort(lessFunc), actual.Sort(lessFunc))
func (es ResourceLogsSlice) Sort(less func(a, b ResourceLogs) bool) ResourceLogsSlice {
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
	return es
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ResourceLogsSlice) MoveAndAppendTo(dest ResourceLogsSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ResourceLogsSlice) RemoveIf(f func(ResourceLogs) bool) {
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.orig = (*es.orig)[:newLen]
}

// ResourceLogs is a collection of logs from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogs struct {
	orig *otlplogs.ResourceLogs
}

func newResourceLogs(orig *otlplogs.ResourceLogs) ResourceLogs {
	return ResourceLogs{orig: orig}
}

// NewResourceLogs creates a new empty ResourceLogs.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewResourceLogs() ResourceLogs {
	return newResourceLogs(&otlplogs.ResourceLogs{})
}

// MoveTo moves all properties from the current struct to dest
// resetting the current instance to its zero value
func (ms ResourceLogs) MoveTo(dest ResourceLogs) {
	*dest.orig = *ms.orig
	*ms.orig = otlplogs.ResourceLogs{}
}

// Resource returns the resource associated with this ResourceLogs.
func (ms ResourceLogs) Resource() Resource {
	return newResource(&(*ms.orig).Resource)
}

// SchemaUrl returns the schemaurl associated with this ResourceLogs.
func (ms ResourceLogs) SchemaUrl() string {
	return (*ms.orig).SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ResourceLogs.
func (ms ResourceLogs) SetSchemaUrl(v string) {
	(*ms.orig).SchemaUrl = v
}

// ScopeLogs returns the ScopeLogs associated with this ResourceLogs.
func (ms ResourceLogs) ScopeLogs() ScopeLogsSlice {
	return newScopeLogsSlice(&(*ms.orig).ScopeLogs)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms ResourceLogs) CopyTo(dest ResourceLogs) {
	ms.Resource().CopyTo(dest.Resource())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.ScopeLogs().CopyTo(dest.ScopeLogs())
}

// ScopeLogsSlice logically represents a slice of ScopeLogs.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewScopeLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeLogsSlice struct {
	// orig points to the slice otlplogs.ScopeLogs field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like EnsureCapacity.
	orig *[]*otlplogs.ScopeLogs
}

func newScopeLogsSlice(orig *[]*otlplogs.ScopeLogs) ScopeLogsSlice {
	return ScopeLogsSlice{orig}
}

// NewScopeLogsSlice creates a ScopeLogsSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewScopeLogsSlice() ScopeLogsSlice {
	orig := []*otlplogs.ScopeLogs(nil)
	return ScopeLogsSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewScopeLogsSlice()".
func (es ScopeLogsSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//   for i := 0; i < es.Len(); i++ {
//       e := es.At(i)
//       ... // Do something with the element
//   }
func (es ScopeLogsSlice) At(ix int) ScopeLogs {
	return newScopeLogs((*es.orig)[ix])
}

// CopyTo copies all elements from the current slice to the dest.
func (es ScopeLogsSlice) CopyTo(dest ScopeLogsSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newScopeLogs((*es.orig)[i]).CopyTo(newScopeLogs((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlplogs.ScopeLogs, srcLen)
	wrappers := make([]*otlplogs.ScopeLogs, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newScopeLogs((*es.orig)[i]).CopyTo(newScopeLogs(wrappers[i]))
	}
	*dest.orig = wrappers
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ScopeLogsSlice can be initialized:
//   es := NewScopeLogsSlice()
//   es.EnsureCapacity(4)
//   for i := 0; i < 4; i++ {
//       e := es.AppendEmpty()
//       // Here should set all the values for e.
//   }
func (es ScopeLogsSlice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlplogs.ScopeLogs, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty ScopeLogs.
// It returns the newly added ScopeLogs.
func (es ScopeLogsSlice) AppendEmpty() ScopeLogs {
	*es.orig = append(*es.orig, &otlplogs.ScopeLogs{})
	return es.At(es.Len() - 1)
}

// Sort sorts the ScopeLogs elements within ScopeLogsSlice given the
// provided less function so that two instances of ScopeLogsSlice
// can be compared.
//
// Returns the same instance to allow nicer code like:
//   lessFunc := func(a, b ScopeLogs) bool {
//     return a.Name() < b.Name() // choose any comparison here
//   }
//   assert.EqualValues(t, expected.Sort(lessFunc), actual.Sort(lessFunc))
func (es ScopeLogsSlice) Sort(less func(a, b ScopeLogs) bool) ScopeLogsSlice {
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
	return es
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ScopeLogsSlice) MoveAndAppendTo(dest ScopeLogsSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ScopeLogsSlice) RemoveIf(f func(ScopeLogs) bool) {
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.orig = (*es.orig)[:newLen]
}

// ScopeLogs is a collection of logs from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewScopeLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeLogs struct {
	orig *otlplogs.ScopeLogs
}

func newScopeLogs(orig *otlplogs.ScopeLogs) ScopeLogs {
	return ScopeLogs{orig: orig}
}

// NewScopeLogs creates a new empty ScopeLogs.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewScopeLogs() ScopeLogs {
	return newScopeLogs(&otlplogs.ScopeLogs{})
}

// MoveTo moves all properties from the current struct to dest
// resetting the current instance to its zero value
func (ms ScopeLogs) MoveTo(dest ScopeLogs) {
	*dest.orig = *ms.orig
	*ms.orig = otlplogs.ScopeLogs{}
}

// Scope returns the scope associated with this ScopeLogs.
func (ms ScopeLogs) Scope() InstrumentationScope {
	return newInstrumentationScope(&(*ms.orig).Scope)
}

// SchemaUrl returns the schemaurl associated with this ScopeLogs.
func (ms ScopeLogs) SchemaUrl() string {
	return (*ms.orig).SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ScopeLogs.
func (ms ScopeLogs) SetSchemaUrl(v string) {
	(*ms.orig).SchemaUrl = v
}

// LogRecords returns the LogRecords associated with this ScopeLogs.
func (ms ScopeLogs) LogRecords() LogRecordSlice {
	return newLogRecordSlice(&(*ms.orig).LogRecords)
}

// CopyTo copies all properties from the current struct to the dest.
func (ms ScopeLogs) CopyTo(dest ScopeLogs) {
	ms.Scope().CopyTo(dest.Scope())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.LogRecords().CopyTo(dest.LogRecords())
}

// LogRecordSlice logically represents a slice of LogRecord.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewLogRecordSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecordSlice struct {
	// orig points to the slice otlplogs.LogRecord field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like EnsureCapacity.
	orig *[]*otlplogs.LogRecord
}

func newLogRecordSlice(orig *[]*otlplogs.LogRecord) LogRecordSlice {
	return LogRecordSlice{orig}
}

// NewLogRecordSlice creates a LogRecordSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewLogRecordSlice() LogRecordSlice {
	orig := []*otlplogs.LogRecord(nil)
	return LogRecordSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewLogRecordSlice()".
func (es LogRecordSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//   for i := 0; i < es.Len(); i++ {
//       e := es.At(i)
//       ... // Do something with the element
//   }
func (es LogRecordSlice) At(ix int) LogRecord {
	return newLogRecord((*es.orig)[ix])
}

// CopyTo copies all elements from the current slice to the dest.
func (es LogRecordSlice) CopyTo(dest LogRecordSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newLogRecord((*es.orig)[i]).CopyTo(newLogRecord((*dest.orig)[i]))
		}
		return
	}
	origs := make([]otlplogs.LogRecord, srcLen)
	wrappers := make([]*otlplogs.LogRecord, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newLogRecord((*es.orig)[i]).CopyTo(newLogRecord(wrappers[i]))
	}
	*dest.orig = wrappers
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new LogRecordSlice can be initialized:
//   es := NewLogRecordSlice()
//   es.EnsureCapacity(4)
//   for i := 0; i < 4; i++ {
//       e := es.AppendEmpty()
//       // Here should set all the values for e.
//   }
func (es LogRecordSlice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlplogs.LogRecord, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty LogRecord.
// It returns the newly added LogRecord.
func (es LogRecordSlice) AppendEmpty() LogRecord {
	*es.orig = append(*es.orig, &otlplogs.LogRecord{})
	return es.At(es.Len() - 1)
}

// Sort sorts the LogRecord elements within LogRecordSlice given the
// provided less function so that two instances of LogRecordSlice
// can be compared.
//
// Returns the same instance to allow nicer code like:
//   lessFunc := func(a, b LogRecord) bool {
//     return a.Name() < b.Name() // choose any comparison here
//   }
//   assert.EqualValues(t, expected.Sort(lessFunc), actual.Sort(lessFunc))
func (es LogRecordSlice) Sort(less func(a, b LogRecord) bool) LogRecordSlice {
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
	return es
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es LogRecordSlice) MoveAndAppendTo(dest LogRecordSlice) {
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es LogRecordSlice) RemoveIf(f func(LogRecord) bool) {
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.orig = (*es.orig)[:newLen]
}

// LogRecord are experimental implementation of OpenTelemetry Log Data Model.

//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogRecord function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecord struct {
	orig *otlplogs.LogRecord
}

func newLogRecord(orig *otlplogs.LogRecord) LogRecord {
	return LogRecord{orig: orig}
}

// NewLogRecord creates a new empty LogRecord.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewLogRecord() LogRecord {
	return newLogRecord(&otlplogs.LogRecord{})
}

// MoveTo moves all properties from the current struct to dest
// resetting the current instance to its zero value
func (ms LogRecord) MoveTo(dest LogRecord) {
	*dest.orig = *ms.orig
	*ms.orig = otlplogs.LogRecord{}
}

// ObservedTimestamp returns the observedtimestamp associated with this LogRecord.
func (ms LogRecord) ObservedTimestamp() Timestamp {
	return Timestamp((*ms.orig).ObservedTimeUnixNano)
}

// SetObservedTimestamp replaces the observedtimestamp associated with this LogRecord.
func (ms LogRecord) SetObservedTimestamp(v Timestamp) {
	(*ms.orig).ObservedTimeUnixNano = uint64(v)
}

// Timestamp returns the timestamp associated with this LogRecord.
func (ms LogRecord) Timestamp() Timestamp {
	return Timestamp((*ms.orig).TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this LogRecord.
func (ms LogRecord) SetTimestamp(v Timestamp) {
	(*ms.orig).TimeUnixNano = uint64(v)
}

// TraceID returns the traceid associated with this LogRecord.
func (ms LogRecord) TraceID() TraceID {
	return TraceID{orig: ((*ms.orig).TraceId)}
}

// SetTraceID replaces the traceid associated with this LogRecord.
func (ms LogRecord) SetTraceID(v TraceID) {
	(*ms.orig).TraceId = v.orig
}

// SpanID returns the spanid associated with this LogRecord.
func (ms LogRecord) SpanID() SpanID {
	return SpanID{orig: ((*ms.orig).SpanId)}
}

// SetSpanID replaces the spanid associated with this LogRecord.
func (ms LogRecord) SetSpanID(v SpanID) {
	(*ms.orig).SpanId = v.orig
}

// Flags returns the flags associated with this LogRecord.
func (ms LogRecord) Flags() uint32 {
	return uint32((*ms.orig).Flags)
}

// SetFlags replaces the flags associated with this LogRecord.
func (ms LogRecord) SetFlags(v uint32) {
	(*ms.orig).Flags = uint32(v)
}

// SeverityText returns the severitytext associated with this LogRecord.
func (ms LogRecord) SeverityText() string {
	return (*ms.orig).SeverityText
}

// SetSeverityText replaces the severitytext associated with this LogRecord.
func (ms LogRecord) SetSeverityText(v string) {
	(*ms.orig).SeverityText = v
}

// SeverityNumber returns the severitynumber associated with this LogRecord.
func (ms LogRecord) SeverityNumber() SeverityNumber {
	return SeverityNumber((*ms.orig).SeverityNumber)
}

// SetSeverityNumber replaces the severitynumber associated with this LogRecord.
func (ms LogRecord) SetSeverityNumber(v SeverityNumber) {
	(*ms.orig).SeverityNumber = otlplogs.SeverityNumber(v)
}

// Body returns the body associated with this LogRecord.
func (ms LogRecord) Body() Value {
	return newValue(&(*ms.orig).Body)
}

// Attributes returns the Attributes associated with this LogRecord.
func (ms LogRecord) Attributes() Map {
	return newMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this LogRecord.
func (ms LogRecord) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this LogRecord.
func (ms LogRecord) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct to the dest.
func (ms LogRecord) CopyTo(dest LogRecord) {
	dest.SetObservedTimestamp(ms.ObservedTimestamp())
	dest.SetTimestamp(ms.Timestamp())
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	dest.SetFlags(ms.Flags())
	dest.SetSeverityText(ms.SeverityText())
	dest.SetSeverityNumber(ms.SeverityNumber())
	ms.Body().CopyTo(dest.Body())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}
