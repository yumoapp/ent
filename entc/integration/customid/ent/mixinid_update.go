// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/mixinid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// MixinIDUpdate is the builder for updating MixinID entities.
type MixinIDUpdate struct {
	config
	hooks    []Hook
	mutation *MixinIDMutation
}

// Where appends a list predicates to the MixinIDUpdate builder.
func (_u *MixinIDUpdate) Where(ps ...predicate.MixinID) *MixinIDUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetSomeField sets the "some_field" field.
func (_u *MixinIDUpdate) SetSomeField(v string) *MixinIDUpdate {
	_u.mutation.SetSomeField(v)
	return _u
}

// SetNillableSomeField sets the "some_field" field if the given value is not nil.
func (_u *MixinIDUpdate) SetNillableSomeField(v *string) *MixinIDUpdate {
	if v != nil {
		_u.SetSomeField(*v)
	}
	return _u
}

// SetMixinField sets the "mixin_field" field.
func (_u *MixinIDUpdate) SetMixinField(v string) *MixinIDUpdate {
	_u.mutation.SetMixinField(v)
	return _u
}

// SetNillableMixinField sets the "mixin_field" field if the given value is not nil.
func (_u *MixinIDUpdate) SetNillableMixinField(v *string) *MixinIDUpdate {
	if v != nil {
		_u.SetMixinField(*v)
	}
	return _u
}

// Mutation returns the MixinIDMutation object of the builder.
func (_u *MixinIDUpdate) Mutation() *MixinIDMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *MixinIDUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *MixinIDUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *MixinIDUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *MixinIDUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *MixinIDUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mixinid.Table, mixinid.Columns, sqlgraph.NewFieldSpec(mixinid.FieldID, field.TypeUUID))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.SomeField(); ok {
		_spec.SetField(mixinid.FieldSomeField, field.TypeString, value)
	}
	if value, ok := _u.mutation.MixinField(); ok {
		_spec.SetField(mixinid.FieldMixinField, field.TypeString, value)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mixinid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// MixinIDUpdateOne is the builder for updating a single MixinID entity.
type MixinIDUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MixinIDMutation
}

// SetSomeField sets the "some_field" field.
func (_u *MixinIDUpdateOne) SetSomeField(v string) *MixinIDUpdateOne {
	_u.mutation.SetSomeField(v)
	return _u
}

// SetNillableSomeField sets the "some_field" field if the given value is not nil.
func (_u *MixinIDUpdateOne) SetNillableSomeField(v *string) *MixinIDUpdateOne {
	if v != nil {
		_u.SetSomeField(*v)
	}
	return _u
}

// SetMixinField sets the "mixin_field" field.
func (_u *MixinIDUpdateOne) SetMixinField(v string) *MixinIDUpdateOne {
	_u.mutation.SetMixinField(v)
	return _u
}

// SetNillableMixinField sets the "mixin_field" field if the given value is not nil.
func (_u *MixinIDUpdateOne) SetNillableMixinField(v *string) *MixinIDUpdateOne {
	if v != nil {
		_u.SetMixinField(*v)
	}
	return _u
}

// Mutation returns the MixinIDMutation object of the builder.
func (_u *MixinIDUpdateOne) Mutation() *MixinIDMutation {
	return _u.mutation
}

// Where appends a list predicates to the MixinIDUpdate builder.
func (_u *MixinIDUpdateOne) Where(ps ...predicate.MixinID) *MixinIDUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *MixinIDUpdateOne) Select(field string, fields ...string) *MixinIDUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated MixinID entity.
func (_u *MixinIDUpdateOne) Save(ctx context.Context) (*MixinID, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *MixinIDUpdateOne) SaveX(ctx context.Context) *MixinID {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *MixinIDUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *MixinIDUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *MixinIDUpdateOne) sqlSave(ctx context.Context) (_node *MixinID, err error) {
	_spec := sqlgraph.NewUpdateSpec(mixinid.Table, mixinid.Columns, sqlgraph.NewFieldSpec(mixinid.FieldID, field.TypeUUID))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MixinID.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mixinid.FieldID)
		for _, f := range fields {
			if !mixinid.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mixinid.FieldID {
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
	if value, ok := _u.mutation.SomeField(); ok {
		_spec.SetField(mixinid.FieldSomeField, field.TypeString, value)
	}
	if value, ok := _u.mutation.MixinField(); ok {
		_spec.SetField(mixinid.FieldMixinField, field.TypeString, value)
	}
	_node = &MixinID{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mixinid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
