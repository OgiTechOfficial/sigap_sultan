-- Create schema prod
CREATE SCHEMA IF NOT EXISTS prod;

-- Set default search path
SET search_path TO prod, public;

-- 1. settings table
CREATE TABLE IF NOT EXISTS prod.settings (
    id SERIAL PRIMARY KEY,
    parent_id INTEGER,
    name VARCHAR(255) NOT NULL,
    value VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 2. tm_province table
CREATE TABLE IF NOT EXISTS prod.tm_province (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 3. tm_city table
CREATE TABLE IF NOT EXISTS prod.tm_city (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    province_id INTEGER REFERENCES prod.tm_province(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    sequence INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 4. tm_position table
CREATE TABLE IF NOT EXISTS prod.tm_position (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 5. tm_menu table
CREATE TABLE IF NOT EXISTS prod.tm_menu (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255),
    icon VARCHAR(255),
    parent_id INTEGER,
    position INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 6. map_jabatan_menu table
CREATE TABLE IF NOT EXISTS prod.map_jabatan_menu (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    position_id INTEGER REFERENCES prod.tm_position(id) ON DELETE CASCADE,
    menu_id INTEGER REFERENCES prod.tm_menu(id) ON DELETE CASCADE,
    read_status INTEGER DEFAULT 1,
    create_status INTEGER DEFAULT 1,
    update_status INTEGER DEFAULT 1,
    delete_status INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 7. tm_user table
CREATE TABLE IF NOT EXISTS prod.tm_user (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 8. tm_profile table
CREATE TABLE IF NOT EXISTS prod.tm_profile (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    user_id INTEGER REFERENCES prod.tm_user(id) ON DELETE CASCADE,
    nama_depan VARCHAR(255),
    nama_belakang VARCHAR(255),
    jabatan VARCHAR(255),
    position_id INTEGER REFERENCES prod.tm_position(id) ON DELETE SET NULL,
    organisasi VARCHAR(255),
    bidang_unit_kerja VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 9. tm_user_forgot_token table
CREATE TABLE IF NOT EXISTS prod.tm_user_forgot_token (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES prod.tm_user(id) ON DELETE CASCADE,
    token VARCHAR(255),
    expired_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- 10. tm_commodity table
CREATE TABLE IF NOT EXISTS prod.tm_commodity (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    parent_id INTEGER REFERENCES prod.tm_commodity(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    unit_id INTEGER REFERENCES prod.settings(id) ON DELETE SET NULL,
    unit_id_neraca INTEGER REFERENCES prod.settings(id) ON DELETE SET NULL,
    sequence INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    assets_relation_id NUMERIC
);

-- 11. tx_commodity_stock table
CREATE TABLE IF NOT EXISTS prod.tx_commodity_stock (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    commodity_id INTEGER REFERENCES prod.tm_commodity(id) ON DELETE CASCADE,
    city_id INTEGER REFERENCES prod.tm_city(id) ON DELETE CASCADE,
    stock NUMERIC DEFAULT 0,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 12. tx_commodity_stock_province table
CREATE TABLE IF NOT EXISTS prod.tx_commodity_stock_province (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    commodity_id INTEGER REFERENCES prod.tm_commodity(id) ON DELETE CASCADE,
    province_id INTEGER REFERENCES prod.tm_province(id) ON DELETE CASCADE,
    stock NUMERIC DEFAULT 0,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 13. tx_commodity_stock_national table
CREATE TABLE IF NOT EXISTS prod.tx_commodity_stock_national (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    commodity_id INTEGER REFERENCES prod.tm_commodity(id) ON DELETE CASCADE,
    stock NUMERIC DEFAULT 0,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 14. tx_commodity_price table
CREATE TABLE IF NOT EXISTS prod.tx_commodity_price (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    commodity_id INTEGER REFERENCES prod.tm_commodity(id) ON DELETE CASCADE,
    city_id INTEGER REFERENCES prod.tm_city(id) ON DELETE CASCADE,
    price NUMERIC DEFAULT 0,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 15. tx_commodity_price_province table
CREATE TABLE IF NOT EXISTS prod.tx_commodity_price_province (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    commodity_id INTEGER REFERENCES prod.tm_commodity(id) ON DELETE CASCADE,
    province_id INTEGER REFERENCES prod.tm_province(id) ON DELETE CASCADE,
    price NUMERIC DEFAULT 0,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 16. tx_commodity_price_national table
CREATE TABLE IF NOT EXISTS prod.tx_commodity_price_national (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    commodity_id INTEGER REFERENCES prod.tm_commodity(id) ON DELETE CASCADE,
    price NUMERIC DEFAULT 0,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 17. tx_file_upload_history table
CREATE TABLE IF NOT EXISTS prod.tx_file_upload_history (
    id SERIAL PRIMARY KEY,
    file_name VARCHAR(255) NOT NULL,
    row_total INTEGER DEFAULT 0,
    status INTEGER DEFAULT 1,
    module_type VARCHAR(255) NOT NULL,
    errors TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 18. tm_jenis_informasi table
CREATE TABLE IF NOT EXISTS prod.tm_jenis_informasi (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    parent_id INTEGER REFERENCES prod.tm_jenis_informasi(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


-- ====================
-- SEED INITIAL DATA
-- ====================

-- Seed settings
INSERT INTO prod.settings (id, parent_id, name, value) VALUES
(1, NULL, 'BASE_URL', NULL),
(2, 1, 'http://localhost:8000', 'http://localhost:8000'),
(3, NULL, 'UNIT', NULL),
(4, 3, 'kg', 'Kilogram'),
(5, 3, 'ton', 'Ton');

-- Seed province
INSERT INTO prod.tm_province (id, name) VALUES
(1, 'Sulawesi Selatan');

-- Seed cities in South Sulawesi
INSERT INTO prod.tm_city (id, province_id, name, sequence) VALUES
(1, 1, 'Makassar', 1),
(2, 1, 'Gowa', 2),
(3, 1, 'Maros', 3),
(4, 1, 'Parepare', 4),
(5, 1, 'Bone', 5),
(6, 1, 'Bulukumba', 6),
(7, 1, 'Bantaeng', 7),
(8, 1, 'Jeneponto', 8),
(9, 1, 'Takalar', 9),
(10, 1, 'Sinjai', 10);

-- Seed positions
INSERT INTO prod.tm_position (id, name) VALUES
(1, 'Admin'),
(2, 'Operator');

-- Seed menus
INSERT INTO prod.tm_menu (id, name, url, icon, position) VALUES
(1, 'Dashboard', '/dashboard', 'dashboard', 1),
(2, 'Upload Log', '/upload-log', 'upload', 2),
(3, 'User Management', '/pengguna', 'users', 3),
(4, 'Position Management', '/jabatan', 'lock', 4);

-- Seed map_jabatan_menu (grant full permissions to Admin)
INSERT INTO prod.map_jabatan_menu (position_id, menu_id, read_status, create_status, update_status, delete_status) VALUES
(1, 1, 1, 1, 1, 1),
(1, 2, 1, 1, 1, 1),
(1, 3, 1, 1, 1, 1),
(1, 4, 1, 1, 1, 1),
(2, 1, 1, 0, 0, 0),
(2, 2, 1, 1, 0, 0);

-- Seed admin user
-- Password is 'admin' -> SHA256 is '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918'
INSERT INTO prod.tm_user (id, username, email, password) VALUES
(1, 'admin', 'admin@sigap.id', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918');

-- Seed profile for admin user
INSERT INTO prod.tm_profile (user_id, nama_depan, nama_belakang, jabatan, position_id, organisasi, bidang_unit_kerja) VALUES
(1, 'Super', 'Admin', 'Administrator', 1, 'Bank Indonesia', 'IT Division');

-- Seed commodities
INSERT INTO prod.tm_commodity (id, name, unit_id, unit_id_neraca, sequence) VALUES
(1, 'Beras', 4, 4, 1),
(2, 'Bawang Merah', 4, 4, 2),
(3, 'Cabai Rawit', 4, 4, 3),
(4, 'Minyak Goreng', 4, 4, 4),
(5, 'Gula Pasir', 4, 4, 5);

-- Adjust sequence values to handle serial pk auto-increment mismatch
SELECT setval('prod.settings_id_seq', (SELECT MAX(id) FROM prod.settings));
SELECT setval('prod.tm_province_id_seq', (SELECT MAX(id) FROM prod.tm_province));
SELECT setval('prod.tm_city_id_seq', (SELECT MAX(id) FROM prod.tm_city));
SELECT setval('prod.tm_position_id_seq', (SELECT MAX(id) FROM prod.tm_position));
SELECT setval('prod.tm_menu_id_seq', (SELECT MAX(id) FROM prod.tm_menu));
SELECT setval('prod.tm_user_id_seq', (SELECT MAX(id) FROM prod.tm_user));
SELECT setval('prod.tm_commodity_id_seq', (SELECT MAX(id) FROM prod.tm_commodity));
