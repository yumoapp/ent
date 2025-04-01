// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/license"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// LicenseUpdate is the builder for updating License entities.
type LicenseUpdate struct {
	config
	hooks     []Hook
	mutation  *LicenseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LicenseUpdate builder.
func (_u *LicenseUpdate) Where(ps ...predicate.License) *LicenseUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetUpdateTime sets the "update_time" field.
func (_u *LicenseUpdate) SetUpdateTime(v time.Time) *LicenseUpdate {
	_u.mutation.SetUpdateTime(v)
	return _u
}

// Mutation returns the LicenseMutation object of the builder.
func (_u *LicenseUpdate) Mutation() *LicenseMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *LicenseUpdate) Save(ctx context.Context) (int, error) {
	_u.defaults()
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *LicenseUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *LicenseUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *LicenseUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_u *LicenseUpdate) defaults() {
	if _, ok := _u.mutation.UpdateTime(); !ok {
		v := license.UpdateDefaultUpdateTime()
		_u.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *LicenseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LicenseUpdate {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *LicenseUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.UpdateTime(); ok {
		_spec.SetField(license.FieldUpdateTime, field.TypeTime, value)
	}
	_spec.AddModifiers(_u.modifiers...)
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// LicenseUpdateOne is the builder for updating a single License entity.
type LicenseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LicenseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (_u *LicenseUpdateOne) SetUpdateTime(v time.Time) *LicenseUpdateOne {
	_u.mutation.SetUpdateTime(v)
	return _u
}

// Mutation returns the LicenseMutation object of the builder.
func (_u *LicenseUpdateOne) Mutation() *LicenseMutation {
	return _u.mutation
}

// Where appends a list predicates to the LicenseUpdate builder.
func (_u *LicenseUpdateOne) Where(ps ...predicate.License) *LicenseUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *LicenseUpdateOne) Select(field string, fields ...string) *LicenseUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated License entity.
func (_u *LicenseUpdateOne) Save(ctx context.Context) (*License, error) {
	_u.defaults()
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *LicenseUpdateOne) SaveX(ctx context.Context) *License {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *LicenseUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *LicenseUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_u *LicenseUpdateOne) defaults() {
	if _, ok := _u.mutation.UpdateTime(); !ok {
		v := license.UpdateDefaultUpdateTime()
		_u.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *LicenseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LicenseUpdateOne {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *LicenseUpdateOne) sqlSave(ctx context.Context) (_node *License, err error) {
	_spec := sqlgraph.NewUpdateSpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "License.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, license.FieldID)
		for _, f := range fields {
			if !license.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != license.FieldID {
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
	if value, ok := _u.mutation.UpdateTime(); ok {
		_spec.SetField(license.FieldUpdateTime, field.TypeTime, value)
	}
	_spec.AddModifiers(_u.modifiers...)
	_node = &License{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{license.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
