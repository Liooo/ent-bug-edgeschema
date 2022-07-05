package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
    "github.com/google/uuid"
)

// Like holds the edge schema definition for the Like edge.
type Like struct {
	ent.Schema
}

func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "tweet_id"),
	}
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Time("liked_at").
			Default(time.Now),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("tweet_id", uuid.UUID{}),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
        // this generates empty primary key
		// edge.To("tweet", Tweet.Type).
		// 	Unique().
		// 	Required().
		// 	Field("tweet_id"),
		// edge.To("user", User.Type).
		// 	Unique().
		// 	Required().
		// 	Field("user_id"),
        // this generates composite primary keys
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
	}
}
