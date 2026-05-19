package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"strings"
)

type TmCityRepository struct {
	Db *pgxpool.Pool
}

func NewTmCityRepository(db *pgxpool.Pool) *TmCityRepository {
	return &TmCityRepository{Db: db}
}

func (repo TmCityRepository) Get(params domain.CityRequestParams) (interface{}, error) {
	var results []models.TmCity
	var splitSortBy []string
	var err error
	var rows pgx.Rows
	sortBy := params.PaginationParams.SortBy

	if sortBy != "" {
		contains := strings.Contains(sortBy, ":")
		if contains {
			splitSortBy = strings.Split(params.PaginationParams.SortBy, ":")
		}
	}

	if len(splitSortBy) > 0 {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.TmCityGet,
			pgx.NamedArgs{
				"page":      params.PaginationParams.Page,
				"limit":     params.PaginationParams.Limit,
				"orderBy":   splitSortBy[0],
				"ascending": splitSortBy[1],
			},
		)
	} else {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.TmCityGet,
			pgx.NamedArgs{
				"page":    params.PaginationParams.Page,
				"limit":   params.PaginationParams.Limit,
				"orderBy": "sequence",
			},
		)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var data models.TmCity
		err = rows.Scan(
			&data.Id.Id,
			&data.ProvinceId,
			&data.Name,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.DeletedAt,
			&data.Sequence,
			&data.AssetsRelationId,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}

	return &results, nil
}

func (repo TmCityRepository) Count() (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCityGetCount,
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmCityRepository) GetByProvince(params domain.CityRequestParams) (interface{}, error) {
	var results []models.TmCity
	var splitSortBy []string
	var err error
	var rows pgx.Rows
	sortBy := params.PaginationParams.SortBy

	if sortBy != "" {
		contains := strings.Contains(sortBy, ":")
		if contains {
			splitSortBy = strings.Split(params.PaginationParams.SortBy, ":")
		}
	}

	if len(splitSortBy) > 0 {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.TmCityGetByProvinceId,
			pgx.NamedArgs{
				"page":       params.PaginationParams.Page,
				"limit":      params.PaginationParams.Limit,
				"orderBy":    splitSortBy[0],
				"ascending":  splitSortBy[1],
				"provinceId": params.ProvinceId,
			},
		)
	} else {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.TmCityGetByProvinceId,
			pgx.NamedArgs{
				"page":       params.PaginationParams.Page,
				"limit":      params.PaginationParams.Limit,
				"provinceId": params.ProvinceId,
			},
		)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var data models.TmCity
		err = rows.Scan(
			&data.Id.Id,
			&data.ProvinceId,
			&data.Name,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.DeletedAt,
			&data.Sequence,
			&data.AssetsRelationId,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}

	return &results, nil
}

func (repo TmCityRepository) CountByProvinceId(params domain.CityRequestParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCityGetByProvinceIdCount,
		pgx.NamedArgs{
			"page":       params.PaginationParams.Page,
			"limit":      params.PaginationParams.Limit,
			"provinceId": params.ProvinceId,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmCityRepository) GetByName(name string) (*models.TmCity, error) {
	var tmCity models.TmCity
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCityGetByName,
		pgx.NamedArgs{
			"name": name,
		},
	).Scan(&tmCity)
	if err != nil {
		return nil, err
	}

	return &tmCity, nil
}

func (repo TmCityRepository) CountByName(name string) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCityGetByNameCount,
		pgx.NamedArgs{
			"name": name,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmCityRepository) GetById(id int) (*models.TmCity, error) {
	var tmCity models.TmCity
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCityGetById,
		pgx.NamedArgs{
			"id": id,
		},
	).Scan(&tmCity)
	if err != nil {
		return nil, err
	}

	return &tmCity, nil
}
