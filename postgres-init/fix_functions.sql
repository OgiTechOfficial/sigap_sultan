DROP FUNCTION IF EXISTS prod.get_harga_by_city_id_avg_child(integer, integer, varchar, varchar) CASCADE;
DROP FUNCTION IF EXISTS prod.get_harga_by_city_id(integer, integer, varchar, varchar) CASCADE;

CREATE OR REPLACE FUNCTION prod.get_harga_by_city_id(c_id integer, comm_id integer, start_date varchar, end_date varchar)
RETURNS TABLE(last_update timestamptz, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT p.last_update, p.price::numeric
    FROM prod.tx_commodity_price p
    WHERE p.city_id = c_id
      AND p.commodity_id = comm_id
      AND p.last_update::date BETWEEN start_date::date AND end_date::date;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prod.get_harga_by_city_id_avg_child(c_id integer, comm_id integer, start_date varchar, end_date varchar)
RETURNS TABLE(last_update timestamptz, price numeric) AS $$
BEGIN
    RETURN QUERY
    SELECT * FROM prod.get_harga_by_city_id(c_id, comm_id, start_date, end_date);
END;
$$ LANGUAGE plpgsql;
