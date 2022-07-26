// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-db/ent/crew"
	"imdb-db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CrewUpdate is the builder for updating Crew entities.
type CrewUpdate struct {
	config
	hooks    []Hook
	mutation *CrewMutation
}

// Where appends a list predicates to the CrewUpdate builder.
func (cu *CrewUpdate) Where(ps ...predicate.Crew) *CrewUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTconst sets the "tconst" field.
func (cu *CrewUpdate) SetTconst(s string) *CrewUpdate {
	cu.mutation.SetTconst(s)
	return cu
}

// SetDirectors sets the "directors" field.
func (cu *CrewUpdate) SetDirectors(s []string) *CrewUpdate {
	cu.mutation.SetDirectors(s)
	return cu
}

// SetWriters sets the "writers" field.
func (cu *CrewUpdate) SetWriters(s []string) *CrewUpdate {
	cu.mutation.SetWriters(s)
	return cu
}

// Mutation returns the CrewMutation object of the builder.
func (cu *CrewUpdate) Mutation() *CrewMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CrewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CrewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CrewUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CrewUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CrewUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CrewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   crew.Table,
			Columns: crew.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: crew.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: crew.FieldTconst,
		})
	}
	if value, ok := cu.mutation.Directors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: crew.FieldDirectors,
		})
	}
	if value, ok := cu.mutation.Writers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: crew.FieldWriters,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{crew.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CrewUpdateOne is the builder for updating a single Crew entity.
type CrewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CrewMutation
}

// SetTconst sets the "tconst" field.
func (cuo *CrewUpdateOne) SetTconst(s string) *CrewUpdateOne {
	cuo.mutation.SetTconst(s)
	return cuo
}

// SetDirectors sets the "directors" field.
func (cuo *CrewUpdateOne) SetDirectors(s []string) *CrewUpdateOne {
	cuo.mutation.SetDirectors(s)
	return cuo
}

// SetWriters sets the "writers" field.
func (cuo *CrewUpdateOne) SetWriters(s []string) *CrewUpdateOne {
	cuo.mutation.SetWriters(s)
	return cuo
}

// Mutation returns the CrewMutation object of the builder.
func (cuo *CrewUpdateOne) Mutation() *CrewMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CrewUpdateOne) Select(field string, fields ...string) *CrewUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Crew entity.
func (cuo *CrewUpdateOne) Save(ctx context.Context) (*Crew, error) {
	var (
		err  error
		node *Crew
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CrewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Crew)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CrewMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CrewUpdateOne) SaveX(ctx context.Context) *Crew {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CrewUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CrewUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CrewUpdateOne) sqlSave(ctx context.Context) (_node *Crew, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   crew.Table,
			Columns: crew.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: crew.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Crew.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, crew.FieldID)
		for _, f := range fields {
			if !crew.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != crew.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: crew.FieldTconst,
		})
	}
	if value, ok := cuo.mutation.Directors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: crew.FieldDirectors,
		})
	}
	if value, ok := cuo.mutation.Writers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: crew.FieldWriters,
		})
	}
	_node = &Crew{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{crew.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
