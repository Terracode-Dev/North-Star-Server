-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS emp_loan_req (
    id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    emp_id BIGINT NOT NULL,
    reason VARCHAR(200),
    amount DECIMAL(10,2),
    status VARCHAR(100),
    declined_by BIGINT,
    decline_reason VARCHAR(150),
    requested_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_changed_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (emp_id) REFERENCES HR_Employee(id),
    FOREIGN KEY (declined_by) REFERENCES HR_Admin(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS emp_loan_req;
-- +goose StatementEnd
