package repositories

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"sigap-sultan-be/src/common"
	"time"
)

type TxFileUploadHistoryRepository struct {
	Db *pgxpool.Pool
}

func NewTxFileUploadHistoryRepository(db *pgxpool.Pool) *TxFileUploadHistoryRepository {
	return &TxFileUploadHistoryRepository{Db: db}
}

//
//func (repo TxFileUploadHistoryRepository) GetByName(name string) (*models.TmCommodity, error) {
//	var tmCommodity models.TmCommodity
//	err := repo.Db.QueryRow(
//		context.Background(),
//		queries.TmCommodityGetByName,
//		pgx.NamedArgs{
//			"name": name,
//		},
//	).Scan(&tmCommodity)
//	if err != nil {
//		return nil, err
//	}
//
//	return &tmCommodity, nil
//}

func (repo *TxFileUploadHistoryRepository) Save(txFileUploadHistory models.TxFileUploadHistory) {
	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		log.Error("QueryErrors:", err)
	}

	_, err = tx.Exec(
		context.Background(),
		queries.PriceTxFileUploadHistoryInsert,
		txFileUploadHistory.FileName,
		txFileUploadHistory.RowTotal,
		txFileUploadHistory.Status,
		txFileUploadHistory.ModuleType,
		txFileUploadHistory.Errors,
		time.Now().Format(common.DATETIME_FORMAT),
	)
	if err != nil {
		_ = tx.Rollback(context.Background())
		log.Fatalf("tx.Exec failed: %v\n", err)
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		log.Error("QueryErrors:", err)
	}
}
