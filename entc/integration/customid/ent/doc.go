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
	"entgo.io/ent/entc/integration/customid/ent/doc"
	"entgo.io/ent/entc/integration/customid/ent/schema"
)

// Doc is the model entity for the Doc schema.
type Doc struct {
	config `json:"-"`
	// ID of the ent.
	ID schema.DocID `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DocQuery when eager-loading is set.
	Edges        DocEdges `json:"edges"`
	doc_children *schema.DocID
	selectValues sql.SelectValues
}

// DocEdges holds the relations/edges for other nodes in the graph.
type DocEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Doc `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Doc `json:"children,omitempty"`
	// Related holds the value of the related edge.
	Related []*Doc `json:"related,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DocEdges) ParentOrErr() (*Doc, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: doc.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e DocEdges) ChildrenOrErr() ([]*Doc, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// RelatedOrErr returns the Related value or an error if the edge
// was not loaded in eager-loading.
func (e DocEdges) RelatedOrErr() ([]*Doc, error) {
	if e.loadedTypes[2] {
		return e.Related, nil
	}
	return nil, &NotLoadedError{edge: "related"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Doc) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case doc.FieldID:
			values[i] = new(schema.DocID)
		case doc.FieldText:
			values[i] = new(sql.NullString)
		case doc.ForeignKeys[0]: // doc_children
			values[i] = &sql.NullScanner{S: new(schema.DocID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Doc fields.
func (_m *Doc) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case doc.FieldID:
			if value, ok := values[i].(*schema.DocID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_m.ID = *value
			}
		case doc.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				_m.Text = value.String
			}
		case doc.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field doc_children", values[i])
			} else if value.Valid {
				_m.doc_children = new(schema.DocID)
				*_m.doc_children = *value.S.(*schema.DocID)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Doc.
// This includes values selected through modifiers, order, etc.
func (_m *Doc) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Doc entity.
func (_m *Doc) QueryParent() *DocQuery {
	return NewDocClient(_m.config).QueryParent(_m)
}

// QueryChildren queries the "children" edge of the Doc entity.
func (_m *Doc) QueryChildren() *DocQuery {
	return NewDocClient(_m.config).QueryChildren(_m)
}

// QueryRelated queries the "related" edge of the Doc entity.
func (_m *Doc) QueryRelated() *DocQuery {
	return NewDocClient(_m.config).QueryRelated(_m)
}

// Update returns a builder for updating this Doc.
// Note that you need to call Doc.Unwrap() before calling this method if this Doc
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Doc) Update() *DocUpdateOne {
	return NewDocClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Doc entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Doc) Unwrap() *Doc {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Doc is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Doc) String() string {
	var builder strings.Builder
	builder.WriteString("Doc(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("text=")
	builder.WriteString(_m.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Docs is a parsable slice of Doc.
type Docs []*Doc
