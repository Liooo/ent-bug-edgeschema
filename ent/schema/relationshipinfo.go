package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
)

// RelationshipInfo holds the schema definition for the RelationshipInfo entity.
type RelationshipInfo struct {
    ent.Schema
}

// Fields of the RelationshipInfo.
func (RelationshipInfo) Fields() []ent.Field {
    return []ent.Field{
    }
}

// Edges of the RelationshipInfo.
func (RelationshipInfo) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("follower_user", User.Type).Ref("info_for_follower").Through("relationship", Relationship.Type),
        edge.From("followee_user", User.Type).Ref("info_for_followee").Through("relationship", Relationship.Type),
    }
}
