// Code generated by ent, DO NOT EDIT.

package ratings

import (
	"imdb-db/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Tconst applies equality check predicate on the "tconst" field. It's identical to TconstEQ.
func Tconst(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTconst), v))
	})
}

// AverageRating applies equality check predicate on the "averageRating" field. It's identical to AverageRatingEQ.
func AverageRating(v float64) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAverageRating), v))
	})
}

// NumVotes applies equality check predicate on the "numVotes" field. It's identical to NumVotesEQ.
func NumVotes(v int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumVotes), v))
	})
}

// TconstEQ applies the EQ predicate on the "tconst" field.
func TconstEQ(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTconst), v))
	})
}

// TconstNEQ applies the NEQ predicate on the "tconst" field.
func TconstNEQ(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTconst), v))
	})
}

// TconstIn applies the In predicate on the "tconst" field.
func TconstIn(vs ...string) predicate.Ratings {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ratings(func(s *sql.Selector) {
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
func TconstNotIn(vs ...string) predicate.Ratings {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ratings(func(s *sql.Selector) {
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
func TconstGT(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTconst), v))
	})
}

// TconstGTE applies the GTE predicate on the "tconst" field.
func TconstGTE(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTconst), v))
	})
}

// TconstLT applies the LT predicate on the "tconst" field.
func TconstLT(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTconst), v))
	})
}

// TconstLTE applies the LTE predicate on the "tconst" field.
func TconstLTE(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTconst), v))
	})
}

// TconstContains applies the Contains predicate on the "tconst" field.
func TconstContains(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTconst), v))
	})
}

// TconstHasPrefix applies the HasPrefix predicate on the "tconst" field.
func TconstHasPrefix(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTconst), v))
	})
}

// TconstHasSuffix applies the HasSuffix predicate on the "tconst" field.
func TconstHasSuffix(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTconst), v))
	})
}

// TconstEqualFold applies the EqualFold predicate on the "tconst" field.
func TconstEqualFold(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTconst), v))
	})
}

// TconstContainsFold applies the ContainsFold predicate on the "tconst" field.
func TconstContainsFold(v string) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTconst), v))
	})
}

// AverageRatingEQ applies the EQ predicate on the "averageRating" field.
func AverageRatingEQ(v float64) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAverageRating), v))
	})
}

// AverageRatingNEQ applies the NEQ predicate on the "averageRating" field.
func AverageRatingNEQ(v float64) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAverageRating), v))
	})
}

// AverageRatingIn applies the In predicate on the "averageRating" field.
func AverageRatingIn(vs ...float64) predicate.Ratings {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ratings(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAverageRating), v...))
	})
}

// AverageRatingNotIn applies the NotIn predicate on the "averageRating" field.
func AverageRatingNotIn(vs ...float64) predicate.Ratings {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ratings(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAverageRating), v...))
	})
}

// AverageRatingGT applies the GT predicate on the "averageRating" field.
func AverageRatingGT(v float64) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAverageRating), v))
	})
}

// AverageRatingGTE applies the GTE predicate on the "averageRating" field.
func AverageRatingGTE(v float64) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAverageRating), v))
	})
}

// AverageRatingLT applies the LT predicate on the "averageRating" field.
func AverageRatingLT(v float64) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAverageRating), v))
	})
}

// AverageRatingLTE applies the LTE predicate on the "averageRating" field.
func AverageRatingLTE(v float64) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAverageRating), v))
	})
}

// NumVotesEQ applies the EQ predicate on the "numVotes" field.
func NumVotesEQ(v int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumVotes), v))
	})
}

// NumVotesNEQ applies the NEQ predicate on the "numVotes" field.
func NumVotesNEQ(v int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumVotes), v))
	})
}

// NumVotesIn applies the In predicate on the "numVotes" field.
func NumVotesIn(vs ...int) predicate.Ratings {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ratings(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumVotes), v...))
	})
}

// NumVotesNotIn applies the NotIn predicate on the "numVotes" field.
func NumVotesNotIn(vs ...int) predicate.Ratings {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ratings(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumVotes), v...))
	})
}

// NumVotesGT applies the GT predicate on the "numVotes" field.
func NumVotesGT(v int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumVotes), v))
	})
}

// NumVotesGTE applies the GTE predicate on the "numVotes" field.
func NumVotesGTE(v int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumVotes), v))
	})
}

// NumVotesLT applies the LT predicate on the "numVotes" field.
func NumVotesLT(v int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumVotes), v))
	})
}

// NumVotesLTE applies the LTE predicate on the "numVotes" field.
func NumVotesLTE(v int) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumVotes), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ratings) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ratings) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
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
func Not(p predicate.Ratings) predicate.Ratings {
	return predicate.Ratings(func(s *sql.Selector) {
		p(s.Not())
	})
}
