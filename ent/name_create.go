// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-db/ent/name"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NameCreate is the builder for creating a Name entity.
type NameCreate struct {
	config
	mutation *NameMutation
	hooks    []Hook
}

// SetTconst sets the "tconst" field.
func (nc *NameCreate) SetTconst(s string) *NameCreate {
	nc.mutation.SetTconst(s)
	return nc
}

// SetPrimaryName sets the "primaryName" field.
func (nc *NameCreate) SetPrimaryName(s string) *NameCreate {
	nc.mutation.SetPrimaryName(s)
	return nc
}

// SetBirthYear sets the "birthYear" field.
func (nc *NameCreate) SetBirthYear(i int) *NameCreate {
	nc.mutation.SetBirthYear(i)
	return nc
}

// SetDeathYear sets the "deathYear" field.
func (nc *NameCreate) SetDeathYear(i int) *NameCreate {
	nc.mutation.SetDeathYear(i)
	return nc
}

// Mutation returns the NameMutation object of the builder.
func (nc *NameCreate) Mutation() *NameMutation {
	return nc.mutation
}

// Save creates the Name in the database.
func (nc *NameCreate) Save(ctx context.Context) (*Name, error) {
	var (
		err  error
		node *Name
	)
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Name)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NameMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NameCreate) SaveX(ctx context.Context) *Name {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NameCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NameCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NameCreate) check() error {
	if _, ok := nc.mutation.Tconst(); !ok {
		return &ValidationError{Name: "tconst", err: errors.New(`ent: missing required field "Name.tconst"`)}
	}
	if _, ok := nc.mutation.PrimaryName(); !ok {
		return &ValidationError{Name: "primaryName", err: errors.New(`ent: missing required field "Name.primaryName"`)}
	}
	if _, ok := nc.mutation.BirthYear(); !ok {
		return &ValidationError{Name: "birthYear", err: errors.New(`ent: missing required field "Name.birthYear"`)}
	}
	if v, ok := nc.mutation.BirthYear(); ok {
		if err := name.BirthYearValidator(v); err != nil {
			return &ValidationError{Name: "birthYear", err: fmt.Errorf(`ent: validator failed for field "Name.birthYear": %w`, err)}
		}
	}
	if _, ok := nc.mutation.DeathYear(); !ok {
		return &ValidationError{Name: "deathYear", err: errors.New(`ent: missing required field "Name.deathYear"`)}
	}
	if v, ok := nc.mutation.DeathYear(); ok {
		if err := name.DeathYearValidator(v); err != nil {
			return &ValidationError{Name: "deathYear", err: fmt.Errorf(`ent: validator failed for field "Name.deathYear": %w`, err)}
		}
	}
	return nil
}

func (nc *NameCreate) sqlSave(ctx context.Context) (*Name, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nc *NameCreate) createSpec() (*Name, *sqlgraph.CreateSpec) {
	var (
		_node = &Name{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: name.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: name.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.Tconst(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldTconst,
		})
		_node.Tconst = value
	}
	if value, ok := nc.mutation.PrimaryName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: name.FieldPrimaryName,
		})
		_node.PrimaryName = value
	}
	if value, ok := nc.mutation.BirthYear(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: name.FieldBirthYear,
		})
		_node.BirthYear = value
	}
	if value, ok := nc.mutation.DeathYear(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: name.FieldDeathYear,
		})
		_node.DeathYear = value
	}
	return _node, _spec
}

// NameCreateBulk is the builder for creating many Name entities in bulk.
type NameCreateBulk struct {
	config
	builders []*NameCreate
}

// Save creates the Name entities in the database.
func (ncb *NameCreateBulk) Save(ctx context.Context) ([]*Name, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Name, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NameMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NameCreateBulk) SaveX(ctx context.Context) []*Name {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NameCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NameCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}