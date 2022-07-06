package schema

import (
	"entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("info_for_follower", RelationshipInfo.Type).Through("follower", Relationship.Type),
        edge.To("info_for_followee", RelationshipInfo.Type).Through("followee", Relationship.Type),
    }
}
