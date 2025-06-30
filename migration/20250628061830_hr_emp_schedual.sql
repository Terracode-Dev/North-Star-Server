-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `HR_EMP_SCHEDUAL` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `emp_id` bigint(20) NOT NULL,
  `monday` bool DEFAULT false COMMENT 'true = working day, false = off day',
  `monday_from` time DEFAULT NULL,
  `monday_to` time DEFAULT NULL,
  `tuesday` bool DEFAULT false,
  `tuesday_from` time DEFAULT NULL,
  `tuesday_to` time DEFAULT NULL,
  `wednesday` bool DEFAULT false,
  `wednesday_from` time DEFAULT NULL,
  `wednesday_to` time DEFAULT NULL,
  `thursday` bool DEFAULT false,
  `thursday_from` time DEFAULT NULL,
  `thursday_to` time DEFAULT NULL,
  `friday` bool DEFAULT false,
  `friday_from` time DEFAULT NULL,
  `friday_to` time DEFAULT NULL,
  `saturday` bool DEFAULT false,
  `saturday_from` time DEFAULT NULL,
  `saturday_to` time DEFAULT NULL,
  `sunday` bool DEFAULT false,
  `sunday_from` time DEFAULT NULL,
  `sunday_to` time DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_emp_schedule` (`emp_id`),
  KEY `emp_id` (`emp_id`),
  CONSTRAINT `HR_EMP_SCHEDUAL_ibfk_1` FOREIGN KEY (`emp_id`) REFERENCES `HR_Employee` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS HR_EMP_SCHEDUAL;
-- +goose StatementEnd
