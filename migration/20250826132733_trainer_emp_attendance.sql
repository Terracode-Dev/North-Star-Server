-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS HR_EMP_ATTENDANCE (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  emp_id bigint(20) NOT NULL,
  attendance_type enum('in','out') NOT NULL,
  create_date timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (id),
  KEY emp_id (emp_id),
  KEY idx_emp_date (emp_id,create_date),
  CONSTRAINT HR_EMP_ATTENDANCE_ibfk_1 FOREIGN KEY (emp_id) REFERENCES HR_Employee (id) ON DELETE CASCADE ON UPDATE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS HR_EMP_ATTENDANCE;
-- +goose StatementEnd
