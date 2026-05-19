package repositories

import (
	"context"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"sigap-sultan-be/src/common"
	"strings"
	"time"
	_ "time"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TmMenuRepository struct {
	Db *pgxpool.Pool
}

func NewTmMenuRepository(db *pgxpool.Pool) *TmMenuRepository {
	return &TmMenuRepository{Db: db}
}

func (repo TmMenuRepository) Get(paginationParams common.PaginationParams) (interface{}, error) {
	var results []models.TmMenuResponse
	rows, err := repo.Db.Query(
		context.Background(),
		queries.TmMenuList,
		pgx.NamedArgs{
			"page":  paginationParams.Page,
			"limit": paginationParams.Limit,
		},
	)

	for rows.Next() {
		var data models.TmMenuResponse
		err = rows.Scan(
			&data.Id.Id,
			&data.ClientId,
			&data.Name,
			&data.Position,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.DeletedAt,
			&data.Url,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}

	return &results, nil
}

func (repo TmMenuRepository) GetCount() (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmMenuListCount,
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmMenuRepository) GetById(id int) (interface{}, error) {
	var tmMenu models.TmMenuResponse
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmMenuGetById,
		pgx.NamedArgs{
			"id": id,
		},
	).Scan(
		&tmMenu.Id.Id,
		&tmMenu.ClientId,
		&tmMenu.Name,
		&tmMenu.Position,
		&tmMenu.CreatedAt,
		&tmMenu.UpdatedAt,
		&tmMenu.DeletedAt,
		&tmMenu.Url,
	)
	if err != nil {
		return nil, err
	}

	return &tmMenu, nil
}

func (repo TmMenuRepository) GetByName(name string, paginationParams common.PaginationParams) (interface{}, error) {
	var tmMenu models.TmMenuResponse
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmMenuGetByName,
		pgx.NamedArgs{
			"name":  "%" + strings.ToLower(name) + "%",
			"page":  paginationParams.Page,
			"limit": paginationParams.Limit,
		},
	).Scan(
		&tmMenu.Id.Id,
		&tmMenu.ClientId,
		&tmMenu.Name,
		&tmMenu.Position,
		&tmMenu.CreatedAt,
		&tmMenu.UpdatedAt,
		&tmMenu.DeletedAt,
		&tmMenu.Url,
	)
	if err != nil {
		return nil, err
	}

	return &tmMenu, nil
}

func (repo TmMenuRepository) GetByNameCount(name string, paginationParams common.PaginationParams) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmMenuGetByNameCount,
		pgx.NamedArgs{
			"name":  "%" + strings.ToLower(name) + "%",
			"page":  paginationParams.Page,
			"limit": paginationParams.Limit,
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmMenuRepository) Insert(request models.TmMenu) (interface{}, error) {
	var err error
	var lastInsertedId int
	now := time.Now()
	request.CreatedAt = &now

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmMenuInsert,
		request.ClientId,
		request.Name,
		request.Position,
		request.Url,
		request.CreatedAt,
	).Scan(&lastInsertedId)
	if err != nil {
		log.Error("QueryErrors:", err)
		return nil, err
	}

	lastTmMenu, err := repo.GetById(lastInsertedId)

	return &lastTmMenu, nil
}

func (repo TmMenuRepository) Update(id int, request models.TmMenu) (interface{}, error) {
	var err error
	now := time.Now()
	request.ClientId = 1
	request.UpdatedAt = &now

	_, err = repo.Db.Exec(
		context.Background(),
		queries.TmMenuUpdate,
		pgx.NamedArgs{
			"id":        id,
			"clientId":  request.ClientId,
			"name":      request.Name,
			"url":       request.Url,
			"position":  request.Position,
			"updatedAt": request.UpdatedAt,
		},
	)
	if err != nil {
		log.Error("QueryErrors:", err)
		return nil, err
	}

	lastTmMenu, err := repo.GetById(id)

	return &lastTmMenu, nil
}

func (repo TmMenuRepository) Delete(id int) (interface{}, error) {
	var err error
	var lastInsertedId int

	_, err = repo.Db.Exec(
		context.Background(),
		queries.TmMenuDeleteMaster,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		log.Error("QueryErrors:", err)
		return nil, err
	}

	_, err = repo.Db.Exec(
		context.Background(),
		queries.TmMenuDeleteChild,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		log.Error("QueryErrors:", err)
		return nil, err
	}

	lastTmMenu, err := repo.GetById(lastInsertedId)

	return &lastTmMenu, nil
}
