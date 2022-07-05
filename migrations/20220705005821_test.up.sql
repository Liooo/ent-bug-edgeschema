-- create "tweets" table
CREATE TABLE "tweets" ("id" uuid NOT NULL, "text" text NOT NULL, PRIMARY KEY ("id"));
-- create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "name" character varying NOT NULL DEFAULT 'Unknown', PRIMARY KEY ("id"));
-- create "likes" table
CREATE TABLE "likes" ("liked_at" timestamptz NOT NULL, "tweet_id" uuid NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY (), CONSTRAINT "likes_tweets_tweet" FOREIGN KEY ("tweet_id") REFERENCES "tweets" ("id") ON DELETE NO ACTION, CONSTRAINT "likes_users_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION);
