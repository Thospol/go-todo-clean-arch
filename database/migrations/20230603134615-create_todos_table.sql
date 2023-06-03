
-- +migrate Up
CREATE TABLE `todo` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT COMMENT 'ไอดี',
  `created_at` timestamptz NOT NULL COMMENT 'วันเวลาที่สร้าง',
  `updated_at` timestamptz NOT NULL COMMENT 'วันเวลาที่อัพเดตล่าสุด',
  `deleted_at` timestamptz DEFAULT null COMMENT 'วันเวลาที่ลบ',
  `title` varchar(255) NOT NULL COMMENT 'หัวข้อ',
  `description` varchar(255) NOT NULL COMMENT 'รายละเอียด'
);

ALTER TABLE `todo` COMMENT = 'todo table';

-- +migrate Down
DROP TABLE IF EXISTS `todo`;
