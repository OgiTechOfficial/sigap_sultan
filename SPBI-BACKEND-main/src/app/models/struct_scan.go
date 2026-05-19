package models

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (c *AssetsResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("AssetsResponse unsupported type")
}

func (c *CommodityPriceTableResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("CommodityPriceTableResponse unsupported type")
}

func (c *StockCommodityByLocationResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

func (c *StockLocationDiffResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

func (c *PrivilegesResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

func (c *PermissionResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

func (c *PriceTableHargaLevelMapData) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

//func (c *PriceTableHargaLevelList) Scan(v interface{}) error {
//	if v == nil {
//		return nil
//	}
//
//	switch tv := v.(type) {
//	case []byte:
//		return json.Unmarshal(tv, &c)
//	}
//
//	return errors.New("unsupported type")
//}

func (c *PriceDiffResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

func (rs *TmCity) ScanRow(rows pgx.Rows) error {
	return rows.Scan(&rs.Id.Id, &rs.ProvinceId, &rs.Name, &rs.CreatedAt, &rs.UpdatedAt, &rs.DeletedAt, &rs.Sequence, &rs.AssetsRelationId)
}

func (c *CityResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

func (c *CityLevelHargaResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("CityLevelHargaResponse unsupported type")
}

func (c *CommodityResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("CommodityResponse unsupported type")
}

func (c *CommodityWithPriceListResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("CommodityWithPriceListResponse unsupported type")
}

func (c *CityWithPriceListResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("CityWithPriceListResponse unsupported type")
}

func (rs *TmCommodity) ScanRow(rows pgx.Rows) error {
	return rows.Scan(&rs.Id.Id, &rs.ClientId, &rs.ParentId, &rs.Code, &rs.Name, &rs.CreatedAt, &rs.UpdatedAt, &rs.DeletedAt)
}

func (c *DetailJenisInformasiResponse) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	switch tv := v.(type) {
	case []byte:
		return json.Unmarshal(tv, &c)
	}

	return errors.New("unsupported type")
}

func (rs *TmMenu) ScanRow(rows pgx.Rows) error {
	return rows.Scan(
		&rs.Id.Id,
		&rs.ClientId,
		&rs.Name,
		&rs.Position,
		&rs.CreatedAt,
		&rs.UpdatedAt,
		&rs.DeletedAt,
	)
}

func (rs *TmMenuResponse) ScanRow(rows pgx.Rows) error {
	return rows.Scan(
		&rs.Id.Id,
		&rs.ClientId,
		&rs.Name,
		&rs.Position,
		&rs.CreatedAt,
		&rs.UpdatedAt,
		&rs.DeletedAt,
	)
}
