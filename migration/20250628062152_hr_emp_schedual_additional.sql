-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `HR_EMP_SCHEDUAL_additional` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `emp_id` bigint(20) NOT NULL,
  `date` date NOT NULL,
  `from_time` time DEFAULT NULL,
  `to_time` time DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_emp_date` (`emp_id`, `date`),
  KEY `emp_id` (`emp_id`),
  KEY `date_idx` (`date`),
  CONSTRAINT `HR_EMP_SCHEDUAL_additional_ibfk_1` FOREIGN KEY (`emp_id`) REFERENCES `HR_Employee` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS HR_EMP_SCHEDUAL_additional;
-- +goose StatementEnd
