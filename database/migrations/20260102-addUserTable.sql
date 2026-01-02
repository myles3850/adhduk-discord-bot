CREATE TABLE IF NOT EXISTS "users" (
	"id" serial NOT NULL UNIQUE,
	"discord_id" bigint NOT NULL UNIQUE,
	"discord_user" varchar(255) NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "todo" (
	"id" serial NOT NULL UNIQUE,
	"channel_message_id" varchar(255) NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" varchar(255),
	"completed" varchar(255) NOT NULL DEFAULT '0',
	"completed_by" bigint NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "messages" (
	"id" serial NOT NULL UNIQUE,
	"message_id" bigint NOT NULL,
	"user_id" bigint NOT NULL,
	"channel" varchar(255) NOT NULL,
	"deleted" varchar(255) NOT NULL DEFAULT '0',
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "reminders" (
	"id" serial NOT NULL UNIQUE,
	"title" varchar(255) NOT NULL,
	"additional_info" varchar(255),
	"user_created" bigint NOT NULL,
	"users_to_tag" bigint NOT NULL,
	"reminder_datetime" timestamp with time zone NOT NULL,
	PRIMARY KEY ("id")
);



ALTER TABLE "todo" ADD CONSTRAINT "todo_fk5" FOREIGN KEY ("completed_by") REFERENCES "users"("id");
ALTER TABLE "messages" ADD CONSTRAINT "messages_fk2" FOREIGN KEY ("user_id") REFERENCES "users"("id");
ALTER TABLE "reminders" ADD CONSTRAINT "reminders_fk3" FOREIGN KEY ("user_created") REFERENCES "users"("id");
