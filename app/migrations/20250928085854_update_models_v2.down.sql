-- reverse: drop "resumes" table
CREATE TABLE "resumes" (
  "id" bigserial NOT NULL,
  "title" text NULL,
  "description" text NULL,
  "category" text NULL,
  PRIMARY KEY ("id")
);
-- reverse: create index "idx_resume_contents_deleted_at" to table: "resume_contents"
DROP INDEX "idx_resume_contents_deleted_at";
-- reverse: create index "idx_resume_contents_category_id" to table: "resume_contents"
DROP INDEX "idx_resume_contents_category_id";
-- reverse: create "resume_contents" table
DROP TABLE "resume_contents";
-- reverse: create index "idx_categories_name" to table: "categories"
DROP INDEX "idx_categories_name";
-- reverse: create index "idx_categories_deleted_at" to table: "categories"
DROP INDEX "idx_categories_deleted_at";
-- reverse: create "categories" table
DROP TABLE "categories";
-- reverse: create index "idx_projects_user_id" to table: "projects"
DROP INDEX "idx_projects_user_id";
-- reverse: create index "idx_projects_deleted_at" to table: "projects"
DROP INDEX "idx_projects_deleted_at";
-- reverse: modify "projects" table
ALTER TABLE "projects" DROP CONSTRAINT "fk_users_projects", DROP COLUMN "deleted_at", DROP COLUMN "updated_at", DROP COLUMN "created_at", ALTER COLUMN "description" TYPE text, ALTER COLUMN "title" TYPE text, ALTER COLUMN "title" DROP NOT NULL;
-- reverse: create index "idx_users_deleted_at" to table: "users"
DROP INDEX "idx_users_deleted_at";
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "deleted_at", DROP COLUMN "updated_at", ALTER COLUMN "email" TYPE text, ALTER COLUMN "email" DROP NOT NULL, ALTER COLUMN "name" TYPE text, ALTER COLUMN "name" DROP NOT NULL;
-- reverse: create index "idx_blogs_status" to table: "blogs"
DROP INDEX "idx_blogs_status";
-- reverse: create index "idx_blogs_slug" to table: "blogs"
DROP INDEX "idx_blogs_slug";
-- reverse: create index "idx_blogs_deleted_at" to table: "blogs"
DROP INDEX "idx_blogs_deleted_at";
-- reverse: modify "blogs" table
ALTER TABLE "blogs" DROP COLUMN "deleted_at", DROP COLUMN "updated_at", DROP COLUMN "meta_description", DROP COLUMN "tags", DROP COLUMN "status", DROP COLUMN "published_at", DROP COLUMN "author", DROP COLUMN "slug", DROP COLUMN "summary", ALTER COLUMN "content" DROP NOT NULL, ALTER COLUMN "title" TYPE text, ALTER COLUMN "title" DROP NOT NULL;
