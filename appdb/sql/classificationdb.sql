DROP DATABASE IF EXISTS classificationdb;
CREATE DATABASE IF NOT EXISTS classificationdb;
USE classificationdb;


-- Table to store database connection details securely
CREATE TABLE database_connections (
    id INT AUTO_INCREMENT PRIMARY KEY,
    host VARCHAR(255) NOT NULL,
    port INT NOT NULL,
    dbusername VARCHAR(100) NOT NULL,
    dbpassword VARCHAR(255) NOT NULL,
    dbname VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_connection UNIQUE (host, port, dbusername, dbname)
);

-- Table to store scan history
CREATE TABLE scan_histories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    database_id INT NOT NULL,
    scan_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    scan_status ENUM('IN_PROGRESS', 'COMPLETED', 'FAILED') DEFAULT 'IN_PROGRESS',
    FOREIGN KEY (database_id) REFERENCES database_connections(id) ON DELETE CASCADE
);

-- Table to store the scanned structure of the databases
CREATE TABLE scanned_tables (
    id INT AUTO_INCREMENT PRIMARY KEY,
    scan_id INT NOT NULL,
    table_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (scan_id) REFERENCES scan_histories(id) ON DELETE CASCADE
);

-- Table to store scan and classification results
CREATE TABLE scan_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    table_id INT NOT NULL,
    column_name VARCHAR(255) NOT NULL,
    information_type VARCHAR(100),
    FOREIGN KEY (table_id) REFERENCES scanned_tables(id) ON DELETE CASCADE
);

-- Table to store the generate api keys
CREATE TABLE api_keys (
    id INT AUTO_INCREMENT PRIMARY KEY,
    api_key VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
