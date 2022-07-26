package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Akas holds the schema definition for the Akas entity.
type Akas struct {
	ent.Schema
}

// Fields of the Akas.
func (Akas) Fields() []ent.Field {
	return []ent.Field{
		field.String("titleId"),
		field.Int("ordering").
			Positive(),
		field.String("title"),
		field.String("region"),
		field.String("language"),
		field.JSON("types", []string{}),
		field.JSON("attributes", []string{}),
		field.Bool("isOriginalTitle"),
	}
}

// Edges of the Akas.
func (Akas) Edges() []ent.Edge {
	return nil
}
