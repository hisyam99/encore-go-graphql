-- create "blogs" table
CREATE TABLE "blogs" (
  "id" bigserial NOT NULL,
  "title" text NULL,
  "content" text NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- create "projects" table
CREATE TABLE "projects" (
  "id" bigserial NOT NULL,
  "title" text NULL,
  "description" text NULL,
  "user_id" bigint NULL,
  PRIMARY KEY ("id")
);
-- create "resumes" table
CREATE TABLE "resumes" (
  "id" bigserial NOT NULL,
  "title" text NULL,
  "description" text NULL,
  "category" text NULL,
  PRIMARY KEY ("id")
);
-- create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "email" text NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
