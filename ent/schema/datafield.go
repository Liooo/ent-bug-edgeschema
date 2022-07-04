package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type DataField struct {
	ent.Schema
}

// Fields of the Field.
func (DataField) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Field.
func (DataField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("records", Record.Type).Through("cells", Cell.Type),
	}
}
