// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/triggers/ent/userauditlog"
	"entgo.io/ent/schema/field"
)

// UserAuditLogCreate is the builder for creating a UserAuditLog entity.
type UserAuditLogCreate struct {
	config
	mutation *UserAuditLogMutation
	hooks    []Hook
}

// SetOperationType sets the "operation_type" field.
func (_c *UserAuditLogCreate) SetOperationType(v string) *UserAuditLogCreate {
	_c.mutation.SetOperationType(v)
	return _c
}

// SetOperationTime sets the "operation_time" field.
func (_c *UserAuditLogCreate) SetOperationTime(v string) *UserAuditLogCreate {
	_c.mutation.SetOperationTime(v)
	return _c
}

// SetOldValue sets the "old_value" field.
func (_c *UserAuditLogCreate) SetOldValue(v string) *UserAuditLogCreate {
	_c.mutation.SetOldValue(v)
	return _c
}

// SetNillableOldValue sets the "old_value" field if the given value is not nil.
func (_c *UserAuditLogCreate) SetNillableOldValue(v *string) *UserAuditLogCreate {
	if v != nil {
		_c.SetOldValue(*v)
	}
	return _c
}

// SetNewValue sets the "new_value" field.
func (_c *UserAuditLogCreate) SetNewValue(v string) *UserAuditLogCreate {
	_c.mutation.SetNewValue(v)
	return _c
}

// SetNillableNewValue sets the "new_value" field if the given value is not nil.
func (_c *UserAuditLogCreate) SetNillableNewValue(v *string) *UserAuditLogCreate {
	if v != nil {
		_c.SetNewValue(*v)
	}
	return _c
}

// Mutation returns the UserAuditLogMutation object of the builder.
func (_c *UserAuditLogCreate) Mutation() *UserAuditLogMutation {
	return _c.mutation
}

// Save creates the UserAuditLog in the database.
func (_c *UserAuditLogCreate) Save(ctx context.Context) (*UserAuditLog, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *UserAuditLogCreate) SaveX(ctx context.Context) *UserAuditLog {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserAuditLogCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserAuditLogCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *UserAuditLogCreate) check() error {
	if _, ok := _c.mutation.OperationType(); !ok {
		return &ValidationError{Name: "operation_type", err: errors.New(`ent: missing required field "UserAuditLog.operation_type"`)}
	}
	if _, ok := _c.mutation.OperationTime(); !ok {
		return &ValidationError{Name: "operation_time", err: errors.New(`ent: missing required field "UserAuditLog.operation_time"`)}
	}
	return nil
}

func (_c *UserAuditLogCreate) sqlSave(ctx context.Context) (*UserAuditLog, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *UserAuditLogCreate) createSpec() (*UserAuditLog, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAuditLog{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(userauditlog.Table, sqlgraph.NewFieldSpec(userauditlog.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.OperationType(); ok {
		_spec.SetField(userauditlog.FieldOperationType, field.TypeString, value)
		_node.OperationType = value
	}
	if value, ok := _c.mutation.OperationTime(); ok {
		_spec.SetField(userauditlog.FieldOperationTime, field.TypeString, value)
		_node.OperationTime = value
	}
	if value, ok := _c.mutation.OldValue(); ok {
		_spec.SetField(userauditlog.FieldOldValue, field.TypeString, value)
		_node.OldValue = value
	}
	if value, ok := _c.mutation.NewValue(); ok {
		_spec.SetField(userauditlog.FieldNewValue, field.TypeString, value)
		_node.NewValue = value
	}
	return _node, _spec
}

// UserAuditLogCreateBulk is the builder for creating many UserAuditLog entities in bulk.
type UserAuditLogCreateBulk struct {
	config
	err      error
	builders []*UserAuditLogCreate
}

// Save creates the UserAuditLog entities in the database.
func (_c *UserAuditLogCreateBulk) Save(ctx context.Context) ([]*UserAuditLog, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*UserAuditLog, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAuditLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *UserAuditLogCreateBulk) SaveX(ctx context.Context) []*UserAuditLog {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserAuditLogCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserAuditLogCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
