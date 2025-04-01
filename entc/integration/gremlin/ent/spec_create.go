// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/spec"
)

// SpecCreate is the builder for creating a Spec entity.
type SpecCreate struct {
	config
	mutation *SpecMutation
	hooks    []Hook
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (_c *SpecCreate) AddCardIDs(ids ...string) *SpecCreate {
	_c.mutation.AddCardIDs(ids...)
	return _c
}

// AddCard adds the "card" edges to the Card entity.
func (_c *SpecCreate) AddCard(v ...*Card) *SpecCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (_c *SpecCreate) Mutation() *SpecMutation {
	return _c.mutation
}

// Save creates the Spec in the database.
func (_c *SpecCreate) Save(ctx context.Context) (*Spec, error) {
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *SpecCreate) SaveX(ctx context.Context) *Spec {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *SpecCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *SpecCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *SpecCreate) check() error {
	return nil
}

func (_c *SpecCreate) gremlinSave(ctx context.Context) (*Spec, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := _c.gremlin().Query()
	if err := _c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Spec{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *SpecCreate) gremlin() *dsl.Traversal {
	v := g.AddV(spec.Label)
	for _, id := range _c.mutation.CardIDs() {
		v.AddE(spec.CardLabel).To(g.V(id)).OutV()
	}
	return v.ValueMap(true)
}

// SpecCreateBulk is the builder for creating many Spec entities in bulk.
type SpecCreateBulk struct {
	config
	err      error
	builders []*SpecCreate
}
