-- Recreate SPBI core transaction tables matching Sentech Golang backend structures
DROP TABLE IF EXISTS dev.tx_commodity_price_national CASCADE;
DROP TABLE IF EXISTS dev.tx_commodity_price_province CASCADE;
DROP TABLE IF EXISTS dev.tx_commodity_price CASCADE;
DROP TABLE IF EXISTS dev.tx_commodity_stock_national CASCADE;
DROP TABLE IF EXISTS dev.tx_commodity_stock_province CASCADE;
DROP TABLE IF EXISTS dev.tx_commodity_stock CASCADE;

-- 1. Recreate tx_commodity_price
CREATE TABLE dev.tx_commodity_price (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    city_id INTEGER REFERENCES dev.tm_city(id) ON DELETE CASCADE,
    commodity_id INTEGER REFERENCES dev.tm_commodity(id) ON DELETE CASCADE,
    commodity_name VARCHAR(255),
    price NUMERIC DEFAULT 0,
    last_update TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 2. Recreate tx_commodity_price_province
CREATE TABLE dev.tx_commodity_price_province (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    province_id INTEGER REFERENCES dev.tm_province(id) ON DELETE CASCADE,
    commodity_id INTEGER REFERENCES dev.tm_commodity(id) ON DELETE CASCADE,
    commodity_name VARCHAR(255),
    price NUMERIC DEFAULT 0,
    last_update TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 3. Recreate tx_commodity_price_national
CREATE TABLE dev.tx_commodity_price_national (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    national_id INTEGER DEFAULT 1,
    commodity_id INTEGER REFERENCES dev.tm_commodity(id) ON DELETE CASCADE,
    commodity_name VARCHAR(255),
    price NUMERIC DEFAULT 0,
    last_update TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 4. Recreate tx_commodity_stock
CREATE TABLE dev.tx_commodity_stock (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    city_id INTEGER REFERENCES dev.tm_city(id) ON DELETE CASCADE,
    city_name VARCHAR(255),
    commodity_id INTEGER REFERENCES dev.tm_commodity(id) ON DELETE CASCADE,
    commodity_name VARCHAR(255),
    ketersediaan NUMERIC DEFAULT 0,
    kebutuhan NUMERIC DEFAULT 0,
    neraca NUMERIC DEFAULT 0,
    last_update TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 5. Recreate tx_commodity_stock_province
CREATE TABLE dev.tx_commodity_stock_province (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    province_id INTEGER REFERENCES dev.tm_province(id) ON DELETE CASCADE,
    province_name VARCHAR(255),
    commodity_id INTEGER REFERENCES dev.tm_commodity(id) ON DELETE CASCADE,
    commodity_name VARCHAR(255),
    ketersediaan NUMERIC DEFAULT 0,
    kebutuhan NUMERIC DEFAULT 0,
    neraca NUMERIC DEFAULT 0,
    last_update TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 6. Recreate tx_commodity_stock_national
CREATE TABLE dev.tx_commodity_stock_national (
    id SERIAL PRIMARY KEY,
    client_id INTEGER DEFAULT 1,
    national_id INTEGER DEFAULT 1,
    national_name VARCHAR(255),
    commodity_id INTEGER REFERENCES dev.tm_commodity(id) ON DELETE CASCADE,
    commodity_name VARCHAR(255),
    ketersediaan NUMERIC DEFAULT 0,
    kebutuhan NUMERIC DEFAULT 0,
    neraca NUMERIC DEFAULT 0,
    last_update TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);


-- ==========================================
-- DEFINE MISSING POSTGRES HELPER FUNCTIONS
-- ==========================================

-- 1. rupiah_format
CREATE OR REPLACE FUNCTION dev.rupiah_format(price numeric)
RETURNS text AS $$
BEGIN
    IF price IS NULL THEN
        RETURN 'Rp 0';
    END IF;
    RETURN 'Rp ' || trim(to_char(price, '999G999G999G999'));
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.rupiah_format(price bigint)
RETURNS text AS $$
BEGIN
    IF price IS NULL THEN
        RETURN 'Rp 0';
    END IF;
    RETURN 'Rp ' || trim(to_char(price, '999G999G999G999'));
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.rupiah_format(price integer)
RETURNS text AS $$
BEGIN
    IF price IS NULL THEN
        RETURN 'Rp 0';
    END IF;
    RETURN 'Rp ' || trim(to_char(price, '999G999G999G999'));
END;
$$ LANGUAGE plpgsql;

-- 2. thousand_format
CREATE OR REPLACE FUNCTION dev.thousand_format(val numeric)
RETURNS text AS $$
BEGIN
    IF val IS NULL THEN
        RETURN '0';
    END IF;
    RETURN trim(to_char(val, '999G999G999G999'));
END;
$$ LANGUAGE plpgsql;

-- 3. province_object
CREATE OR REPLACE FUNCTION dev.province_object(prov_id integer)
RETURNS json AS $$
DECLARE
    res json;
BEGIN
    SELECT json_build_object('id', id, 'name', name)
    INTO res
    FROM dev.tm_province
    WHERE id = prov_id;
    RETURN res;
END;
$$ LANGUAGE plpgsql;

-- 4. city_object
CREATE OR REPLACE FUNCTION dev.city_object(ct_id integer)
RETURNS json AS $$
DECLARE
    res json;
BEGIN
    SELECT json_build_object('id', id, 'name', name)
    INTO res
    FROM dev.tm_city
    WHERE id = ct_id;
    RETURN res;
END;
$$ LANGUAGE plpgsql;

-- 5. get_level_harga
CREATE OR REPLACE FUNCTION dev.get_level_harga(prov_id integer, comm_id integer, sel_date varchar)
RETURNS TABLE(commodity_price_id integer, client_id integer, id integer, commodity_id integer, price numeric, date date) AS $$
BEGIN
    RETURN QUERY
    SELECT p.id, p.client_id, p.city_id, p.commodity_id, p.price::numeric, p.last_update::date
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date = sel_date::date;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.get_level_harga(prov_id integer, comm_id integer, sel_date date)
RETURNS TABLE(commodity_price_id integer, client_id integer, id integer, commodity_id integer, price numeric, date date) AS $$
BEGIN
    RETURN QUERY
    SELECT p.id, p.client_id, p.city_id, p.commodity_id, p.price::numeric, p.last_update::date
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date = sel_date;
END;
$$ LANGUAGE plpgsql;

-- 6. get_level_harga_min / max / range
CREATE OR REPLACE FUNCTION dev.get_level_harga_min(prov_id integer, comm_id integer, sel_date varchar)
RETURNS numeric AS $$
DECLARE
    res numeric;
BEGIN
    SELECT COALESCE(MIN(p.price), 0)::numeric INTO res
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date = sel_date::date;
    RETURN res;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.get_level_harga_max(prov_id integer, comm_id integer, sel_date varchar)
RETURNS numeric AS $$
DECLARE
    res numeric;
BEGIN
    SELECT COALESCE(MAX(p.price), 0)::numeric INTO res
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date = sel_date::date;
    RETURN res;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.get_level_harga_price_range(prov_id integer, comm_id integer, sel_date varchar)
RETURNS numeric AS $$
DECLARE
    min_p numeric;
    max_p numeric;
BEGIN
    min_p := dev.get_level_harga_min(prov_id, comm_id, sel_date);
    max_p := dev.get_level_harga_max(prov_id, comm_id, sel_date);
    RETURN max_p - min_p;
END;
$$ LANGUAGE plpgsql;

-- 7. get_level_harga_province
CREATE OR REPLACE FUNCTION dev.get_level_harga_province(comm_id integer, sel_date varchar)
RETURNS TABLE(province_id integer, client_id integer, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT c.province_id, p.client_id, COALESCE(ROUND(AVG(p.price)), 0)::numeric
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND p.last_update::date = sel_date::date
    GROUP BY c.province_id, p.client_id;
END;
$$ LANGUAGE plpgsql;

-- 8. price_province_cr
CREATE OR REPLACE FUNCTION dev.price_province_cr(comm_id integer, prov_id integer, sel_date varchar)
RETURNS TABLE(commodity_price_province_id integer, province_id integer, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT 1 as commodity_price_province_id, prov_id, COALESCE(ROUND(AVG(p.price)), 0)::numeric
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date = sel_date::date;
END;
$$ LANGUAGE plpgsql;

-- 9. price_province_avg_child
CREATE OR REPLACE FUNCTION dev.price_province_avg_child(comm_id integer, prov_id integer, sel_date varchar)
RETURNS TABLE(province_id integer, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT prov_id, COALESCE(ROUND(AVG(p.price)), 0)::numeric
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date = sel_date::date;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.price_province_avg_child(comm_id integer, prov_id integer, sel_date date)
RETURNS TABLE(province_id integer, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT prov_id, COALESCE(ROUND(AVG(p.price)), 0)::numeric
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date = sel_date;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.price_province_avg_child(comm_id integer, prov_id integer)
RETURNS TABLE(province_id integer, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT prov_id, COALESCE(ROUND(AVG(p.price)), 0)::numeric
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.price_province_avg_child(comm_id integer, prov_id integer, start_date varchar, end_date varchar)
RETURNS TABLE(province_id integer, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT prov_id, COALESCE(ROUND(AVG(p.price)), 0)::numeric
    FROM dev.tx_commodity_price p
    JOIN dev.tm_city c ON c.id = p.city_id
    WHERE p.commodity_id = comm_id
      AND c.province_id = prov_id
      AND p.last_update::date BETWEEN start_date::date AND end_date::date;
END;
$$ LANGUAGE plpgsql;

-- 10. get_level_harga_avg_child / mins / max / range
CREATE OR REPLACE FUNCTION dev.get_level_harga_avg_child(prov_id integer, comm_id integer, sel_date varchar)
RETURNS TABLE(commodity_price_id integer, client_id integer, id integer, commodity_id integer, price numeric, date date) AS $$
BEGIN
    RETURN QUERY SELECT * FROM dev.get_level_harga(prov_id, comm_id, sel_date);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.get_level_harga_min_avg_child(prov_id integer, comm_id integer, sel_date varchar)
RETURNS numeric AS $$
BEGIN
    RETURN dev.get_level_harga_min(prov_id, comm_id, sel_date);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.get_level_harga_max_avg_child(prov_id integer, comm_id integer, sel_date varchar)
RETURNS numeric AS $$
BEGIN
    RETURN dev.get_level_harga_max(prov_id, comm_id, sel_date);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.get_level_harga_price_range_avg_child(prov_id integer, comm_id integer, sel_date varchar)
RETURNS numeric AS $$
BEGIN
    RETURN dev.get_level_harga_price_range(prov_id, comm_id, sel_date);
END;
$$ LANGUAGE plpgsql;

-- 11. get_level_stock_cr
CREATE OR REPLACE FUNCTION dev.get_level_stock_cr(prov_id integer, comm_id integer, start_date varchar, end_date varchar)
RETURNS TABLE(commodity_price_id integer, client_id integer, id integer, commodity_id integer, stock numeric, kebutuhan numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT s.id, s.client_id, s.city_id, s.commodity_id, COALESCE(s.ketersediaan, 0)::numeric, COALESCE(s.kebutuhan, 0)::numeric
    FROM dev.tx_commodity_stock s
    JOIN dev.tm_city c ON c.id = s.city_id
    WHERE s.commodity_id = comm_id
      AND c.province_id = prov_id
      AND s.last_update::date BETWEEN start_date::date AND end_date::date;
END;
$$ LANGUAGE plpgsql;

-- 12. get_level_stock_province_cr
CREATE OR REPLACE FUNCTION dev.get_level_stock_province_cr(comm_id integer, start_date varchar, end_date varchar)
RETURNS TABLE(province_id integer, client_id integer, stock numeric, kebutuhan numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT s.province_id, s.client_id, COALESCE(s.ketersediaan, 0)::numeric, COALESCE(s.kebutuhan, 0)::numeric
    FROM dev.tx_commodity_stock_province s
    WHERE s.commodity_id = comm_id
      AND s.last_update::date BETWEEN start_date::date AND end_date::date;
END;
$$ LANGUAGE plpgsql;

-- 13. neraca_defisit_cr / rentan / waspada / aman
CREATE OR REPLACE FUNCTION dev.neraca_defisit_cr(comm_id integer, prov_id integer, start_date varchar, end_date varchar)
RETURNS integer AS $$
DECLARE
    cnt integer;
BEGIN
    SELECT COUNT(*)::integer INTO cnt
    FROM dev.get_level_stock_cr(prov_id, comm_id, start_date, end_date) a
    WHERE (CASE WHEN a.kebutuhan = 0 THEN 100 ELSE (a.stock / a.kebutuhan)*100 END) < 0 OR a.stock < 0;
    RETURN COALESCE(cnt, 0);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.neraca_rentan_cr(comm_id integer, prov_id integer, start_date varchar, end_date varchar)
RETURNS integer AS $$
DECLARE
    cnt integer;
BEGIN
    SELECT COUNT(*)::integer INTO cnt
    FROM dev.get_level_stock_cr(prov_id, comm_id, start_date, end_date) a
    WHERE (CASE WHEN a.kebutuhan = 0 THEN 100 ELSE (a.stock / a.kebutuhan)*100 END) BETWEEN 0 AND 46;
    RETURN COALESCE(cnt, 0);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.neraca_waspada_cr(comm_id integer, prov_id integer, start_date varchar, end_date varchar)
RETURNS integer AS $$
DECLARE
    cnt integer;
BEGIN
    SELECT COUNT(*)::integer INTO cnt
    FROM dev.get_level_stock_cr(prov_id, comm_id, start_date, end_date) a
    WHERE (CASE WHEN a.kebutuhan = 0 THEN 100 ELSE (a.stock / a.kebutuhan)*100 END) BETWEEN 47 AND 80;
    RETURN COALESCE(cnt, 0);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION dev.neraca_aman_cr(comm_id integer, prov_id integer, start_date varchar, end_date varchar)
RETURNS integer AS $$
DECLARE
    cnt integer;
BEGIN
    SELECT COUNT(*)::integer INTO cnt
    FROM dev.get_level_stock_cr(prov_id, comm_id, start_date, end_date) a
    WHERE (CASE WHEN a.kebutuhan = 0 THEN 100 ELSE (a.stock / a.kebutuhan)*100 END) >= 81;
    RETURN COALESCE(cnt, 0);
END;
$$ LANGUAGE plpgsql;


-- ===========================================
-- SEED PREMIUM LIVE DATA FOR GIS DASHBOARD
-- ===========================================

-- Seed tx_commodity_price (Last 10 days for trends, plus today's main prices)
-- Today is 2026-05-19
INSERT INTO dev.tx_commodity_price (client_id, city_id, commodity_id, commodity_name, price, last_update)
SELECT 
    1, 
    c.id, 
    comm.id, 
    comm.name,
    -- Realistic prices: Beras=14000, Bawang=32000, Cabai=45000, Minyak=18000, Gula=16000 with a small variation per city and date offset
    CASE 
        WHEN comm.id = 1 THEN 13500 + (c.id * 120) + (d.day_offset * 50)
        WHEN comm.id = 2 THEN 30000 + (c.id * 350) - (d.day_offset * 100)
        WHEN comm.id = 3 THEN 42000 + (c.id * 500) + (d.day_offset * 150)
        WHEN comm.id = 4 THEN 17000 + (c.id * 150) + (d.day_offset * 30)
        ELSE 15000 + (c.id * 100) - (d.day_offset * 20)
    END,
    ('2026-05-19'::date - (d.day_offset || ' days')::interval)::timestamp
FROM dev.tm_city c
CROSS JOIN dev.tm_commodity comm
CROSS JOIN (SELECT generate_series(0, 9) AS day_offset) d;

-- Seed tx_commodity_price_province
INSERT INTO dev.tx_commodity_price_province (client_id, province_id, commodity_id, commodity_name, price, last_update)
SELECT 
    1, 
    1, 
    comm.id, 
    comm.name,
    ROUND(AVG(p.price)),
    p.last_update
FROM dev.tx_commodity_price p
JOIN dev.tm_city c ON c.id = p.city_id
JOIN dev.tm_commodity comm ON comm.id = p.commodity_id
GROUP BY comm.id, comm.name, p.last_update;

-- Seed tx_commodity_price_national
INSERT INTO dev.tx_commodity_price_national (client_id, national_id, commodity_id, commodity_name, price, last_update)
SELECT 
    1, 
    1, 
    comm.id, 
    comm.name,
    ROUND(AVG(p.price) * 1.05), -- National average slightly higher
    p.last_update
FROM dev.tx_commodity_price p
JOIN dev.tm_city c ON c.id = p.city_id
JOIN dev.tm_commodity comm ON comm.id = p.commodity_id
GROUP BY comm.id, comm.name, p.last_update;

-- Seed tx_commodity_stock
-- Needs to generate variations so the map has green, yellow, orange, and red cities
-- Green (Aman) = Stock ratio >= 81%
-- Yellow (Waspada) = Stock ratio 47% - 80%
-- Orange (Rentan) = Stock ratio 0% - 46%
-- Red (Defisit) = Stock ratio < 0 or Stock < 0 (We will trigger some defisits by setting negative stock/neraca)
INSERT INTO dev.tx_commodity_stock (client_id, city_id, city_name, commodity_id, commodity_name, ketersediaan, kebutuhan, neraca, last_update)
SELECT 
    1, 
    c.id, 
    c.name,
    comm.id, 
    comm.name,
    -- Ketersediaan
    CASE 
        WHEN c.id IN (1, 3, 5) THEN 1200 + (c.id * 100) -- High stock (Aman)
        WHEN c.id IN (2, 4, 6) THEN 750 + (c.id * 50)  -- Medium stock (Waspada)
        WHEN c.id IN (7, 8) THEN 350 + (c.id * 20)     -- Low stock (Rentan)
        ELSE 100                                       -- Critically low (Defisit)
    END,
    -- Kebutuhan
    1000 + (c.id * 40),
    -- Neraca (Ketersediaan - Kebutuhan)
    CASE 
        WHEN c.id IN (1, 3, 5) THEN (1200 + (c.id * 100)) - (1000 + (c.id * 40))
        WHEN c.id IN (2, 4, 6) THEN (750 + (c.id * 50)) - (1000 + (c.id * 40))
        WHEN c.id IN (7, 8) THEN (350 + (c.id * 20)) - (1000 + (c.id * 40))
        ELSE -800
    END,
    ('2026-05-19'::date - (d.day_offset || ' days')::interval)::timestamp
FROM dev.tm_city c
CROSS JOIN dev.tm_commodity comm
CROSS JOIN (SELECT generate_series(0, 9) AS day_offset) d;

-- Seed tx_commodity_stock_province
INSERT INTO dev.tx_commodity_stock_province (client_id, province_id, province_name, commodity_id, commodity_name, ketersediaan, kebutuhan, neraca, last_update)
SELECT 
    1, 
    1, 
    'Sulawesi Selatan',
    comm.id, 
    comm.name,
    SUM(s.ketersediaan),
    SUM(s.kebutuhan),
    SUM(s.neraca),
    s.last_update
FROM dev.tx_commodity_stock s
JOIN dev.tm_city c ON c.id = s.city_id
JOIN dev.tm_commodity comm ON comm.id = s.commodity_id
GROUP BY comm.id, comm.name, s.last_update;

-- Seed tx_commodity_stock_national
INSERT INTO dev.tx_commodity_stock_national (client_id, national_id, national_name, commodity_id, commodity_name, ketersediaan, kebutuhan, neraca, last_update)
SELECT 
    1, 
    1, 
    'Nasional',
    comm.id, 
    comm.name,
    SUM(s.ketersediaan) * 15,
    SUM(s.kebutuhan) * 15,
    SUM(s.neraca) * 15,
    s.last_update
FROM dev.tx_commodity_stock s
JOIN dev.tm_city c ON c.id = s.city_id
JOIN dev.tm_commodity comm ON comm.id = s.commodity_id
GROUP BY comm.id, comm.name, s.last_update;
