CREATE TABLE IF NOT EXISTS role (
  id BIGINT unsigned AUTO_INCREMENT NOT NULL,
  name VARCHAR(255) NOT NULL,
  CONSTRAINT `role_PK` PRIMARY KEY (id)
);
