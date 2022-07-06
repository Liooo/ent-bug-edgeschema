package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Relationship holds the schema definition for the Relationship entity.
type Relationship struct {
    ent.Schema
}

// Fields of the Relationship.
func (Relationship) Fields() []ent.Field {
    return []ent.Field{
        field.Int("follower_user_id"),
        field.Int("followee_user_id"),
        field.Int("relationship_info_id"),
    }
}

// Edges of the Relationship.
func (Relationship) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("follower", User.Type).Required().Unique().Field("follower_user_id"),
        edge.To("followee", User.Type).Required().Unique().Field("followee_user_id"),
        edge.To("relationship_info", RelationshipInfo.Type).Required().Unique().Field("relationship_info_id"),
    }
}
