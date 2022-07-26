// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdb-db/ent/predicate"
	"imdb-db/ent/title"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TitleUpdate is the builder for updating Title entities.
type TitleUpdate struct {
	config
	hooks    []Hook
	mutation *TitleMutation
}

// Where appends a list predicates to the TitleUpdate builder.
func (tu *TitleUpdate) Where(ps ...predicate.Title) *TitleUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTconst sets the "tconst" field.
func (tu *TitleUpdate) SetTconst(s string) *TitleUpdate {
	tu.mutation.SetTconst(s)
	return tu
}

// SetTitleType sets the "titleType" field.
func (tu *TitleUpdate) SetTitleType(s string) *TitleUpdate {
	tu.mutation.SetTitleType(s)
	return tu
}

// SetPrimaryTitle sets the "primaryTitle" field.
func (tu *TitleUpdate) SetPrimaryTitle(s string) *TitleUpdate {
	tu.mutation.SetPrimaryTitle(s)
	return tu
}

// SetOriginalTitle sets the "originalTitle" field.
func (tu *TitleUpdate) SetOriginalTitle(s string) *TitleUpdate {
	tu.mutation.SetOriginalTitle(s)
	return tu
}

// SetIsAdult sets the "isAdult" field.
func (tu *TitleUpdate) SetIsAdult(b bool) *TitleUpdate {
	tu.mutation.SetIsAdult(b)
	return tu
}

// SetNillableIsAdult sets the "isAdult" field if the given value is not nil.
func (tu *TitleUpdate) SetNillableIsAdult(b *bool) *TitleUpdate {
	if b != nil {
		tu.SetIsAdult(*b)
	}
	return tu
}

// SetStartYear sets the "startYear" field.
func (tu *TitleUpdate) SetStartYear(i int) *TitleUpdate {
	tu.mutation.ResetStartYear()
	tu.mutation.SetStartYear(i)
	return tu
}

// AddStartYear adds i to the "startYear" field.
func (tu *TitleUpdate) AddStartYear(i int) *TitleUpdate {
	tu.mutation.AddStartYear(i)
	return tu
}

// SetEndYear sets the "endYear" field.
func (tu *TitleUpdate) SetEndYear(i int) *TitleUpdate {
	tu.mutation.ResetEndYear()
	tu.mutation.SetEndYear(i)
	return tu
}

// AddEndYear adds i to the "endYear" field.
func (tu *TitleUpdate) AddEndYear(i int) *TitleUpdate {
	tu.mutation.AddEndYear(i)
	return tu
}

// SetRuntimeMinutes sets the "runtimeMinutes" field.
func (tu *TitleUpdate) SetRuntimeMinutes(i int) *TitleUpdate {
	tu.mutation.ResetRuntimeMinutes()
	tu.mutation.SetRuntimeMinutes(i)
	return tu
}

// AddRuntimeMinutes adds i to the "runtimeMinutes" field.
func (tu *TitleUpdate) AddRuntimeMinutes(i int) *TitleUpdate {
	tu.mutation.AddRuntimeMinutes(i)
	return tu
}

// SetGenre sets the "genre" field.
func (tu *TitleUpdate) SetGenre(s []string) *TitleUpdate {
	tu.mutation.SetGenre(s)
	return tu
}

// Mutation returns the TitleMutation object of the builder.
func (tu *TitleUpdate) Mutation() *TitleMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TitleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TitleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TitleUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TitleUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TitleUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TitleUpdate) check() error {
	if v, ok := tu.mutation.StartYear(); ok {
		if err := title.StartYearValidator(v); err != nil {
			return &ValidationError{Name: "startYear", err: fmt.Errorf(`ent: validator failed for field "Title.startYear": %w`, err)}
		}
	}
	if v, ok := tu.mutation.EndYear(); ok {
		if err := title.EndYearValidator(v); err != nil {
			return &ValidationError{Name: "endYear", err: fmt.Errorf(`ent: validator failed for field "Title.endYear": %w`, err)}
		}
	}
	if v, ok := tu.mutation.RuntimeMinutes(); ok {
		if err := title.RuntimeMinutesValidator(v); err != nil {
			return &ValidationError{Name: "runtimeMinutes", err: fmt.Errorf(`ent: validator failed for field "Title.runtimeMinutes": %w`, err)}
		}
	}
	return nil
}

func (tu *TitleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   title.Table,
			Columns: title.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: title.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldTconst,
		})
	}
	if value, ok := tu.mutation.TitleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldTitleType,
		})
	}
	if value, ok := tu.mutation.PrimaryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldPrimaryTitle,
		})
	}
	if value, ok := tu.mutation.OriginalTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldOriginalTitle,
		})
	}
	if value, ok := tu.mutation.IsAdult(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: title.FieldIsAdult,
		})
	}
	if value, ok := tu.mutation.StartYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldStartYear,
		})
	}
	if value, ok := tu.mutation.AddedStartYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldStartYear,
		})
	}
	if value, ok := tu.mutation.EndYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldEndYear,
		})
	}
	if value, ok := tu.mutation.AddedEndYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldEndYear,
		})
	}
	if value, ok := tu.mutation.RuntimeMinutes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldRuntimeMinutes,
		})
	}
	if value, ok := tu.mutation.AddedRuntimeMinutes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldRuntimeMinutes,
		})
	}
	if value, ok := tu.mutation.Genre(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: title.FieldGenre,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{title.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TitleUpdateOne is the builder for updating a single Title entity.
type TitleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TitleMutation
}

// SetTconst sets the "tconst" field.
func (tuo *TitleUpdateOne) SetTconst(s string) *TitleUpdateOne {
	tuo.mutation.SetTconst(s)
	return tuo
}

// SetTitleType sets the "titleType" field.
func (tuo *TitleUpdateOne) SetTitleType(s string) *TitleUpdateOne {
	tuo.mutation.SetTitleType(s)
	return tuo
}

// SetPrimaryTitle sets the "primaryTitle" field.
func (tuo *TitleUpdateOne) SetPrimaryTitle(s string) *TitleUpdateOne {
	tuo.mutation.SetPrimaryTitle(s)
	return tuo
}

// SetOriginalTitle sets the "originalTitle" field.
func (tuo *TitleUpdateOne) SetOriginalTitle(s string) *TitleUpdateOne {
	tuo.mutation.SetOriginalTitle(s)
	return tuo
}

// SetIsAdult sets the "isAdult" field.
func (tuo *TitleUpdateOne) SetIsAdult(b bool) *TitleUpdateOne {
	tuo.mutation.SetIsAdult(b)
	return tuo
}

// SetNillableIsAdult sets the "isAdult" field if the given value is not nil.
func (tuo *TitleUpdateOne) SetNillableIsAdult(b *bool) *TitleUpdateOne {
	if b != nil {
		tuo.SetIsAdult(*b)
	}
	return tuo
}

// SetStartYear sets the "startYear" field.
func (tuo *TitleUpdateOne) SetStartYear(i int) *TitleUpdateOne {
	tuo.mutation.ResetStartYear()
	tuo.mutation.SetStartYear(i)
	return tuo
}

// AddStartYear adds i to the "startYear" field.
func (tuo *TitleUpdateOne) AddStartYear(i int) *TitleUpdateOne {
	tuo.mutation.AddStartYear(i)
	return tuo
}

// SetEndYear sets the "endYear" field.
func (tuo *TitleUpdateOne) SetEndYear(i int) *TitleUpdateOne {
	tuo.mutation.ResetEndYear()
	tuo.mutation.SetEndYear(i)
	return tuo
}

// AddEndYear adds i to the "endYear" field.
func (tuo *TitleUpdateOne) AddEndYear(i int) *TitleUpdateOne {
	tuo.mutation.AddEndYear(i)
	return tuo
}

// SetRuntimeMinutes sets the "runtimeMinutes" field.
func (tuo *TitleUpdateOne) SetRuntimeMinutes(i int) *TitleUpdateOne {
	tuo.mutation.ResetRuntimeMinutes()
	tuo.mutation.SetRuntimeMinutes(i)
	return tuo
}

// AddRuntimeMinutes adds i to the "runtimeMinutes" field.
func (tuo *TitleUpdateOne) AddRuntimeMinutes(i int) *TitleUpdateOne {
	tuo.mutation.AddRuntimeMinutes(i)
	return tuo
}

// SetGenre sets the "genre" field.
func (tuo *TitleUpdateOne) SetGenre(s []string) *TitleUpdateOne {
	tuo.mutation.SetGenre(s)
	return tuo
}

// Mutation returns the TitleMutation object of the builder.
func (tuo *TitleUpdateOne) Mutation() *TitleMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TitleUpdateOne) Select(field string, fields ...string) *TitleUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Title entity.
func (tuo *TitleUpdateOne) Save(ctx context.Context) (*Title, error) {
	var (
		err  error
		node *Title
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TitleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Title)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TitleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TitleUpdateOne) SaveX(ctx context.Context) *Title {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TitleUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TitleUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TitleUpdateOne) check() error {
	if v, ok := tuo.mutation.StartYear(); ok {
		if err := title.StartYearValidator(v); err != nil {
			return &ValidationError{Name: "startYear", err: fmt.Errorf(`ent: validator failed for field "Title.startYear": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.EndYear(); ok {
		if err := title.EndYearValidator(v); err != nil {
			return &ValidationError{Name: "endYear", err: fmt.Errorf(`ent: validator failed for field "Title.endYear": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.RuntimeMinutes(); ok {
		if err := title.RuntimeMinutesValidator(v); err != nil {
			return &ValidationError{Name: "runtimeMinutes", err: fmt.Errorf(`ent: validator failed for field "Title.runtimeMinutes": %w`, err)}
		}
	}
	return nil
}

func (tuo *TitleUpdateOne) sqlSave(ctx context.Context) (_node *Title, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   title.Table,
			Columns: title.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: title.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Title.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, title.FieldID)
		for _, f := range fields {
			if !title.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != title.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Tconst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldTconst,
		})
	}
	if value, ok := tuo.mutation.TitleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldTitleType,
		})
	}
	if value, ok := tuo.mutation.PrimaryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldPrimaryTitle,
		})
	}
	if value, ok := tuo.mutation.OriginalTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldOriginalTitle,
		})
	}
	if value, ok := tuo.mutation.IsAdult(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: title.FieldIsAdult,
		})
	}
	if value, ok := tuo.mutation.StartYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldStartYear,
		})
	}
	if value, ok := tuo.mutation.AddedStartYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldStartYear,
		})
	}
	if value, ok := tuo.mutation.EndYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldEndYear,
		})
	}
	if value, ok := tuo.mutation.AddedEndYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldEndYear,
		})
	}
	if value, ok := tuo.mutation.RuntimeMinutes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldRuntimeMinutes,
		})
	}
	if value, ok := tuo.mutation.AddedRuntimeMinutes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: title.FieldRuntimeMinutes,
		})
	}
	if value, ok := tuo.mutation.Genre(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: title.FieldGenre,
		})
	}
	_node = &Title{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{title.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}