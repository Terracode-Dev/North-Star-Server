-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_Create_Services(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    category VARCHAR(200) NOT NULL,
    value VARCHAR(200) NOT NULL,
    updated_by BIGINT,
    FOREIGN KEY (updated_by) REFERENCES HR_Admin(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_Create_Services;
-- +goose StatementEnd
