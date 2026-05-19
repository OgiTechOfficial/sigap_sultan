package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"sigap-sultan-be/src/common"
	"strings"
)

type TmCommodityRepository struct {
	Db               *pgxpool.Pool
	AssetsRepository *AssetsRepository
}

func NewTmCommodityRepository(db *pgxpool.Pool, assetsRepository *AssetsRepository) *TmCommodityRepository {
	return &TmCommodityRepository{Db: db, AssetsRepository: assetsRepository}
}

func (repo TmCommodityRepository) Get(params domain.TmCommodityRequestParam) (interface{}, error) {
	var results []models.TmCommodity
	var rows pgx.Rows
	var err error

	if params.ModuleType == "neraca" {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.TmCommodityListNeraca,
		)
	} else {
		rows, err = repo.Db.Query(
			context.Background(),
			queries.TmCommodityList,
		)
	}

	if err != nil {
		log.Fatal("TmCommodityRepository.QueryErrors:", err)
		return nil, err
	}

	for rows.Next() {
		var result models.TmCommodity
		err = rows.Scan(
			&result.Id.Id,
			&result.ClientId,
			&result.ParentId,
			&result.Code,
			&result.Name,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
			&result.AssetsRelationId,
			&result.Sequence,
			&result.UnitId,
			&result.UnitIdNeraca,
			&result.UnitName,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}

		results = append(results, result)
	}

	return &results, nil
}

func (repo TmCommodityRepository) GetGrouping() (interface{}, error) {
	var results [][]string
	var tmCommoditiesGroup map[string][]string
	tmCommoditiesGroup = map[string][]string{}

	//var tmCommodities []models.TmCommodity

	rows, err := repo.Db.Query(
		context.Background(),
		queries.TmCommodityParentListOrderBySequence,
	)
	if err != nil {
		log.Fatal("TmCommodityRepository.QueryErrors:", err)
		return nil, err
	}

	for rows.Next() {
		var tmCommodityParent models.TmCommodity
		err = rows.Scan(
			&tmCommodityParent.Id.Id,
			&tmCommodityParent.ClientId,
			&tmCommodityParent.ParentId,
			&tmCommodityParent.Code,
			&tmCommodityParent.Name,
			&tmCommodityParent.CreatedAt,
			&tmCommodityParent.UpdatedAt,
			&tmCommodityParent.DeletedAt,
			&tmCommodityParent.AssetsRelationId,
			&tmCommodityParent.Sequence,
		)
		if err != nil {
			fmt.Printf("Scan error: %v\n", err)
		}
		//key = append(key, tmCommodityParent.Name)

		commodityByParentIdRows, errGetCommodityByParentId := repo.Db.Query(
			context.Background(),
			queries.TmCommodityGetByParentId,
			pgx.NamedArgs{
				"id": tmCommodityParent.Id.Id,
			},
		)
		if errGetCommodityByParentId != nil {
			log.Fatal("TmCommodityRepository.QueryErrors:", err)
			return nil, err
		}

		var child []string
		var tmCommodityChild models.TmCommodity
		for commodityByParentIdRows.Next() {
			err := commodityByParentIdRows.Scan(
				&tmCommodityChild.Id.Id,
				&tmCommodityChild.ClientId,
				&tmCommodityChild.ParentId,
				&tmCommodityChild.Code,
				&tmCommodityChild.Name,
				&tmCommodityChild.CreatedAt,
				&tmCommodityChild.UpdatedAt,
				&tmCommodityChild.DeletedAt,
				&tmCommodityChild.AssetsRelationId,
				&tmCommodityChild.Sequence,
			)
			if err != nil {
				return nil, err
			}

			child = append(child, tmCommodityChild.Name)
		}

		tmCommoditiesGroup[tmCommodityParent.Name] = child
		//var a []string = tmCommoditiesGroup[0]
		//tmCommoditiesGroup[len(tmCommoditiesGroup)]

		//tmCommodities = append(tmCommodities, result)

		results = append(results, child)
	}

	return &results, nil
}

func (repo TmCommodityRepository) GetById(id int32) (*models.TmCommodity, error) {
	var tmCommodity models.TmCommodity
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityGetById,
		pgx.NamedArgs{
			"id": id,
		},
	).Scan(
		&tmCommodity.Id.Id,
		&tmCommodity.ClientId,
		&tmCommodity.ParentId,
		&tmCommodity.Code,
		&tmCommodity.Name,
		&tmCommodity.CreatedAt,
		&tmCommodity.UpdatedAt,
		&tmCommodity.DeletedAt,
		&tmCommodity.AssetsRelationId,
		&tmCommodity.Sequence,
		&tmCommodity.UnitId,
		&tmCommodity.UnitIdNeraca,
		&tmCommodity.UnitName,
	)
	if err != nil {
		return nil, err
	}

	return &tmCommodity, nil
}

func (repo TmCommodityRepository) GetByName(name string) (*models.TmCommodity, error) {
	var tmCommodity models.TmCommodity
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityGetByName,
		pgx.NamedArgs{
			"name": "%" + strings.ToLower(name) + "%",
		},
	).Scan(
		&tmCommodity.Id.Id,
		&tmCommodity.ClientId,
		&tmCommodity.ParentId,
		&tmCommodity.Code,
		&tmCommodity.Name,
		&tmCommodity.CreatedAt,
		&tmCommodity.UpdatedAt,
		&tmCommodity.DeletedAt,
		&tmCommodity.AssetsRelationId,
		&tmCommodity.Sequence,
		&tmCommodity.UnitId,
		&tmCommodity.UnitIdNeraca,
		&tmCommodity.UnitName,
	)
	if err != nil {
		return nil, err
	}

	return &tmCommodity, nil
}

func (repo TmCommodityRepository) GetByEqualName(name string) (*models.TmCommodity, error) {
	var tmCommodity models.TmCommodity
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityGetByEqualName,
		pgx.NamedArgs{
			"name": strings.ToLower(name),
		},
	).Scan(
		&tmCommodity.Id.Id,
		&tmCommodity.ClientId,
		&tmCommodity.ParentId,
		&tmCommodity.Code,
		&tmCommodity.Name,
		&tmCommodity.CreatedAt,
		&tmCommodity.UpdatedAt,
		&tmCommodity.DeletedAt,
		&tmCommodity.AssetsRelationId,
		&tmCommodity.Sequence,
		&tmCommodity.UnitId,
		&tmCommodity.UnitIdNeraca,
		&tmCommodity.UnitName,
	)
	if err != nil {
		return nil, err
	}

	return &tmCommodity, nil
}

func (repo TmCommodityRepository) Insert(request domain.TmCommodityParam) (interface{}, error) {
	var err error
	var lastCommodityInsertedId *int32
	var parentInsertId *int32
	var assetId *int32
	var assetParamName string
	var tmCommodity *models.TmCommodity

	nameSplit := strings.Split(request.Name, " ")
	if len(nameSplit) > 2 {
		assetParamName = nameSplit[0] + " " + nameSplit[1] + " " + nameSplit[2]
	} else if len(nameSplit) == 2 {
		assetParamName = nameSplit[0] + " " + nameSplit[1]
	} else if len(nameSplit) == 1 {
		assetParamName = nameSplit[0]
	} else {
		return nil, errors.New("commodity Name must be required")
	}

	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		log.Error("TmCommodityRepository.QueryErrors:", err)
	}

	assetData, errGetByName := repo.AssetsRepository.GetByName(assetParamName)
	if errGetByName != nil {
		return nil, errGetByName
	}
	if assetData != nil {
		assetsResponse := assetData.(*models.AssetsResponse)
		assetId = assetsResponse.Id.Id
	}

	tmCommodity, err = repo.GetByName(request.Name)
	if err != nil {
		if err.Error() != "no rows in result set" {
			_ = tx.Rollback(context.Background())
			log.Error("TmCommodityRepository.QueryErrors:", err)
			return nil, err
		}
	}
	if tmCommodity != nil {
		return tmCommodity, nil
	}

	if len(nameSplit) > 1 {
		parentCommodity, errGetByName := repo.GetByName(nameSplit[0])
		if errGetByName != nil {
			// INSERT PARENT
			if errGetByName.Error() == "no rows in result set" {
				//var data []interface{}
				parentInsertId, err = repo.doInsertParent(
					tx,
					map[string]any{
						"name":            nameSplit[0],
						"assetRelationId": assetId,
					},
				)

				// INSERT CHILD
				lastCommodityInsertedId, err = repo.doInsertChild(
					tx,
					map[string]any{
						"parentId":        parentInsertId,
						"name":            request.Name,
						"assetRelationId": assetId,
					},
				)
				if err != nil {
					_ = tx.Rollback(context.Background())
					return nil, err
				}
			} else {
				_ = tx.Rollback(context.Background())
				log.Error("TmCommodityRepository.QueryErrors:", errGetByName)
				return nil, errGetByName
			}
		} else {
			// INSERT CHILD
			lastCommodityInsertedId, err = repo.doInsertChild(
				tx,
				map[string]any{
					"parentId":        parentCommodity.Id.Id,
					"name":            request.Name,
					"assetRelationId": assetId,
				},
			)
			if err != nil {
				_ = tx.Rollback(context.Background())
				return nil, err
			}
		}

		tmCommodity, err = repo.GetById(*lastCommodityInsertedId)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return nil, err
		}
	} else {
		// INSERT CHILD
		_, err := repo.GetByName(request.Name)
		if err != nil {
			// INSERT PARENT
			if err.Error() == "no rows in result set" {
				//var data []interface{}
				parentInsertId, err = repo.doInsertParent(
					tx,
					map[string]any{
						"name":            request.Name,
						"assetRelationId": assetId,
					},
				)

				tmCommodity, err = repo.GetById(*parentInsertId)
				if err != nil {
					_ = tx.Rollback(context.Background())
					return nil, err
				}
			} else {
				_ = tx.Rollback(context.Background())
				log.Error("TmCommodityRepository.QueryErrors:", err)
				return nil, err
			}
		}

		//lastCommodityInsertedId, err = repo.doInsertChild(
		//	tx,
		//	map[string]any{
		//		"parentId":        parentCommodity.ParentId,
		//		"name":            request.Name,
		//		"assetRelationId": assetId,
		//	},
		//)
		//if err != nil {
		//	_ = tx.Rollback(context.Background())
		//	return nil, err
		//}
	}

	_ = tx.Commit(context.Background())

	return tmCommodity, nil
}

func (repo TmCommodityRepository) IsParent(id int32) (*bool, error) {
	var result bool
	err := repo.Db.QueryRow(
		context.Background(),
		queries.TmCommodityIsParent,
		pgx.NamedArgs{
			"id": id,
		},
	).Scan(
		&result,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *TmCommodityRepository) doInsertChild(tx pgx.Tx, data map[string]any) (*int32, error) {
	var err error
	var lastInsertedId int32

	err = tx.QueryRow(
		context.Background(),
		queries.TmCommodityInsertChildWhenNotExists,
		pgx.NamedArgs{
			"parentId":        data["parentId"],
			"name":            data["name"],
			"createdAt":       common.GetDateTimeNow(),
			"assetRelationId": data["assetRelationId"],
		},
	).Scan(&lastInsertedId)
	if err != nil {
		log.Error("QueryErrors:", err)
		_ = tx.Rollback(context.Background())
		return nil, err
	}
	_ = tx.Commit(context.Background())

	return &lastInsertedId, nil

	//_, err = tx.Exec(
	//	context.Background(),
	//	queries.TmCommodityInsertChildWhenNotExists,
	//	data,
	//).Scan(&lastInsertedId)
	//if err != nil {
	//	log.Fatalf("tx.Exec failed: %v\n", err)
	//}
	//_ = tx.Commit(context.Background())
	//
	//if err != nil {
	//	log.Error("TmCommodityRepository.QueryErrors:", err)
	//}
}

func (repo *TmCommodityRepository) doInsertParent(tx pgx.Tx, data map[string]any) (*int32, error) {
	var err error
	var lastInsertedId int32

	err = tx.QueryRow(
		context.Background(),
		queries.TmCommodityInsertParentWhenNotExists,
		pgx.NamedArgs{
			"name":            data["name"],
			"createdAt":       common.GetDateTimeNow(),
			"assetRelationId": data["assetRelationId"],
		},
	).Scan(&lastInsertedId)
	if err != nil {
		log.Error("QueryErrors:", err)
		_ = tx.Rollback(context.Background())
		return nil, err
	}
	//_ = tx.Commit(context.Background())

	return &lastInsertedId, nil
}

//func (repo *TmCommodityRepository) DoInsert(tmCommodity *models.TmCommodity, parentId *int32) {
//	var data []interface{}
//	tx, err := repo.Db.Begin(context.Background())
//	if err != nil {
//		log.Error("TmCommodityRepository.QueryErrors:", err)
//	}
//
//	data = append(data, 1, tmCommodity.Name, tmCommodity.CreatedAt, common.GetDateTimeNow())
//
//	if parentId != nil {
//		data = append(data, parentId)
//	} else {
//
//	}
//	_, err = tx.Exec(
//		context.Background(),
//		queries.PriceTxCommodityInsert,
//	)
//	if err != nil {
//		log.Fatalf("tx.Exec failed: %v\n", err)
//	}
//	_ = tx.Commit(context.Background())
//
//	if err != nil {
//		log.Error("TmCommodityRepository.QueryErrors:", err)
//	}
//}
