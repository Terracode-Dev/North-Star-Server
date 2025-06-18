-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS HR_Trainer_Emp (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    trainer_id BIGINT NOT NULL,
    employee_id BIGINT NOT NULL,
    attendee_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (employee_id) REFERENCES HR_Employee(id),
    FOREIGN KEY (attendee_id) REFERENCES door_lock_users(attendee_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS HR_Trainer_Emp;
-- +goose StatementEnd
