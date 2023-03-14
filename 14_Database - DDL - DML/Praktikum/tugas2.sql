-- 1. Untuk membuat database baru dengan nama "alta_online_shop", kita dapat menggunakan perintah SQL berikut:
CREATE DATABASE alta_online_shop;

use alta_online_shop;

-- 2a. Create table user:
CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  address VARCHAR(255),
  birth_date DATE,
  status_user VARCHAR(50),
  gender CHAR(1),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2b. Create table product, product_type, operator, product_description, payment_method:
CREATE TABLE products (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  product_type_id INT,
  product_description_id INT,
  price INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE product_types (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE operators (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE product_descriptions (
  id INT PRIMARY KEY AUTO_INCREMENT,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE payment_methods (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2c. Create table transaction, transaction_detail:
CREATE TABLE transactions (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT,
  operator_id INT,
  payment_method_id INT,
  total_amount INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE transaction_details (
  id INT PRIMARY KEY AUTO_INCREMENT,
  transaction_id INT,
  product_id INT,
  quantity INT,
  total_price INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 3. Create tabel kurir dengan field id, name, created_at, updated_at:
CREATE TABLE kurir (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 4. Tambahkan ongkos_dasar column di tabel kurir:
ALTER TABLE kurir ADD ongkos_dasar INT;

-- 5. Rename tabel kurir menjadi shipping:
ALTER TABLE kurir RENAME TO shipping;

-- 6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan:
DROP TABLE shipping;

-- 7a. 1-to-1: payment method description.
CREATE TABLE payment_method_descriptions (
  id INT PRIMARY KEY AUTO_INCREMENT,
  description TEXT,
  payment_method_id INT,
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 7b. 1-to-many: user dengan alamat.
CREATE TABLE alamats (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT,
  nama_jalan VARCHAR(255),
  kota VARCHAR(255),
  provinsi VARCHAR(255),
  kode_pos VARCHAR(10),
  FOREIGN KEY (user_id) REFERENCES users(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 7c. many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE user_payment_method_details (
  user_id INT,
  payment_method_id INT,
  PRIMARY KEY (user_id, payment_method_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);