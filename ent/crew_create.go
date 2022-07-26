// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-db/ent/crew"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CrewCreate is the builder for creating a Crew entity.
type CrewCreate struct {
	config
	mutation *CrewMutation
	hooks    []Hook
}

// SetTconst sets the "tconst" field.
func (cc *CrewCreate) SetTconst(s string) *CrewCreate {
	cc.mutation.SetTconst(s)
	return cc
}

// SetDirectors sets the "directors" field.
func (cc *CrewCreate) SetDirectors(s []string) *CrewCreate {
	cc.mutation.SetDirectors(s)
	return cc
}

// SetWriters sets the "writers" field.
func (cc *CrewCreate) SetWriters(s []string) *CrewCreate {
	cc.mutation.SetWriters(s)
	return cc
}

// Mutation returns the CrewMutation object of the builder.
func (cc *CrewCreate) Mutation() *CrewMutation {
	return cc.mutation
}

// Save creates the Crew in the database.
func (cc *CrewCreate) Save(ctx context.Context) (*Crew, error) {
	var (
		err  error
		node *Crew
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CrewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (cc *CrewCreate) SaveX(ctx context.Context) *Crew {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CrewCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CrewCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CrewCreate) check() error {
	if _, ok := cc.mutation.Tconst(); !ok {
		return &ValidationError{Name: "tconst", err: errors.New(`ent: missing required field "Crew.tconst"`)}
	}
	if _, ok := cc.mutation.Directors(); !ok {
		return &ValidationError{Name: "directors", err: errors.New(`ent: missing required field "Crew.directors"`)}
	}
	if _, ok := cc.mutation.Writers(); !ok {
		return &ValidationError{Name: "writers", err: errors.New(`ent: missing required field "Crew.writers"`)}
	}
	return nil
}

func (cc *CrewCreate) sqlSave(ctx context.Context) (*Crew, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CrewCreate) createSpec() (*Crew, *sqlgraph.CreateSpec) {
	var (
		_node = &Crew{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: crew.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: crew.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Tconst(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: crew.FieldTconst,
		})
		_node.Tconst = value
	}
	if value, ok := cc.mutation.Directors(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: crew.FieldDirectors,
		})
		_node.Directors = value
	}
	if value, ok := cc.mutation.Writers(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: crew.FieldWriters,
		})
		_node.Writers = value
	}
	return _node, _spec
}

// CrewCreateBulk is the builder for creating many Crew entities in bulk.
type CrewCreateBulk struct {
	config
	builders []*CrewCreate
}

// Save creates the Crew entities in the database.
func (ccb *CrewCreateBulk) Save(ctx context.Context) ([]*Crew, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Crew, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CrewMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CrewCreateBulk) SaveX(ctx context.Context) []*Crew {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CrewCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CrewCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
