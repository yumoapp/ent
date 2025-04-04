// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/ent/role"
	"entgo.io/ent/entc/integration/ent/schema"
	"entgo.io/ent/entc/integration/gremlin/ent/fieldtype"
	"github.com/google/uuid"
)

// FieldType is the model entity for the FieldType schema.
type FieldType struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Int holds the value of the "int" field.
	Int int `json:"int,omitempty"`
	// Int8 holds the value of the "int8" field.
	Int8 int8 `json:"int8,omitempty"`
	// Int16 holds the value of the "int16" field.
	Int16 int16 `json:"int16,omitempty"`
	// Int32 holds the value of the "int32" field.
	Int32 int32 `json:"int32,omitempty"`
	// Int64 holds the value of the "int64" field.
	Int64 int64 `json:"int64,omitempty"`
	// OptionalInt holds the value of the "optional_int" field.
	OptionalInt int `json:"optional_int,omitempty"`
	// OptionalInt8 holds the value of the "optional_int8" field.
	OptionalInt8 int8 `json:"optional_int8,omitempty"`
	// OptionalInt16 holds the value of the "optional_int16" field.
	OptionalInt16 int16 `json:"optional_int16,omitempty"`
	// OptionalInt32 holds the value of the "optional_int32" field.
	OptionalInt32 int32 `json:"optional_int32,omitempty"`
	// OptionalInt64 holds the value of the "optional_int64" field.
	OptionalInt64 int64 `json:"optional_int64,omitempty"`
	// NillableInt holds the value of the "nillable_int" field.
	NillableInt *int `json:"nillable_int,omitempty"`
	// NillableInt8 holds the value of the "nillable_int8" field.
	NillableInt8 *int8 `json:"nillable_int8,omitempty"`
	// NillableInt16 holds the value of the "nillable_int16" field.
	NillableInt16 *int16 `json:"nillable_int16,omitempty"`
	// NillableInt32 holds the value of the "nillable_int32" field.
	NillableInt32 *int32 `json:"nillable_int32,omitempty"`
	// NillableInt64 holds the value of the "nillable_int64" field.
	NillableInt64 *int64 `json:"nillable_int64,omitempty"`
	// ValidateOptionalInt32 holds the value of the "validate_optional_int32" field.
	ValidateOptionalInt32 int32 `json:"validate_optional_int32,omitempty"`
	// OptionalUint holds the value of the "optional_uint" field.
	OptionalUint uint `json:"optional_uint,omitempty"`
	// OptionalUint8 holds the value of the "optional_uint8" field.
	OptionalUint8 uint8 `json:"optional_uint8,omitempty"`
	// OptionalUint16 holds the value of the "optional_uint16" field.
	OptionalUint16 uint16 `json:"optional_uint16,omitempty"`
	// OptionalUint32 holds the value of the "optional_uint32" field.
	OptionalUint32 uint32 `json:"optional_uint32,omitempty"`
	// OptionalUint64 holds the value of the "optional_uint64" field.
	OptionalUint64 uint64 `json:"optional_uint64,omitempty"`
	// State holds the value of the "state" field.
	State fieldtype.State `json:"state,omitempty"`
	// OptionalFloat holds the value of the "optional_float" field.
	OptionalFloat float64 `json:"optional_float,omitempty"`
	// OptionalFloat32 holds the value of the "optional_float32" field.
	OptionalFloat32 float32 `json:"optional_float32,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Decimal holds the value of the "decimal" field.
	Decimal float64 `json:"decimal,omitempty"`
	// LinkOther holds the value of the "link_other" field.
	LinkOther *schema.Link `json:"link_other,omitempty"`
	// LinkOtherFunc holds the value of the "link_other_func" field.
	LinkOtherFunc *schema.Link `json:"link_other_func,omitempty"`
	// MAC holds the value of the "mac" field.
	MAC schema.MAC `json:"mac,omitempty"`
	// StringArray holds the value of the "string_array" field.
	StringArray schema.Strings `json:"string_array,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// StringScanner holds the value of the "string_scanner" field.
	StringScanner *schema.StringScanner `json:"string_scanner,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration time.Duration `json:"duration,omitempty"`
	// Dir holds the value of the "dir" field.
	Dir http.Dir `json:"dir,omitempty"`
	// Ndir holds the value of the "ndir" field.
	Ndir *http.Dir `json:"ndir,omitempty"`
	// Str holds the value of the "str" field.
	Str sql.NullString `json:"str,omitempty"`
	// NullStr holds the value of the "null_str" field.
	NullStr *sql.NullString `json:"null_str,omitempty"`
	// Link holds the value of the "link" field.
	Link schema.Link `json:"link,omitempty"`
	// NullLink holds the value of the "null_link" field.
	NullLink *schema.Link `json:"null_link,omitempty"`
	// Active holds the value of the "active" field.
	Active schema.Status `json:"active,omitempty"`
	// NullActive holds the value of the "null_active" field.
	NullActive *schema.Status `json:"null_active,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted *sql.NullBool `json:"deleted,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *sql.NullTime `json:"deleted_at,omitempty"`
	// RawData holds the value of the "raw_data" field.
	RawData []byte `json:"raw_data,omitempty"`
	// Sensitive holds the value of the "sensitive" field.
	Sensitive []byte `json:"-"`
	// IP holds the value of the "ip" field.
	IP net.IP `json:"ip,omitempty"`
	// NullInt64 holds the value of the "null_int64" field.
	NullInt64 *sql.NullInt64 `json:"null_int64,omitempty"`
	// SchemaInt holds the value of the "schema_int" field.
	SchemaInt schema.Int `json:"schema_int,omitempty"`
	// SchemaInt8 holds the value of the "schema_int8" field.
	SchemaInt8 schema.Int8 `json:"schema_int8,omitempty"`
	// SchemaInt64 holds the value of the "schema_int64" field.
	SchemaInt64 schema.Int64 `json:"schema_int64,omitempty"`
	// SchemaFloat holds the value of the "schema_float" field.
	SchemaFloat schema.Float64 `json:"schema_float,omitempty"`
	// SchemaFloat32 holds the value of the "schema_float32" field.
	SchemaFloat32 schema.Float32 `json:"schema_float32,omitempty"`
	// NullFloat holds the value of the "null_float" field.
	NullFloat *sql.NullFloat64 `json:"null_float,omitempty"`
	// Role holds the value of the "role" field.
	Role role.Role `json:"role,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority role.Priority `json:"priority,omitempty"`
	// OptionalUUID holds the value of the "optional_uuid" field.
	OptionalUUID uuid.UUID `json:"optional_uuid,omitempty"`
	// NillableUUID holds the value of the "nillable_uuid" field.
	NillableUUID *uuid.UUID `json:"nillable_uuid,omitempty"`
	// Strings holds the value of the "strings" field.
	Strings []string `json:"strings,omitempty"`
	// Pair holds the value of the "pair" field.
	Pair schema.Pair `json:"pair,omitempty"`
	// NilPair holds the value of the "nil_pair" field.
	NilPair *schema.Pair `json:"nil_pair,omitempty"`
	// Vstring holds the value of the "vstring" field.
	Vstring schema.VString `json:"vstring,omitempty"`
	// Triple holds the value of the "triple" field.
	Triple schema.Triple `json:"triple,omitempty"`
	// BigInt holds the value of the "big_int" field.
	BigInt schema.BigInt `json:"big_int,omitempty"`
	// PasswordOther holds the value of the "password_other" field.
	PasswordOther schema.Password `json:"-"`
}

// FromResponse scans the gremlin response data into FieldType.
func (_m *FieldType) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m struct {
		ID                    string                `json:"id,omitempty"`
		Int                   int                   `json:"int,omitempty"`
		Int8                  int8                  `json:"int8,omitempty"`
		Int16                 int16                 `json:"int16,omitempty"`
		Int32                 int32                 `json:"int32,omitempty"`
		Int64                 int64                 `json:"int64,omitempty"`
		OptionalInt           int                   `json:"optional_int,omitempty"`
		OptionalInt8          int8                  `json:"optional_int8,omitempty"`
		OptionalInt16         int16                 `json:"optional_int16,omitempty"`
		OptionalInt32         int32                 `json:"optional_int32,omitempty"`
		OptionalInt64         int64                 `json:"optional_int64,omitempty"`
		NillableInt           *int                  `json:"nillable_int,omitempty"`
		NillableInt8          *int8                 `json:"nillable_int8,omitempty"`
		NillableInt16         *int16                `json:"nillable_int16,omitempty"`
		NillableInt32         *int32                `json:"nillable_int32,omitempty"`
		NillableInt64         *int64                `json:"nillable_int64,omitempty"`
		ValidateOptionalInt32 int32                 `json:"validate_optional_int32,omitempty"`
		OptionalUint          uint                  `json:"optional_uint,omitempty"`
		OptionalUint8         uint8                 `json:"optional_uint8,omitempty"`
		OptionalUint16        uint16                `json:"optional_uint16,omitempty"`
		OptionalUint32        uint32                `json:"optional_uint32,omitempty"`
		OptionalUint64        uint64                `json:"optional_uint64,omitempty"`
		State                 fieldtype.State       `json:"state,omitempty"`
		OptionalFloat         float64               `json:"optional_float,omitempty"`
		OptionalFloat32       float32               `json:"optional_float32,omitempty"`
		Text                  string                `json:"text,omitempty"`
		Datetime              int64                 `json:"datetime,omitempty"`
		Decimal               float64               `json:"decimal,omitempty"`
		LinkOther             *schema.Link          `json:"link_other,omitempty"`
		LinkOtherFunc         *schema.Link          `json:"link_other_func,omitempty"`
		MAC                   schema.MAC            `json:"mac,omitempty"`
		StringArray           schema.Strings        `json:"string_array,omitempty"`
		Password              string                `json:"password,omitempty"`
		StringScanner         *schema.StringScanner `json:"string_scanner,omitempty"`
		Duration              time.Duration         `json:"duration,omitempty"`
		Dir                   http.Dir              `json:"dir,omitempty"`
		Ndir                  *http.Dir             `json:"ndir,omitempty"`
		Str                   sql.NullString        `json:"str,omitempty"`
		NullStr               *sql.NullString       `json:"null_str,omitempty"`
		Link                  schema.Link           `json:"link,omitempty"`
		NullLink              *schema.Link          `json:"null_link,omitempty"`
		Active                schema.Status         `json:"active,omitempty"`
		NullActive            *schema.Status        `json:"null_active,omitempty"`
		Deleted               *sql.NullBool         `json:"deleted,omitempty"`
		DeletedAt             *sql.NullTime         `json:"deleted_at,omitempty"`
		RawData               []byte                `json:"raw_data,omitempty"`
		Sensitive             []byte                `json:"sensitive,omitempty"`
		IP                    net.IP                `json:"ip,omitempty"`
		NullInt64             *sql.NullInt64        `json:"null_int64,omitempty"`
		SchemaInt             schema.Int            `json:"schema_int,omitempty"`
		SchemaInt8            schema.Int8           `json:"schema_int8,omitempty"`
		SchemaInt64           schema.Int64          `json:"schema_int64,omitempty"`
		SchemaFloat           schema.Float64        `json:"schema_float,omitempty"`
		SchemaFloat32         schema.Float32        `json:"schema_float32,omitempty"`
		NullFloat             *sql.NullFloat64      `json:"null_float,omitempty"`
		Role                  role.Role             `json:"role,omitempty"`
		Priority              role.Priority         `json:"priority,omitempty"`
		OptionalUUID          uuid.UUID             `json:"optional_uuid,omitempty"`
		NillableUUID          *uuid.UUID            `json:"nillable_uuid,omitempty"`
		Strings               []string              `json:"strings,omitempty"`
		Pair                  schema.Pair           `json:"pair,omitempty"`
		NilPair               *schema.Pair          `json:"nil_pair,omitempty"`
		Vstring               schema.VString        `json:"vstring,omitempty"`
		Triple                schema.Triple         `json:"triple,omitempty"`
		BigInt                schema.BigInt         `json:"big_int,omitempty"`
		PasswordOther         schema.Password       `json:"password_other,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	_m.ID = scan_m.ID
	_m.Int = scan_m.Int
	_m.Int8 = scan_m.Int8
	_m.Int16 = scan_m.Int16
	_m.Int32 = scan_m.Int32
	_m.Int64 = scan_m.Int64
	_m.OptionalInt = scan_m.OptionalInt
	_m.OptionalInt8 = scan_m.OptionalInt8
	_m.OptionalInt16 = scan_m.OptionalInt16
	_m.OptionalInt32 = scan_m.OptionalInt32
	_m.OptionalInt64 = scan_m.OptionalInt64
	_m.NillableInt = scan_m.NillableInt
	_m.NillableInt8 = scan_m.NillableInt8
	_m.NillableInt16 = scan_m.NillableInt16
	_m.NillableInt32 = scan_m.NillableInt32
	_m.NillableInt64 = scan_m.NillableInt64
	_m.ValidateOptionalInt32 = scan_m.ValidateOptionalInt32
	_m.OptionalUint = scan_m.OptionalUint
	_m.OptionalUint8 = scan_m.OptionalUint8
	_m.OptionalUint16 = scan_m.OptionalUint16
	_m.OptionalUint32 = scan_m.OptionalUint32
	_m.OptionalUint64 = scan_m.OptionalUint64
	_m.State = scan_m.State
	_m.OptionalFloat = scan_m.OptionalFloat
	_m.OptionalFloat32 = scan_m.OptionalFloat32
	_m.Text = scan_m.Text
	_m.Datetime = time.Unix(0, scan_m.Datetime)
	_m.Decimal = scan_m.Decimal
	_m.LinkOther = scan_m.LinkOther
	_m.LinkOtherFunc = scan_m.LinkOtherFunc
	_m.MAC = scan_m.MAC
	_m.StringArray = scan_m.StringArray
	_m.Password = scan_m.Password
	_m.StringScanner = scan_m.StringScanner
	_m.Duration = scan_m.Duration
	_m.Dir = scan_m.Dir
	_m.Ndir = scan_m.Ndir
	_m.Str = scan_m.Str
	_m.NullStr = scan_m.NullStr
	_m.Link = scan_m.Link
	_m.NullLink = scan_m.NullLink
	_m.Active = scan_m.Active
	_m.NullActive = scan_m.NullActive
	_m.Deleted = scan_m.Deleted
	_m.DeletedAt = scan_m.DeletedAt
	_m.RawData = scan_m.RawData
	_m.Sensitive = scan_m.Sensitive
	_m.IP = scan_m.IP
	_m.NullInt64 = scan_m.NullInt64
	_m.SchemaInt = scan_m.SchemaInt
	_m.SchemaInt8 = scan_m.SchemaInt8
	_m.SchemaInt64 = scan_m.SchemaInt64
	_m.SchemaFloat = scan_m.SchemaFloat
	_m.SchemaFloat32 = scan_m.SchemaFloat32
	_m.NullFloat = scan_m.NullFloat
	_m.Role = scan_m.Role
	_m.Priority = scan_m.Priority
	_m.OptionalUUID = scan_m.OptionalUUID
	_m.NillableUUID = scan_m.NillableUUID
	_m.Strings = scan_m.Strings
	_m.Pair = scan_m.Pair
	_m.NilPair = scan_m.NilPair
	_m.Vstring = scan_m.Vstring
	_m.Triple = scan_m.Triple
	_m.BigInt = scan_m.BigInt
	_m.PasswordOther = scan_m.PasswordOther
	return nil
}

// Update returns a builder for updating this FieldType.
// Note that you need to call FieldType.Unwrap() before calling this method if this FieldType
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *FieldType) Update() *FieldTypeUpdateOne {
	return NewFieldTypeClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the FieldType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *FieldType) Unwrap() *FieldType {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: FieldType is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *FieldType) String() string {
	var builder strings.Builder
	builder.WriteString("FieldType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("int=")
	builder.WriteString(fmt.Sprintf("%v", _m.Int))
	builder.WriteString(", ")
	builder.WriteString("int8=")
	builder.WriteString(fmt.Sprintf("%v", _m.Int8))
	builder.WriteString(", ")
	builder.WriteString("int16=")
	builder.WriteString(fmt.Sprintf("%v", _m.Int16))
	builder.WriteString(", ")
	builder.WriteString("int32=")
	builder.WriteString(fmt.Sprintf("%v", _m.Int32))
	builder.WriteString(", ")
	builder.WriteString("int64=")
	builder.WriteString(fmt.Sprintf("%v", _m.Int64))
	builder.WriteString(", ")
	builder.WriteString("optional_int=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalInt))
	builder.WriteString(", ")
	builder.WriteString("optional_int8=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalInt8))
	builder.WriteString(", ")
	builder.WriteString("optional_int16=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalInt16))
	builder.WriteString(", ")
	builder.WriteString("optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalInt32))
	builder.WriteString(", ")
	builder.WriteString("optional_int64=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalInt64))
	builder.WriteString(", ")
	if v := _m.NillableInt; v != nil {
		builder.WriteString("nillable_int=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := _m.NillableInt8; v != nil {
		builder.WriteString("nillable_int8=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := _m.NillableInt16; v != nil {
		builder.WriteString("nillable_int16=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := _m.NillableInt32; v != nil {
		builder.WriteString("nillable_int32=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := _m.NillableInt64; v != nil {
		builder.WriteString("nillable_int64=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("validate_optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", _m.ValidateOptionalInt32))
	builder.WriteString(", ")
	builder.WriteString("optional_uint=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalUint))
	builder.WriteString(", ")
	builder.WriteString("optional_uint8=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalUint8))
	builder.WriteString(", ")
	builder.WriteString("optional_uint16=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalUint16))
	builder.WriteString(", ")
	builder.WriteString("optional_uint32=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalUint32))
	builder.WriteString(", ")
	builder.WriteString("optional_uint64=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalUint64))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", _m.State))
	builder.WriteString(", ")
	builder.WriteString("optional_float=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalFloat))
	builder.WriteString(", ")
	builder.WriteString("optional_float32=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalFloat32))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(_m.Text)
	builder.WriteString(", ")
	builder.WriteString("datetime=")
	builder.WriteString(_m.Datetime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("decimal=")
	builder.WriteString(fmt.Sprintf("%v", _m.Decimal))
	builder.WriteString(", ")
	builder.WriteString("link_other=")
	builder.WriteString(fmt.Sprintf("%v", _m.LinkOther))
	builder.WriteString(", ")
	builder.WriteString("link_other_func=")
	builder.WriteString(fmt.Sprintf("%v", _m.LinkOtherFunc))
	builder.WriteString(", ")
	builder.WriteString("mac=")
	builder.WriteString(fmt.Sprintf("%v", _m.MAC))
	builder.WriteString(", ")
	builder.WriteString("string_array=")
	builder.WriteString(fmt.Sprintf("%v", _m.StringArray))
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	if v := _m.StringScanner; v != nil {
		builder.WriteString("string_scanner=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", _m.Duration))
	builder.WriteString(", ")
	builder.WriteString("dir=")
	builder.WriteString(fmt.Sprintf("%v", _m.Dir))
	builder.WriteString(", ")
	if v := _m.Ndir; v != nil {
		builder.WriteString("ndir=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("str=")
	builder.WriteString(fmt.Sprintf("%v", _m.Str))
	builder.WriteString(", ")
	if v := _m.NullStr; v != nil {
		builder.WriteString("null_str=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("link=")
	builder.WriteString(fmt.Sprintf("%v", _m.Link))
	builder.WriteString(", ")
	if v := _m.NullLink; v != nil {
		builder.WriteString("null_link=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", _m.Active))
	builder.WriteString(", ")
	if v := _m.NullActive; v != nil {
		builder.WriteString("null_active=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := _m.Deleted; v != nil {
		builder.WriteString("deleted=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", _m.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("raw_data=")
	builder.WriteString(fmt.Sprintf("%v", _m.RawData))
	builder.WriteString(", ")
	builder.WriteString("sensitive=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(fmt.Sprintf("%v", _m.IP))
	builder.WriteString(", ")
	builder.WriteString("null_int64=")
	builder.WriteString(fmt.Sprintf("%v", _m.NullInt64))
	builder.WriteString(", ")
	builder.WriteString("schema_int=")
	builder.WriteString(fmt.Sprintf("%v", _m.SchemaInt))
	builder.WriteString(", ")
	builder.WriteString("schema_int8=")
	builder.WriteString(fmt.Sprintf("%v", _m.SchemaInt8))
	builder.WriteString(", ")
	builder.WriteString("schema_int64=")
	builder.WriteString(fmt.Sprintf("%v", _m.SchemaInt64))
	builder.WriteString(", ")
	builder.WriteString("schema_float=")
	builder.WriteString(fmt.Sprintf("%v", _m.SchemaFloat))
	builder.WriteString(", ")
	builder.WriteString("schema_float32=")
	builder.WriteString(fmt.Sprintf("%v", _m.SchemaFloat32))
	builder.WriteString(", ")
	builder.WriteString("null_float=")
	builder.WriteString(fmt.Sprintf("%v", _m.NullFloat))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", _m.Role))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", _m.Priority))
	builder.WriteString(", ")
	builder.WriteString("optional_uuid=")
	builder.WriteString(fmt.Sprintf("%v", _m.OptionalUUID))
	builder.WriteString(", ")
	if v := _m.NillableUUID; v != nil {
		builder.WriteString("nillable_uuid=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("strings=")
	builder.WriteString(fmt.Sprintf("%v", _m.Strings))
	builder.WriteString(", ")
	builder.WriteString("pair=")
	builder.WriteString(fmt.Sprintf("%v", _m.Pair))
	builder.WriteString(", ")
	if v := _m.NilPair; v != nil {
		builder.WriteString("nil_pair=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("vstring=")
	builder.WriteString(fmt.Sprintf("%v", _m.Vstring))
	builder.WriteString(", ")
	builder.WriteString("triple=")
	builder.WriteString(fmt.Sprintf("%v", _m.Triple))
	builder.WriteString(", ")
	builder.WriteString("big_int=")
	builder.WriteString(fmt.Sprintf("%v", _m.BigInt))
	builder.WriteString(", ")
	builder.WriteString("password_other=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// FieldTypes is a parsable slice of FieldType.
type FieldTypes []*FieldType

// FromResponse scans the gremlin response data into FieldTypes.
func (_m *FieldTypes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m []struct {
		ID                    string                `json:"id,omitempty"`
		Int                   int                   `json:"int,omitempty"`
		Int8                  int8                  `json:"int8,omitempty"`
		Int16                 int16                 `json:"int16,omitempty"`
		Int32                 int32                 `json:"int32,omitempty"`
		Int64                 int64                 `json:"int64,omitempty"`
		OptionalInt           int                   `json:"optional_int,omitempty"`
		OptionalInt8          int8                  `json:"optional_int8,omitempty"`
		OptionalInt16         int16                 `json:"optional_int16,omitempty"`
		OptionalInt32         int32                 `json:"optional_int32,omitempty"`
		OptionalInt64         int64                 `json:"optional_int64,omitempty"`
		NillableInt           *int                  `json:"nillable_int,omitempty"`
		NillableInt8          *int8                 `json:"nillable_int8,omitempty"`
		NillableInt16         *int16                `json:"nillable_int16,omitempty"`
		NillableInt32         *int32                `json:"nillable_int32,omitempty"`
		NillableInt64         *int64                `json:"nillable_int64,omitempty"`
		ValidateOptionalInt32 int32                 `json:"validate_optional_int32,omitempty"`
		OptionalUint          uint                  `json:"optional_uint,omitempty"`
		OptionalUint8         uint8                 `json:"optional_uint8,omitempty"`
		OptionalUint16        uint16                `json:"optional_uint16,omitempty"`
		OptionalUint32        uint32                `json:"optional_uint32,omitempty"`
		OptionalUint64        uint64                `json:"optional_uint64,omitempty"`
		State                 fieldtype.State       `json:"state,omitempty"`
		OptionalFloat         float64               `json:"optional_float,omitempty"`
		OptionalFloat32       float32               `json:"optional_float32,omitempty"`
		Text                  string                `json:"text,omitempty"`
		Datetime              int64                 `json:"datetime,omitempty"`
		Decimal               float64               `json:"decimal,omitempty"`
		LinkOther             *schema.Link          `json:"link_other,omitempty"`
		LinkOtherFunc         *schema.Link          `json:"link_other_func,omitempty"`
		MAC                   schema.MAC            `json:"mac,omitempty"`
		StringArray           schema.Strings        `json:"string_array,omitempty"`
		Password              string                `json:"password,omitempty"`
		StringScanner         *schema.StringScanner `json:"string_scanner,omitempty"`
		Duration              time.Duration         `json:"duration,omitempty"`
		Dir                   http.Dir              `json:"dir,omitempty"`
		Ndir                  *http.Dir             `json:"ndir,omitempty"`
		Str                   sql.NullString        `json:"str,omitempty"`
		NullStr               *sql.NullString       `json:"null_str,omitempty"`
		Link                  schema.Link           `json:"link,omitempty"`
		NullLink              *schema.Link          `json:"null_link,omitempty"`
		Active                schema.Status         `json:"active,omitempty"`
		NullActive            *schema.Status        `json:"null_active,omitempty"`
		Deleted               *sql.NullBool         `json:"deleted,omitempty"`
		DeletedAt             *sql.NullTime         `json:"deleted_at,omitempty"`
		RawData               []byte                `json:"raw_data,omitempty"`
		Sensitive             []byte                `json:"sensitive,omitempty"`
		IP                    net.IP                `json:"ip,omitempty"`
		NullInt64             *sql.NullInt64        `json:"null_int64,omitempty"`
		SchemaInt             schema.Int            `json:"schema_int,omitempty"`
		SchemaInt8            schema.Int8           `json:"schema_int8,omitempty"`
		SchemaInt64           schema.Int64          `json:"schema_int64,omitempty"`
		SchemaFloat           schema.Float64        `json:"schema_float,omitempty"`
		SchemaFloat32         schema.Float32        `json:"schema_float32,omitempty"`
		NullFloat             *sql.NullFloat64      `json:"null_float,omitempty"`
		Role                  role.Role             `json:"role,omitempty"`
		Priority              role.Priority         `json:"priority,omitempty"`
		OptionalUUID          uuid.UUID             `json:"optional_uuid,omitempty"`
		NillableUUID          *uuid.UUID            `json:"nillable_uuid,omitempty"`
		Strings               []string              `json:"strings,omitempty"`
		Pair                  schema.Pair           `json:"pair,omitempty"`
		NilPair               *schema.Pair          `json:"nil_pair,omitempty"`
		Vstring               schema.VString        `json:"vstring,omitempty"`
		Triple                schema.Triple         `json:"triple,omitempty"`
		BigInt                schema.BigInt         `json:"big_int,omitempty"`
		PasswordOther         schema.Password       `json:"password_other,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	for _, v := range scan_m {
		node := &FieldType{ID: v.ID}
		node.Int = v.Int
		node.Int8 = v.Int8
		node.Int16 = v.Int16
		node.Int32 = v.Int32
		node.Int64 = v.Int64
		node.OptionalInt = v.OptionalInt
		node.OptionalInt8 = v.OptionalInt8
		node.OptionalInt16 = v.OptionalInt16
		node.OptionalInt32 = v.OptionalInt32
		node.OptionalInt64 = v.OptionalInt64
		node.NillableInt = v.NillableInt
		node.NillableInt8 = v.NillableInt8
		node.NillableInt16 = v.NillableInt16
		node.NillableInt32 = v.NillableInt32
		node.NillableInt64 = v.NillableInt64
		node.ValidateOptionalInt32 = v.ValidateOptionalInt32
		node.OptionalUint = v.OptionalUint
		node.OptionalUint8 = v.OptionalUint8
		node.OptionalUint16 = v.OptionalUint16
		node.OptionalUint32 = v.OptionalUint32
		node.OptionalUint64 = v.OptionalUint64
		node.State = v.State
		node.OptionalFloat = v.OptionalFloat
		node.OptionalFloat32 = v.OptionalFloat32
		node.Text = v.Text
		node.Datetime = time.Unix(0, v.Datetime)
		node.Decimal = v.Decimal
		node.LinkOther = v.LinkOther
		node.LinkOtherFunc = v.LinkOtherFunc
		node.MAC = v.MAC
		node.StringArray = v.StringArray
		node.Password = v.Password
		node.StringScanner = v.StringScanner
		node.Duration = v.Duration
		node.Dir = v.Dir
		node.Ndir = v.Ndir
		node.Str = v.Str
		node.NullStr = v.NullStr
		node.Link = v.Link
		node.NullLink = v.NullLink
		node.Active = v.Active
		node.NullActive = v.NullActive
		node.Deleted = v.Deleted
		node.DeletedAt = v.DeletedAt
		node.RawData = v.RawData
		node.Sensitive = v.Sensitive
		node.IP = v.IP
		node.NullInt64 = v.NullInt64
		node.SchemaInt = v.SchemaInt
		node.SchemaInt8 = v.SchemaInt8
		node.SchemaInt64 = v.SchemaInt64
		node.SchemaFloat = v.SchemaFloat
		node.SchemaFloat32 = v.SchemaFloat32
		node.NullFloat = v.NullFloat
		node.Role = v.Role
		node.Priority = v.Priority
		node.OptionalUUID = v.OptionalUUID
		node.NillableUUID = v.NillableUUID
		node.Strings = v.Strings
		node.Pair = v.Pair
		node.NilPair = v.NilPair
		node.Vstring = v.Vstring
		node.Triple = v.Triple
		node.BigInt = v.BigInt
		node.PasswordOther = v.PasswordOther
		*_m = append(*_m, node)
	}
	return nil
}
