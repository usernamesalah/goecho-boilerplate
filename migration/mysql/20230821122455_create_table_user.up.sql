BEGIN;

CREATE TABLE IF NOT EXISTS user (
  id BIGINT unsigned AUTO_INCREMENT NOT NULL,
  role_id BIGINT NOT NULL,
  uid VARCHAR(27) NOT NULL,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(40) NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT `user_PK` PRIMARY KEY (id),
  CONSTRAINT `user_UID` UNIQUE KEY (uid),
  FOREIGN KEY (role_id) REFERENCES role(id),
);


CREATE TABLE IF NOT EXISTS profile (
  id BIGINT unsigned AUTO_INCREMENT NOT NULL,
  user_id BIGINT NOT NULL,
  address TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT `profile_PK` PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES user(id),
);


COMMIT;