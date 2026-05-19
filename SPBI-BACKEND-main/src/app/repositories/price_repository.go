package repositories

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"sigap-sultan-be/src/app/repositories/queries/price"
	"sigap-sultan-be/src/common"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jmoiron/sqlx"
)

type PriceRepository struct {
	Db *pgxpool.Pool
}

func NewPriceRepository(db *pgxpool.Pool) *PriceRepository {
	return &PriceRepository{Db: db}
}

func (repo *PriceRepository) SavePriceCity(prices []*models.PriceCity) {
	var buffer bytes.Buffer
	query := price.PriceTxCommodityCityInsert

	for idx, price := range prices {
		if idx == 0 {
			buffer.WriteString(
				"\n( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.CityId)) + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Price + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n",
			)
		} else {
			buffer.WriteString(
				",( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.CityId)) + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Price + "'," +
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

func (repo *PriceRepository) SavePriceProvince(prices []*models.PriceProvince) {
	var buffer bytes.Buffer
	query := price.PriceTxCommodityProvinceInsert

	for idx, price := range prices {
		if idx == 0 {
			buffer.WriteString(
				"\n( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.ProvinceId)) + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Price + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n",
			)
		} else {
			buffer.WriteString(
				",( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.ProvinceId)) + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Price + "'," +
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

func (repo *PriceRepository) SavePriceNational(prices []*models.PriceNational) {
	var buffer bytes.Buffer
	query := price.PriceTxCommodityNationalInsert

	for idx, price := range prices {
		if idx == 0 {
			buffer.WriteString(
				"\n( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.NationalId)) + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Price + "'," +
					"'" + *price.LastUpdate + "'," +
					"'" + common.GetDateTimeNow() + "'" +
					")\n",
			)
		} else {
			buffer.WriteString(
				",( " +
					"'1'," +
					"'" + strconv.Itoa(int(price.NationalId)) + "'," +
					"'" + strconv.Itoa(int(price.CommodityId)) + "'," +
					"'" + price.CommodityName + "'," +
					"'" + price.Price + "'," +
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

func (repo *PriceRepository) GetLevelHarga(queryParams domain.PriceGetRepoParamsNew) (interface{}, error) {
	var results models.PriceTableHargaLevelMap
	err := repo.Db.QueryRow(
		context.Background(),
		price.PricePerubahanHargaMap,
		pgx.NamedArgs{
			"provinceId":   73,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(
		&results.PriceLevel,
		&results.PriceMin,
		&results.PriceMax,
		&results.PriceDiff,
		&results.PriceBarCategory,
	)

	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	var tiers []models.PriceTier

	if results.PriceDiff != nil {
		if *results.PriceDiff > 0 {
			i := 1
			var currentPrice = int32(0)
			var nextPrice = int32(0)

			for i <= 5 {
				if i == 1 {
					currentPrice = *results.PriceMin + *results.PriceBarCategory
					tiers = append(tiers, models.PriceTier{
						Title:                "Sangat Rendah",
						PriceMin:             *results.PriceMin,
						PriceMinRupiahFormat: common.ThousandFormat(*results.PriceMin),
						PriceMax:             currentPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(currentPrice),
						Color:                "#208245",
					})
					nextPrice = currentPrice + 1
				} else if i == 2 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Rendah",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#27AE65",
					})
					nextPrice++
				} else if i == 3 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Sedang",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#FFFF00",
					})
					nextPrice++
				} else if i == 4 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Tinggi",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#C0392B",
					})
					nextPrice++
				} else if i == 5 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Sangat Tinggi",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#7C2019",
					})
				}

				i++
			}
		} else {
			tiers = append(tiers, models.PriceTier{
				Title:                "Sangat Tinggi",
				PriceMin:             *results.PriceMin,
				PriceMinRupiahFormat: common.ThousandFormat(*results.PriceMin),
				PriceMax:             *results.PriceMax,
				PriceMaxRupiahFormat: common.ThousandFormat(*results.PriceMax),
				Color:                "#7C2019",
			})
		}
	}

	results.Tier = tiers

	return &results, nil
}

func (repo *PriceRepository) GetLevelHargaNew(queryParams domain.PriceGetRepoParamsNew) (interface{}, error) {
	var err error
	var commodityIsParent bool
	var results models.PriceTableHargaLevelNew

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PricePerubahanHargaMapNewAvgChild,
			pgx.NamedArgs{
				"provinceId":   73,
				"commodityId":  queryParams.CommodityId,
				"selectedDate": queryParams.SelectedDate,
			},
		).Scan(
			&results.PriceProvince,
			&results.CityPrice,
			&results.PriceMin,
			&results.PriceMax,
			&results.PriceDiff,
			&results.PriceBarCategory,
		)
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PricePerubahanHargaMapNew,
			pgx.NamedArgs{
				"provinceId":   73,
				"commodityId":  queryParams.CommodityId,
				"selectedDate": queryParams.SelectedDate,
			},
		).Scan(
			&results.PriceProvince,
			&results.CityPrice,
			&results.PriceMin,
			&results.PriceMax,
			&results.PriceDiff,
			&results.PriceBarCategory,
		)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	}

	var tiers []models.PriceTier

	if results.PriceDiff != nil {
		if *results.PriceDiff > 0 {
			i := 1
			var currentPrice = int32(0)
			var nextPrice = int32(0)

			for i <= 5 {
				if i == 1 {
					currentPrice = *results.PriceMin + *results.PriceBarCategory
					tiers = append(tiers, models.PriceTier{
						Title:                "Sangat Rendah",
						PriceMin:             *results.PriceMin,
						PriceMinRupiahFormat: common.ThousandFormat(*results.PriceMin),
						PriceMax:             currentPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(currentPrice),
						Color:                "#208245",
					})
					nextPrice = currentPrice + 1
				} else if i == 2 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Rendah",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#27AE65",
					})
					nextPrice++
				} else if i == 3 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Sedang",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#FFFF00",
					})
					nextPrice++
				} else if i == 4 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Tinggi",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#C0392B",
					})
					nextPrice++
				} else if i == 5 {
					currentPrice = nextPrice
					nextPrice = currentPrice + (*results.PriceBarCategory - 1)
					tiers = append(tiers, models.PriceTier{
						Title:                "Sangat Tinggi",
						PriceMin:             currentPrice,
						PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
						PriceMax:             nextPrice,
						PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
						Color:                "#7C2019",
					})
				}

				i++
			}
		} else {
			tiers = append(tiers, models.PriceTier{
				Title:                "Sangat Tinggi",
				PriceMin:             *results.PriceMin,
				PriceMinRupiahFormat: common.ThousandFormat(*results.PriceMin),
				PriceMax:             *results.PriceMax,
				PriceMaxRupiahFormat: common.ThousandFormat(*results.PriceMax),
				Color:                "#7C2019",
			})
		}
	}
	results.Tier = tiers

	return &results, nil
}

func (repo *PriceRepository) GetLevelHargaList(queryParams domain.PriceGetRepoParamsNew) (interface{}, error) {
	log.Info(queryParams)
	var result models.PriceTableHargaLevelList
	var err error
	var sortBy string
	var sortByRank string
	var query string
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if queryParams.PaginationParams.SortBy != "" {
		split := strings.Split(queryParams.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if commodityIsParent {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceLevelHargaListWithOrderAvgChild
				} else {
					query = price.PriceLevelHargaListWithOrderDescAvgChild
				}
			} else {
				query = price.PriceLevelHargaListAvgChild
			}
		} else {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceLevelHargaListWithOrder
				} else {
					query = price.PriceLevelHargaListWithOrderDesc
				}
			} else {
				query = price.PriceLevelHargaList
			}
		}

		//if sortByRank != "" {
		//	if sortByRank == "asc" {
		//		err = repo.Db.QueryRow(
		//			context.Background(),
		//			query,
		//			pgx.NamedArgs{
		//				"nationalId":   1,
		//				"provinceId":   73,
		//				"commodityId":  queryParams.CommodityId,
		//				"selectedDate": queryParams.SelectedDate,
		//				"page":         queryParams.PaginationParams.Page,
		//				"limit":        queryParams.PaginationParams.Limit,
		//				"sortBy":       sortBy,
		//				"sortByRank":   sortByRank,
		//			},
		//		).Scan(
		//			&result.PriceProvince,
		//			&result.PriceCity,
		//		)
		//	} else {
		//		err = repo.Db.QueryRow(
		//			context.Background(),
		//			query,
		//			pgx.NamedArgs{
		//				"nationalId":   1,
		//				"provinceId":   73,
		//				"commodityId":  queryParams.CommodityId,
		//				"selectedDate": queryParams.SelectedDate,
		//				"page":         queryParams.PaginationParams.Page,
		//				"limit":        queryParams.PaginationParams.Limit,
		//				"sortBy":       sortBy,
		//				"sortByRank":   sortByRank,
		//			},
		//		).Scan(
		//			&result.PriceProvince,
		//			&result.PriceCity,
		//		)
		//	}
		//} else {
		//	err = repo.Db.QueryRow(
		//		context.Background(),
		//		query,
		//		pgx.NamedArgs{
		//			"nationalId":   1,
		//			"provinceId":   73,
		//			"commodityId":  queryParams.CommodityId,
		//			"selectedDate": queryParams.SelectedDate,
		//			"page":         queryParams.PaginationParams.Page,
		//			"limit":        queryParams.PaginationParams.Limit,
		//			"sortBy":       sortBy,
		//			"sortByRank":   sortByRank,
		//		},
		//	).Scan(
		//		&result.PriceProvince,
		//		&result.PriceCity,
		//	)
		//}
	} else {
		if commodityIsParent {
			query = price.PriceLevelHargaListAvgChild
		} else {
			query = price.PriceLevelHargaList
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"nationalId":   1,
			"provinceId":   73,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
			"sortBy":       sortBy,
			"sortByRank":   sortByRank,
		},
	).Scan(
		&result.PriceProvince,
		&result.PriceCity,
	)
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}
	if result.PriceProvince == nil || result.PriceCity == nil {
		return nil, errors.New("no rows in result set")
	}

	var priceDiffLevelHargaData interface{}
	if commodityIsParent {
		priceDiffLevelHargaData, err = repo.GetPriceDiffLevelHargaAvgChild(queryParams)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	} else {
		priceDiffLevelHargaData, err = repo.GetPriceDiffLevelHarga(queryParams)
		if err != nil {
			log.Error("err")
			log.Error(err)
			return nil, err
		}
	}

	var priceDiffLevelHarga *models.PriceDiffLevelHarga
	priceDiffLevelHarga = priceDiffLevelHargaData.(*models.PriceDiffLevelHarga)

	var tiers []models.PriceTier
	if priceDiffLevelHarga.PriceDiff != nil {
		generatePriceTier(priceDiffLevelHarga, &tiers)
	}

	//if result.PriceProvince != nil {
	//	if result.PriceProvince.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.PriceProvince.Province.Assets.AssetsLocation + "/" + result.PriceProvince.Province.Assets.AssetsName
	//		result.PriceProvince.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	var dataLists []models.PriceTableHargaLevelListData
	for _, row := range *result.PriceCity {
		//if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
		//	//baseUrl := os.Getenv("BASE_URL")
		//	baseUrl := "https://project.bi.sentech.id/api/v1/stg"
		//	assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
		//	row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
		//}

		for _, v := range tiers {
			if row.Price != nil {
				if *row.Price >= v.PriceMin && *row.Price <= v.PriceMax {
					title := strings.ReplaceAll(strings.ToLower(v.Title), " ", "-")
					row.Tier = &title
				}
			}

			if result.PriceProvince.Price != nil {
				if *result.PriceProvince.Price >= v.PriceMin && *result.PriceProvince.Price <= v.PriceMax {
					title := strings.ReplaceAll(strings.ToLower(v.Title), " ", "-")
					result.PriceProvince.Tier = title
				}
			}
		}

		dataLists = append(dataLists, row)
	}
	result.PriceCity = &dataLists

	return &result, nil
}

func (repo PriceRepository) GetLast5DaysPriceByCityId(params domain.PriceLast5DaysRepoParams) (interface{}, error) {
	var results models.PriceLast5DaysCityByIdResponse

	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceLast5DaysPriceByCityId,
		pgx.NamedArgs{
			"cityId":    params.CityId,
			"startDate": params.StartDate,
			"endDate":   params.EndDate,
		},
	).Scan(
		&results.City,
		&results.Commodities,
	)

	if err != nil {
		log.Error("QueryErrors GetLast5DaysPriceByCityId:", err)
		return nil, err
	}

	return &results, nil
}

func (repo *PriceRepository) GetLast5DaysPriceByCityIdCount(params domain.PriceLast5DaysRepoParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceLast5DaysPriceByCityIdCount,
		pgx.NamedArgs{
			"cityId": params.CityId,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo PriceRepository) GetLast5DaysPriceByCommodityId(params domain.PriceLast5DaysRepoParams) (interface{}, error) {
	var results models.PriceLast5DaysCommodityByIdResponse
	var commodityIsParent bool
	var err error

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceLast5DaysPriceByCommodityIdNewAvgChild,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&results.Commodities,
			&results.City,
		)
		if err != nil {
			log.Error("QueryErrors GetLast5DaysPriceByCommodityId:", err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceLast5DaysPriceByCommodityIdNew,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&results.Commodities,
			&results.City,
		)
		if err != nil {
			log.Error("QueryErrors GetLast5DaysPriceByCommodityId:", err)
			return nil, err
		}
	}

	return &results, nil
}

func (repo *PriceRepository) GetLast5DaysPriceByCommodityIdCount(params domain.PriceLast5DaysRepoParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceLast5DaysPriceByCommodityIdCount,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetCompareBySulsel(queryParams domain.PriceGetCompareProvinceParams) (interface{}, error) {
	var result models.PriceTableCompareProvinceMap
	var err error
	var commodityIsParent bool
	var query string

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		query = price.PriceGetCompareBySulselAvgChild
	} else {
		query = price.PriceGetCompareBySulsel
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(
		&result.Summary,
		&result.Commodity,
		&result.PriceProvince,
		&result.PriceLevel,
		&result.PriceTier,
		&result.PriceTierCode,
	)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	if result.PriceProvince == nil || result.PriceLevel == nil {
		return nil, errors.New("no rows in result set")
	}

	//for _, row := range *result.PriceLevel {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *PriceRepository) GetCompareBySulselList(queryParams domain.PriceGetCompareProvinceParams) (interface{}, error) {
	var result models.PriceTableCompareProvinceList
	var rows pgx.Rows
	var err error
	var sortBy string
	var sortByRank string
	var commodityIsParent bool
	var query string

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if queryParams.PaginationParams.SortBy != "" {
		split := strings.Split(queryParams.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if commodityIsParent {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetCompareBySulselListWithOrderAscAvgChild
				} else {
					query = price.PriceGetCompareBySulselListWithOrderDescAvgChild
				}
			} else {
				query = price.PriceGetCompareBySulselListAvgChild
			}
		} else {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetCompareBySulselListWithOrderAsc
				} else {
					query = price.PriceGetCompareBySulselListWithOrderDesc
				}
			} else {
				query = price.PriceGetCompareBySulselList
			}
		}
	} else {
		if commodityIsParent {
			query = price.PriceGetCompareBySulselListAvgChild
		} else {
			query = price.PriceGetCompareBySulselList
		}
	}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
			"sortBy":       sortBy,
			"sortByRank":   sortByRank,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Commodity,
			&result.PriceProvince,
			&result.PriceCity,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	if result.PriceProvince == nil || result.PriceCity == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.PriceProvince != nil {
	//	if result.PriceProvince.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.PriceProvince.Province.Assets.AssetsLocation + "/" + result.PriceProvince.Province.Assets.AssetsName
	//		result.PriceProvince.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.PriceCity {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *PriceRepository) GetCompareProvinceCityHistory(queryParams domain.PriceDiffByCityAndCommodityParams) (interface{}, error) {
	var result models.PriceTableCompareProvinceCityHistory
	var err error
	var realLastUpdate interface{}
	var dateFormat = "2006-01-02"
	var newDate *common.GenerateNewDate
	var query string

	var startDate time.Time
	startDate, err = time.Parse(dateFormat, queryParams.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate time.Time
	endDate, err = time.Parse(dateFormat, queryParams.EndDate)
	if err != nil {
		return nil, err
	}

	//realLastUpdate, err = repo.GetLastUpdate(
	//	domain.PriceLastUpdateParams{
	//		ProvinceId:   "73",
	//		CityId:       queryParams.CityId,
	//		CommodityId:  queryParams.CommodityId,
	//		SelectedDate: queryParams.EndDate,
	//	},
	//)
	//if err != nil {
	//	return nil, err
	//}

	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvinceAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		} else {
			realLastUpdate, err = repo.GetLastUpdateAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		}
	} else {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvince(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		} else {
			realLastUpdate, err = repo.GetLastUpdate(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		}
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	queryParams.StartDate = newDate.StartDate
	queryParams.EndDate = *lastUpdate

	if commodityIsParent {
		if queryParams.CityId == "false" {
			query = price.PriceGetCompareBySulselCityHistoryProvinceAvgChild
		} else {
			query = price.PriceGetCompareBySulselCityHistoryAvgChild
		}
	} else {
		if queryParams.CityId == "false" {
			query = price.PriceGetCompareBySulselCityHistoryProvince
		} else {
			query = price.PriceGetCompareBySulselCityHistory
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
			"provinceId":  queryParams.ProvinceId,
			"cityId":      queryParams.CityId,
			"startDate":   queryParams.StartDate,
			"endDate":     queryParams.EndDate,
		},
	).Scan(
		&result.Commodity,
		&result.City,
		&result.PriceDiff,
	)
	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetCompareProvinceCommodityHistory(queryParams domain.PriceGetCompareProvinceCommodityHistoryParams) (interface{}, error) {
	var result models.PriceTableCompareProvinceCommodityHistory
	var rows pgx.Rows
	var err error
	var query string
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetCompareBySulselCommodityHistoryHigherAvgChild
		case "same":
			query = price.PriceGetCompareBySulselCommodityHistorySameAvgChild
		case "lower":
			query = price.PriceGetCompareBySulselCommodityHistoryLowerAvgChild
		}
	} else {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetCompareBySulselCommodityHistoryHigher
		case "same":
			query = price.PriceGetCompareBySulselCommodityHistorySame
		case "lower":
			query = price.PriceGetCompareBySulselCommodityHistoryLower
		}
	}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   "73",
			"selectedDate": queryParams.SelectedDate,
		},
	)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Commodity,
			&result.PriceDiff,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *PriceRepository) GetCompareByNational(queryParams domain.PriceGetCompareNationalParams) (interface{}, error) {
	var result models.PriceTableCompareNationalMap
	var err error
	var query string
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		query = price.PriceGetCompareByNationalAvgChild
	} else {
		query = price.PriceGetCompareByNational
	}

	rows, err := repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"nationalId":   queryParams.NationalId,
			"selectedDate": queryParams.SelectedDate,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Summary,
			&result.Commodity,
			&result.PriceNational,
			&result.PriceProvince,
			&result.PriceLevel,
			&result.PriceTier,
			&result.PriceTierCode,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *PriceRepository) GetCompareByNationalList(queryParams domain.PriceGetCompareNationalParams) (interface{}, error) {
	var result models.PriceTableCompareNationalList
	var rows pgx.Rows
	var err error
	var sortBy string
	var sortByRank string
	var commodityIsParent bool
	var query string

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if queryParams.PaginationParams.SortBy != "" {
		split := strings.Split(queryParams.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if commodityIsParent {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetCompareByNationalListWithOrderAscAvgChild
				} else {
					query = price.PriceGetCompareByNationalListWithOrderDescAvgChild
				}
			} else {
				query = price.PriceGetCompareByNationalListAvgChild
			}
		} else {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetCompareByNationalListWithOrderAsc
				} else {
					query = price.PriceGetCompareByNationalListWithOrderDesc
				}
			} else {
				query = price.PriceGetCompareByNationalList
			}
		}
	} else {
		if commodityIsParent {
			query = price.PriceGetCompareByNationalListAvgChild
		} else {
			query = price.PriceGetCompareByNationalList
		}
	}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"nationalId":   queryParams.NationalId,
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
			"sortBy":       sortBy,
			"sortByRank":   sortByRank,
		},
	)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Commodity,
			&result.PriceNational,
			&result.PriceProvince,
			&result.CityPrice,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	if result.PriceProvince == nil || result.CityPrice == nil {
		return nil, errors.New("no rows in result set")
	}

	//if result.PriceProvince != nil {
	//	if result.PriceProvince.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.PriceProvince.Province.Assets.AssetsLocation + "/" + result.PriceProvince.Province.Assets.AssetsName
	//		result.PriceProvince.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	//for _, row := range *result.CityPrice {
	//	if row.City.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := row.City.Assets.AssetsLocation + "/" + row.City.Assets.AssetsName
	//		row.City.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *PriceRepository) GetCompareNationalCityHistory(queryParams domain.PriceDiffByCityAndCommodityParams) (interface{}, error) {
	var result models.PriceTableCompareProvinceCityHistory
	var realLastUpdate interface{}
	var dateFormat = "2006-01-02"
	var newDate *common.GenerateNewDate
	var err error
	var query string

	var startDate time.Time
	startDate, err = time.Parse(dateFormat, queryParams.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate time.Time
	endDate, err = time.Parse(dateFormat, queryParams.EndDate)
	if err != nil {
		return nil, err
	}

	//realLastUpdate, err = repo.GetLastUpdate(
	//	domain.PriceLastUpdateParams{
	//		ProvinceId:   "73",
	//		CityId:       queryParams.CityId,
	//		CommodityId:  queryParams.CommodityId,
	//		SelectedDate: queryParams.EndDate,
	//	},
	//)
	//if err != nil {
	//	return nil, err
	//}

	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvinceAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		} else {
			realLastUpdate, err = repo.GetLastUpdateAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		}
	} else {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvince(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		} else {
			realLastUpdate, err = repo.GetLastUpdate(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
			if err != nil {
				return nil, err
			}
		}
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	queryParams.StartDate = newDate.StartDate
	queryParams.EndDate = *lastUpdate

	if commodityIsParent {
		if queryParams.CityId == "false" {
			query = price.PriceGetCompareByNationalCityHistoryProvinceAvgChild
		} else {
			query = price.PriceGetCompareByNationalCityHistoryAvgChild
		}
	} else {
		if queryParams.CityId == "false" {
			query = price.PriceGetCompareByNationalCityHistoryProvince
		} else {
			query = price.PriceGetCompareByNationalCityHistory
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
			"provinceId":  queryParams.ProvinceId,
			"cityId":      queryParams.CityId,
			"startDate":   queryParams.StartDate,
			"endDate":     queryParams.EndDate,
			"nationalId":  1,
		},
	).Scan(
		&result.Commodity,
		&result.City,
		&result.PriceDiff,
	)
	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetCompareNationalCommodityHistory(queryParams domain.PriceGetCompareProvinceCommodityHistoryParams) (interface{}, error) {
	var result models.PriceTableCompareProvinceCommodityHistory
	var err error
	var query string
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetCompareByNationalCommodityHistoryHigherAvgChild
		case "same":
			query = price.PriceGetCompareByNationalCommodityHistorySameAvgChild
		case "lower":
			query = price.PriceGetCompareByNationalCommodityHistoryLowerAvgChild
		}
	} else {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetCompareByNationalCommodityHistoryHigher
		case "same":
			query = price.PriceGetCompareByNationalCommodityHistorySame
		case "lower":
			query = price.PriceGetCompareByNationalCommodityHistoryLower
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   73,
			"selectedDate": queryParams.SelectedDate,
			"nationalId":   1,
		},
	).Scan(
		&result.Commodity,
		&result.PriceDiff,
	)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) Get(params domain.PriceListRepoParams) (interface{}, error) {
	var results models.PriceListByCommodityAndCityResponse
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

	//realLastUpdate, err = repo.GetLastUpdate(
	//	domain.PriceLastUpdateParams{
	//		ProvinceId:   "73",
	//		CityId:       params.CityId,
	//		CommodityId:  params.CommodityId,
	//		SelectedDate: params.EndDate,
	//	},
	//)
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

	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		realLastUpdate, err = repo.GetLastUpdateProvinceAvgChild(
			domain.PriceLastUpdateParams{
				ProvinceId:   "73",
				CommodityId:  params.CommodityId,
				SelectedDate: params.EndDate,
			},
		)
		if err != nil {
			return nil, err
		}
	} else {
		realLastUpdate, err = repo.GetLastUpdateProvince(
			domain.PriceLastUpdateParams{
				ProvinceId:   "73",
				CommodityId:  params.CommodityId,
				SelectedDate: params.EndDate,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	params.StartDate = newDate.StartDate
	params.EndDate = *lastUpdate

	if commodityIsParent {
		err = repo.Db.QueryRow(
			context.Background(),
			queries.PriceGetListByCommodityIdAndCityIdAvgChild,
			pgx.NamedArgs{
				"page":        params.PaginationParams.Page,
				"limit":       params.PaginationParams.Limit,
				"cityId":      params.CityId,
				"commodityId": params.CommodityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&results.Commodities,
			&results.City,
			&results.PriceDiff,
			&results.PriceDiffFormat,
			&results.PriceList,
		)
		if err != nil {
			log.Error("QueryErrors Get:", err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			queries.PriceGetListByCommodityIdAndCityId,
			pgx.NamedArgs{
				"page":        params.PaginationParams.Page,
				"limit":       params.PaginationParams.Limit,
				"cityId":      params.CityId,
				"commodityId": params.CommodityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&results.Commodities,
			&results.City,
			&results.PriceDiff,
			&results.PriceDiffFormat,
			&results.PriceList,
		)
		if err != nil {
			log.Error("QueryErrors Get:", err)
			return nil, err
		}
	}

	return &results, nil
}

func (repo *PriceRepository) GetCount(params domain.PriceListRepoParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.PriceGetListByCommodityIdAndCityIdCount,
		pgx.NamedArgs{
			"cityId":      params.CityId,
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

func (repo *PriceRepository) GetForProvince(params domain.PriceListRepoParams) (interface{}, error) {
	var results models.PriceListByCommodityAndCityResponse
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

	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		realLastUpdate, err = repo.GetLastUpdateProvinceAvgChild(
			domain.PriceLastUpdateParams{
				ProvinceId:   "73",
				CommodityId:  params.CommodityId,
				SelectedDate: params.EndDate,
			},
		)
		if err != nil {
			return nil, err
		}
	} else {
		realLastUpdate, err = repo.GetLastUpdateProvince(
			domain.PriceLastUpdateParams{
				ProvinceId:   "73",
				CommodityId:  params.CommodityId,
				SelectedDate: params.EndDate,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	params.StartDate = newDate.StartDate
	params.EndDate = *lastUpdate

	if commodityIsParent {
		err = repo.Db.QueryRow(
			context.Background(),
			queries.PriceGetListByCommodityIdAndProvinceIdAvgChild,
			pgx.NamedArgs{
				"page":        params.PaginationParams.Page,
				"limit":       params.PaginationParams.Limit,
				"provinceId":  params.ProvinceId,
				"commodityId": params.CommodityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&results.Commodities,
			&results.City,
			&results.PriceDiff,
			&results.PriceDiffFormat,
			&results.PriceList,
		)
		if err != nil {
			log.Error("QueryErrors Get:", err)
			return nil, err
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			queries.PriceGetListByCommodityIdAndProvinceId,
			pgx.NamedArgs{
				"page":        params.PaginationParams.Page,
				"limit":       params.PaginationParams.Limit,
				"provinceId":  params.ProvinceId,
				"commodityId": params.CommodityId,
				"startDate":   params.StartDate,
				"endDate":     params.EndDate,
			},
		).Scan(
			&results.Commodities,
			&results.City,
			&results.PriceDiff,
			&results.PriceDiffFormat,
			&results.PriceList,
		)
		if err != nil {
			log.Error("QueryErrors Get:", err)
			return nil, err
		}
	}

	return &results, nil
}

func (repo *PriceRepository) GetCountProvince(params domain.PriceListRepoParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.PriceGetListByCommodityIdAndProvinceIdCount,
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

func (repo *PriceRepository) GetMtm(queryParams domain.PricePerubahanParams) (interface{}, error) {
	var result models.PriceMtmResponse
	var commodityIsParent bool
	var err error

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		var priceExists bool
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceExists,
			pgx.NamedArgs{
				"commodityId": queryParams.CommodityId,
				"startDate":   queryParams.SelectedDate,
				"endDate":     queryParams.SelectedDate,
			},
		).Scan(&priceExists)

		if !priceExists {
			err = repo.Db.QueryRow(
				context.Background(),
				price.PriceGetMtmAvgChild,
				pgx.NamedArgs{
					"commodityId":  queryParams.CommodityId,
					"provinceId":   queryParams.ProvinceId,
					"selectedDate": queryParams.SelectedDate,
				},
			).Scan(
				&result.Summary,
				&result.Commodity,
				&result.PriceProvince,
				&result.PriceLevel,
				&result.PriceTier,
				&result.PriceTierCode,
			)
			if err != nil {
				log.Error("QueryErrors GetCompareBySulsel:", err)
				return nil, err
			}
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				price.PriceGetMtm,
				pgx.NamedArgs{
					"commodityId":  queryParams.CommodityId,
					"provinceId":   queryParams.ProvinceId,
					"selectedDate": queryParams.SelectedDate,
				},
			).Scan(
				&result.Summary,
				&result.Commodity,
				&result.PriceProvince,
				&result.PriceLevel,
				&result.PriceTier,
				&result.PriceTierCode,
			)
			if err != nil {
				log.Error("QueryErrors GetCompareBySulsel:", err)
				return nil, err
			}
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceGetMtm,
			pgx.NamedArgs{
				"commodityId":  queryParams.CommodityId,
				"provinceId":   queryParams.ProvinceId,
				"selectedDate": queryParams.SelectedDate,
			},
		).Scan(
			&result.Summary,
			&result.Commodity,
			&result.PriceProvince,
			&result.PriceLevel,
			&result.PriceTier,
			&result.PriceTierCode,
		)
		if err != nil {
			log.Error("QueryErrors GetCompareBySulsel:", err)
			return nil, err
		}
	}

	return &result, nil
}

func (repo *PriceRepository) GetMtmList(queryParams domain.PricePerubahanParams) (interface{}, error) {
	var err error
	var result models.PriceMtmListResponse
	var commodityIsParent bool
	var sortBy string
	var sortByRank string
	var query string

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if queryParams.PaginationParams.SortBy != "" {
		split := strings.Split(queryParams.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if commodityIsParent {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetMtmListWithOrderAvgChild
				} else {
					query = price.PriceGetMtmListWithOrderDescAvgChild
				}
			} else {
				query = price.PriceGetMtmListAvgChild
			}
		} else {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetMtmListWithOrder
				} else {
					query = price.PriceGetMtmListWithOrderDesc
				}
			} else {
				query = price.PriceGetMtmList
			}
		}
	} else {
		if commodityIsParent {
			query = price.PriceGetMtmListAvgChild
		} else {
			query = price.PriceGetMtmList
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"nationalId":   1,
			"provinceId":   73,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
			"sortBy":       sortBy,
			"sortByRank":   sortByRank,
		},
	).Scan(
		&result.Commodity,
		&result.PriceProvince,
		&result.PriceCity,
	)
	if err != nil {
		log.Error("err")
		log.Error(err)
		return nil, err
	}

	//for rows.Next() {
	//	err = rows.Scan(
	//		&result.Commodity,
	//		&result.PriceProvince,
	//		&result.PriceCity,
	//	)
	//	if err != nil {
	//		fmt.Printf("Scan error: %v\n", err)
	//	}
	//}

	//if result.PriceProvince != nil {
	//	if result.PriceProvince.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.PriceProvince.Province.Assets.AssetsLocation + "/" + result.PriceProvince.Province.Assets.AssetsName
	//		result.PriceProvince.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *PriceRepository) GetMtmCityHistory(queryParams domain.PriceMtmCityHistoryParams) (interface{}, error) {
	var result models.PriceMtmProvinceHistoryResponse
	var err error
	var realLastUpdate interface{}
	var dateFormat = "2006-01-02"
	var newDate *common.GenerateNewDate
	var query string
	var commodityIsParent bool

	var startDate time.Time
	startDate, err = time.Parse(dateFormat, queryParams.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate time.Time
	endDate, err = time.Parse(dateFormat, queryParams.EndDate)
	if err != nil {
		return nil, err
	}

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvinceAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		} else {
			realLastUpdate, err = repo.GetLastUpdateAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		}
	} else {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvince(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		} else {
			realLastUpdate, err = repo.GetLastUpdate(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		}
	}
	if err != nil {
		return nil, err
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	queryParams.StartDate = newDate.StartDate
	queryParams.EndDate = *lastUpdate

	if commodityIsParent {
		if queryParams.CityId == "false" {
			query = price.PriceGetMtmCityHistoryProvinceAvgChild
		} else {
			query = price.PriceGetMtmCityHistoryAvgChild
		}
	} else {
		if queryParams.CityId == "false" {
			query = price.PriceGetMtmCityHistoryProvince
		} else {
			query = price.PriceGetMtmCityHistory
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"cityId":      queryParams.CityId,
			"provinceId":  queryParams.ProvinceId,
			"commodityId": queryParams.CommodityId,
			"startDate":   queryParams.PeriodDateParam.StartDate,
			"endDate":     queryParams.PeriodDateParam.EndDate,
		},
	).Scan(
		&result.Commodity,
		&result.Province,
		&result.CommodityInflations,
	)
	if err != nil {
		log.Error("QueryErrors GetMtmCityHistory:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetMtmCommodityHistory(queryParams domain.PriceMtmCommodityHistoryParams) (interface{}, error) {
	var result models.PriceMtmCommodityHistoryResponse
	var err error
	var query string
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetMtmCommodityHistoryHigherAvgChild
		case "same":
			query = price.PriceGetMtmCommodityHistorySameAvgChild
		case "lower":
			query = price.PriceGetMtmCommodityHistoryLowerAvgChild
		}
	} else {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetMtmCommodityHistoryHigher
		case "same":
			query = price.PriceGetMtmCommodityHistorySame
		case "lower":
			query = price.PriceGetMtmCommodityHistoryLower
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   73,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(
		&result.Commodity,
		&result.CityInflations,
	)
	if err != nil {
		log.Error("QueryErrors GetMtmCityHistory:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetYtd(queryParams domain.PricePerubahanParams) (interface{}, error) {
	var result models.PriceMtmResponse
	var commodityIsParent bool
	var query string
	var err error

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		var priceExists bool
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceExists,
			pgx.NamedArgs{
				"commodityId": queryParams.CommodityId,
				"startDate":   queryParams.SelectedDate,
				"endDate":     queryParams.SelectedDate,
			},
		).Scan(&priceExists)

		if !priceExists {
			query = price.PriceGetYtdAvgChild
		} else {
			query = price.PriceGetYtd
		}
	} else {
		query = price.PriceGetYtd
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(
		&result.Summary,
		&result.Commodity,
		&result.PriceProvince,
		&result.PriceLevel,
		&result.PriceTier,
		&result.PriceTierCode,
	)
	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetYtdList(queryParams domain.PricePerubahanParams) (interface{}, error) {
	var result models.PriceMtmListResponse
	var commodityIsParent bool
	var query string
	var err error
	var rows pgx.Rows
	var sortBy string
	var sortByRank string

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if queryParams.PaginationParams.SortBy != "" {
		split := strings.Split(queryParams.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if commodityIsParent {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetYtdListWithOrderAvgChild
				} else {
					query = price.PriceGetYtdListWithOrderDescAvgChild
				}
			} else {
				query = price.PriceGetYtdListAvgChild
			}
		} else {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetYtdListWithOrder
				} else {
					query = price.PriceGetYtdListWithOrderDesc
				}
			} else {
				query = price.PriceGetYtdList
			}
		}
	} else {
		if commodityIsParent {
			query = price.PriceGetYtdListAvgChild
		} else {
			query = price.PriceGetYtdList
		}
	}

	//if commodityIsParent {
	//	query = price.PriceGetYtdListAvgChild
	//} else {
	//	query = price.PriceGetYtdList
	//}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
			"sortBy":       sortBy,
			"sortByRank":   sortByRank,
		},
	)
	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Commodity,
			&result.PriceProvince,
			&result.PriceCity,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	//if result.PriceProvince != nil {
	//	if result.PriceProvince.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.PriceProvince.Province.Assets.AssetsLocation + "/" + result.PriceProvince.Province.Assets.AssetsName
	//		result.PriceProvince.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *PriceRepository) GetYtdCityHistory(queryParams domain.PriceMtmCityHistoryParams) (interface{}, error) {
	var result models.PriceMtmProvinceHistoryResponse
	var err error
	var realLastUpdate interface{}
	var dateFormat = "2006-01-02"
	var newDate *common.GenerateNewDate
	var query string
	var commodityIsParent bool

	var startDate time.Time
	startDate, err = time.Parse(dateFormat, queryParams.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate time.Time
	endDate, err = time.Parse(dateFormat, queryParams.EndDate)
	if err != nil {
		return nil, err
	}

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvinceAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		} else {
			realLastUpdate, err = repo.GetLastUpdateAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		}
	} else {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvince(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		} else {
			realLastUpdate, err = repo.GetLastUpdate(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		}
	}
	if err != nil {
		return nil, err
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	queryParams.StartDate = newDate.StartDate
	queryParams.EndDate = *lastUpdate

	if commodityIsParent {
		if queryParams.CityId == "false" {
			query = price.PriceGetYtdCityHistoryProvinceAvgChild
		} else {
			query = price.PriceGetYtdCityHistoryAvgChild
		}
	} else {
		if queryParams.CityId == "false" {
			query = price.PriceGetYtdCityHistoryProvince
		} else {
			query = price.PriceGetYtdCityHistory
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"cityId":      queryParams.CityId,
			"provinceId":  queryParams.ProvinceId,
			"commodityId": queryParams.CommodityId,
			"startDate":   queryParams.PeriodDateParam.StartDate,
			"endDate":     queryParams.PeriodDateParam.EndDate,
		},
	).Scan(
		&result.Commodity,
		&result.Province,
		&result.CommodityInflations,
	)
	if err != nil {
		log.Error("QueryErrors GetMtmCityHistory:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetYtdCommodityHistory(queryParams domain.PriceMtmCommodityHistoryParams) (interface{}, error) {
	var result models.PriceMtmCommodityHistoryResponse
	var query string
	var err error
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetYtdCommodityHistoryHigherAvgChild
		case "same":
			query = price.PriceGetYtdCommodityHistorySameAvgChild
		case "lower":
			query = price.PriceGetYtdCommodityHistoryLowerAvgChild
		}
	} else {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetYtdCommodityHistoryHigher
		case "same":
			query = price.PriceGetYtdCommodityHistorySame
		case "lower":
			query = price.PriceGetYtdCommodityHistoryLower
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   73,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(
		&result.Commodity,
		&result.CityInflations,
	)
	if err != nil {
		log.Error("QueryErrors GetMtmCityHistory:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetYty(queryParams domain.PricePerubahanParams) (interface{}, error) {
	var result models.PriceMtmResponse
	var commodityIsParent bool
	var query string
	var err error

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		var priceExists bool
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceExists,
			pgx.NamedArgs{
				"commodityId": queryParams.CommodityId,
				"startDate":   queryParams.SelectedDate,
				"endDate":     queryParams.SelectedDate,
			},
		).Scan(&priceExists)

		if !priceExists {
			query = price.PriceGetYtyAvgChild
		} else {
			query = price.PriceGetYty
		}
	} else {
		query = price.PriceGetYty
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(
		&result.Summary,
		&result.Commodity,
		&result.PriceProvince,
		&result.PriceLevel,
		&result.PriceTier,
		&result.PriceTierCode,
	)
	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}
	//rows, err := repo.Db.Query(
	//	context.Background(),
	//	price.PriceGetYty,
	//	pgx.NamedArgs{
	//		"commodityId":  queryParams.CommodityId,
	//		"provinceId":   queryParams.ProvinceId,
	//		"selectedDate": queryParams.SelectedDate,
	//	},
	//)
	//
	//if err != nil {
	//	log.Error("QueryErrors GetCompareBySulsel:", err)
	//	return nil, err
	//}
	//
	//for rows.Next() {
	//	err = rows.Scan(
	//		&result.Summary,
	//		&result.Commodity,
	//		&result.PriceProvince,
	//		&result.PriceLevel,
	//		&result.PriceTier,
	//		&result.PriceTierCode,
	//	)
	//	if err != nil {
	//		fmt.Printf("Scan error: %v\n", err)
	//	}
	//}
	//
	return &result, nil
}

func (repo *PriceRepository) GetYtyList(queryParams domain.PricePerubahanParams) (interface{}, error) {
	var result models.PriceMtmListResponse
	var commodityIsParent bool
	var query string
	var err error
	var rows pgx.Rows
	var sortBy string
	var sortByRank string

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if queryParams.PaginationParams.SortBy != "" {
		split := strings.Split(queryParams.PaginationParams.SortBy, ":")
		if len(split) > 1 {
			sortBy = split[0]
			sortByRank = split[1]
		} else {
			sortBy = split[0]
		}

		if commodityIsParent {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetYtyListWithOrderAvgChild
				} else {
					query = price.PriceGetYtyListWithOrderDescAvgChild
				}
			} else {
				query = price.PriceGetYtyListAvgChild
			}
		} else {
			if sortByRank != "" {
				if sortByRank == "asc" {
					query = price.PriceGetYtyListWithOrder
				} else {
					query = price.PriceGetYtyListWithOrderDesc
				}
			} else {
				query = price.PriceGetYtyList
			}
		}
	} else {
		if commodityIsParent {
			query = price.PriceGetYtyListAvgChild
		} else {
			query = price.PriceGetYtyList
		}
	}

	//if commodityIsParent {
	//	query = price.PriceGetYtyListAvgChild
	//} else {
	//	query = price.PriceGetYtyList
	//}

	rows, err = repo.Db.Query(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   queryParams.ProvinceId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
			"sortBy":       sortBy,
			"sortByRank":   sortByRank,
		},
	)

	for rows.Next() {
		err = rows.Scan(
			&result.Commodity,
			&result.PriceProvince,
			&result.PriceCity,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	//rows, err := repo.Db.Query(
	//	context.Background(),
	//	price.PriceGetYtyList,
	//	pgx.NamedArgs{
	//		"commodityId":  queryParams.CommodityId,
	//		"provinceId":   queryParams.ProvinceId,
	//		"selectedDate": queryParams.SelectedDate,
	//	},
	//)
	//
	//if err != nil {
	//	log.Error("QueryErrors GetCompareBySulsel:", err)
	//	return nil, err
	//}
	//
	//for rows.Next() {
	//	err = rows.Scan(
	//		&result.Commodity,
	//		&result.PriceProvince,
	//		&result.PriceCity,
	//	)
	//	if err != nil {
	//		fmt.Printf("Scan error: %v\n", err)
	//	}
	//}

	//if result.PriceProvince != nil {
	//	if result.PriceProvince.Province.Assets.AssetsLocationType == common.ASSETS_TYPE_DIRECTORY {
	//		//baseUrl := os.Getenv("BASE_URL")
	//		baseUrl := "https://project.bi.sentech.id/api/v1/stg"
	//		assetsLocation := result.PriceProvince.Province.Assets.AssetsLocation + "/" + result.PriceProvince.Province.Assets.AssetsName
	//		result.PriceProvince.Province.Assets.AssetsUrl = baseUrl + "/assets?assets_location=" + assetsLocation
	//	}
	//}

	return &result, nil
}

func (repo *PriceRepository) GetYtyCityHistory(queryParams domain.PriceMtmCityHistoryParams) (interface{}, error) {
	var result models.PriceMtmProvinceHistoryResponse
	var err error
	var realLastUpdate interface{}
	var dateFormat = "2006-01-02"
	var newDate *common.GenerateNewDate
	var query string
	var commodityIsParent bool

	var startDate time.Time
	startDate, err = time.Parse(dateFormat, queryParams.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate time.Time
	endDate, err = time.Parse(dateFormat, queryParams.EndDate)
	if err != nil {
		return nil, err
	}

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvinceAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		} else {
			realLastUpdate, err = repo.GetLastUpdateAvgChild(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		}
	} else {
		if queryParams.CityId == "false" {
			realLastUpdate, err = repo.GetLastUpdateProvince(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		} else {
			realLastUpdate, err = repo.GetLastUpdate(
				domain.PriceLastUpdateParams{
					ProvinceId:   "73",
					CityId:       queryParams.CityId,
					CommodityId:  queryParams.CommodityId,
					SelectedDate: queryParams.EndDate,
				},
			)
		}
	}
	if err != nil {
		return nil, err
	}

	lastUpdate := realLastUpdate.(*string)
	newDate, err = generateNewDate(startDate, endDate, lastUpdate)
	if err != nil {
		return nil, err
	}

	queryParams.StartDate = newDate.StartDate
	queryParams.EndDate = *lastUpdate

	if commodityIsParent {
		if queryParams.CityId == "false" {
			query = price.PriceGetYtyCityHistoryProvinceAvgChild
		} else {
			query = price.PriceGetYtyCityHistoryAvgChild
		}
	} else {
		if queryParams.CityId == "false" {
			query = price.PriceGetYtyCityHistoryProvince
		} else {
			query = price.PriceGetYtyCityHistory
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"cityId":      queryParams.CityId,
			"provinceId":  queryParams.ProvinceId,
			"commodityId": queryParams.CommodityId,
			"startDate":   queryParams.PeriodDateParam.StartDate,
			"endDate":     queryParams.PeriodDateParam.EndDate,
		},
	).Scan(
		&result.Commodity,
		&result.Province,
		&result.CommodityInflations,
	)
	if err != nil {
		log.Error("QueryErrors GetMtmCityHistory:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetYtyCommodityHistory(queryParams domain.PriceMtmCommodityHistoryParams) (interface{}, error) {
	var result models.PriceMtmCommodityHistoryResponse
	var query string
	var err error
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": queryParams.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetYtyCommodityHistoryHigherAvgChild
		case "same":
			query = price.PriceGetYtyCommodityHistorySameAvgChild
		case "lower":
			query = price.PriceGetYtyCommodityHistoryLowerAvgChild
		}
	} else {
		switch queryParams.Status {
		case "higher":
			query = price.PriceGetYtyCommodityHistoryHigher
		case "same":
			query = price.PriceGetYtyCommodityHistorySame
		case "lower":
			query = price.PriceGetYtyCommodityHistoryLower
		}
	}

	err = repo.Db.QueryRow(
		context.Background(),
		query,
		pgx.NamedArgs{
			"commodityId":  queryParams.CommodityId,
			"provinceId":   73,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(
		&result.Commodity,
		&result.CityInflations,
	)
	if err != nil {
		log.Error("QueryErrors GetMtmCityHistory:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetPriceReport(params domain.PriceListRepoParams) (interface{}, error) {
	var result []map[string]string
	var data []models.ReportPriceModel
	var rows pgx.Rows
	var err error

	if params.CommodityId != "" {
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
				queries.ReportPriceByCommodityAvgChild,
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
				queries.ReportPriceByCommodity,
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
			queries.ReportPriceByCity,
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
		var reportPriceData models.ReportPriceModel
		err = rows.Scan(
			&reportPriceData.Title,
			&reportPriceData.Prices,
		)

		data = append(data, reportPriceData)
	}

	i := 0
	for _, v := range data {
		result = append(result, map[string]string{
			"title": v.Title,
		})

		i++
	}

	var response []map[string]string
	idx := 0
	for _, v := range result {
		if v["title"] == data[idx].Title {
			for _, pricesArrValue := range data[idx].Prices {
				for priceMapKey, priceMapValue := range pricesArrValue {
					v[priceMapKey] = priceMapValue
				}
			}
		}
		response = append(response, v)
		idx++
	}

	return &response, nil
}

func (repo *PriceRepository) GetPriceReportDownload(params domain.PriceListRepoParams) (interface{}, error) {
	var result []map[string]string
	var data []models.ReportPriceModel
	var rows pgx.Rows
	var err error

	if params.CommodityId != "" {
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
				queries.ReportPriceByCommodityAvgChildDownload,
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
				queries.ReportPriceByCommodityDownload,
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
			queries.ReportPriceByCityDownload,
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
		var reportPriceData models.ReportPriceModel
		err = rows.Scan(
			&reportPriceData.Title,
			&reportPriceData.Prices,
		)

		data = append(data, reportPriceData)
	}

	if data == nil {
		//responseEmpty := make(map[string]interface{})
		return nil, nil
	}

	i := 0
	for _, v := range data {
		result = append(result, map[string]string{
			"title": v.Title,
		})

		i++
	}

	var response []map[string]string
	idx := 0
	for _, v := range result {
		if v["title"] == data[idx].Title {
			for _, pricesArrValue := range data[idx].Prices {
				for priceMapKey, priceMapValue := range pricesArrValue {
					v[priceMapKey] = priceMapValue
				}
			}
		}
		response = append(response, v)
		idx++
	}

	return &response, nil
}

func (repo *PriceRepository) GetPriceReportCount(params domain.PriceListRepoParams) (*int, error) {
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

	if params.CommodityId != "" {
		if commodityIsParent {
			err = repo.Db.QueryRow(
				context.Background(),
				queries.ReportPriceByCommodityCountAvgChild,
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
				queries.ReportPriceByCommodityCount,
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
			queries.ReportPriceByCityCount,
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

func (repo *PriceRepository) GetPriceDiffLevelHarga(queryParams domain.PriceGetRepoParamsNew) (interface{}, error) {
	var result models.PriceDiffLevelHarga
	rows, err := repo.Db.Query(
		context.Background(),
		price.PriceDiffLevelHarga,
		pgx.NamedArgs{
			"provinceId":   73,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
		},
	)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.PriceMin,
			&result.PriceMax,
			&result.PriceDiff,
			&result.PriceBarCategory,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *PriceRepository) GetPriceDiffLevelHargaAvgChild(queryParams domain.PriceGetRepoParamsNew) (interface{}, error) {
	var result models.PriceDiffLevelHarga
	rows, err := repo.Db.Query(
		context.Background(),
		price.PriceDiffLevelHargaAvgChild,
		pgx.NamedArgs{
			"provinceId":   73,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
			"page":         queryParams.PaginationParams.Page,
			"limit":        queryParams.PaginationParams.Limit,
		},
	)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.PriceMin,
			&result.PriceMax,
			&result.PriceDiff,
			&result.PriceBarCategory,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}

func (repo *PriceRepository) GetLastUpdate(queryParams domain.PriceLastUpdateParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceGetLatestDateAvail,
		pgx.NamedArgs{
			"provinceId":   queryParams.ProvinceId,
			"cityId":       queryParams.CityId,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(&result)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetLastUpdateAvgChild(queryParams domain.PriceLastUpdateParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceGetLatestDateAvailAvgChild,
		pgx.NamedArgs{
			"provinceId":   queryParams.ProvinceId,
			"cityId":       queryParams.CityId,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(&result)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetLastUpdateProvince(queryParams domain.PriceLastUpdateParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceGetLatestDateProvinceAvail,
		pgx.NamedArgs{
			"provinceId":   queryParams.ProvinceId,
			"cityId":       queryParams.CityId,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(&result)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetLastUpdateProvinceAvgChild(queryParams domain.PriceLastUpdateParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceGetLatestDateProvinceAvailAvgChild,
		pgx.NamedArgs{
			"provinceId":   queryParams.ProvinceId,
			"cityId":       queryParams.CityId,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(&result)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetLastUpdateNational(queryParams domain.PriceLastUpdateParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceGetLatestDateNationalAvail,
		pgx.NamedArgs{
			"nationalId":   queryParams.NationalId,
			"provinceId":   queryParams.ProvinceId,
			"cityId":       queryParams.CityId,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(&result)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) GetLastUpdateNationalAvgChild(queryParams domain.PriceLastUpdateParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceGetLatestDateNationalAvailAvgChild,
		pgx.NamedArgs{
			"nationalId":   queryParams.NationalId,
			"provinceId":   queryParams.ProvinceId,
			"cityId":       queryParams.CityId,
			"commodityId":  queryParams.CommodityId,
			"selectedDate": queryParams.SelectedDate,
		},
	).Scan(&result)

	if err != nil {
		log.Error("QueryErrors GetCompareBySulsel:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *PriceRepository) Exist(params domain.PriceListRepoParams) (interface{}, error) {
	var err error
	var result map[string]bool
	var commodityIsParent bool

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"commodityId": params.CommodityId,
		},
	).Scan(&commodityIsParent)

	if commodityIsParent {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceExistParent,
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
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceExist,
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
	}

	return &result, nil
}

func (repo *PriceRepository) LatestDateExist(params domain.PriceListRepoParams) (interface{}, error) {
	var err error
	var result string
	var commodityIsParent bool
	var latestDateExistChild interface{}

	if params.CommodityId != "" {
		err = repo.Db.QueryRow(
			context.Background(),
			queries.TmCommodityIsParent,
			pgx.NamedArgs{
				"commodityId": params.CommodityId,
			},
		).Scan(&commodityIsParent)

		if commodityIsParent {
			latestDateExistChild, err = repo.LatestDateExistChild(params)
		} else {
			err = repo.Db.QueryRow(
				context.Background(),
				price.PriceLatestDateExist,
				pgx.NamedArgs{
					"provinceId":  params.ProvinceId,
					"commodityId": params.CommodityId,
				},
			).Scan(&result)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err = repo.Db.QueryRow(
			context.Background(),
			price.PriceLatestDateExistWithoutCommodity,
			pgx.NamedArgs{
				"provinceId": params.ProvinceId,
			},
		).Scan(&result)
		if err != nil {
			return nil, err
		}
	}

	if latestDateExistChild != nil {
		return latestDateExistChild.(*string), nil
	}

	return &result, nil
}

func (repo *PriceRepository) LatestDateExistChild(params domain.PriceListRepoParams) (interface{}, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		price.PriceLatestDateExistChild,
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

func generateNewDate(startDate time.Time, endDate time.Time, lastUpdate *string) (*common.GenerateNewDate, error) {
	var err error
	var lastUpdateTime time.Time
	var dateFormat = "2006-01-02"
	var params = common.GenerateNewDate{}

	lastUpdateTime, err = time.Parse(dateFormat, *lastUpdate)
	if err != nil {
		return nil, err
	}

	subForStartDate := startDate.Sub(endDate)
	newStartDate := lastUpdateTime.Add(subForStartDate)
	params.StartDate = newStartDate.Format(dateFormat)
	params.EndDate = lastUpdateTime.Format(dateFormat)

	return &params, nil
}

func generateNewDateNeraca(startDate time.Time, endDate time.Time, lastUpdate *string) (*common.GenerateNewDate, error) {
	var err error
	var lastUpdateTime time.Time
	var dateFormat = "2006-01-02"
	var dateDiff = 0
	var params = common.GenerateNewDate{}

	lastUpdateTime, err = time.Parse(dateFormat, *lastUpdate)
	if err != nil {
		return nil, err
	}

	startMonth, _ := strconv.Atoi(common.MonthWithLeadingZero(int(startDate.Month())))
	endMonth, _ := strconv.Atoi(common.MonthWithLeadingZero(int(endDate.Month())))

	if startDate.Year() != endDate.Year() {
		dateDiff = endDate.Year() - startDate.Year()
		params.StartDate = lastUpdateTime.AddDate(-dateDiff, 0, 0).Format(dateFormat)
	} else if startDate.Month() != endDate.Month() {
		dateDiff = endMonth - startMonth
		params.StartDate = lastUpdateTime.AddDate(0, -dateDiff, 0).Format(dateFormat)
	}

	return &params, nil
}

func generatePriceTier(priceDiffLevelHarga *models.PriceDiffLevelHarga, tiers *[]models.PriceTier) {
	if *priceDiffLevelHarga.PriceDiff > 0 {
		i := 1
		var currentPrice = int32(0)
		var nextPrice = int32(0)

		for i <= 5 {
			if i == 1 {
				currentPrice = *priceDiffLevelHarga.PriceMin + *priceDiffLevelHarga.PriceBarCategory
				*tiers = append(*tiers, models.PriceTier{
					Title:                "Sangat Rendah",
					PriceMin:             *priceDiffLevelHarga.PriceMin,
					PriceMinRupiahFormat: common.ThousandFormat(*priceDiffLevelHarga.PriceMin),
					PriceMax:             currentPrice,
					PriceMaxRupiahFormat: common.ThousandFormat(currentPrice),
					Color:                "#208245",
				})
				nextPrice = currentPrice + 1
			} else if i == 2 {
				currentPrice = nextPrice
				nextPrice = currentPrice + (*priceDiffLevelHarga.PriceBarCategory - 1)
				*tiers = append(*tiers, models.PriceTier{
					Title:                "Rendah",
					PriceMin:             currentPrice,
					PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
					PriceMax:             nextPrice,
					PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
					Color:                "#27AE65",
				})
				nextPrice++
			} else if i == 3 {
				currentPrice = nextPrice
				nextPrice = currentPrice + (*priceDiffLevelHarga.PriceBarCategory - 1)
				*tiers = append(*tiers, models.PriceTier{
					Title:                "Sedang",
					PriceMin:             currentPrice,
					PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
					PriceMax:             nextPrice,
					PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
					Color:                "#FFFF00",
				})
				nextPrice++
			} else if i == 4 {
				currentPrice = nextPrice
				nextPrice = currentPrice + (*priceDiffLevelHarga.PriceBarCategory - 1)
				*tiers = append(*tiers, models.PriceTier{
					Title:                "Tinggi",
					PriceMin:             currentPrice,
					PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
					PriceMax:             nextPrice,
					PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
					Color:                "#C0392B",
				})
				nextPrice++
			} else if i == 5 {
				currentPrice = nextPrice
				nextPrice = currentPrice + (*priceDiffLevelHarga.PriceBarCategory - 1)
				*tiers = append(*tiers, models.PriceTier{
					Title:                "Sangat Tinggi",
					PriceMin:             currentPrice,
					PriceMinRupiahFormat: common.ThousandFormat(currentPrice),
					PriceMax:             nextPrice,
					PriceMaxRupiahFormat: common.ThousandFormat(nextPrice),
					Color:                "#7C2019",
				})
			}

			i++
		}
	} else {
		*tiers = append(*tiers, models.PriceTier{
			Title:                "Sangat Tinggi",
			PriceMin:             *priceDiffLevelHarga.PriceMin,
			PriceMinRupiahFormat: common.ThousandFormat(*priceDiffLevelHarga.PriceMin),
			PriceMax:             *priceDiffLevelHarga.PriceMax,
			PriceMaxRupiahFormat: common.ThousandFormat(*priceDiffLevelHarga.PriceMax),
			Color:                "#7C2019",
		})
	}
}

func (repo *PriceRepository) GetHistory(queryParams domain.HistoryRequestParams) (interface{}, error) {
	var results []models.TxFileUploadHistory
	rows, err := repo.Db.Query(
		context.Background(),
		price.HistoryQuery,
		pgx.NamedArgs{
			"module": queryParams.Module,
			"page":   queryParams.PaginationParams.Page,
			"limit":  queryParams.PaginationParams.Limit,
			"search": "%" + *queryParams.Search + "%",
		},
	)

	for rows.Next() {
		var data models.TxFileUploadHistory
		err = rows.Scan(
			&data.Id.Id,
			&data.FileName,
			&data.RowTotal,
			&data.Status,
			&data.ModuleType,
			&data.Errors,
			&data.AuditRail.CreatedAt,
			&data.AuditRail.UpdatedAt,
			&data.AuditRail.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}

	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (repo *PriceRepository) GetCountHistory(params domain.HistoryRequestParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		price.HistoryCountQuery,
		pgx.NamedArgs{
			"module": params.Module,
			"search": "%" + *params.Search + "%",
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
