package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Principals holds the schema definition for the Principals entity.
type Principals struct {
	ent.Schema
}

// Fields of the Principals.
func (Principals) Fields() []ent.Field {
	return []ent.Field{
		field.String("tconst"),
		field.Int("ordering").
			Positive(),
		field.String("nconst"),
		field.String("category"),
		field.String("job"),
		field.String("characters"),
	}
}

// Edges of the Principals.
func (Principals) Edges() []ent.Edge {
	return nil
}
