ALTER TABLE "user_devices" DROP CONSTRAINT IF EXISTS "user_devices_user_id_fkey";

ALTER TABLE "user_devices" DROP CONSTRAINT IF EXISTS "user_devices_device_id_fkey";


DROP TABLE IF EXISTS "user_devices";
DROP TABLE IF EXISTS "devices";