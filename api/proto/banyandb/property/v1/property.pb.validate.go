// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: banyandb/property/v1/property.proto

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

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Metadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Metadata with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetadataMultiError, or nil
// if none found.
func (m *Metadata) ValidateAll() error {
	return m.validate(true)
}

func (m *Metadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetContainer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetadataValidationError{
					field:  "Container",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetadataValidationError{
					field:  "Container",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContainer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataValidationError{
				field:  "Container",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Id

	if len(errors) > 0 {
		return MetadataMultiError(errors)
	}

	return nil
}

// MetadataMultiError is an error wrapping multiple validation errors returned
// by Metadata.ValidateAll() if the designated constraints aren't met.
type MetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetadataMultiError) AllErrors() []error { return m }

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

// Validate checks the field values on Property with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Property) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Property with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PropertyMultiError, or nil
// if none found.
func (m *Property) ValidateAll() error {
	return m.validate(true)
}

func (m *Property) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTags() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PropertyValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PropertyValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PropertyValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PropertyMultiError(errors)
	}

	return nil
}

// PropertyMultiError is an error wrapping multiple validation errors returned
// by Property.ValidateAll() if the designated constraints aren't met.
type PropertyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PropertyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PropertyMultiError) AllErrors() []error { return m }

// PropertyValidationError is the validation error returned by
// Property.Validate if the designated constraints aren't met.
type PropertyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertyValidationError) ErrorName() string { return "PropertyValidationError" }

// Error satisfies the builtin error interface
func (e PropertyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProperty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertyValidationError{}
