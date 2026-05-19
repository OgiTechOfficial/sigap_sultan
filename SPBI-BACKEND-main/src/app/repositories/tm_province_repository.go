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

type TmProvinceRepository struct {
	Db *pgxpool.Pool
}

func NewTmProvinceRepository(db *pgxpool.Pool) *TmProvinceRepository {
	return &TmProvinceRepository{Db: db}
}

func (repo TmProvinceRepository) Get(params domain.ProvinceRequestParams) (interface{}, error) {
	var results []models.TmProvince
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
			queries.TmProvinceGet,
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
			queries.TmProvinceGet,
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
		var data models.TmProvince
		err = rows.Scan(
			&data.Id.Id,
			&data.Id,
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

func (repo TmProvinceRepository) Count() (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmProvinceGetCount,
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmProvinceRepository) GetByProvince(params domain.ProvinceRequestParams) (interface{}, error) {
	var results []models.TmProvince
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
			queries.TmProvinceGetById,
			pgx.NamedArgs{
				"page":      params.PaginationParams.Page,
				"limit":     params.PaginationParams.Limit,
				"orderBy":   splitSortBy[0],
				"ascending": splitSortBy[1],
				"id":        params.Id,
			},
		)
	} else {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.TmProvinceGetById,
			pgx.NamedArgs{
				"id": params.Id,
			},
		)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var data models.TmProvince
		err = rows.Scan(
			&data.Id.Id,
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

func (repo TmProvinceRepository) CountById(params domain.ProvinceRequestParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmProvinceGetByIdCount,
		pgx.NamedArgs{
			"page":  params.PaginationParams.Page,
			"limit": params.PaginationParams.Limit,
			"id":    params.Id,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmProvinceRepository) GetByName(name string) (*models.TmProvince, error) {
	var data models.TmProvince
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmProvinceGetByName,
		pgx.NamedArgs{
			"name": "%" + strings.ToLower(name) + "%",
		},
	).Scan(
		&data.Id.Id,
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

	return &data, nil
}

func (repo TmProvinceRepository) CountByName(name string) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmProvinceGetByNameCount,
		pgx.NamedArgs{
			"name": "%" + strings.ToLower(name) + "%",
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
