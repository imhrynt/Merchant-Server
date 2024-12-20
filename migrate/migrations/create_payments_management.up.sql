CREATE TABLE IF NOT EXISTS providers (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `provider` VARCHAR(255) NOT NULL UNIQUE,
    `api_key` VARCHAR(255) NOT NULL UNIQUE,
    `private_key` VARCHAR(255) NULL UNIQUE,
    `merchant_code` VARCHAR(255) NOT NULL UNIQUE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,,
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS payments (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT UNSIGNED NULL,
    `provider_id` INT UNSIGNED NOT NULL,
    `total_cost` DOUBLE NOT NULL,
    `payment_url` TEXT NOT NULL,
    `status` ENUM('PAID', 'UNPAID') DEFAULT 'UNPAID',
    `payment_at` TIMESTAMP NULL,
    `expired_at` TIMESTAMP NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_payments_users_id` FOREIGN KEY(`user_id`) REFERENCES users(`id`) 
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_payments_providers_id` FOREIGN KEY(`provider_id`) REFERENCES providers(`id`) 
        ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

