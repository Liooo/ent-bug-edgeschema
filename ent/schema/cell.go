package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ----- cell.go

// Cell holds the schema definition for the CellText entity.
type Cell struct {
	ent.Schema
}

// Fields of the Cell.
func (Cell) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("data_field_id", uuid.UUID{}),
		field.UUID("record_id", uuid.UUID{}),
	}
}

func (Cell) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("data_field_id", "record_id"),
	}
}

// Edges of the Cell.
func (Cell) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("record", Record.Type).Unique().Required().Field("record_id"),
		edge.To("data_field", DataField.Type).Unique().Required().Field("data_field_id"),
	}
}
