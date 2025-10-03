-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS emp_link (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    emp_data JSON NOT NULL,
    preset_id BIGINT(20) NOT NULL,
    is_approved BOOLEAN NOT NULL DEFAULT FALSE,
    create_date timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (id),
    FOREIGN KEY (preset_id) REFERENCES Admin_Presets (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS emp_link;
-- +goose StatementEnd
