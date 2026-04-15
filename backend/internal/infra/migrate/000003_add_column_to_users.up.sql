ALTER TABLE users ADD COLUMN department_id UUID NOT NULL;

ALTER TABLE users
  ADD CONSTRAINT fk_users_department
    FOREIGN KEY (department_id)
    REFERENCES departments(department_id)
    ON DELETE RESTRICT
    ON UPDATE CASCADE;