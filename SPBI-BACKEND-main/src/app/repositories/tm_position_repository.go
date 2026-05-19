package repositories

import (
	"context"
	"fmt"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"sigap-sultan-be/src/common"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TmPositionRepository struct {
	Db *pgxpool.Pool
}

func NewTmPositionRepository(db *pgxpool.Pool) *TmPositionRepository {
	return &TmPositionRepository{Db: db}
}

func (repo TmPositionRepository) Get(paginationParams common.PaginationParams) (interface{}, error) {
	var results []models.TmPositionResponse
	rows, err := repo.Db.Query(
		context.Background(),
		queries.TmPositionList,
		pgx.NamedArgs{
			"page":  paginationParams.Page,
			"limit": paginationParams.Limit,
		},
	)

	fmt.Println("page", paginationParams.Page)
	fmt.Println("limit", paginationParams.Limit)

	for rows.Next() {
		var data models.TmPositionResponse
		err = rows.Scan(
			&data.Id.Id,
			&data.ClientId,
			&data.Name,
			&data.Privileges,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}

	return &results, nil
}

func (repo TmPositionRepository) Count() (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmPositionCount,
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmPositionRepository) GetById(id int) (interface{}, error) {
	var tmPosition models.TmPositionResponse
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmPositionGetById,
		pgx.NamedArgs{
			"id": id,
		},
	).Scan(
		&tmPosition.Id.Id,
		&tmPosition.ClientId,
		&tmPosition.Name,
		&tmPosition.Privileges,
		&tmPosition.CreatedAt,
		&tmPosition.UpdatedAt,
		&tmPosition.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &tmPosition, nil
}

func (repo TmPositionRepository) GetByName(name string) (interface{}, error) {
	var tmPosition models.TmPositionResponse
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmPositionGetByName,
		pgx.NamedArgs{
			"name": "%" + strings.ToLower(name) + "%",
		},
	).Scan(
		&tmPosition.Id.Id,
		&tmPosition.ClientId,
		&tmPosition.Name,
		&tmPosition.Privileges,
		&tmPosition.CreatedAt,
		&tmPosition.UpdatedAt,
		&tmPosition.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &tmPosition, nil
}

func (repo TmPositionRepository) CountByName(name string) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmPositionGetByNameCount,
		pgx.NamedArgs{
			"name": "%" + strings.ToLower(name) + "%",
		},
	).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo TmPositionRepository) Insert(privilegeRequest models.PrivilegesRequest) (interface{}, error) {
	var err error
	var lastInsertedId int
	var tmPositionResponses []*models.TmPositionResponse
	var tx pgx.Tx

	tx, err = repo.Db.Begin(context.Background())
	// var tmPosition models.TmPosition
	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmPositionInsert,
		1,
		privilegeRequest.Position,
		time.Now().Format(common.DATETIME_FORMAT),
	).Scan(&lastInsertedId)
	if err != nil {
		_ = tx.Rollback(context.Background())
		log.Error("QueryErrors:", err)
		return nil, err
	}

	for _, privilege := range privilegeRequest.Privileges {
		if err != nil {
			log.Error("tx.Exec failed: %v\n", err)
			_ = tx.Rollback(context.Background())
			return nil, err
		}

		_, err = tx.Exec(
			context.Background(),
			queries.MapJabatanMenuInsert,
			1,
			lastInsertedId,
			privilege.MenuId,
			privilege.Permissions.Create,
			privilege.Permissions.Read,
			privilege.Permissions.Update,
			privilege.Permissions.Delete,
			time.Now().Format(common.DATETIME_FORMAT),
		)
	}

	if err != nil {
		_ = tx.Rollback(context.Background())
		log.Error("QueryErrors:", err)
		return nil, err
	}

	_ = tx.Commit(context.Background())

	lastTmPosition, errGetById := repo.GetById(lastInsertedId)
	if errGetById != nil {
		_ = tx.Rollback(context.Background())
		return nil, errGetById
	}

	var tmPositionResponse *models.TmPositionResponse
	tmPositionResponse = lastTmPosition.(*models.TmPositionResponse)
	tmPositionResponses = append(tmPositionResponses, tmPositionResponse)

	return &tmPositionResponses, nil
}

func (repo TmPositionRepository) Update(id int, tmPosition models.TmPosition) (interface{}, error) {
	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		log.Error("QueryErrors:", err)
	}

	_, err = tx.Exec(
		context.Background(),
		queries.TmPositionUpdate,
		pgx.NamedArgs{
			"id":         id,
			"name":       tmPosition.Name,
			"created_at": time.Now().Format(common.DATETIME_FORMAT),
		},
	)
	if err != nil {
		log.Error("tx.Exec failed: %v\n", err)
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		log.Error("QueryErrors:", err)
	}

	return &tmPosition, nil
}

func (repo TmPositionRepository) Delete(id int) (interface{}, error) {
	var err error

	result, err := repo.Db.Exec(
		context.Background(),
		queries.TmPositionSoftDelete,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		log.Error("QueryErrors:", err)
		return nil, err
	}

	return &result, nil
}
