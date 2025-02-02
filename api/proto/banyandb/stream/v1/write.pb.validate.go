// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: banyandb/stream/v1/write.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ElementValue with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ElementValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ElementValue with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ElementValueMultiError, or
// nil if none found.
func (m *ElementValue) ValidateAll() error {
	return m.validate(true)
}

func (m *ElementValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ElementId

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ElementValueValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ElementValueValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ElementValueValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTagFamilies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ElementValueValidationError{
						field:  fmt.Sprintf("TagFamilies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ElementValueValidationError{
						field:  fmt.Sprintf("TagFamilies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ElementValueValidationError{
					field:  fmt.Sprintf("TagFamilies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ElementValueMultiError(errors)
	}

	return nil
}

// ElementValueMultiError is an error wrapping multiple validation errors
// returned by ElementValue.ValidateAll() if the designated constraints aren't met.
type ElementValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ElementValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ElementValueMultiError) AllErrors() []error { return m }

// ElementValueValidationError is the validation error returned by
// ElementValue.Validate if the designated constraints aren't met.
type ElementValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ElementValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ElementValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ElementValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ElementValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ElementValueValidationError) ErrorName() string { return "ElementValueValidationError" }

// Error satisfies the builtin error interface
func (e ElementValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sElementValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ElementValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ElementValueValidationError{}

// Validate checks the field values on WriteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WriteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WriteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WriteRequestMultiError, or
// nil if none found.
func (m *WriteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMetadata() == nil {
		err := WriteRequestValidationError{
			field:  "Metadata",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WriteRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WriteRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WriteRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetElement() == nil {
		err := WriteRequestValidationError{
			field:  "Element",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetElement()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WriteRequestValidationError{
					field:  "Element",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WriteRequestValidationError{
					field:  "Element",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetElement()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WriteRequestValidationError{
				field:  "Element",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WriteRequestMultiError(errors)
	}

	return nil
}

// WriteRequestMultiError is an error wrapping multiple validation errors
// returned by WriteRequest.ValidateAll() if the designated constraints aren't met.
type WriteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteRequestMultiError) AllErrors() []error { return m }

// WriteRequestValidationError is the validation error returned by
// WriteRequest.Validate if the designated constraints aren't met.
type WriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteRequestValidationError) ErrorName() string { return "WriteRequestValidationError" }

// Error satisfies the builtin error interface
func (e WriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteRequestValidationError{}

// Validate checks the field values on WriteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WriteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WriteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WriteResponseMultiError, or
// nil if none found.
func (m *WriteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return WriteResponseMultiError(errors)
	}

	return nil
}

// WriteResponseMultiError is an error wrapping multiple validation errors
// returned by WriteResponse.ValidateAll() if the designated constraints
// aren't met.
type WriteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteResponseMultiError) AllErrors() []error { return m }

// WriteResponseValidationError is the validation error returned by
// WriteResponse.Validate if the designated constraints aren't met.
type WriteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteResponseValidationError) ErrorName() string { return "WriteResponseValidationError" }

// Error satisfies the builtin error interface
func (e WriteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteResponseValidationError{}

// Validate checks the field values on InternalWriteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InternalWriteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InternalWriteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InternalWriteRequestMultiError, or nil if none found.
func (m *InternalWriteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InternalWriteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShardId

	// no validation rules for SeriesHash

	if all {
		switch v := interface{}(m.GetRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalWriteRequestValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalWriteRequestValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalWriteRequestValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return InternalWriteRequestMultiError(errors)
	}

	return nil
}

// InternalWriteRequestMultiError is an error wrapping multiple validation
// errors returned by InternalWriteRequest.ValidateAll() if the designated
// constraints aren't met.
type InternalWriteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InternalWriteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InternalWriteRequestMultiError) AllErrors() []error { return m }

// InternalWriteRequestValidationError is the validation error returned by
// InternalWriteRequest.Validate if the designated constraints aren't met.
type InternalWriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InternalWriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InternalWriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InternalWriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InternalWriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InternalWriteRequestValidationError) ErrorName() string {
	return "InternalWriteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InternalWriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInternalWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InternalWriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InternalWriteRequestValidationError{}
