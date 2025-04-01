// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/note"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/schema/field"
)

// NoteCreate is the builder for creating a Note entity.
type NoteCreate struct {
	config
	mutation *NoteMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetText sets the "text" field.
func (_c *NoteCreate) SetText(v string) *NoteCreate {
	_c.mutation.SetText(v)
	return _c
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_c *NoteCreate) SetNillableText(v *string) *NoteCreate {
	if v != nil {
		_c.SetText(*v)
	}
	return _c
}

// SetID sets the "id" field.
func (_c *NoteCreate) SetID(v schema.NoteID) *NoteCreate {
	_c.mutation.SetID(v)
	return _c
}

// SetNillableID sets the "id" field if the given value is not nil.
func (_c *NoteCreate) SetNillableID(v *schema.NoteID) *NoteCreate {
	if v != nil {
		_c.SetID(*v)
	}
	return _c
}

// SetParentID sets the "parent" edge to the Note entity by ID.
func (_c *NoteCreate) SetParentID(id schema.NoteID) *NoteCreate {
	_c.mutation.SetParentID(id)
	return _c
}

// SetNillableParentID sets the "parent" edge to the Note entity by ID if the given value is not nil.
func (_c *NoteCreate) SetNillableParentID(id *schema.NoteID) *NoteCreate {
	if id != nil {
		_c = _c.SetParentID(*id)
	}
	return _c
}

// SetParent sets the "parent" edge to the Note entity.
func (_c *NoteCreate) SetParent(v *Note) *NoteCreate {
	return _c.SetParentID(v.ID)
}

// AddChildIDs adds the "children" edge to the Note entity by IDs.
func (_c *NoteCreate) AddChildIDs(ids ...schema.NoteID) *NoteCreate {
	_c.mutation.AddChildIDs(ids...)
	return _c
}

// AddChildren adds the "children" edges to the Note entity.
func (_c *NoteCreate) AddChildren(v ...*Note) *NoteCreate {
	ids := make([]schema.NoteID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddChildIDs(ids...)
}

// Mutation returns the NoteMutation object of the builder.
func (_c *NoteCreate) Mutation() *NoteMutation {
	return _c.mutation
}

// Save creates the Note in the database.
func (_c *NoteCreate) Save(ctx context.Context) (*Note, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *NoteCreate) SaveX(ctx context.Context) *Note {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *NoteCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *NoteCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *NoteCreate) defaults() {
	if _, ok := _c.mutation.ID(); !ok {
		v := note.DefaultID()
		_c.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *NoteCreate) check() error {
	if v, ok := _c.mutation.ID(); ok {
		if err := note.IDValidator(string(v)); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Note.id": %w`, err)}
		}
	}
	return nil
}

func (_c *NoteCreate) sqlSave(ctx context.Context) (*Note, error) {
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
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(schema.NoteID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Note.ID type: %T", _spec.ID.Value)
		}
	}
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *NoteCreate) createSpec() (*Note, *sqlgraph.CreateSpec) {
	var (
		_node = &Note{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(note.Table, sqlgraph.NewFieldSpec(note.FieldID, field.TypeString))
	)
	_spec.OnConflict = _c.conflict
	if id, ok := _c.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := _c.mutation.Text(); ok {
		_spec.SetField(note.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if nodes := _c.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.ParentTable,
			Columns: []string{note.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.note_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.ChildrenTable,
			Columns: []string{note.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Note.Create().
//		SetText(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NoteUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (_c *NoteCreate) OnConflict(opts ...sql.ConflictOption) *NoteUpsertOne {
	_c.conflict = opts
	return &NoteUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *NoteCreate) OnConflictColumns(columns ...string) *NoteUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &NoteUpsertOne{
		create: _c,
	}
}

type (
	// NoteUpsertOne is the builder for "upsert"-ing
	//  one Note node.
	NoteUpsertOne struct {
		create *NoteCreate
	}

	// NoteUpsert is the "OnConflict" setter.
	NoteUpsert struct {
		*sql.UpdateSet
	}
)

// SetText sets the "text" field.
func (u *NoteUpsert) SetText(v string) *NoteUpsert {
	u.Set(note.FieldText, v)
	return u
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *NoteUpsert) UpdateText() *NoteUpsert {
	u.SetExcluded(note.FieldText)
	return u
}

// ClearText clears the value of the "text" field.
func (u *NoteUpsert) ClearText() *NoteUpsert {
	u.SetNull(note.FieldText)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(note.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NoteUpsertOne) UpdateNewValues() *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(note.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Note.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NoteUpsertOne) Ignore() *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NoteUpsertOne) DoNothing() *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NoteCreate.OnConflict
// documentation for more info.
func (u *NoteUpsertOne) Update(set func(*NoteUpsert)) *NoteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NoteUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *NoteUpsertOne) SetText(v string) *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *NoteUpsertOne) UpdateText() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateText()
	})
}

// ClearText clears the value of the "text" field.
func (u *NoteUpsertOne) ClearText() *NoteUpsertOne {
	return u.Update(func(s *NoteUpsert) {
		s.ClearText()
	})
}

// Exec executes the query.
func (u *NoteUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NoteCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NoteUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NoteUpsertOne) ID(ctx context.Context) (id schema.NoteID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NoteUpsertOne.ID is not supported by MySQL driver. Use NoteUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NoteUpsertOne) IDX(ctx context.Context) schema.NoteID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NoteCreateBulk is the builder for creating many Note entities in bulk.
type NoteCreateBulk struct {
	config
	err      error
	builders []*NoteCreate
	conflict []sql.ConflictOption
}

// Save creates the Note entities in the database.
func (_c *NoteCreateBulk) Save(ctx context.Context) ([]*Note, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Note, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NoteMutation)
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
					spec.OnConflict = _c.conflict
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
func (_c *NoteCreateBulk) SaveX(ctx context.Context) []*Note {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *NoteCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *NoteCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Note.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NoteUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (_c *NoteCreateBulk) OnConflict(opts ...sql.ConflictOption) *NoteUpsertBulk {
	_c.conflict = opts
	return &NoteUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *NoteCreateBulk) OnConflictColumns(columns ...string) *NoteUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &NoteUpsertBulk{
		create: _c,
	}
}

// NoteUpsertBulk is the builder for "upsert"-ing
// a bulk of Note nodes.
type NoteUpsertBulk struct {
	create *NoteCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(note.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NoteUpsertBulk) UpdateNewValues() *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(note.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Note.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NoteUpsertBulk) Ignore() *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NoteUpsertBulk) DoNothing() *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NoteCreateBulk.OnConflict
// documentation for more info.
func (u *NoteUpsertBulk) Update(set func(*NoteUpsert)) *NoteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NoteUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *NoteUpsertBulk) SetText(v string) *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *NoteUpsertBulk) UpdateText() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.UpdateText()
	})
}

// ClearText clears the value of the "text" field.
func (u *NoteUpsertBulk) ClearText() *NoteUpsertBulk {
	return u.Update(func(s *NoteUpsert) {
		s.ClearText()
	})
}

// Exec executes the query.
func (u *NoteUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NoteCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NoteCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NoteUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
