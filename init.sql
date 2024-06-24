CREATE DATABASE IF NOT EXISTS bank;

USE bank;

CREATE TABLE IF NOT EXISTS accounts (
    account_id INT AUTO_INCREMENT PRIMARY KEY,
    document_number VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS operation_types (
    operation_type_id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    account_id INT,
    operation_type_id INT,
    amount DECIMAL(10, 2),
    event_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id),
    FOREIGN KEY (operation_type_id) REFERENCES operation_types(operation_type_id)
);

INSERT INTO operation_types (description) VALUES ('Normal Purchase');
INSERT INTO operation_types (description) VALUES ('Purchase with installments');
INSERT INTO operation_types (description) VALUES ('Withdrawal');
INSERT INTO operation_types (description) VALUES ('Credit Voucher');
