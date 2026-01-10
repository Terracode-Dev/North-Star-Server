-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits 
DROP COLUMN benifits,
DROP COLUMN benifits_from,
DROP COLUMN benifits_till;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Benifits 
ADD COLUMN benefits TEXT NOT NULL,
ADD COLUMN benifits_from DATE NOT NULL,
ADD COLUMN benifits_till DATE NOT NULL;

-- +goose StatementEnd
