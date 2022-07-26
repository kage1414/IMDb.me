package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Crew holds the schema definition for the Crew entity.
type Crew struct {
	ent.Schema
}

// Fields of the Crew.
func (Crew) Fields() []ent.Field {
	return []ent.Field{
		field.String("tconst"),
		field.JSON("directors", []string{}),
		field.JSON("writers", []string{}),
	}
}

// Edges of the Crew.
func (Crew) Edges() []ent.Edge {
	return nil
}
