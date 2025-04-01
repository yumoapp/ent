// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/ent/schema/task"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority task.Priority `json:"priority,omitempty"`
	// Priorities holds the value of the "priorities" field.
	Priorities map[string]task.Priority `json:"priorities,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Name holds the value of the "name" field.
	//
	// Deprecated: Field "name" was marked as deprecated in the schema.
	Name string `json:"name,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// OrderOption holds the value of the "order_option" field.
	OrderOption int `json:"order_option,omitempty"`
	// Op holds the value of the "op" field.
	Op string `json:"op,omitempty"`
}

// FromResponse scans the gremlin response data into Task.
func (_m *Task) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m struct {
		ID          string                   `json:"id,omitempty"`
		Priority    task.Priority            `json:"priority,omitempty"`
		Priorities  map[string]task.Priority `json:"priorities,omitempty"`
		CreatedAt   int64                    `json:"created_at,omitempty"`
		Name        string                   `json:"name,omitempty"`
		Owner       string                   `json:"owner,omitempty"`
		Order       int                      `json:"order,omitempty"`
		OrderOption int                      `json:"order_option,omitempty"`
		Op          string                   `json:"op,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	_m.ID = scan_m.ID
	_m.Priority = scan_m.Priority
	_m.Priorities = scan_m.Priorities
	v2 := time.Unix(0, scan_m.CreatedAt)
	_m.CreatedAt = &v2
	_m.Name = scan_m.Name
	_m.Owner = scan_m.Owner
	_m.Order = scan_m.Order
	_m.OrderOption = scan_m.OrderOption
	_m.Op = scan_m.Op
	return nil
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Task) Update() *TaskUpdateOne {
	return NewTaskClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Task) Unwrap() *Task {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", _m.Priority))
	builder.WriteString(", ")
	builder.WriteString("priorities=")
	builder.WriteString(fmt.Sprintf("%v", _m.Priorities))
	builder.WriteString(", ")
	if v := _m.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteString(", ")
	builder.WriteString("owner=")
	builder.WriteString(_m.Owner)
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", _m.Order))
	builder.WriteString(", ")
	builder.WriteString("order_option=")
	builder.WriteString(fmt.Sprintf("%v", _m.OrderOption))
	builder.WriteString(", ")
	builder.WriteString("op=")
	builder.WriteString(_m.Op)
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

// FromResponse scans the gremlin response data into Tasks.
func (_m *Tasks) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_m []struct {
		ID          string                   `json:"id,omitempty"`
		Priority    task.Priority            `json:"priority,omitempty"`
		Priorities  map[string]task.Priority `json:"priorities,omitempty"`
		CreatedAt   int64                    `json:"created_at,omitempty"`
		Name        string                   `json:"name,omitempty"`
		Owner       string                   `json:"owner,omitempty"`
		Order       int                      `json:"order,omitempty"`
		OrderOption int                      `json:"order_option,omitempty"`
		Op          string                   `json:"op,omitempty"`
	}
	if err := vmap.Decode(&scan_m); err != nil {
		return err
	}
	for _, v := range scan_m {
		node := &Task{ID: v.ID}
		node.Priority = v.Priority
		node.Priorities = v.Priorities
		v2 := time.Unix(0, v.CreatedAt)
		node.CreatedAt = &v2
		node.Name = v.Name
		node.Owner = v.Owner
		node.Order = v.Order
		node.OrderOption = v.OrderOption
		node.Op = v.Op
		*_m = append(*_m, node)
	}
	return nil
}
