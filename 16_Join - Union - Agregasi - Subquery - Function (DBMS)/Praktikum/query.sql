-- DML
-- 1a.
INSERT INTO operators (name, created_at, updated_at) VALUES
("Operator 1", "2023-03-16", "2023-03-16"),
("Operator 2", "2023-03-16", "2023-03-16"),
("Operator 3", "2023-03-16", "2023-03-16"),
("Operator 4", "2023-03-16", "2023-03-16"),
("Operator 5", "2023-03-16", "2023-03-16");

-- 1b.
INSERT INTO product_types (name, created_at, updated_at) VALUES
("Product Type 1", "2023-03-16", "2023-03-16"),
("Product Type 2", "2023-03-16", "2023-03-16"),
("Product Type 3", "2023-03-16", "2023-03-16");

-- 1cde.
INSERT INTO products (operator_id, name, product_type_id, product_description_id, price, status, created_at, updated_at) VALUES
(3, "Product 1C 1", 1, 1, 100000, 1, "2023-03-16", "2023-03-16"),
(3, "Product 1C 2", 1, 1, 100000, 2, "2023-03-16", "2023-03-16"),
(1, "Product 1D 1", 2, 1, 110000, 1, "2023-03-16", "2023-03-16"),
(1, "Product 1D 2", 2, 1, 110000, 2, "2023-03-16", "2023-03-16"),
(1, "Product 1D 3", 2, 1, 110000, 3, "2023-03-16", "2023-03-16"),
(4, "Product 1E 1", 3, 1, 120000, 1, "2023-03-16", "2023-03-16"),
(4, "Product 1E 2", 3, 1, 120000, 2, "2023-03-16", "2023-03-16"),
(4, "Product 1E 3", 3, 1, 120000, 3, "2023-03-16", "2023-03-16");

-- 1f.
INSERT INTO product_descriptions (name, created_at, updated_at) VALUES
(1, "Product Description 1", "2023-03-16", "2023-03-16"),
(2, "Product Description 2", "2023-03-16", "2023-03-16"),
(3, "Product Description 3", "2023-03-16", "2023-03-16"),
(4, "Product Description 4", "2023-03-16", "2023-03-16"),
(5, "Product Description 5", "2023-03-16", "2023-03-16");

-- 1g.
INSERT INTO payment_methods (name, status, created_at, updated_at) VALUES
(1, "Payment Method 1", 1, "2023-03-16", "2023-03-16"),
(2, "Payment Method 2", 1, "2023-03-16", "2023-03-16"),
(3, "Payment Method 3", 1, "2023-03-16", "2023-03-16");

-- 1h.
INSERT INTO users (name, address, birth_date, status_user, gender, created_at, updated_at) VALUES
("User A", "Alamat A", "2002-09-1", 1, "M", "2023-03-16", "2023-03-16"),
("User B", "Alamat B", "2002-09-2", 1, "F", "2023-03-16", "2023-03-16"),
("User C", "Alamat C", "2002-09-3", 1, "M", "2023-03-16", "2023-03-16"),
("User D", "Alamat D", "2002-09-4", 1, "F", "2023-03-16", "2023-03-16"),
("User E", "Alamat E", "2002-09-5", 1, "M", "2023-03-16", "2023-03-16");

-- 1i.
INSERT INTO transactions (user_id, payment_method_id, total_quantity, total_amount, status, created_at, updated_at) VALUES
(1, 1, 100000, 1, "Paid", "2023-03-16", "2023-03-16"),
(2, 1, 110000, 1, "Unpaid", "2023-03-16", "2023-03-16"),
(3, 2, 110000, 1, "Canceled", "2023-03-16", "2023-03-16");

-- 1j.
INSERT INTO transaction_details (transaction_id, product_id, quantity, total_price, status, created_at, updated_at) VALUES
(1, 1, 1, 100000, "Paid", "2023-03-16", "2023-03-16"),
(1, 3, 1, 110000, "Unpaid", "2023-03-16", "2023-03-16"),
(3, 3, 1, 110000, "Canceled", "2023-03-16", "2023-03-16");

-- 2a.
SELECT name FROM users WHERE gender = 'M';

-- 2b.
SELECT * FROM products WHERE id = 3;

-- 2c.
SELECT * FROM users WHERE created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';

-- 2d.
SELECT COUNT(*) as female_user_count FROM users WHERE gender = 'F';

-- 2e.
SELECT * FROM users ORDER BY name ASC;

-- 2f.
SELECT * FROM products LIMIT 5;

-- 3a.
UPDATE products SET name = 'product dummy' WHERE id = 1;

-- 3b.
UPDATE transaction_details SET quantity = 3 WHERE product_id = 1;

-- 4a.
DELETE FROM products WHERE id = 1;

-- 4b.
DELETE FROM products WHERE product_type_id = 1;

-- Join, Union, Sub query, Function
-- 1.
SELECT * FROM transactions WHERE user_id = 1 
UNION 
SELECT * FROM transactions WHERE user_id = 2;

-- 2.
SELECT SUM(total_amount) as jumlah_harga FROM transactions t 
JOIN transaction_details td ON t.id = td.transaction_id 
WHERE t.user_id = 1;

-- 3.
SELECT COUNT(*) as total_transaksi FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
WHERE p.product_type_id = 2;

-- 4.
SELECT p.*, pt.name as product_type_name FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

-- 5.
SELECT t.*, p.name as product_name, u.name as user_name FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
JOIN users u ON t.user_id = u.id;

-- 6.
DELIMITER $$
CREATE FUNCTION delete_transaction_detail_trigger() RETURNS TRIGGER
BEGIN
   DELETE FROM transaction_details
   WHERE transaction_id = OLD.transaction_id;
   RETURN OLD;
END $$
DELIMITER ;

CREATE TRIGGER delete_transaction_detail
AFTER DELETE ON transactions
FOR EACH ROW
CALL delete_transaction_detail_trigger();

-- 7.
DELIMITER $$
CREATE FUNCTION update_total_qty_trigger() RETURNS TRIGGER
BEGIN
   UPDATE transactions
   SET total_quantity = total_quantity - OLD.total_quantity
   WHERE id = OLD.id;
   RETURN OLD;
END;
DELIMITER ;

CREATE TRIGGER update_total_qty
AFTER DELETE ON transaction_details
FOR EACH ROW
CALL update_total_qty_trigger();

-- 8.
SELECT * FROM products WHERE id NOT IN (
   SELECT product_id
   FROM transaction_details
);
