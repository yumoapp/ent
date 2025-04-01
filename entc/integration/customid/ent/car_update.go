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
	"entgo.io/ent/entc/integration/customid/ent/car"
	"entgo.io/ent/entc/integration/customid/ent/pet"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (_u *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetBeforeID sets the "before_id" field.
func (_u *CarUpdate) SetBeforeID(v float64) *CarUpdate {
	_u.mutation.ResetBeforeID()
	_u.mutation.SetBeforeID(v)
	return _u
}

// SetNillableBeforeID sets the "before_id" field if the given value is not nil.
func (_u *CarUpdate) SetNillableBeforeID(v *float64) *CarUpdate {
	if v != nil {
		_u.SetBeforeID(*v)
	}
	return _u
}

// AddBeforeID adds value to the "before_id" field.
func (_u *CarUpdate) AddBeforeID(v float64) *CarUpdate {
	_u.mutation.AddBeforeID(v)
	return _u
}

// ClearBeforeID clears the value of the "before_id" field.
func (_u *CarUpdate) ClearBeforeID() *CarUpdate {
	_u.mutation.ClearBeforeID()
	return _u
}

// SetAfterID sets the "after_id" field.
func (_u *CarUpdate) SetAfterID(v float64) *CarUpdate {
	_u.mutation.ResetAfterID()
	_u.mutation.SetAfterID(v)
	return _u
}

// SetNillableAfterID sets the "after_id" field if the given value is not nil.
func (_u *CarUpdate) SetNillableAfterID(v *float64) *CarUpdate {
	if v != nil {
		_u.SetAfterID(*v)
	}
	return _u
}

// AddAfterID adds value to the "after_id" field.
func (_u *CarUpdate) AddAfterID(v float64) *CarUpdate {
	_u.mutation.AddAfterID(v)
	return _u
}

// ClearAfterID clears the value of the "after_id" field.
func (_u *CarUpdate) ClearAfterID() *CarUpdate {
	_u.mutation.ClearAfterID()
	return _u
}

// SetModel sets the "model" field.
func (_u *CarUpdate) SetModel(v string) *CarUpdate {
	_u.mutation.SetModel(v)
	return _u
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (_u *CarUpdate) SetNillableModel(v *string) *CarUpdate {
	if v != nil {
		_u.SetModel(*v)
	}
	return _u
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (_u *CarUpdate) SetOwnerID(id string) *CarUpdate {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (_u *CarUpdate) SetNillableOwnerID(id *string) *CarUpdate {
	if id != nil {
		_u = _u.SetOwnerID(*id)
	}
	return _u
}

// SetOwner sets the "owner" edge to the Pet entity.
func (_u *CarUpdate) SetOwner(v *Pet) *CarUpdate {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the CarMutation object of the builder.
func (_u *CarUpdate) Mutation() *CarMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the Pet entity.
func (_u *CarUpdate) ClearOwner() *CarUpdate {
	_u.mutation.ClearOwner()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *CarUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *CarUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CarUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *CarUpdate) check() error {
	if v, ok := _u.mutation.BeforeID(); ok {
		if err := car.BeforeIDValidator(v); err != nil {
			return &ValidationError{Name: "before_id", err: fmt.Errorf(`ent: validator failed for field "Car.before_id": %w`, err)}
		}
	}
	if v, ok := _u.mutation.AfterID(); ok {
		if err := car.AfterIDValidator(v); err != nil {
			return &ValidationError{Name: "after_id", err: fmt.Errorf(`ent: validator failed for field "Car.after_id": %w`, err)}
		}
	}
	return nil
}

func (_u *CarUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.BeforeID(); ok {
		_spec.SetField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if value, ok := _u.mutation.AddedBeforeID(); ok {
		_spec.AddField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if _u.mutation.BeforeIDCleared() {
		_spec.ClearField(car.FieldBeforeID, field.TypeFloat64)
	}
	if value, ok := _u.mutation.AfterID(); ok {
		_spec.SetField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if value, ok := _u.mutation.AddedAfterID(); ok {
		_spec.AddField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if _u.mutation.AfterIDCleared() {
		_spec.ClearField(car.FieldAfterID, field.TypeFloat64)
	}
	if value, ok := _u.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetBeforeID sets the "before_id" field.
func (_u *CarUpdateOne) SetBeforeID(v float64) *CarUpdateOne {
	_u.mutation.ResetBeforeID()
	_u.mutation.SetBeforeID(v)
	return _u
}

// SetNillableBeforeID sets the "before_id" field if the given value is not nil.
func (_u *CarUpdateOne) SetNillableBeforeID(v *float64) *CarUpdateOne {
	if v != nil {
		_u.SetBeforeID(*v)
	}
	return _u
}

// AddBeforeID adds value to the "before_id" field.
func (_u *CarUpdateOne) AddBeforeID(v float64) *CarUpdateOne {
	_u.mutation.AddBeforeID(v)
	return _u
}

// ClearBeforeID clears the value of the "before_id" field.
func (_u *CarUpdateOne) ClearBeforeID() *CarUpdateOne {
	_u.mutation.ClearBeforeID()
	return _u
}

// SetAfterID sets the "after_id" field.
func (_u *CarUpdateOne) SetAfterID(v float64) *CarUpdateOne {
	_u.mutation.ResetAfterID()
	_u.mutation.SetAfterID(v)
	return _u
}

// SetNillableAfterID sets the "after_id" field if the given value is not nil.
func (_u *CarUpdateOne) SetNillableAfterID(v *float64) *CarUpdateOne {
	if v != nil {
		_u.SetAfterID(*v)
	}
	return _u
}

// AddAfterID adds value to the "after_id" field.
func (_u *CarUpdateOne) AddAfterID(v float64) *CarUpdateOne {
	_u.mutation.AddAfterID(v)
	return _u
}

// ClearAfterID clears the value of the "after_id" field.
func (_u *CarUpdateOne) ClearAfterID() *CarUpdateOne {
	_u.mutation.ClearAfterID()
	return _u
}

// SetModel sets the "model" field.
func (_u *CarUpdateOne) SetModel(v string) *CarUpdateOne {
	_u.mutation.SetModel(v)
	return _u
}

// SetNillableModel sets the "model" field if the given value is not nil.
func (_u *CarUpdateOne) SetNillableModel(v *string) *CarUpdateOne {
	if v != nil {
		_u.SetModel(*v)
	}
	return _u
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (_u *CarUpdateOne) SetOwnerID(id string) *CarUpdateOne {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (_u *CarUpdateOne) SetNillableOwnerID(id *string) *CarUpdateOne {
	if id != nil {
		_u = _u.SetOwnerID(*id)
	}
	return _u
}

// SetOwner sets the "owner" edge to the Pet entity.
func (_u *CarUpdateOne) SetOwner(v *Pet) *CarUpdateOne {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the CarMutation object of the builder.
func (_u *CarUpdateOne) Mutation() *CarMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the Pet entity.
func (_u *CarUpdateOne) ClearOwner() *CarUpdateOne {
	_u.mutation.ClearOwner()
	return _u
}

// Where appends a list predicates to the CarUpdate builder.
func (_u *CarUpdateOne) Where(ps ...predicate.Car) *CarUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Car entity.
func (_u *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CarUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *CarUpdateOne) check() error {
	if v, ok := _u.mutation.BeforeID(); ok {
		if err := car.BeforeIDValidator(v); err != nil {
			return &ValidationError{Name: "before_id", err: fmt.Errorf(`ent: validator failed for field "Car.before_id": %w`, err)}
		}
	}
	if v, ok := _u.mutation.AfterID(); ok {
		if err := car.AfterIDValidator(v); err != nil {
			return &ValidationError{Name: "after_id", err: fmt.Errorf(`ent: validator failed for field "Car.after_id": %w`, err)}
		}
	}
	return nil
}

func (_u *CarUpdateOne) sqlSave(ctx context.Context) (_node *Car, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
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
	if value, ok := _u.mutation.BeforeID(); ok {
		_spec.SetField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if value, ok := _u.mutation.AddedBeforeID(); ok {
		_spec.AddField(car.FieldBeforeID, field.TypeFloat64, value)
	}
	if _u.mutation.BeforeIDCleared() {
		_spec.ClearField(car.FieldBeforeID, field.TypeFloat64)
	}
	if value, ok := _u.mutation.AfterID(); ok {
		_spec.SetField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if value, ok := _u.mutation.AddedAfterID(); ok {
		_spec.AddField(car.FieldAfterID, field.TypeFloat64, value)
	}
	if _u.mutation.AfterIDCleared() {
		_spec.ClearField(car.FieldAfterID, field.TypeFloat64)
	}
	if value, ok := _u.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Car{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
