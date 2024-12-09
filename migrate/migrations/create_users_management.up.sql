CREATE TABLE IF NOT EXISTS roles (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `role` VARCHAR(255) NOT NULL UNIQUE,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `roles` (`role`) VALUES 
    ('ADMIN'),
    ('OPERATOR'),
    ('USER');

CREATE TABLE IF NOT EXISTS users (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `role_id` INT UNSIGNED NOT NULL DEFAULT 4,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `phone` VARCHAR(255) NOT NULL UNIQUE,
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `img_url` VARCHAR(255) NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_users_roles_id` FOREIGN KEY(`role_id`) REFERENCES roles(`id`) ON DELETE CASCADE ON UPDATE CASCADE 
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX index_username ON users (username);

INSERT INTO `users` (`role_id`, `name`, `email`, `phone`, `username`, `password`) VALUES 
    (1, 'amirullah heryanto muslan', 'imhrynt@gmail.com', '082131802323', 'admin1234', '$2a$12$m1hf7lKypWnvitLLf7GloOI7AX/eqBqD6u0/y/H/zrUpqIlZKLTwm');

CREATE TABLE IF NOT EXISTS balances (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `balance` DOUBLE NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_balances_users_id` FOREIGN KEY(`user_id`) REFERENCES users(`id`) ON DELETE CASCADE ON UPDATE CASCADE 
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `balances` (`user_id`) VALUES 
    (1);
