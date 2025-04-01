// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/media"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/schema/field"
)

// MediaUpdate is the builder for updating Media entities.
type MediaUpdate struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where appends a list predicates to the MediaUpdate builder.
func (_u *MediaUpdate) Where(ps ...predicate.Media) *MediaUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetSource sets the "source" field.
func (_u *MediaUpdate) SetSource(v string) *MediaUpdate {
	_u.mutation.SetSource(v)
	return _u
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (_u *MediaUpdate) SetNillableSource(v *string) *MediaUpdate {
	if v != nil {
		_u.SetSource(*v)
	}
	return _u
}

// ClearSource clears the value of the "source" field.
func (_u *MediaUpdate) ClearSource() *MediaUpdate {
	_u.mutation.ClearSource()
	return _u
}

// SetSourceURI sets the "source_uri" field.
func (_u *MediaUpdate) SetSourceURI(v string) *MediaUpdate {
	_u.mutation.SetSourceURI(v)
	return _u
}

// SetNillableSourceURI sets the "source_uri" field if the given value is not nil.
func (_u *MediaUpdate) SetNillableSourceURI(v *string) *MediaUpdate {
	if v != nil {
		_u.SetSourceURI(*v)
	}
	return _u
}

// ClearSourceURI clears the value of the "source_uri" field.
func (_u *MediaUpdate) ClearSourceURI() *MediaUpdate {
	_u.mutation.ClearSourceURI()
	return _u
}

// SetText sets the "text" field.
func (_u *MediaUpdate) SetText(v string) *MediaUpdate {
	_u.mutation.SetText(v)
	return _u
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_u *MediaUpdate) SetNillableText(v *string) *MediaUpdate {
	if v != nil {
		_u.SetText(*v)
	}
	return _u
}

// ClearText clears the value of the "text" field.
func (_u *MediaUpdate) ClearText() *MediaUpdate {
	_u.mutation.ClearText()
	return _u
}

// Mutation returns the MediaMutation object of the builder.
func (_u *MediaUpdate) Mutation() *MediaMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *MediaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *MediaUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *MediaUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *MediaUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *MediaUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Source(); ok {
		_spec.SetField(media.FieldSource, field.TypeString, value)
	}
	if _u.mutation.SourceCleared() {
		_spec.ClearField(media.FieldSource, field.TypeString)
	}
	if value, ok := _u.mutation.SourceURI(); ok {
		_spec.SetField(media.FieldSourceURI, field.TypeString, value)
	}
	if _u.mutation.SourceURICleared() {
		_spec.ClearField(media.FieldSourceURI, field.TypeString)
	}
	if value, ok := _u.mutation.Text(); ok {
		_spec.SetField(media.FieldText, field.TypeString, value)
	}
	if _u.mutation.TextCleared() {
		_spec.ClearField(media.FieldText, field.TypeString)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// MediaUpdateOne is the builder for updating a single Media entity.
type MediaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MediaMutation
}

// SetSource sets the "source" field.
func (_u *MediaUpdateOne) SetSource(v string) *MediaUpdateOne {
	_u.mutation.SetSource(v)
	return _u
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (_u *MediaUpdateOne) SetNillableSource(v *string) *MediaUpdateOne {
	if v != nil {
		_u.SetSource(*v)
	}
	return _u
}

// ClearSource clears the value of the "source" field.
func (_u *MediaUpdateOne) ClearSource() *MediaUpdateOne {
	_u.mutation.ClearSource()
	return _u
}

// SetSourceURI sets the "source_uri" field.
func (_u *MediaUpdateOne) SetSourceURI(v string) *MediaUpdateOne {
	_u.mutation.SetSourceURI(v)
	return _u
}

// SetNillableSourceURI sets the "source_uri" field if the given value is not nil.
func (_u *MediaUpdateOne) SetNillableSourceURI(v *string) *MediaUpdateOne {
	if v != nil {
		_u.SetSourceURI(*v)
	}
	return _u
}

// ClearSourceURI clears the value of the "source_uri" field.
func (_u *MediaUpdateOne) ClearSourceURI() *MediaUpdateOne {
	_u.mutation.ClearSourceURI()
	return _u
}

// SetText sets the "text" field.
func (_u *MediaUpdateOne) SetText(v string) *MediaUpdateOne {
	_u.mutation.SetText(v)
	return _u
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_u *MediaUpdateOne) SetNillableText(v *string) *MediaUpdateOne {
	if v != nil {
		_u.SetText(*v)
	}
	return _u
}

// ClearText clears the value of the "text" field.
func (_u *MediaUpdateOne) ClearText() *MediaUpdateOne {
	_u.mutation.ClearText()
	return _u
}

// Mutation returns the MediaMutation object of the builder.
func (_u *MediaUpdateOne) Mutation() *MediaMutation {
	return _u.mutation
}

// Where appends a list predicates to the MediaUpdate builder.
func (_u *MediaUpdateOne) Where(ps ...predicate.Media) *MediaUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *MediaUpdateOne) Select(field string, fields ...string) *MediaUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Media entity.
func (_u *MediaUpdateOne) Save(ctx context.Context) (*Media, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *MediaUpdateOne) SaveX(ctx context.Context) *Media {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *MediaUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *MediaUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *MediaUpdateOne) sqlSave(ctx context.Context) (_node *Media, err error) {
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv2: missing "Media.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, media.FieldID)
		for _, f := range fields {
			if !media.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
			}
			if f != media.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Source(); ok {
		_spec.SetField(media.FieldSource, field.TypeString, value)
	}
	if _u.mutation.SourceCleared() {
		_spec.ClearField(media.FieldSource, field.TypeString)
	}
	if value, ok := _u.mutation.SourceURI(); ok {
		_spec.SetField(media.FieldSourceURI, field.TypeString, value)
	}
	if _u.mutation.SourceURICleared() {
		_spec.ClearField(media.FieldSourceURI, field.TypeString)
	}
	if value, ok := _u.mutation.Text(); ok {
		_spec.SetField(media.FieldText, field.TypeString, value)
	}
	if _u.mutation.TextCleared() {
		_spec.ClearField(media.FieldText, field.TypeString)
	}
	_node = &Media{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
