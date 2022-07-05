// Code generated by ent, DO NOT EDIT.

package like

import (
	"time"
)

const (
	// Label holds the string label denoting the like type in the database.
	Label = "like"
	// FieldLikedAt holds the string denoting the liked_at field in the database.
	FieldLikedAt = "liked_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTweetID holds the string denoting the tweet_id field in the database.
	FieldTweetID = "tweet_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeTweet holds the string denoting the tweet edge name in mutations.
	EdgeTweet = "tweet"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// TweetFieldID holds the string denoting the ID field of the Tweet.
	TweetFieldID = "id"
	// Table holds the table name of the like in the database.
	Table = "likes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "likes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// TweetTable is the table that holds the tweet relation/edge.
	TweetTable = "likes"
	// TweetInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	TweetInverseTable = "tweets"
	// TweetColumn is the table column denoting the tweet relation/edge.
	TweetColumn = "tweet_id"
)

// Columns holds all SQL columns for like fields.
var Columns = []string{
	FieldLikedAt,
	FieldUserID,
	FieldTweetID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultLikedAt holds the default value on creation for the "liked_at" field.
	DefaultLikedAt func() time.Time
)
