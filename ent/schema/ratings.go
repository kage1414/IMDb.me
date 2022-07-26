package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Ratings holds the schema definition for the Ratings entity.
type Ratings struct {
	ent.Schema
}

// Fields of the Ratings.
func (Ratings) Fields() []ent.Field {
	return []ent.Field{
		field.String("tconst"),
		field.Float("averageRating").
			Positive(),
		field.Int("numVotes"),
	}
}

// Edges of the Ratings.
func (Ratings) Edges() []ent.Edge {
	return nil
}
