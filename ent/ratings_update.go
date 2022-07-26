// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-db/ent/predicate"
	"imdb-db/ent/ratings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RatingsUpdate is the builder for updating Ratings entities.
type RatingsUpdate struct {
	config
	hooks    []Hook
	mutation *RatingsMutation
}

// Where appends a list predicates to the RatingsUpdate builder.
func (ru *RatingsUpdate) Where(ps ...predicate.Ratings) *RatingsUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetTconst sets the "tconst" field.
func (ru *RatingsUpdate) SetTconst(s string) *RatingsUpdate {
	ru.mutation.SetTconst(s)
	return ru
}

// SetAverageRating sets the "averageRating" field.
func (ru *RatingsUpdate) SetAverageRating(f float64) *RatingsUpdate {
	ru.mutation.ResetAverageRating()
	ru.mutation.SetAverageRating(f)
	return ru
}

// AddAverageRating adds f to the "averageRating" field.
func (ru *RatingsUpdate) AddAverageRating(f float64) *RatingsUpdate {
	ru.mutation.AddAverageRating(f)
	return ru
}

// SetNumVotes sets the "numVotes" field.
func (ru *RatingsUpdate) SetNumVotes(i int) *RatingsUpdate {
	ru.mutation.ResetNumVotes()
	ru.mutation.SetNumVotes(i)
	return ru
}

// AddNumVotes adds i to the "numVotes" field.
func (ru *RatingsUpdate) AddNumVotes(i int) *RatingsUpdate {
	ru.mutation.AddNumVotes(i)
	return ru
}

// Mutation returns the RatingsMutation object of the builder.
func (ru *RatingsUpdate) Mutation() *RatingsMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RatingsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RatingsUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RatingsUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RatingsUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RatingsUpdate) check() error {
	if v, ok := ru.mutation.AverageRating(); ok {
		if err := ratings.AverageRatingValidator(v); err != nil {
			return &ValidationError{Name: "averageRating", err: fmt.Errorf(`ent: validator failed for field "Ratings.averageRating": %w`, err)}
		}
	}
	return nil
}

func (ru *RatingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ratings.Table,
			Columns: ratings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ratings.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ratings.FieldTconst,
		})
	}
	if value, ok := ru.mutation.AverageRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ratings.FieldAverageRating,
		})
	}
	if value, ok := ru.mutation.AddedAverageRating(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ratings.FieldAverageRating,
		})
	}
	if value, ok := ru.mutation.NumVotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ratings.FieldNumVotes,
		})
	}
	if value, ok := ru.mutation.AddedNumVotes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ratings.FieldNumVotes,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ratings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RatingsUpdateOne is the builder for updating a single Ratings entity.
type RatingsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RatingsMutation
}

// SetTconst sets the "tconst" field.
func (ruo *RatingsUpdateOne) SetTconst(s string) *RatingsUpdateOne {
	ruo.mutation.SetTconst(s)
	return ruo
}

// SetAverageRating sets the "averageRating" field.
func (ruo *RatingsUpdateOne) SetAverageRating(f float64) *RatingsUpdateOne {
	ruo.mutation.ResetAverageRating()
	ruo.mutation.SetAverageRating(f)
	return ruo
}

// AddAverageRating adds f to the "averageRating" field.
func (ruo *RatingsUpdateOne) AddAverageRating(f float64) *RatingsUpdateOne {
	ruo.mutation.AddAverageRating(f)
	return ruo
}

// SetNumVotes sets the "numVotes" field.
func (ruo *RatingsUpdateOne) SetNumVotes(i int) *RatingsUpdateOne {
	ruo.mutation.ResetNumVotes()
	ruo.mutation.SetNumVotes(i)
	return ruo
}

// AddNumVotes adds i to the "numVotes" field.
func (ruo *RatingsUpdateOne) AddNumVotes(i int) *RatingsUpdateOne {
	ruo.mutation.AddNumVotes(i)
	return ruo
}

// Mutation returns the RatingsMutation object of the builder.
func (ruo *RatingsUpdateOne) Mutation() *RatingsMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RatingsUpdateOne) Select(field string, fields ...string) *RatingsUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Ratings entity.
func (ruo *RatingsUpdateOne) Save(ctx context.Context) (*Ratings, error) {
	var (
		err  error
		node *Ratings
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Ratings)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RatingsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RatingsUpdateOne) SaveX(ctx context.Context) *Ratings {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RatingsUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RatingsUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RatingsUpdateOne) check() error {
	if v, ok := ruo.mutation.AverageRating(); ok {
		if err := ratings.AverageRatingValidator(v); err != nil {
			return &ValidationError{Name: "averageRating", err: fmt.Errorf(`ent: validator failed for field "Ratings.averageRating": %w`, err)}
		}
	}
	return nil
}

func (ruo *RatingsUpdateOne) sqlSave(ctx context.Context) (_node *Ratings, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ratings.Table,
			Columns: ratings.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ratings.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ratings.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ratings.FieldID)
		for _, f := range fields {
			if !ratings.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ratings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ratings.FieldTconst,
		})
	}
	if value, ok := ruo.mutation.AverageRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ratings.FieldAverageRating,
		})
	}
	if value, ok := ruo.mutation.AddedAverageRating(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: ratings.FieldAverageRating,
		})
	}
	if value, ok := ruo.mutation.NumVotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ratings.FieldNumVotes,
		})
	}
	if value, ok := ruo.mutation.AddedNumVotes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ratings.FieldNumVotes,
		})
	}
	_node = &Ratings{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ratings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
