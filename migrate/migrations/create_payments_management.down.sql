ALTER TABLE users DROP FOREIGN KEY fk_payments_users_id;
ALTER TABLE providers DROP FOREIGN KEY fk_payments_providers_id;
DROP TABLE IF EXISTS providers;
DROP TABLE IF EXISTS payments;