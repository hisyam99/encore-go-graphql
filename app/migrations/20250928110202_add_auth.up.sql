-- modify "users" table
-- Add columns that can be NULL initially
ALTER TABLE "users" 
ADD COLUMN "password" character varying(255) NULL, 
ADD COLUMN "role" character varying(20) NOT NULL DEFAULT 'VIEWER', 
ADD COLUMN "is_active" boolean NOT NULL DEFAULT true, 
ADD COLUMN "last_login" timestamptz NULL;

-- Set default password for existing users (they will need to reset)
UPDATE "users" SET "password" = '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi' WHERE "password" IS NULL;

-- Now make password column NOT NULL
ALTER TABLE "users" ALTER COLUMN "password" SET NOT NULL;
