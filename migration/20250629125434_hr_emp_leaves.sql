-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `HR_EMP_LEAVES` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `emp_id` bigint(20) NOT NULL,
    `leave_type` varchar(50) NOT NULL,
    `leave_date` date NOT NULL,
    `reason` text NOT NULL,
    `create_date` timestamp NULL DEFAULT current_timestamp(),
    `added_by` bigint(20) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `emp_id` (`emp_id`),
    KEY `added_by` (`added_by`),
    CONSTRAINT `HR_EMP_LEAVES_ibfk_1` FOREIGN KEY (`emp_id`) REFERENCES `HR_Employee` (`id`) ON DELETE CASCADE,
    CONSTRAINT `HR_EMP_LEAVES_ibfk_2` FOREIGN KEY (`added_by`) REFERENCES `HR_Admin` (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS HR_EMP_LEAVES;
-- +goose StatementEnd
