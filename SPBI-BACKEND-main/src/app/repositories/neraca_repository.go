package repositories

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"sigap-sultan-be/src/app/repositories/queries/neraca"
	"sigap-sultan-be/src/app/repositories/queries/price"
	"sigap-sultan-be/src/common"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NeracaRepository struct {
	Db              *pgxpool.Pool
	PriceRepository *PriceRepository
}

func NewNeracaRepository(db *pgxpool.Pool, priceRepository *PriceRepository) *NeracaRepository {
	return &NeracaRepository{Db: db, PriceRepository: priceRepository}
}

func (repo *NeracaRepository) SaveNeracaCity(prices []*models.NeracaCity) {
	var buffer bytes.Buffer
	query := queries.NeracaTxCommodityStockCityInsert

	for idx, price := range prices {
		if idx == 0 {
			buffer.WriteString(
				"\n( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.CityId)) + "'," +
					"'" + price.CityName + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Ketersediaan + "'," +
					"'" + price.Kebutuhan + "'," +
					"'" + price.Neraca + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n",
			)
		} else {
			buffer.WriteString(
				",( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.CityId)) + "'," +
					"'" + price.CityName + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Ketersediaan + "'," +
					"'" + price.Kebutuhan + "'," +
					"'" + price.Neraca + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n ",
			)
		}
	}

	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		log.Error("QueryErrors:", err)
	}

	_, err = tx.Exec(context.Background(), query+buffer.String())
	if err != nil {
		log.Fatalf("tx.Exec failed: %v\n", err)
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		log.Error("QueryErrors:", err)
	}
}

func (repo *NeracaRepository) SaveNeracaProvince(prices []*models.NeracaProvince) {
	var buffer bytes.Buffer
	query := queries.NeracaTxCommodityStockProvinceInsert

	for idx, price := range prices {
		if idx == 0 {
			buffer.WriteString(
				"\n( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.ProvinceId)) + "'," +
					"'" + price.ProvinceName + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Ketersediaan + "'," +
					"'" + price.Kebutuhan + "'," +
					"'" + price.Neraca + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n",
			)
		} else {
			buffer.WriteString(
				",( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.ProvinceId)) + "'," +
					"'" + price.ProvinceName + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Ketersediaan + "'," +
					"'" + price.Kebutuhan + "'," +
					"'" + price.Neraca + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n ",
			)
		}
	}

	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		log.Error("QueryErrors:", err)
	}

	_, err = tx.Exec(context.Background(), query+buffer.String())
	if err != nil {
		log.Fatalf("tx.Exec failed: %v\n", err)
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		log.Error("QueryErrors:", err)
	}
}

func (repo *NeracaRepository) SaveNeracaNational(prices []*models.NeracaNational) {
	var buffer bytes.Buffer
	query := queries.NeracaTxCommodityStockNationalInsert

	for idx, price := range prices {
		if idx == 0 {
			buffer.WriteString(
				"\n( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.NationalId)) + "'," +
					"'" + price.NatinoalName + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Ketersediaan + "'," +
					"'" + price.Kebutuhan + "'," +
					"'" + price.Neraca + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n",
			)
		} else {
			buffer.WriteString(
				",( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.NationalId)) + "'," +
					"'" + price.NatinoalName + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Ketersediaan + "'," +
					"'" + price.Kebutuhan + "'," +
					"'" + price.Neraca + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n ",
			)
		}
	}

	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		log.Error("QueryErrors:", err)
	}

	_, err = tx.Exec(context.Background(), query+buffer.String())
	if err != nil {
		log.Fatalf("tx.Exec failed: %v\n", err)
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		log.Error("QueryErrors:", err)
	}
}

func (repo *NeracaRepository) GetStockAkhir(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	var results models.NeracaStokAkhirMapResponse
	err = repo.Db.QueryRow(
		context.Background(),
		neraca.NeracaStokAkhirMap,
		pgx.NamedArgs{
			"provinceId":  73,
			"commodityId": params.CommodityId,
			"startDate":   startDate,
			"endDate":     endDate,
		},
	).Scan(
		&results.Unit,
		&results.Commodity,
		&results.ProvinceStock,
		&results.CityStock,
		&results.NeracaTierLevel,
		&results.StockTierCode,
	)

	var summary models.NeracaStokAkhirSummary
	err = repo.Db.QueryRow(
		context.Background(),
		neraca.NeraacStokAkhirSummary,
		pgx.NamedArgs{
			"provinceId":  73,
			"commodityId": params.CommodityId,
			"startDate":   startDate,
			"endDate":     endDate,
		},
	).Scan(
		&summary,
	)

	results.Summary = &summary

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	//if results.ProvinceStock != nil {
	//	if results.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := results.ProvinceStock.Province.Assets.AssetsLocation + "/" + results.ProvinceStock.Province.Assets.AssetsName
	//		results.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//if results.CityStock != nil {
	//	for _, row := range *results.CityStock {
	//		if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//			//baseUrl := os.Getenv("BASE_URL")
	//			baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//			assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//			row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//		}
	//	}
	//}

	return &results, nil
}

func (repo *NeracaRepository) GetStockAkhirListByCommodity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var sortBy string
	var sortByRank string
	var result models.NeracaStokAkhirListResponse
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	if params.PaginationParams.SortBy != "" {
		split := strings.Split(params.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if sortByRank != "" {
			if sortByRank == "asc" {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaStokAkhirListByCommodityAsc,
					pgx.NamedArgs{
						"provinceId":  73,
						"commodityId": params.CommodityId,
						"startDate":   startDate,
						"endDate":     endDate,
						"page":        params.PaginationParams.Page,
						"limit":       params.PaginationParams.Limit,
						"sortBy":      sortBy,
						"sortByRank":  sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			} else {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaStokAkhirListByCommodityDesc,
					pgx.NamedArgs{
						"provinceId":  73,
						"commodityId": params.CommodityId,
						"startDate":   startDate,
						"endDate":     endDate,
						"page":        params.PaginationParams.Page,
						"limit":       params.PaginationParams.Limit,
						"sortBy":      sortBy,
						"sortByRank":  sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaStokAkhirListByCommodityAsc,
				pgx.NamedArgs{
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   startDate,
					"endDate":     endDate,
					"page":        params.PaginationParams.Page,
					"limit":       params.PaginationParams.Limit,
					"sortBy":      sortBy,
					"sortByRank":  sortByRank,
				},
			).Scan(
				&result.Unit,
				&result.Commodities,
				&result.ProvinceStock,
				&result.CityStock,
				&result.NeracaTierLevel,
			)
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirListByCommodity,
			pgx.NamedArgs{
				"provinceId":  73,
				"commodityId": params.CommodityId,
				"startDate":   startDate,
				"endDate":     endDate,
				"page":        params.PaginationParams.Page,
				"limit":       params.PaginationParams.Limit,
			},
		).Scan(
			&result.Unit,
			&result.Commodities,
			&result.ProvinceStock,
			&result.CityStock,
			&result.NeracaTierLevel,
		)
	}
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if result.ProvinceStock == nil || result.CityStock == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.ProvinceStock != nil {
	//	if result.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.ProvinceStock.Province.Assets.AssetsLocation + "/" + result.ProvinceStock.Province.Assets.AssetsName
	//		result.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *NeracaRepository) GetStockAkhirByCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaStokAkhirCommodityHistoryResponse
	var rows pgx.Rows
	var err error
	var query string
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))
	//var realLastUpdate interface{}
	//var dateFormat = "2006-01-02"
	//var newDate *common.GenerateNewDate

	//var startDate time.Time
	//startDate, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var endDate time.Time
	//endDate, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.SelectedDate = params.EndDate
	//realLastUpdate, err = repo.GetLastUpdate(params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//lastUpdate := realLastUpdate.(*string)
	//newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.StartDate = newDate.StartDate
	//params.EndDate = *lastUpdate

	switch params.Status {
	case "defisit":
		query = neraca.NeracaStokAkhirCommodityHistoryDefisit
	case "rentan":
		query = neraca.NeracaStokAkhirCommodityHistoryRentan
	case "waspada":
		query = neraca.NeracaStokAkhirCommodityHistoryWaspada
	default:
		query = neraca.NeracaStokAkhirCommodityHistoryAman
	}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
			"provinceId":  params.ProvinceId,
			"startDate":   startDate,
			"endDate":     endDate,
			"status":      params.Status,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Unit,
			&result.Commodities,
			&result.CityStock,
			&result.NeracaTierLevel,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetStockAkhirByCommodityCityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaStokAkhirCityHistoryResponse
	var rows pgx.Rows
	var err error
	var realLastUpdate interface{}
	var dateFormat = "2006-01-02"
	var newDate *common.GenerateNewDate

	var startDate time.Time
	startDate, err = time.Parse(dateFormat, params.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate time.Time
	endDate, err = time.Parse(dateFormat, params.EndDate)
	if err != nil {
		return nil, err
	}

	params.SelectedDate = params.EndDate
	realLastUpdate, err = repo.GetLastUpdate(params)
	if err != nil {
		return nil, err
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	params.StartDate = newDate.StartDate
	params.EndDate = *lastUpdate

	rows, err = repo.Db.Query(
		context.Background(),
		neraca.NeracaStokAkhirCityHistory,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
			"provinceId":  params.ProvinceId,
			"cityId":      params.CityId,
			"startDate":   params.StartDate,
			"endDate":     params.EndDate,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Unit,
			&result.Commodities,
			&result.City,
			&result.NeracaListResponse,
			&result.NeracaTierLevel,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetKetersediaanByCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaKetersediaanCommodityHistoryResponse
	//Unit            string                                   `json:"unit"`
	//Commodities     *CommodityResponse                       `json:"commodity"`
	//CityStock       *[]NeracaCitKetersediaanWithDiffResponse `json:"ketersediaanDiff"`
	//NeracaTierLevel *NeracaTierLevel                         `json:"stockTier"`

	//var result models.NeracaStokAkhirCommodityHistoryResponse
	//Unit            string                     `json:"unit"`
	//Commodities     *CommodityResponse         `json:"commodity"`
	//CityStock       *[]NeracaCityStockResponse `json:"cityStock"`
	//NeracaTierLevel *NeracaTierLevel           `json:"stockTier"`

	var rows pgx.Rows
	var err error
	var query string
	//var realLastUpdate interface{}
	//var dateFormat = "2006-01-02"
	//var newDate *common.GenerateNewDate

	//var startDate time.Time
	//startDate, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var endDate time.Time
	//endDate, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.SelectedDate = params.EndDate
	//realLastUpdate, err = repo.GetLastUpdate(params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//lastUpdate := realLastUpdate.(*string)
	//newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.StartDate = newDate.StartDate
	//params.EndDate = *lastUpdate

	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	switch params.Status {
	case "menurun":
		query = neraca.NeracaKetersediaanCommodityHistoryMenurun
	case "stabil":
		query = neraca.NeracaKetersediaanCommodityHistoryStabil
	case "meningkat":
		query = neraca.NeracaKetersediaanCommodityHistoryMeningkat
	default:
		return nil, errors.New("Status yang tersedia: menurun, stabil, meningkat. Status yang dikirim: " + params.Status)
	}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
			"provinceId":  params.ProvinceId,
			"startDate":   startDate,
			"endDate":     endDate,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Unit,
			&result.Commodities,
			&result.CityStock,
			&result.NeracaTierLevel,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetKetersediaanByCommodityCityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaKetersediaanCommodityCityHistoryResponse
	var rows pgx.Rows
	var err error
	//var realLastUpdate interface{}
	//var dateFormat = "2006-01-02"
	//var newDate *common.GenerateNewDate
	//
	//var startDate time.Time
	//startDate, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var endDate time.Time
	//endDate, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.SelectedDate = params.EndDate
	//realLastUpdate, err = repo.GetLastUpdate(params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//lastUpdate := realLastUpdate.(*string)
	//newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.StartDate = newDate.StartDate
	//params.EndDate = *lastUpdate

	err = common.NeracaPeriodDateGenerate(&params.StartDate, &params.EndDate)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	rows, err = repo.Db.Query(
		context.Background(),
		neraca.NeracaKetersediaanCommodityCityHistory,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
			"provinceId":  params.ProvinceId,
			"cityId":      params.CityId,
			"startDate":   params.StartDate,
			"endDate":     params.EndDate,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Unit,
			&result.Commodities,
			//&result.CityStock,
			&result.NeracaListResponse,
			&result.NeracaTierLevel,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetKebutuhanByCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaKebutuhanCommodityHistoryResponse
	var rows pgx.Rows
	var err error
	var query string
	//var realLastUpdate interface{}
	//var dateFormat = "2006-01-02"
	//var newDate *common.GenerateNewDate

	//var startDate time.Time
	//startDate, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var endDate time.Time
	//endDate, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.SelectedDate = params.EndDate
	//realLastUpdate, err = repo.GetLastUpdate(params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//lastUpdate := realLastUpdate.(*string)
	//newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.StartDate = newDate.StartDate
	//params.EndDate = *lastUpdate

	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	switch params.Status {
	case "menurun":
		query = neraca.NeracaKebutuhanCommodityHistoryMenurun
	case "stabil":
		query = neraca.NeracaKebutuhanCommodityHistoryStabil
	case "meningkat":
		query = neraca.NeracaKebutuhanCommodityHistoryMeningkat
	default:
		return nil, errors.New("Status yang tersedia: menurun, stabil, meningkat. Status yang dikirim: " + params.Status)
	}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
			"provinceId":  params.ProvinceId,
			"startDate":   startDate,
			"endDate":     endDate,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Unit,
			&result.Commodities,
			&result.CityStock,
			&result.NeracaTierLevel,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetKebutuhanByCommodityCityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaKetersediaanCommodityCityHistoryResponse
	var rows pgx.Rows
	var err error
	//var realLastUpdate interface{}
	//var dateFormat = "2006-01-02"
	//var newDate *common.GenerateNewDate
	//
	//var startDate time.Time
	//startDate, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var endDate time.Time
	//endDate, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.SelectedDate = params.EndDate
	//realLastUpdate, err = repo.GetLastUpdate(params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//lastUpdate := realLastUpdate.(*string)
	//newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.StartDate = newDate.StartDate
	//params.EndDate = *lastUpdate

	err = common.NeracaPeriodDateGenerate(&params.StartDate, &params.EndDate)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	rows, err = repo.Db.Query(
		context.Background(),
		neraca.NeracaKebutuhanCommodityCityHistory,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
			"provinceId":  params.ProvinceId,
			"cityId":      params.CityId,
			"startDate":   params.StartDate,
			"endDate":     params.EndDate,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Unit,
			&result.Commodities,
			&result.NeracaListResponse,
			&result.NeracaTierLevel,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetLastUpdate(queryParams domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result string
	var err error

	if queryParams.CityId == "73" {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaGetLatestDateAvailProvince,
			pgx.NamedArgs{
				"provinceId":   "73",
				"cityId":       queryParams.CityId,
				"commodityId":  queryParams.CommodityId,
				"selectedDate": queryParams.SelectedDate,
			},
		).Scan(&result)
		if err != nil {
			log.Error("QueryErrors GetLastUpdate:", err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaGetLatestDateAvail,
			pgx.NamedArgs{
				"provinceId":   "73",
				"cityId":       queryParams.CityId,
				"commodityId":  queryParams.CommodityId,
				"selectedDate": queryParams.SelectedDate,
			},
		).Scan(&result)
		if err != nil {
			log.Error("QueryErrors GetLastUpdate:", err)
			return nil, err
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetKetersediaanMap(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	var results models.NeracaKetersediaanMapResponse
	err = repo.Db.QueryRow(
		context.Background(),
		neraca.NeracaKetersediaanMap,
		pgx.NamedArgs{
			"provinceId":  73,
			"commodityId": params.CommodityId,
			"startDate":   startDate,
			"endDate":     endDate,
		},
	).Scan(
		&results.Unit,
		&results.Commodities,
		&results.Summary,
		&results.ProvinceStock,
		&results.CityStock,
		&results.NeracaTierLevel,
		&results.StockTierCode,
	)

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	//if results.ProvinceStock != nil {
	//	if results.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := results.ProvinceStock.Province.Assets.AssetsLocation + "/" + results.ProvinceStock.Province.Assets.AssetsName
	//		results.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *results.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &results, nil
}

func (repo *NeracaRepository) GetKetersediaanListByCommodity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var sortBy string
	var sortByRank string
	var result models.NeracaKetersediaanListResponse
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	if params.PaginationParams.SortBy != "" {
		split := strings.Split(params.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if sortByRank != "" {
			if sortByRank == "asc" {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaKetersediaanListByCommodityAsc,
					pgx.NamedArgs{
						"provinceId":  73,
						"commodityId": params.CommodityId,
						"startDate":   startDate,
						"endDate":     endDate,
						"page":        params.PaginationParams.Page,
						"limit":       params.PaginationParams.Limit,
						"sortBy":      sortBy,
						"sortByRank":  sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
					&result.StockTierCode,
				)
			} else {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaKetersediaanListByCommodityDesc,
					pgx.NamedArgs{
						"provinceId":  73,
						"commodityId": params.CommodityId,
						"startDate":   startDate,
						"endDate":     endDate,
						"page":        params.PaginationParams.Page,
						"limit":       params.PaginationParams.Limit,
						"sortBy":      sortBy,
						"sortByRank":  sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
					&result.StockTierCode,
				)
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaKetersediaanListByCommodityAsc,
				pgx.NamedArgs{
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   startDate,
					"endDate":     endDate,
					"page":        params.PaginationParams.Page,
					"limit":       params.PaginationParams.Limit,
					"sortBy":      sortBy,
					"sortByRank":  sortByRank,
				},
			).Scan(
				&result.Unit,
				&result.Commodities,
				&result.ProvinceStock,
				&result.CityStock,
				&result.NeracaTierLevel,
				&result.StockTierCode,
			)
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaKetersediaanListByCommodity,
			pgx.NamedArgs{
				"provinceId":  73,
				"commodityId": params.CommodityId,
				"startDate":   startDate,
				"endDate":     endDate,
				"page":        params.PaginationParams.Page,
				"limit":       params.PaginationParams.Limit,
			},
		).Scan(
			&result.Unit,
			&result.Commodities,
			&result.ProvinceStock,
			&result.CityStock,
			&result.NeracaTierLevel,
			&result.StockTierCode,
		)
	}
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if result.ProvinceStock == nil || result.CityStock == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.ProvinceStock != nil {
	//	if result.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.ProvinceStock.Province.Assets.AssetsLocation + "/" + result.ProvinceStock.Province.Assets.AssetsName
	//		result.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *NeracaRepository) GetKebutuhanMap(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	var results models.NeracaKetersediaanMapResponse
	err = repo.Db.QueryRow(
		context.Background(),
		neraca.NeracaKebutuhanMap,
		pgx.NamedArgs{
			"provinceId":  73,
			"commodityId": params.CommodityId,
			"startDate":   startDate,
			"endDate":     endDate,
		},
	).Scan(
		&results.Unit,
		&results.Commodities,
		&results.Summary,
		&results.ProvinceStock,
		&results.CityStock,
		&results.NeracaTierLevel,
		&results.StockTierCode,
	)

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	//if results.ProvinceStock != nil {
	//	if results.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := results.ProvinceStock.Province.Assets.AssetsLocation + "/" + results.ProvinceStock.Province.Assets.AssetsName
	//		results.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *results.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &results, nil
}

func (repo *NeracaRepository) GetKebutuhanListByCommodity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var sortBy string
	var sortByRank string
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	var result models.NeracaKetersediaanListResponse

	if params.PaginationParams.SortBy != "" {
		split := strings.Split(params.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if sortByRank != "" {
			if sortByRank == "asc" {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaKebutuhanListByCommodityAsc,
					pgx.NamedArgs{
						"provinceId":  73,
						"commodityId": params.CommodityId,
						"startDate":   startDate,
						"endDate":     endDate,
						"page":        params.PaginationParams.Page,
						"limit":       params.PaginationParams.Limit,
						"sortBy":      sortBy,
						"sortByRank":  sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
					&result.StockTierCode,
				)
			} else {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaKebutuhanListByCommodityDesc,
					pgx.NamedArgs{
						"provinceId":  73,
						"commodityId": params.CommodityId,
						"startDate":   startDate,
						"endDate":     endDate,
						"page":        params.PaginationParams.Page,
						"limit":       params.PaginationParams.Limit,
						"sortBy":      sortBy,
						"sortByRank":  sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
					&result.StockTierCode,
				)
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaKebutuhanListByCommodityAsc,
				pgx.NamedArgs{
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   startDate,
					"endDate":     endDate,
					"page":        params.PaginationParams.Page,
					"limit":       params.PaginationParams.Limit,
					"sortBy":      sortBy,
					"sortByRank":  sortByRank,
				},
			).Scan(
				&result.Unit,
				&result.Commodities,
				&result.ProvinceStock,
				&result.CityStock,
				&result.NeracaTierLevel,
				&result.StockTierCode,
			)
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaKebutuhanListByCommodity,
			pgx.NamedArgs{
				"provinceId":  73,
				"commodityId": params.CommodityId,
				"startDate":   startDate,
				"endDate":     endDate,
				"page":        params.PaginationParams.Page,
				"limit":       params.PaginationParams.Limit,
			},
		).Scan(
			&result.Unit,
			&result.Commodities,
			&result.ProvinceStock,
			&result.CityStock,
			&result.NeracaTierLevel,
			&result.StockTierCode,
		)
	}
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if result.ProvinceStock == nil || result.CityStock == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.ProvinceStock != nil {
	//	if result.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.ProvinceStock.Province.Assets.AssetsLocation + "/" + result.ProvinceStock.Province.Assets.AssetsName
	//		result.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *NeracaRepository) GetStockAkhirByCityAndCommodityId(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaStokAkhirCityCommodityResponse
	var err error

	if params.CityId == "73" {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirByCommodityProvinceHistory,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
				"provinceId":  params.ProvinceId,
				"cityId":      params.CityId,
			},
		).Scan(
			&result.Unit,
			&result.UnitDiff,
			&result.Commodities,
			&result.City,
			//&result.StockDiff,
			&result.Stock,
			&result.NeracaTierLevel,
		)
		if err != nil {
			log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirByCommodityCityHistory,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
				"provinceId":  params.ProvinceId,
				"cityId":      params.CityId,
			},
		).Scan(
			&result.Unit,
			&result.UnitDiff,
			&result.Commodities,
			&result.City,
			//&result.StockDiff,
			&result.Stock,
			&result.NeracaTierLevel,
		)
		if err != nil {
			log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
			return nil, err
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) GetStockAkhirByCityAndCommodityChart(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaStokAkhirCityCommodityChartResponse
	var err error
	//var realLastUpdate interface{}
	//var dateFormat = "2006-01-02"
	//var newDate *common.GenerateNewDate

	//var startDate time.Time
	//startDate, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var endDate time.Time
	//endDate, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}

	//params.SelectedDate = params.EndDate
	//realLastUpdate, err = repo.GetLastUpdate(params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//lastUpdate := realLastUpdate.(*string)
	//newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	//if err != nil {
	//	return nil, err
	//}

	//params.StartDate = newDate.StartDate
	//params.EndDate = *lastUpdate

	err = common.NeracaPeriodDateGenerate(&params.StartDate, &params.EndDate)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	if params.CityId == "73" {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirByCommodityCityHistoryChart,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
				"provinceId":  params.ProvinceId,
				"cityId":      7371,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&result.Unit,
			&result.UnitDiff,
			&result.Commodities,
			&result.City,
			//&result.StockDiff,
			&result.Stock,
			&result.NeracaTierLevel,
		)
		if err != nil {
			log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirByCommodityCityHistoryChart,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
				"provinceId":  params.ProvinceId,
				"cityId":      params.CityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&result.Unit,
			&result.UnitDiff,
			&result.Commodities,
			&result.City,
			//&result.StockDiff,
			&result.Stock,
			&result.NeracaTierLevel,
		)
		if err != nil {
			log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
			return nil, err
		}
	}

	return &result, nil
}

func (repo *NeracaRepository) CompareWithPriceCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result models.NeracaCompareWithPriceCommodityHistory
	//var startEndDate *common.StartEndDateDomain
	var err error
	//var realLastUpdate interface{}
	//var dateFormat = "2006-01-02"
	var latestDateInterface interface{}
	var commodityIsParent bool
	//var latestDateTime time.Time
	//var startDateTime time.Time
	//var endDateTime time.Time
	//var commodityIsParent bool
	//var newDate *common.GenerateNewDate
	//
	//var startDate time.Time
	//startDate, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var endDate time.Time
	//endDate, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.SelectedDate = params.EndDate
	//realLastUpdate, err = repo.GetLastUpdate(params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//lastUpdate := realLastUpdate.(*string)
	//newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//params.StartDate = newDate.StartDate
	//params.EndDate = *lastUpdate

	//params.StartDate = common.TimeToDate(common.StartOfMonth(periodDateTime))
	//
	//periodDateTime, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}
	//params.EndDate = common.TimeToDate(common.EndOfMonth(periodDateTime))

	err = common.NeracaPeriodDateGenerate(&params.StartDate, &params.EndDate)
	if err != nil {
		log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
		return nil, err
	}

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		var priceExists bool
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceExists,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(&priceExists)

		if !priceExists {
			if params.CityId == "73" {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaCompareWithPriceCommodityHistoryProvinceAvgChild,
					pgx.NamedArgs{
						"commodityId": params.CommodityId,
						"provinceId":  params.ProvinceId,
						"cityId":      params.CityId,
						"startDate":   params.StartDate,
						"endDate":     params.EndDate,
					},
				).Scan(
					&result.Unit,
					&result.UnitDiff,
					&result.Commodities,
					&result.CityStock,
					&result.NeracaListWithPriceResponse,
					&result.NeracaTierLevel,
				)
				if err != nil {
					log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
					return nil, err
				}
			} else {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaCompareWithPriceCommodityHistoryAvgChild,
					pgx.NamedArgs{
						"commodityId": params.CommodityId,
						"provinceId":  params.ProvinceId,
						"cityId":      params.CityId,
						"startDate":   params.StartDate,
						"endDate":     params.EndDate,
					},
				).Scan(
					&result.Unit,
					&result.UnitDiff,
					&result.Commodities,
					&result.CityStock,
					&result.NeracaListWithPriceResponse,
					&result.NeracaTierLevel,
				)
				if err != nil {
					log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
					return nil, err
				}
			}
		}
	} else {
		if params.CityId == "73" {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaCompareWithPriceCommodityHistoryProvince,
				pgx.NamedArgs{
					"commodityId": params.CommodityId,
					"provinceId":  params.ProvinceId,
					"cityId":      params.CityId,
					"startDate":   params.StartDate,
					"endDate":     params.EndDate,
				},
			).Scan(
				&result.Unit,
				&result.UnitDiff,
				&result.Commodities,
				&result.CityStock,
				&result.NeracaListWithPriceResponse,
				&result.NeracaTierLevel,
			)
			if err != nil {
				log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
				return nil, err
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaCompareWithPriceCommodityHistory,
				pgx.NamedArgs{
					"commodityId": params.CommodityId,
					"provinceId":  params.ProvinceId,
					"cityId":      params.CityId,
					"startDate":   params.StartDate,
					"endDate":     params.EndDate,
				},
			).Scan(
				&result.Unit,
				&result.UnitDiff,
				&result.Commodities,
				&result.CityStock,
				&result.NeracaListWithPriceResponse,
				&result.NeracaTierLevel,
			)
			if err != nil {
				log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
				return nil, err
			}
		}
	}

	//err = repo.Db.QueryRow(
	//	context.Background(),
	//	queries.TmCommodityIsParent,
	//	pgx.NamedArgs{
	//		"commodityId": params.CommodityId,
	//	},
	//).Scan(&commodityIsParent)
	//
	//if commodityIsParent {
	//
	//}

	//startDateCurrent := time.Date(time.Now().Year(), 01, 01, 0, 0, 0, 0, time.UTC).Format(dateFormat)
	//endDateCurrent := time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC).Format(dateFormat)

	if !commodityIsParent {
		latestDateInterface, err = repo.PriceRepository.LatestDateExist(
			domain.PriceListRepoParams{
				ProvinceId:  params.ProvinceId,
				CommodityId: params.CommodityId,
			},
		)
		if latestDateInterface == nil {
			result.CityStock = &[]models.NeracaCityStockResponse{}
			result.NeracaListWithPriceResponse = &[]models.NeracaListWithPriceResponse{}
			return &result, nil
		}
	} else {
		latestDateInterface, err = repo.PriceRepository.LatestDateExistChild(
			domain.PriceListRepoParams{
				ProvinceId:  params.ProvinceId,
				CommodityId: params.CommodityId,
			},
		)
		if latestDateInterface == nil {
			result.CityStock = &[]models.NeracaCityStockResponse{}
			result.NeracaListWithPriceResponse = &[]models.NeracaListWithPriceResponse{}
			return &result, nil
		}
	}

	//latestDateTime, err = time.Parse(dateFormat, *latestDateInterface.(*string))
	//if err != nil {
	//	return nil, err
	//}
	//
	//startDateTime, err = time.Parse(dateFormat, params.StartDate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//endDateTime, err = time.Parse(dateFormat, params.EndDate)
	//if err != nil {
	//	return nil, err
	//}

	//latestDateTime = 2024-09-29 00:00:00 +0000
	//endDateTime = 2024-07-31 00:00:00 +0000
	//between := common.TimeIsBetween(latestDateTime, startDateTime, endDateTime)
	//latestDateTime.Equal(endDateTime)
	//if between {
	//	if params.CityId == "73" {
	//		err = repo.Db.QueryRow(
	//			context.Background(),
	//			neraca.NeracaCompareWithPriceCommodityHistory,
	//			pgx.NamedArgs{
	//				"commodityId": params.CommodityId,
	//				"provinceId":  params.ProvinceId,
	//				"cityId":      7371,
	//				"startDate":   params.StartDate,
	//				"endDate":     params.EndDate,
	//			},
	//		).Scan(
	//			&result.Unit,
	//			&result.UnitDiff,
	//			&result.Commodities,
	//			&result.CityStock,
	//			&result.NeracaListWithPriceResponse,
	//			&result.NeracaTierLevel,
	//		)
	//		if err != nil {
	//			log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
	//			return nil, err
	//		}
	//	} else {
	//		err = repo.Db.QueryRow(
	//			context.Background(),
	//			neraca.NeracaCompareWithPriceCommodityHistory,
	//			pgx.NamedArgs{
	//				"commodityId": params.CommodityId,
	//				"provinceId":  params.ProvinceId,
	//				"cityId":      params.CityId,
	//				"startDate":   params.StartDate,
	//				"endDate":     params.EndDate,
	//			},
	//		).Scan(
	//			&result.Unit,
	//			&result.UnitDiff,
	//			&result.Commodities,
	//			&result.CityStock,
	//			&result.NeracaListWithPriceResponse,
	//			&result.NeracaTierLevel,
	//		)
	//		if err != nil {
	//			log.Error("QueryErrors GetStockAkhirByCommodityCityHistory:", err)
	//			return nil, err
	//		}
	//	}
	//}

	return &result, nil
}

func (repo *NeracaRepository) GetStockAkhirByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time
	var results models.NeracaStokAkhirByCityMapResponse

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	if params.CityId == "73" {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirByCityMapProvince,
			pgx.NamedArgs{
				"provinceId": 73,
				"cityId":     params.CityId,
				"startDate":  startDate,
				"endDate":    endDate,
			},
		).Scan(
			&results.Unit,
			&results.City,
			&results.Summary,
			&results.ProvinceStock,
			&results.CommodityStock,
			&results.NeracaTierLevel,
			&results.StockTierCode,
		)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirByCityMap,
			pgx.NamedArgs{
				"provinceId": 73,
				"cityId":     params.CityId,
				"startDate":  startDate,
				"endDate":    endDate,
			},
		).Scan(
			&results.Unit,
			&results.City,
			&results.Summary,
			&results.ProvinceStock,
			&results.CommodityStock,
			&results.NeracaTierLevel,
			&results.StockTierCode,
		)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	}

	//if results.ProvinceStock != nil {
	//	if results.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := results.ProvinceStock.Province.Assets.AssetsLocation + "/" + results.ProvinceStock.Province.Assets.AssetsName
	//		results.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//if results.CommodityStock != nil {
	//	for _, row := range *results.CommodityStock {
	//		if row.Commodity.Assets != nil {
	//			if row.Commodity.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//				//baseUrl := os.Getenv("BASE_URL")
	//				baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//				assetsLocation := row.Commodity.Assets.AssetsLocation + "/" + row.Commodity.Assets.AssetsName
	//				row.Commodity.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//			}
	//		} else {
	//			log.Error("Commodity asset is nil, ", row.Commodity.Id)
	//		}
	//	}
	//} else {
	//	results.CommodityStock = &[]models.NeracaCommodityStockResponse{}
	//}

	if results.CommodityStock == nil {
		results.CommodityStock = &[]models.NeracaCommodityStockResponse{}
	}

	return &results, nil
}

func (repo *NeracaRepository) GetStockAkhirListByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var sortBy string
	var sortByRank string
	var result models.NeracaStokAkhirListResponse

	if params.PaginationParams.SortBy != "" {
		split := strings.Split(params.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if sortByRank != "" {
			if sortByRank == "asc" {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaStokAkhirListByCommodityAsc,
					pgx.NamedArgs{
						"provinceId":   73,
						"commodityId":  params.CommodityId,
						"selectedDate": params.SelectedDate,
						"page":         params.PaginationParams.Page,
						"limit":        params.PaginationParams.Limit,
						"sortBy":       sortBy,
						"sortByRank":   sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			} else {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaStokAkhirListByCommodityDesc,
					pgx.NamedArgs{
						"provinceId":   73,
						"commodityId":  params.CommodityId,
						"selectedDate": params.SelectedDate,
						"page":         params.PaginationParams.Page,
						"limit":        params.PaginationParams.Limit,
						"sortBy":       sortBy,
						"sortByRank":   sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaStokAkhirListByCommodityAsc,
				pgx.NamedArgs{
					"provinceId":   73,
					"commodityId":  params.CommodityId,
					"selectedDate": params.SelectedDate,
					"page":         params.PaginationParams.Page,
					"limit":        params.PaginationParams.Limit,
					"sortBy":       sortBy,
					"sortByRank":   sortByRank,
				},
			).Scan(
				&result.Unit,
				&result.Commodities,
				&result.ProvinceStock,
				&result.CityStock,
				&result.NeracaTierLevel,
			)
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirListByCommodity,
			pgx.NamedArgs{
				"provinceId":   73,
				"commodityId":  params.CommodityId,
				"selectedDate": params.SelectedDate,
				"page":         params.PaginationParams.Page,
				"limit":        params.PaginationParams.Limit,
			},
		).Scan(
			&result.Unit,
			&result.Commodities,
			&result.ProvinceStock,
			&result.CityStock,
			&result.NeracaTierLevel,
		)
	}
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if result.ProvinceStock == nil || result.CityStock == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.ProvinceStock != nil {
	//	if result.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.ProvinceStock.Province.Assets.AssetsLocation + "/" + result.ProvinceStock.Province.Assets.AssetsName
	//		result.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *NeracaRepository) GetKetersediaanByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time
	var results models.NeracaKetersediaanByCityMapResponse

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	if params.CityId == "73" {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaKetersediaanByCityMapProvince,
			pgx.NamedArgs{
				"provinceId":   73,
				"cityId":       params.CityId,
				"selectedDate": params.SelectedDate,
				"startDate":    startDate,
				"endDate":      endDate,
			},
		).Scan(
			&results.Unit,
			&results.City,
			&results.Summary,
			&results.ProvinceStock,
			&results.CommodityStock,
			&results.NeracaTierLevel,
			&results.StockTierCode,
		)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaKetersediaanByCityMap,
			pgx.NamedArgs{
				"provinceId": 73,
				"cityId":     params.CityId,
				"startDate":  startDate,
				"endDate":    endDate,
			},
		).Scan(
			&results.Unit,
			&results.City,
			&results.Summary,
			&results.ProvinceStock,
			&results.CommodityStock,
			&results.NeracaTierLevel,
			&results.StockTierCode,
		)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	}

	//if results.ProvinceStock != nil {
	//	if results.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := results.ProvinceStock.Province.Assets.AssetsLocation + "/" + results.ProvinceStock.Province.Assets.AssetsName
	//		results.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *results.CommodityStock {
	//	if row.Commodity.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.Commodity.Assets.AssetsLocation + "/" + row.Commodity.Assets.AssetsName
	//		row.Commodity.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &results, nil
}

func (repo *NeracaRepository) GetKetersediaanListByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var sortBy string
	var sortByRank string
	var result models.NeracaStokAkhirListResponse

	if params.PaginationParams.SortBy != "" {
		split := strings.Split(params.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if sortByRank != "" {
			if sortByRank == "asc" {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaKetersediaanListByCommodityAsc,
					pgx.NamedArgs{
						"provinceId":   73,
						"commodityId":  params.CommodityId,
						"selectedDate": params.SelectedDate,
						"page":         params.PaginationParams.Page,
						"limit":        params.PaginationParams.Limit,
						"sortBy":       sortBy,
						"sortByRank":   sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			} else {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaStokAkhirListByCommodityDesc,
					pgx.NamedArgs{
						"provinceId":   73,
						"commodityId":  params.CommodityId,
						"selectedDate": params.SelectedDate,
						"page":         params.PaginationParams.Page,
						"limit":        params.PaginationParams.Limit,
						"sortBy":       sortBy,
						"sortByRank":   sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaStokAkhirListByCommodityAsc,
				pgx.NamedArgs{
					"provinceId":   73,
					"commodityId":  params.CommodityId,
					"selectedDate": params.SelectedDate,
					"page":         params.PaginationParams.Page,
					"limit":        params.PaginationParams.Limit,
					"sortBy":       sortBy,
					"sortByRank":   sortByRank,
				},
			).Scan(
				&result.Unit,
				&result.Commodities,
				&result.ProvinceStock,
				&result.CityStock,
				&result.NeracaTierLevel,
			)
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirListByCommodity,
			pgx.NamedArgs{
				"provinceId":   73,
				"commodityId":  params.CommodityId,
				"selectedDate": params.SelectedDate,
				"page":         params.PaginationParams.Page,
				"limit":        params.PaginationParams.Limit,
			},
		).Scan(
			&result.Unit,
			&result.Commodities,
			&result.ProvinceStock,
			&result.CityStock,
			&result.NeracaTierLevel,
		)
	}
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if result.ProvinceStock == nil || result.CityStock == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.ProvinceStock != nil {
	//	if result.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.ProvinceStock.Province.Assets.AssetsLocation + "/" + result.ProvinceStock.Province.Assets.AssetsName
	//		result.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *NeracaRepository) GetKebutuhanByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var dateFormat = "2006-01-02"
	var startDate string
	var endDate string
	var selectedDateTime time.Time
	var results models.NeracaKetersediaanByCityMapResponse

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	startDate = common.TimeToDate(common.StartOfMonth(selectedDateTime))

	selectedDateTime, err = time.Parse(dateFormat, params.SelectedDate)
	if err != nil {
		return nil, err
	}
	endDate = common.TimeToDate(common.EndOfMonth(selectedDateTime))

	if params.CityId == "73" {
		err := repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaKebutuhanByCityMapProvince,
			pgx.NamedArgs{
				"provinceId": 73,
				"cityId":     params.CityId,
				"startDate":  startDate,
				"endDate":    endDate,
			},
		).Scan(
			&results.Unit,
			&results.City,
			&results.Summary,
			&results.ProvinceStock,
			&results.CommodityStock,
			&results.NeracaTierLevel,
			&results.StockTierCode,
		)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaKebutuhanByCityMap,
			pgx.NamedArgs{
				"provinceId": 73,
				"cityId":     params.CityId,
				"startDate":  startDate,
				"endDate":    endDate,
			},
		).Scan(
			&results.Unit,
			&results.City,
			&results.Summary,
			&results.ProvinceStock,
			&results.CommodityStock,
			&results.NeracaTierLevel,
			&results.StockTierCode,
		)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	}

	//if results.ProvinceStock != nil {
	//	if results.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := results.ProvinceStock.Province.Assets.AssetsLocation + "/" + results.ProvinceStock.Province.Assets.AssetsName
	//		results.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *results.CommodityStock {
	//	if row.Commodity.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.Commodity.Assets.AssetsLocation + "/" + row.Commodity.Assets.AssetsName
	//		row.Commodity.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &results, nil
}

func (repo *NeracaRepository) GetKebutuhanListByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var err error
	var sortBy string
	var sortByRank string
	var result models.NeracaStokAkhirListResponse

	if params.PaginationParams.SortBy != "" {
		split := strings.Split(params.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if sortByRank != "" {
			if sortByRank == "asc" {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaKetersediaanListByCommodityAsc,
					pgx.NamedArgs{
						"provinceId":   73,
						"commodityId":  params.CommodityId,
						"selectedDate": params.SelectedDate,
						"page":         params.PaginationParams.Page,
						"limit":        params.PaginationParams.Limit,
						"sortBy":       sortBy,
						"sortByRank":   sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			} else {
				err = repo.Db.QueryRow(
					context.Background(),
					neraca.NeracaStokAkhirListByCommodityDesc,
					pgx.NamedArgs{
						"provinceId":   73,
						"commodityId":  params.CommodityId,
						"selectedDate": params.SelectedDate,
						"page":         params.PaginationParams.Page,
						"limit":        params.PaginationParams.Limit,
						"sortBy":       sortBy,
						"sortByRank":   sortByRank,
					},
				).Scan(
					&result.Unit,
					&result.Commodities,
					&result.ProvinceStock,
					&result.CityStock,
					&result.NeracaTierLevel,
				)
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				neraca.NeracaStokAkhirListByCommodityAsc,
				pgx.NamedArgs{
					"provinceId":   73,
					"commodityId":  params.CommodityId,
					"selectedDate": params.SelectedDate,
					"page":         params.PaginationParams.Page,
					"limit":        params.PaginationParams.Limit,
					"sortBy":       sortBy,
					"sortByRank":   sortByRank,
				},
			).Scan(
				&result.Unit,
				&result.Commodities,
				&result.ProvinceStock,
				&result.CityStock,
				&result.NeracaTierLevel,
			)
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			neraca.NeracaStokAkhirListByCommodity,
			pgx.NamedArgs{
				"provinceId":   73,
				"commodityId":  params.CommodityId,
				"selectedDate": params.SelectedDate,
				"page":         params.PaginationParams.Page,
				"limit":        params.PaginationParams.Limit,
			},
		).Scan(
			&result.Unit,
			&result.Commodities,
			&result.ProvinceStock,
			&result.CityStock,
			&result.NeracaTierLevel,
		)
	}
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	if result.ProvinceStock == nil || result.CityStock == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.ProvinceStock != nil {
	//	if result.ProvinceStock.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.ProvinceStock.Province.Assets.AssetsLocation + "/" + result.ProvinceStock.Province.Assets.AssetsName
	//		result.ProvinceStock.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.CityStock {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *NeracaRepository) Exist(params domain.NeracaStokAkhirListRequestParams) (interface{}, error) {
	var result map[string]bool
	err := repo.Db.QueryRow(
		context.Background(),
		neraca.NeracaExist,
		pgx.NamedArgs{
			"provinceId":  params.ProvinceId,
			"commodityId": params.CommodityId,
			"startDate":   params.StartDate,
			"endDate":     params.EndDate,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *NeracaRepository) LatestDateExist(params domain.PriceListRepoParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		neraca.NeracaLatestDateExist,
		pgx.NamedArgs{
			"provinceId":  params.ProvinceId,
			"commodityId": params.CommodityId,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *NeracaRepository) GetNeracaReport(params domain.NeracaListParams) (interface{}, error) {
	//var result []map[string]string
	var result []map[string]interface{}
	var data []models.ReportNeracaModel
	var rows pgx.Rows
	var err error

	if params.CommodityId != "" && params.CommodityId != "0" {
		var commodityIsParent bool

		err = repo.Db.QueryRow(
			context.Background(),
			queries.TmCommodityIsParent,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
			},
		).Scan(&commodityIsParent)

		if commodityIsParent {
			rows, err = repo.Db.Query(
				context.Background(),
				queries.ReportNeracaByCommodityAvgChild,
				pgx.NamedArgs{
					"page":        params.PaginationParams.Page,
					"limit":       params.PaginationParams.Limit,
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   params.StartDate + " 00:00:00",
					"endDate":     params.EndDate + " 23:59:59",
				},
			)
		} else {
			rows, err = repo.Db.Query(
				context.Background(),
				queries.ReportNeracaByCommodity,
				pgx.NamedArgs{
					"page":        params.PaginationParams.Page,
					"limit":       params.PaginationParams.Limit,
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   params.StartDate + " 00:00:00",
					"endDate":     params.EndDate + " 23:59:59",
				},
			)
		}
	} else {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.ReportNeracaByCity,
			pgx.NamedArgs{
				"page":      params.PaginationParams.Page,
				"limit":     params.PaginationParams.Limit,
				"cityId":    params.CityId,
				"startDate": params.StartDate + " 00:00:00",
				"endDate":   params.EndDate + " 23:59:59",
			},
		)
	}

	if err != nil {
		log.Error("QueryErrors Get:", err)
		return nil, err
	}

	for rows.Next() {
		var reportNeracaData models.ReportNeracaModel
		err = rows.Scan(
			&reportNeracaData.Title,
			&reportNeracaData.Stocks,
		)

		data = append(data, reportNeracaData)
	}

	if data == nil {
		responseEmpty := make(map[string]interface{})
		return responseEmpty, nil
	}

	i := 0
	for _, v := range data {
		result = append(result, map[string]interface{}{
			"title": v.Title,
		})

		i++
	}

	var response []map[string]interface{}
	idx := 0
	for _, v := range result {
		if v["title"] == data[idx].Title {
			stocksMap := make(map[string]map[string]int)
			for stocksMapKey, stocksMapValues := range data[idx].Stocks {
				stocksMapPeriodValue := make(map[string]int)
				for _, stocksMapValue := range stocksMapValues {
					for stockKey, stockValue := range stocksMapValue {
						stocksMapPeriodValue[stockKey] = stockValue
					}
				}

				stocksMap[stocksMapKey] = stocksMapPeriodValue
				v["stocks"] = stocksMap
			}
		}

		response = append(response, v)
		idx++
	}

	return &response, nil
}

func (repo *NeracaRepository) GetNeracaReportDownload(params domain.NeracaListParams) (interface{}, error) {
	//var result []map[string]string
	var result []map[string]interface{}
	var data []models.ReportNeracaModel
	var rows pgx.Rows
	var err error

	if params.CommodityId != "" && params.CommodityId != "0" {
		var commodityIsParent bool

		err = repo.Db.QueryRow(
			context.Background(),
			queries.TmCommodityIsParent,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
			},
		).Scan(&commodityIsParent)

		if commodityIsParent {
			rows, err = repo.Db.Query(
				context.Background(),
				queries.ReportNeracaByCommodityAvgChildDwonload,
				pgx.NamedArgs{
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   params.StartDate + " 00:00:00",
					"endDate":     params.EndDate + " 23:59:59",
				},
			)
		} else {
			rows, err = repo.Db.Query(
				context.Background(),
				queries.ReportNeracaByCommodityDownload,
				pgx.NamedArgs{
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   params.StartDate + " 00:00:00",
					"endDate":     params.EndDate + " 23:59:59",
				},
			)
		}
	} else {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.ReportNeracaByCityDownload,
			pgx.NamedArgs{
				"cityId":    params.CityId,
				"startDate": params.StartDate + " 00:00:00",
				"endDate":   params.EndDate + " 23:59:59",
			},
		)
	}

	if err != nil {
		log.Error("QueryErrors Get:", err)
		return nil, err
	}

	for rows.Next() {
		var reportNeracaData models.ReportNeracaModel
		err = rows.Scan(
			&reportNeracaData.Title,
			&reportNeracaData.Stocks,
		)

		data = append(data, reportNeracaData)
	}

	if data == nil {
		//responseEmpty := make(map[string]interface{})
		return nil, errors.New("Data Tidak Ditemukan")
	}

	i := 0
	for _, v := range data {
		result = append(result, map[string]interface{}{
			"title": v.Title,
		})

		i++
	}

	var response []map[string]interface{}
	idx := 0

	// for _, v := range result {
	// 	if v["title"] == data[idx].Title {
	// 		stocksMap := make(map[string]map[string]int)
	// 		for stocksMapKey, stocksMapValues := range data[idx].Stocks {
	// 			stocksMapPeriodValue := make(map[string]int)
	// 			for _, stocksMapValue := range stocksMapValues {
	// 				for stockKey, stockValue := range stocksMapValue {
	// 					stocksMapPeriodValue[stockKey] = stockValue
	// 				}
	// 				fmt.Println(stocksMapPeriodValue)
	// 			}

	// 			stocksMap[stocksMapKey] = stocksMapPeriodValue
	// 			v["stocks"] = stocksMap
	// 		}
	// 	}

	// 	response = append(response, v)
	// 	idx++
	// }
	for i := range result {
		if result[i]["title"] == data[idx].Title {
			stocksMap := make(map[string]map[string]int)
			for stocksMapKey, stocksMapValues := range data[idx].Stocks {
				stocksMapPeriodValue := make(map[string]int)

				keys := make([]string, 0, len(stocksMapValues))

				// Simpan key dengan format YYYYMM
				formattedStocksMap := make(map[string]int)

				for _, stocksMapValue := range stocksMapValues {
					for stockKey, stockValue := range stocksMapValue {
						// Ubah MMYYYY menjadi YYYYMM
						if len(stockKey) == 6 { // Pastikan format MMYYYY
							newKey := stockKey[2:6] + stockKey[0:2] // YYYYMM
							formattedStocksMap[newKey] = stockValue
							keys = append(keys, newKey)
						}
					}
				}

				// Sort keys secara ASC
				sort.Strings(keys)

				// Simpan hasil yang sudah terurut ke stocksMapPeriodValue
				for _, key := range keys {
					stocksMapPeriodValue[key] = formattedStocksMap[key]
				}

				// Simpan ke dalam map utama
				stocksMap[stocksMapKey] = stocksMapPeriodValue
			}

			// Update stocks ke dalam result tanpa mengubah urutan array
			result[i]["stocks"] = stocksMap
		}

		// Tambahkan hasil ke response tanpa mengubah urutan
		response = append(response, result[i])
		idx++
	}

	return &response, nil
}

func (repo *NeracaRepository) GetNeracaReportCount(params domain.NeracaListParams) (*int, error) {
	var result int
	var commodityIsParent bool
	var err error

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
		},
	).Scan(&commodityIsParent)

	if params.CommodityId != "" && params.CommodityId != "0" {
		if commodityIsParent {
			err = repo.Db.QueryRow(
				context.Background(),
				queries.ReportNeracaByCommodityCountAvgChild,
				pgx.NamedArgs{
					"page":        params.PaginationParams.Page,
					"limit":       params.PaginationParams.Limit,
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   params.StartDate + " 00:00:00",
					"endDate":     params.EndDate + " 23:59:59",
				},
			).Scan(&result)
			if err != nil {
				return nil, err
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				queries.ReportNeracaByCommodityCount,
				pgx.NamedArgs{
					"page":        params.PaginationParams.Page,
					"limit":       params.PaginationParams.Limit,
					"provinceId":  73,
					"commodityId": params.CommodityId,
					"startDate":   params.StartDate + " 00:00:00",
					"endDate":     params.EndDate + " 23:59:59",
				},
			).Scan(&result)
			if err != nil {
				return nil, err
			}
		}

		return &result, nil
	} else {
		err := repo.Db.QueryRow(
			context.Background(),
			queries.ReportNeracaByCityCount,
			pgx.NamedArgs{
				"page":      params.PaginationParams.Page,
				"limit":     params.PaginationParams.Limit,
				"cityId":    params.CityId,
				"startDate": params.StartDate + " 00:00:00",
				"endDate":   params.EndDate + " 23:59:59",
			},
		).Scan(&result)
		if err != nil {
			return nil, err
		}

		return &result, nil
	}
}

//func (repo *NeracaRepository) Get(queryParams domain.NeracaGetRepoParams) interface{} {
//	var result models.NeracaTableResponse
//	rows, err := repo.Db.Query(
//		context.Background(),
//		queries.NeracaPerubahanHarga,
//		pgx.NamedArgs{
//			"commodityType": queryParams.CommodityType,
//			"startDate":     queryParams.StartDate,
//			"endDate":       queryParams.EndDate,
//		},
//	)
//
//	if err != nil {
//		log.Error("QueryErrors:", err)
//	}
//
//	for rows.Next() {
//		err = rows.Scan(&result.StockCommodity, &result.StockDiff)
//		if err != nil {
//			fmt.Printf("Scan error: %v\n", err)
//		}
//	}
//
//	return &result
//}
