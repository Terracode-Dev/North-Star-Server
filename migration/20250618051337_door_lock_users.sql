-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS door_lock_users (
    attendee_id BIGINT NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    branch_id BIGINT NOT NULL,
    subscription_id BIGINT,
    last_in TIMESTAMP NULL,
    last_out TIMESTAMP NULL,
    current_state VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    role VARCHAR(50),
    phone VARCHAR(20),
    nic VARCHAR(20),
    avatar_url TEXT,

    FOREIGN KEY (branch_id) REFERENCES HR_Branch(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS door_lock_users;
-- +goose StatementEnd
