-- modify "blogs" table
ALTER TABLE "blogs" ALTER COLUMN "title" TYPE character varying(255), ALTER COLUMN "title" SET NOT NULL, ALTER COLUMN "content" SET NOT NULL, ADD COLUMN "summary" character varying(500) NULL, ADD COLUMN "slug" character varying(255) NOT NULL, ADD COLUMN "author" character varying(255) NULL, ADD COLUMN "published_at" timestamptz NULL, ADD COLUMN "status" text NULL DEFAULT 'draft', ADD COLUMN "tags" jsonb NULL, ADD COLUMN "meta_description" character varying(160) NULL, ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "deleted_at" timestamptz NULL;
-- create index "idx_blogs_deleted_at" to table: "blogs"
CREATE INDEX "idx_blogs_deleted_at" ON "blogs" ("deleted_at");
-- create index "idx_blogs_slug" to table: "blogs"
CREATE UNIQUE INDEX "idx_blogs_slug" ON "blogs" ("slug");
-- create index "idx_blogs_status" to table: "blogs"
CREATE INDEX "idx_blogs_status" ON "blogs" ("status");
-- modify "users" table
ALTER TABLE "users" ALTER COLUMN "name" TYPE character varying(255), ALTER COLUMN "name" SET NOT NULL, ALTER COLUMN "email" TYPE character varying(255), ALTER COLUMN "email" SET NOT NULL, ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "deleted_at" timestamptz NULL;
-- create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- modify "projects" table
ALTER TABLE "projects" ALTER COLUMN "title" TYPE character varying(255), ALTER COLUMN "title" SET NOT NULL, ALTER COLUMN "description" TYPE character varying(1000), ADD COLUMN "created_at" timestamptz NULL, ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "deleted_at" timestamptz NULL, ADD CONSTRAINT "fk_users_projects" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- create index "idx_projects_deleted_at" to table: "projects"
CREATE INDEX "idx_projects_deleted_at" ON "projects" ("deleted_at");
-- create index "idx_projects_user_id" to table: "projects"
CREATE INDEX "idx_projects_user_id" ON "projects" ("user_id");
-- create "categories" table
CREATE TABLE "categories" (
  "id" bigserial NOT NULL,
  "name" character varying(100) NOT NULL,
  "description" character varying(500) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_categories_deleted_at" to table: "categories"
CREATE INDEX "idx_categories_deleted_at" ON "categories" ("deleted_at");
-- create index "idx_categories_name" to table: "categories"
CREATE UNIQUE INDEX "idx_categories_name" ON "categories" ("name");
-- create "resume_contents" table
CREATE TABLE "resume_contents" (
  "id" bigserial NOT NULL,
  "title" character varying(255) NOT NULL,
  "description" character varying(500) NULL,
  "detail" text NULL,
  "category_id" bigint NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_categories_resume_contents" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_resume_contents_category_id" to table: "resume_contents"
CREATE INDEX "idx_resume_contents_category_id" ON "resume_contents" ("category_id");
-- create index "idx_resume_contents_deleted_at" to table: "resume_contents"
CREATE INDEX "idx_resume_contents_deleted_at" ON "resume_contents" ("deleted_at");
-- drop "resumes" table
DROP TABLE "resumes";
