-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "last_login", DROP COLUMN "is_active", DROP COLUMN "role", DROP COLUMN "password";
