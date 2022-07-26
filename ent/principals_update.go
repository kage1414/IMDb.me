// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-db/ent/predicate"
	"imdb-db/ent/principals"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrincipalsUpdate is the builder for updating Principals entities.
type PrincipalsUpdate struct {
	config
	hooks    []Hook
	mutation *PrincipalsMutation
}

// Where appends a list predicates to the PrincipalsUpdate builder.
func (pu *PrincipalsUpdate) Where(ps ...predicate.Principals) *PrincipalsUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetTconst sets the "tconst" field.
func (pu *PrincipalsUpdate) SetTconst(s string) *PrincipalsUpdate {
	pu.mutation.SetTconst(s)
	return pu
}

// SetOrdering sets the "ordering" field.
func (pu *PrincipalsUpdate) SetOrdering(i int) *PrincipalsUpdate {
	pu.mutation.ResetOrdering()
	pu.mutation.SetOrdering(i)
	return pu
}

// AddOrdering adds i to the "ordering" field.
func (pu *PrincipalsUpdate) AddOrdering(i int) *PrincipalsUpdate {
	pu.mutation.AddOrdering(i)
	return pu
}

// SetNconst sets the "nconst" field.
func (pu *PrincipalsUpdate) SetNconst(s string) *PrincipalsUpdate {
	pu.mutation.SetNconst(s)
	return pu
}

// SetCategory sets the "category" field.
func (pu *PrincipalsUpdate) SetCategory(s string) *PrincipalsUpdate {
	pu.mutation.SetCategory(s)
	return pu
}

// SetJob sets the "job" field.
func (pu *PrincipalsUpdate) SetJob(s string) *PrincipalsUpdate {
	pu.mutation.SetJob(s)
	return pu
}

// SetCharacters sets the "characters" field.
func (pu *PrincipalsUpdate) SetCharacters(s string) *PrincipalsUpdate {
	pu.mutation.SetCharacters(s)
	return pu
}

// Mutation returns the PrincipalsMutation object of the builder.
func (pu *PrincipalsUpdate) Mutation() *PrincipalsMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PrincipalsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrincipalsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrincipalsUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrincipalsUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrincipalsUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PrincipalsUpdate) check() error {
	if v, ok := pu.mutation.Ordering(); ok {
		if err := principals.OrderingValidator(v); err != nil {
			return &ValidationError{Name: "ordering", err: fmt.Errorf(`ent: validator failed for field "Principals.ordering": %w`, err)}
		}
	}
	return nil
}

func (pu *PrincipalsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   principals.Table,
			Columns: principals.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: principals.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldTconst,
		})
	}
	if value, ok := pu.mutation.Ordering(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: principals.FieldOrdering,
		})
	}
	if value, ok := pu.mutation.AddedOrdering(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: principals.FieldOrdering,
		})
	}
	if value, ok := pu.mutation.Nconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldNconst,
		})
	}
	if value, ok := pu.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldCategory,
		})
	}
	if value, ok := pu.mutation.Job(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldJob,
		})
	}
	if value, ok := pu.mutation.Characters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldCharacters,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{principals.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PrincipalsUpdateOne is the builder for updating a single Principals entity.
type PrincipalsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PrincipalsMutation
}

// SetTconst sets the "tconst" field.
func (puo *PrincipalsUpdateOne) SetTconst(s string) *PrincipalsUpdateOne {
	puo.mutation.SetTconst(s)
	return puo
}

// SetOrdering sets the "ordering" field.
func (puo *PrincipalsUpdateOne) SetOrdering(i int) *PrincipalsUpdateOne {
	puo.mutation.ResetOrdering()
	puo.mutation.SetOrdering(i)
	return puo
}

// AddOrdering adds i to the "ordering" field.
func (puo *PrincipalsUpdateOne) AddOrdering(i int) *PrincipalsUpdateOne {
	puo.mutation.AddOrdering(i)
	return puo
}

// SetNconst sets the "nconst" field.
func (puo *PrincipalsUpdateOne) SetNconst(s string) *PrincipalsUpdateOne {
	puo.mutation.SetNconst(s)
	return puo
}

// SetCategory sets the "category" field.
func (puo *PrincipalsUpdateOne) SetCategory(s string) *PrincipalsUpdateOne {
	puo.mutation.SetCategory(s)
	return puo
}

// SetJob sets the "job" field.
func (puo *PrincipalsUpdateOne) SetJob(s string) *PrincipalsUpdateOne {
	puo.mutation.SetJob(s)
	return puo
}

// SetCharacters sets the "characters" field.
func (puo *PrincipalsUpdateOne) SetCharacters(s string) *PrincipalsUpdateOne {
	puo.mutation.SetCharacters(s)
	return puo
}

// Mutation returns the PrincipalsMutation object of the builder.
func (puo *PrincipalsUpdateOne) Mutation() *PrincipalsMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PrincipalsUpdateOne) Select(field string, fields ...string) *PrincipalsUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Principals entity.
func (puo *PrincipalsUpdateOne) Save(ctx context.Context) (*Principals, error) {
	var (
		err  error
		node *Principals
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrincipalsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Principals)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PrincipalsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrincipalsUpdateOne) SaveX(ctx context.Context) *Principals {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PrincipalsUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrincipalsUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PrincipalsUpdateOne) check() error {
	if v, ok := puo.mutation.Ordering(); ok {
		if err := principals.OrderingValidator(v); err != nil {
			return &ValidationError{Name: "ordering", err: fmt.Errorf(`ent: validator failed for field "Principals.ordering": %w`, err)}
		}
	}
	return nil
}

func (puo *PrincipalsUpdateOne) sqlSave(ctx context.Context) (_node *Principals, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   principals.Table,
			Columns: principals.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: principals.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Principals.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, principals.FieldID)
		for _, f := range fields {
			if !principals.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != principals.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldTconst,
		})
	}
	if value, ok := puo.mutation.Ordering(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: principals.FieldOrdering,
		})
	}
	if value, ok := puo.mutation.AddedOrdering(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: principals.FieldOrdering,
		})
	}
	if value, ok := puo.mutation.Nconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldNconst,
		})
	}
	if value, ok := puo.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldCategory,
		})
	}
	if value, ok := puo.mutation.Job(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldJob,
		})
	}
	if value, ok := puo.mutation.Characters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: principals.FieldCharacters,
		})
	}
	_node = &Principals{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{principals.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
