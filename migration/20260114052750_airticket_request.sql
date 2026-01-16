-- +goose Up
-- +goose StatementBegin
CREATE TABLE emp_airticket_req(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    passenger_name VARCHAR(100) NOT NULL,
    passenger_email VARCHAR(100) NOT NULL,
    passport_number VARCHAR(50) NOT NULL,
    departure_date DATE NOT NULL,
    return_date DATE NOT NULL,
    departure_city VARCHAR(100) NOT NULL,
    arrival_city VARCHAR(100) NOT NULL,
    reason TEXT NOT NULL,
    emp_id BIGINT NOT NULL,
    branch_id BIGINT NOT NULL,
    status ENUM('pending', 'approved', 'rejected') NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (emp_id) REFERENCES HR_Employee(id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES HR_Branch(id) ON DELETE CASCADE,
    
    INDEX idx_emp_id (emp_id),
    INDEX idx_branch_id (branch_id),
    INDEX idx_status (status),
    INDEX idx_emp_branch (emp_id, branch_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS emp_airticket_req;
-- +goose StatementEnd
