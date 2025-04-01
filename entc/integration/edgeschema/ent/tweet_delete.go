// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/schema/field"
)

// TweetDelete is the builder for deleting a Tweet entity.
type TweetDelete struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetDelete builder.
func (_d *TweetDelete) Where(ps ...predicate.Tweet) *TweetDelete {
	_d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (_d *TweetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, _d.sqlExec, _d.mutation, _d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *TweetDelete) ExecX(ctx context.Context) int {
	n, err := _d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (_d *TweetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tweet.Table, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	if ps := _d.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, _d.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	_d.mutation.done = true
	return affected, err
}

// TweetDeleteOne is the builder for deleting a single Tweet entity.
type TweetDeleteOne struct {
	_d *TweetDelete
}

// Where appends a list predicates to the TweetDelete builder.
func (_d *TweetDeleteOne) Where(ps ...predicate.Tweet) *TweetDeleteOne {
	_d._d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query.
func (_d *TweetDeleteOne) Exec(ctx context.Context) error {
	n, err := _d._d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tweet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *TweetDeleteOne) ExecX(ctx context.Context) {
	if err := _d.Exec(ctx); err != nil {
		panic(err)
	}
}
