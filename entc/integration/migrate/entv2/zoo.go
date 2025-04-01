// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv2/zoo"
)

// Zoo is the model entity for the Zoo schema.
type Zoo struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Zoo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case zoo.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Zoo fields.
func (_m *Zoo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case zoo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Zoo.
// This includes values selected through modifiers, order, etc.
func (_m *Zoo) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// Update returns a builder for updating this Zoo.
// Note that you need to call Zoo.Unwrap() before calling this method if this Zoo
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Zoo) Update() *ZooUpdateOne {
	return NewZooClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Zoo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Zoo) Unwrap() *Zoo {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Zoo is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Zoo) String() string {
	var builder strings.Builder
	builder.WriteString("Zoo(")
	builder.WriteString(fmt.Sprintf("id=%v", _m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Zoos is a parsable slice of Zoo.
type Zoos []*Zoo
