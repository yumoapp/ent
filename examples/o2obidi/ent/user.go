// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/o2obidi/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	user_spouse  *int
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Spouse holds the value of the spouse edge.
	Spouse *User `json:"spouse,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SpouseOrErr returns the Spouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SpouseOrErr() (*User, error) {
	if e.Spouse != nil {
		return e.Spouse, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "spouse"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldAge:
			values[i] = new(sql.NullInt64)
		case user.FieldName:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // user_spouse
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (_m *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case user.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				_m.Age = int(value.Int64)
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				_m.Name = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_spouse", value)
			} else if value.Valid {
				_m.user_spouse = new(int)
				*_m.user_spouse = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (_m *User) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QuerySpouse queries the "spouse" edge of the User entity.
func (_m *User) QuerySpouse() *UserQuery {
	return NewUserClient(_m.config).QuerySpouse(_m)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *User) Update() *UserUpdateOne {
	return NewUserClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *User) Unwrap() *User {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", _m.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
