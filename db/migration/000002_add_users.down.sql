ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "user_id_currency_key";
DROP INDEX IF EXISTS "users_id_idx";
DROP INDEX IF EXISTS "accounts_user_id_idx";
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_user_id_fkey";
ALTER TABLE IF EXISTS "accounts" DROP COLUMN IF EXISTS "user_id";
DROP TABLE IF EXISTS "users";