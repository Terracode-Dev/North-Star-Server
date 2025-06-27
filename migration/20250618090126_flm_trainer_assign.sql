-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS FLM_trainer_assign (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    trainer_id BIGINT NOT NULL,
    branch_id BIGINT NOT NULL,
    client_id BIGINT NOT NULL,
    flm_id BIGINT NOT NULL,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `from` TIMESTAMP NOT NULL,
    `to` TIMESTAMP NOT NULL,

    FOREIGN KEY (flm_id) REFERENCES FLM_users(id),
    FOREIGN KEY (client_id) REFERENCES door_lock_users(attendee_id),
    FOREIGN KEY (trainer_id) REFERENCES door_lock_users(attendee_id),
    FOREIGN KEY (branch_id) REFERENCES HR_Branch(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS FLM_trainer_assign;
-- +goose StatementEnd
