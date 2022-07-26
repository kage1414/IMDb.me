package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Name holds the schema definition for the Name entity.
type Name struct {
	ent.Schema
}

// Fields of the Name.
func (Name) Fields() []ent.Field {
	return []ent.Field{
		field.String("tconst"),
		field.String("primaryName"),
		field.Int("birthYear").
			Positive(),
		field.Int("deathYear").
			Positive(),
		field.JSON("primaryProfession", []string{}),
	}
}

// Edges of the Name.
func (Name) Edges() []ent.Edge {
	return nil
}
