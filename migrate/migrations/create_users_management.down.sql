ALTER TABLE users DROP FOREIGN KEY fk_users_roles_id;
ALTER TABLE balances DROP FOREIGN KEY fk_balances_users_id;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS balances;