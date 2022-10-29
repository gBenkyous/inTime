package schema

import "entgo.io/ent"

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return nil
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return nil
}
