package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Title holds the schema definition for the Title entity.
type Title struct {
	ent.Schema
}

// Fields of the Title.
func (Title) Fields() []ent.Field {
	return []ent.Field{
		field.String("tconst"),
		field.String("titleType"),
		field.String("primaryTitle"),
		field.String("originalTitle"),
		field.Bool("isAdult").
			Default(false),
		field.Int("startYear").
			Positive(),
		field.Int("endYear").
			Positive(),
		field.Int("runtimeMinutes").
			Positive(),
		field.JSON("genre", []string{}),
	}
}

// Edges of the Title.
func (Title) Edges() []ent.Edge {
	return nil
}
