-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_EMP_Benifits(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    leave_status BOOLEAN NOT NULL,
    leave_type VARCHAR(100) NOT NULL,
    leave_count INT NOT NULL,
    health_insurance VARCHAR(100) NOT NULL,
    insurance_from DATE NOT NULL,
    insurance_till DATE NOT NULL,
    retainment_plan VARCHAR(200) NOT NULL,
    retainment_plan_from DATE NOT NULL,
    retainment_plan_till DATE NOT NULL,
    benifits VARCHAR(200) NOT NULL,
    benifits_from DATE NOT NULL,
    benifits_till DATE NOT NULL,
    updated_by BIGINT,
    employee_id BIGINT NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES HR_Employee(id),
    FOREIGN KEY (updated_by) REFERENCES HR_Admin(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_EMP_Benifits;
-- +goose StatementEnd
