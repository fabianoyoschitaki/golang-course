CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;

DROP TABLE IF EXISTS user_followers;
CREATE TABLE user_followers (
    -- it must exist in users table. on delete cascade, if user is deleted, all rows within this table are also deleted
    user_id INT NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    follower_user_id INT NOT NULL, FOREIGN KEY (follower_user_id) REFERENCES users(id) ON DELETE CASCADE,

    PRIMARY KEY(user_id, follower_user_id)
) ENGINE=INNODB;