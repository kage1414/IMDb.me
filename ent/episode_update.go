// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-db/ent/episode"
	"imdb-db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EpisodeUpdate is the builder for updating Episode entities.
type EpisodeUpdate struct {
	config
	hooks    []Hook
	mutation *EpisodeMutation
}

// Where appends a list predicates to the EpisodeUpdate builder.
func (eu *EpisodeUpdate) Where(ps ...predicate.Episode) *EpisodeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetTconst sets the "tconst" field.
func (eu *EpisodeUpdate) SetTconst(s string) *EpisodeUpdate {
	eu.mutation.SetTconst(s)
	return eu
}

// SetParentTconst sets the "parentTconst" field.
func (eu *EpisodeUpdate) SetParentTconst(s string) *EpisodeUpdate {
	eu.mutation.SetParentTconst(s)
	return eu
}

// SetSeasonNumber sets the "seasonNumber" field.
func (eu *EpisodeUpdate) SetSeasonNumber(i int) *EpisodeUpdate {
	eu.mutation.ResetSeasonNumber()
	eu.mutation.SetSeasonNumber(i)
	return eu
}

// AddSeasonNumber adds i to the "seasonNumber" field.
func (eu *EpisodeUpdate) AddSeasonNumber(i int) *EpisodeUpdate {
	eu.mutation.AddSeasonNumber(i)
	return eu
}

// SetEpisodeNumber sets the "episodeNumber" field.
func (eu *EpisodeUpdate) SetEpisodeNumber(i int) *EpisodeUpdate {
	eu.mutation.ResetEpisodeNumber()
	eu.mutation.SetEpisodeNumber(i)
	return eu
}

// AddEpisodeNumber adds i to the "episodeNumber" field.
func (eu *EpisodeUpdate) AddEpisodeNumber(i int) *EpisodeUpdate {
	eu.mutation.AddEpisodeNumber(i)
	return eu
}

// Mutation returns the EpisodeMutation object of the builder.
func (eu *EpisodeUpdate) Mutation() *EpisodeMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EpisodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EpisodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EpisodeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EpisodeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EpisodeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EpisodeUpdate) check() error {
	if v, ok := eu.mutation.SeasonNumber(); ok {
		if err := episode.SeasonNumberValidator(v); err != nil {
			return &ValidationError{Name: "seasonNumber", err: fmt.Errorf(`ent: validator failed for field "Episode.seasonNumber": %w`, err)}
		}
	}
	if v, ok := eu.mutation.EpisodeNumber(); ok {
		if err := episode.EpisodeNumberValidator(v); err != nil {
			return &ValidationError{Name: "episodeNumber", err: fmt.Errorf(`ent: validator failed for field "Episode.episodeNumber": %w`, err)}
		}
	}
	return nil
}

func (eu *EpisodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   episode.Table,
			Columns: episode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: episode.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldTconst,
		})
	}
	if value, ok := eu.mutation.ParentTconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldParentTconst,
		})
	}
	if value, ok := eu.mutation.SeasonNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldSeasonNumber,
		})
	}
	if value, ok := eu.mutation.AddedSeasonNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldSeasonNumber,
		})
	}
	if value, ok := eu.mutation.EpisodeNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldEpisodeNumber,
		})
	}
	if value, ok := eu.mutation.AddedEpisodeNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldEpisodeNumber,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EpisodeUpdateOne is the builder for updating a single Episode entity.
type EpisodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EpisodeMutation
}

// SetTconst sets the "tconst" field.
func (euo *EpisodeUpdateOne) SetTconst(s string) *EpisodeUpdateOne {
	euo.mutation.SetTconst(s)
	return euo
}

// SetParentTconst sets the "parentTconst" field.
func (euo *EpisodeUpdateOne) SetParentTconst(s string) *EpisodeUpdateOne {
	euo.mutation.SetParentTconst(s)
	return euo
}

// SetSeasonNumber sets the "seasonNumber" field.
func (euo *EpisodeUpdateOne) SetSeasonNumber(i int) *EpisodeUpdateOne {
	euo.mutation.ResetSeasonNumber()
	euo.mutation.SetSeasonNumber(i)
	return euo
}

// AddSeasonNumber adds i to the "seasonNumber" field.
func (euo *EpisodeUpdateOne) AddSeasonNumber(i int) *EpisodeUpdateOne {
	euo.mutation.AddSeasonNumber(i)
	return euo
}

// SetEpisodeNumber sets the "episodeNumber" field.
func (euo *EpisodeUpdateOne) SetEpisodeNumber(i int) *EpisodeUpdateOne {
	euo.mutation.ResetEpisodeNumber()
	euo.mutation.SetEpisodeNumber(i)
	return euo
}

// AddEpisodeNumber adds i to the "episodeNumber" field.
func (euo *EpisodeUpdateOne) AddEpisodeNumber(i int) *EpisodeUpdateOne {
	euo.mutation.AddEpisodeNumber(i)
	return euo
}

// Mutation returns the EpisodeMutation object of the builder.
func (euo *EpisodeUpdateOne) Mutation() *EpisodeMutation {
	return euo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EpisodeUpdateOne) Select(field string, fields ...string) *EpisodeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Episode entity.
func (euo *EpisodeUpdateOne) Save(ctx context.Context) (*Episode, error) {
	var (
		err  error
		node *Episode
	)
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EpisodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Episode)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EpisodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EpisodeUpdateOne) SaveX(ctx context.Context) *Episode {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EpisodeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EpisodeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EpisodeUpdateOne) check() error {
	if v, ok := euo.mutation.SeasonNumber(); ok {
		if err := episode.SeasonNumberValidator(v); err != nil {
			return &ValidationError{Name: "seasonNumber", err: fmt.Errorf(`ent: validator failed for field "Episode.seasonNumber": %w`, err)}
		}
	}
	if v, ok := euo.mutation.EpisodeNumber(); ok {
		if err := episode.EpisodeNumberValidator(v); err != nil {
			return &ValidationError{Name: "episodeNumber", err: fmt.Errorf(`ent: validator failed for field "Episode.episodeNumber": %w`, err)}
		}
	}
	return nil
}

func (euo *EpisodeUpdateOne) sqlSave(ctx context.Context) (_node *Episode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   episode.Table,
			Columns: episode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: episode.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Episode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, episode.FieldID)
		for _, f := range fields {
			if !episode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != episode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldTconst,
		})
	}
	if value, ok := euo.mutation.ParentTconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldParentTconst,
		})
	}
	if value, ok := euo.mutation.SeasonNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldSeasonNumber,
		})
	}
	if value, ok := euo.mutation.AddedSeasonNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldSeasonNumber,
		})
	}
	if value, ok := euo.mutation.EpisodeNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldEpisodeNumber,
		})
	}
	if value, ok := euo.mutation.AddedEpisodeNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: episode.FieldEpisodeNumber,
		})
	}
	_node = &Episode{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
