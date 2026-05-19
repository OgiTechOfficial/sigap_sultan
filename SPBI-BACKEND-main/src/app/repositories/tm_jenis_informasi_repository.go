package repositories

import (
	"context"
	_ "github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"sigap-sultan-be/src/common"
	"strings"
	_ "time"
)

type TmJenisInformasiRepository struct {
	Db *pgxpool.Pool
}

func NewTmJenisInformasiRepository(db *pgxpool.Pool) *TmJenisInformasiRepository {
	return &TmJenisInformasiRepository{Db: db}
}

func (repo TmJenisInformasiRepository) Get(paginationParams common.PaginationParams) (interface{}, error) {
	var results []models.TmJenisInformasiResponse
	rows, err := repo.Db.Query(
		context.Background(),
		queries.TmJenisInformasiList,
		pgx.NamedArgs{
			"page":  paginationParams.Page,
			"limit": paginationParams.Limit,
		},
	)

	for rows.Next() {
		var data models.TmJenisInformasiResponse
		err = rows.Scan(
			&data.Id.Id,
			&data.ClientId,
			&data.ParentId,
			&data.Name,
			&data.DetailJenisInformasi,
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

func (repo TmJenisInformasiRepository) GetById(id int) (interface{}, error) {
	var tmJenisInformasi models.TmJenisInformasiResponse
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmJenisInformasiGetById,
		pgx.NamedArgs{
			"id": id,
		},
	).Scan(
		&tmJenisInformasi.Id.Id,
		&tmJenisInformasi.ClientId,
		&tmJenisInformasi.Name,
		&tmJenisInformasi.DetailJenisInformasi,
		&tmJenisInformasi.CreatedAt,
		&tmJenisInformasi.UpdatedAt,
		&tmJenisInformasi.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &tmJenisInformasi, nil
}

func (repo TmJenisInformasiRepository) GetByName(name string) (interface{}, error) {
	var tmJenisInformasi models.TmJenisInformasiResponse
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmJenisInformasiGetByName,
		pgx.NamedArgs{
			"name": "%" + strings.ToLower(name) + "%",
		},
	).Scan(
		&tmJenisInformasi.Id.Id,
		&tmJenisInformasi.ClientId,
		&tmJenisInformasi.Name,
		&tmJenisInformasi.DetailJenisInformasi,
		&tmJenisInformasi.CreatedAt,
		&tmJenisInformasi.UpdatedAt,
		&tmJenisInformasi.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &tmJenisInformasi, nil
}

//func (repo TmJenisInformasiRepository) Insert(request models.JenisInformasiRequest) (interface{}, error) {
//
//}

//func (repo TmJenisInformasiRepository) Insert(request models.PrivilegesRequest) (interface{}, error) {
//	var err error
//	var lastInsertedId int
//	var tmJenisInformasi models.TmJenisInformasi
//	tmJenisInformasi.ClientId = 1
//	tmJenisInformasi.Name = *request.JenisInformasi
//
//	now := time.Now()
//	tmJenisInformasi.UpdatedAt = &now
//
//	//tx, err := repo.Db.Begin(context.Background())
//	//if err != nil {
//	//	log.Error("QueryErrors:", err)
//	//}
//
//	//_, err = tx.Exec(
//	//	context.Background(),
//	//	queries.TmJenisInformasiInsert,
//	//	tmJenisInformasi.ClientId,
//	//	tmJenisInformasi.Name,
//	//	time.Now().Format(common.DATETIME_FORMAT),
//	//)
//	//if err != nil {
//	//	log.Fatalf("tx.Exec failed: %v\n", err)
//	//}
//	//_ = tx.Commit(context.Background())
//
//	//if err != nil {
//	//	log.Error("QueryErrors:", err)
//	//}
//
//	err = repo.Db.QueryRow(
//		context.Background(),
//		queries.TmJenisInformasiInsert,
//		tmJenisInformasi.ClientId,
//		tmJenisInformasi.Name,
//		time.Now().Format(common.DATETIME_FORMAT),
//	).Scan(&lastInsertedId)
//	if err != nil {
//		log.Error("QueryErrors:", err)
//		return nil, err
//	}
//
//	for _, v := range *request.Privileges.Permissions {
//		tx, err := repo.Db.Begin(context.Background())
//		if err != nil {
//			log.Error("QueryErrors:", err)
//		}
//
//		_, err = tx.Exec(
//			context.Background(),
//			queries.MapJabatanJenisInformasiInsert,
//			tmJenisInformasi.ClientId,
//			lastInsertedId,
//			request.Privileges.JenisInformasiId,
//			v.UploadPriceCity,
//			v.Read,
//			v.Update,
//			v.Delete,
//			time.Now().Format(common.DATETIME_FORMAT),
//		)
//		if err != nil {
//			log.Fatalf("tx.Exec failed: %v\n", err)
//		}
//		_ = tx.Commit(context.Background())
//
//		if err != nil {
//			log.Error("QueryErrors:", err)
//		}
//	}
//
//	lastTmJenisInformasi, err := repo.GetById(lastInsertedId)
//
//	return &lastTmJenisInformasi, nil
//}
//
//func (repo TmJenisInformasiRepository) Update(tmJenisInformasi models.TmJenisInformasi) (interface{}, error) {
//	tx, err := repo.Db.Begin(context.Background())
//	if err != nil {
//		log.Error("QueryErrors:", err)
//	}
//
//	_, err = tx.Exec(
//		context.Background(),
//		queries.TmJenisInformasiUpdate,
//		tmJenisInformasi.ClientId,
//		tmJenisInformasi.Name,
//		time.Now().Format(common.DATETIME_FORMAT),
//	)
//	if err != nil {
//		log.Fatalf("tx.Exec failed: %v\n", err)
//	}
//	_ = tx.Commit(context.Background())
//
//	if err != nil {
//		log.Error("QueryErrors:", err)
//	}
//
//	return &tmJenisInformasi, nil
//}
