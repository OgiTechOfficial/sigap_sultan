package repositories

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"strings"
)

type AssetsRepository struct {
	Db *pgxpool.Pool
}

func NewAssetsRepository(db *pgxpool.Pool) *AssetsRepository {
	return &AssetsRepository{Db: db}
}

func (repo *AssetsRepository) GetByName(name string) (interface{}, error) {
	var result models.AssetsResponse
	rows, err := repo.Db.Query(
		context.Background(),
		queries.AssetGetSearchByName,
		pgx.NamedArgs{
			"name": "%" + strings.ToLower(name) + "%",
		},
	)

	if err != nil {
		log.Error("repositories.AssetsRepository.GetByName QueryErrors:", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&result.Id.Id,
			&result.AssetsType,
			&result.AssetsLocation,
			&result.AssetsLocationType,
			&result.AssetsMediaType,
			&result.AssetsExt,
			&result.AssetsName,
			&result.AssetsUrl,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
	}

	return &result, nil
}
