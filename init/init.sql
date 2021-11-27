DROP DATABASE IF EXISTS restapiforgo;
CREATE DATABASE restapiforgo;
CREATE USER minexora WITH PASSWORD '1234';
ALTER ROLE minexora SET client_encoding TO 'utf8';
ALTER ROLE minexora SET default_transaction_isolation TO 'read committed';
ALTER ROLE minexora SET timezone TO 'UTC';
GRANT ALL PRIVILEGES ON DATABASE restapiforgo TO minexora;
