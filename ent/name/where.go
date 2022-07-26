// Code generated by ent, DO NOT EDIT.

package name

import (
	"imdb-db/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Tconst applies equality check predicate on the "tconst" field. It's identical to TconstEQ.
func Tconst(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTconst), v))
	})
}

// PrimaryName applies equality check predicate on the "primaryName" field. It's identical to PrimaryNameEQ.
func PrimaryName(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryName), v))
	})
}

// BirthYear applies equality check predicate on the "birthYear" field. It's identical to BirthYearEQ.
func BirthYear(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthYear), v))
	})
}

// DeathYear applies equality check predicate on the "deathYear" field. It's identical to DeathYearEQ.
func DeathYear(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeathYear), v))
	})
}

// TconstEQ applies the EQ predicate on the "tconst" field.
func TconstEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTconst), v))
	})
}

// TconstNEQ applies the NEQ predicate on the "tconst" field.
func TconstNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTconst), v))
	})
}

// TconstIn applies the In predicate on the "tconst" field.
func TconstIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTconst), v...))
	})
}

// TconstNotIn applies the NotIn predicate on the "tconst" field.
func TconstNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTconst), v...))
	})
}

// TconstGT applies the GT predicate on the "tconst" field.
func TconstGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTconst), v))
	})
}

// TconstGTE applies the GTE predicate on the "tconst" field.
func TconstGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTconst), v))
	})
}

// TconstLT applies the LT predicate on the "tconst" field.
func TconstLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTconst), v))
	})
}

// TconstLTE applies the LTE predicate on the "tconst" field.
func TconstLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTconst), v))
	})
}

// TconstContains applies the Contains predicate on the "tconst" field.
func TconstContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTconst), v))
	})
}

// TconstHasPrefix applies the HasPrefix predicate on the "tconst" field.
func TconstHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTconst), v))
	})
}

// TconstHasSuffix applies the HasSuffix predicate on the "tconst" field.
func TconstHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTconst), v))
	})
}

// TconstEqualFold applies the EqualFold predicate on the "tconst" field.
func TconstEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTconst), v))
	})
}

// TconstContainsFold applies the ContainsFold predicate on the "tconst" field.
func TconstContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTconst), v))
	})
}

// PrimaryNameEQ applies the EQ predicate on the "primaryName" field.
func PrimaryNameEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameNEQ applies the NEQ predicate on the "primaryName" field.
func PrimaryNameNEQ(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameIn applies the In predicate on the "primaryName" field.
func PrimaryNameIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrimaryName), v...))
	})
}

// PrimaryNameNotIn applies the NotIn predicate on the "primaryName" field.
func PrimaryNameNotIn(vs ...string) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrimaryName), v...))
	})
}

// PrimaryNameGT applies the GT predicate on the "primaryName" field.
func PrimaryNameGT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameGTE applies the GTE predicate on the "primaryName" field.
func PrimaryNameGTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameLT applies the LT predicate on the "primaryName" field.
func PrimaryNameLT(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameLTE applies the LTE predicate on the "primaryName" field.
func PrimaryNameLTE(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameContains applies the Contains predicate on the "primaryName" field.
func PrimaryNameContains(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameHasPrefix applies the HasPrefix predicate on the "primaryName" field.
func PrimaryNameHasPrefix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameHasSuffix applies the HasSuffix predicate on the "primaryName" field.
func PrimaryNameHasSuffix(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameEqualFold applies the EqualFold predicate on the "primaryName" field.
func PrimaryNameEqualFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrimaryName), v))
	})
}

// PrimaryNameContainsFold applies the ContainsFold predicate on the "primaryName" field.
func PrimaryNameContainsFold(v string) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrimaryName), v))
	})
}

// BirthYearEQ applies the EQ predicate on the "birthYear" field.
func BirthYearEQ(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthYear), v))
	})
}

// BirthYearNEQ applies the NEQ predicate on the "birthYear" field.
func BirthYearNEQ(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBirthYear), v))
	})
}

// BirthYearIn applies the In predicate on the "birthYear" field.
func BirthYearIn(vs ...int) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBirthYear), v...))
	})
}

// BirthYearNotIn applies the NotIn predicate on the "birthYear" field.
func BirthYearNotIn(vs ...int) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBirthYear), v...))
	})
}

// BirthYearGT applies the GT predicate on the "birthYear" field.
func BirthYearGT(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBirthYear), v))
	})
}

// BirthYearGTE applies the GTE predicate on the "birthYear" field.
func BirthYearGTE(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBirthYear), v))
	})
}

// BirthYearLT applies the LT predicate on the "birthYear" field.
func BirthYearLT(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBirthYear), v))
	})
}

// BirthYearLTE applies the LTE predicate on the "birthYear" field.
func BirthYearLTE(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBirthYear), v))
	})
}

// DeathYearEQ applies the EQ predicate on the "deathYear" field.
func DeathYearEQ(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeathYear), v))
	})
}

// DeathYearNEQ applies the NEQ predicate on the "deathYear" field.
func DeathYearNEQ(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeathYear), v))
	})
}

// DeathYearIn applies the In predicate on the "deathYear" field.
func DeathYearIn(vs ...int) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeathYear), v...))
	})
}

// DeathYearNotIn applies the NotIn predicate on the "deathYear" field.
func DeathYearNotIn(vs ...int) predicate.Name {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Name(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeathYear), v...))
	})
}

// DeathYearGT applies the GT predicate on the "deathYear" field.
func DeathYearGT(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeathYear), v))
	})
}

// DeathYearGTE applies the GTE predicate on the "deathYear" field.
func DeathYearGTE(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeathYear), v))
	})
}

// DeathYearLT applies the LT predicate on the "deathYear" field.
func DeathYearLT(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeathYear), v))
	})
}

// DeathYearLTE applies the LTE predicate on the "deathYear" field.
func DeathYearLTE(v int) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeathYear), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Name) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Name) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Name) predicate.Name {
	return predicate.Name(func(s *sql.Selector) {
		p(s.Not())
	})
}
