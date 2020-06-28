// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Item item
//
// swagger:model item
type Item struct {

	// completed
	Completed bool `json:"completed,omitempty"`

	// description
	// Required: true
	// Min Length: 1
	// Pattern: xyx-*-xyz
	Description string `json:"description"`

	// description2
	// Enum: [a b]
	Description2 string `json:"description2,omitempty"`

	// id
	// Read Only: true
	// Maximum: 10
	ID int64 `json:"id,omitempty"`

	// slice
	// Read Only: true
	Slice []*ItemSliceItems0 `json:"slice"`
}

// Validate validates this item
// TODO yy: remove general case
func (m *Item) Validate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", string(m.Description)); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(m.Description), 1); err != nil {
		return err
	}

	if err := validate.Pattern("description", "body", string(m.Description), `xyx-*-xyz`); err != nil {
		return err
	}

	return nil
}

var itemTypeDescription2PropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["a","b"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemTypeDescription2PropEnum = append(itemTypeDescription2PropEnum, v)
	}
}

const (

	// ItemDescription2A captures enum value "a"
	ItemDescription2A string = "a"

	// ItemDescription2B captures enum value "b"
	ItemDescription2B string = "b"
)

// prop value enum
func (m *Item) validateDescription2Enum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemTypeDescription2PropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Item) validateDescription2(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Description2) { // not required
		return nil
	}

	// value enum
	if err := m.validateDescription2Enum("description2", "body", m.Description2); err != nil {
		return err
	}

	return nil
}

func (m *Item) validateID(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	// WIP yy: Implement in validate pkg and replace temp code with this commented code
	// if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
	//    return err
	// }
	// READONLY HERE primitive
	if err := func(ctx context.Context, path string, in string, data interface{}) error {
		if v := ctx.Value("operation-type"); v != nil {
			if s, ok := v.(string); ok {
				if s != "Request" {
					// pass not request
					return nil
				}
				if !swag.IsZero(data) {
					return errors.New(400, fmt.Sprintf("ReadOnly field %v found in %v", in, path))
				}
			}
			return nil
		}
		// operation type not set so skip validating
		return nil
	}(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MaximumInt("id", "body", int64(m.ID), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *Item) validateSlice(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Slice) { // not required
		return nil
	}

	// READONLY For slice.
	// WIP yy: Implement in validate pkg and replace temp code with this commented code
	// if err := validate.ReadOnly(ctx, "slice", "body", m.Slice); err != nil {
	//    return err
	// }

	for i := 0; i < len(m.Slice); i++ {
		if swag.IsZero(m.Slice[i]) { // not required
			continue
		}

		if m.Slice[i] != nil {
			if err := m.Slice[i].Validate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ItemSliceItems0 item slice items0
//
// swagger:model ItemSliceItems0
type ItemSliceItems0 struct {

	// slice item content
	SliceItemContent *ItemSliceItems0SliceItemContent `json:"sliceItemContent,omitempty"`

	// slice item name read only
	// Required: true
	// Read Only: true
	SliceItemNameReadOnly string `json:"sliceItemNameReadOnly"`
}

// Validate validates this item slice items0
// TODO yy: remove general case
func (m *ItemSliceItems0) Validate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.validateSliceItemContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSliceItemNameReadOnly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSliceItems0) validateSliceItemContent(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SliceItemContent) { // not required
		return nil
	}

	if m.SliceItemContent != nil {
		if err := m.SliceItemContent.Validate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sliceItemContent")
			}
			return err
		}
	}

	return nil
}

func (m *ItemSliceItems0) validateSliceItemNameReadOnly(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.RequiredString("sliceItemNameReadOnly", "body", string(m.SliceItemNameReadOnly)); err != nil {
		return err
	}

	// WIP yy: Implement in validate pkg and replace temp code with this commented code
	// if err := validate.ReadOnly(ctx, "sliceItemNameReadOnly", "body", string(m.SliceItemNameReadOnly)); err != nil {
	//    return err
	// }
	// READONLY HERE primitive
	if err := func(ctx context.Context, path string, in string, data interface{}) error {
		if v := ctx.Value("operation-type"); v != nil {
			if s, ok := v.(string); ok {
				if s != "Request" {
					// pass not request
					return nil
				}
				if !swag.IsZero(data) {
					return errors.New(400, fmt.Sprintf("ReadOnly field %v found in %v", in, path))
				}
			}
			return nil
		}
		// operation type not set so skip validating
		return nil
	}(ctx, "sliceItemNameReadOnly", "body", m.SliceItemNameReadOnly); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemSliceItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemSliceItems0) UnmarshalBinary(b []byte) error {
	var res ItemSliceItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ItemSliceItems0SliceItemContent item slice items0 slice item content
//
// swagger:model ItemSliceItems0SliceItemContent
type ItemSliceItems0SliceItemContent struct {

	// age
	Age int64 `json:"Age,omitempty"`

	// content name read only
	// Read Only: true
	ContentNameReadOnly string `json:"ContentNameReadOnly,omitempty"`
}

// Validate validates this item slice items0 slice item content
// TODO yy: remove general case
func (m *ItemSliceItems0SliceItemContent) Validate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentNameReadOnly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSliceItems0SliceItemContent) validateContentNameReadOnly(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ContentNameReadOnly) { // not required
		return nil
	}

	// WIP yy: Implement in validate pkg and replace temp code with this commented code
	// if err := validate.ReadOnly(ctx, "sliceItemContent"+"."+"ContentNameReadOnly", "body", string(m.ContentNameReadOnly)); err != nil {
	//    return err
	// }
	// READONLY HERE primitive
	if err := func(ctx context.Context, path string, in string, data interface{}) error {
		if v := ctx.Value("operation-type"); v != nil {
			if s, ok := v.(string); ok {
				if s != "Request" {
					// pass not request
					return nil
				}
				if !swag.IsZero(data) {
					return errors.New(400, fmt.Sprintf("ReadOnly field %v found in %v", in, path))
				}
			}
			return nil
		}
		// operation type not set so skip validating
		return nil
	}(ctx, "sliceItemContent"+"."+"ContentNameReadOnly", "body", m.ContentNameReadOnly); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemSliceItems0SliceItemContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemSliceItems0SliceItemContent) UnmarshalBinary(b []byte) error {
	var res ItemSliceItems0SliceItemContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
