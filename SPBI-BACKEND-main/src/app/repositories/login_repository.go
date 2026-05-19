package repositories

import (
	"context"
	"fmt"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories/queries"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginRepository struct {
	Db *pgxpool.Pool
}

func NewLoginRepository(db *pgxpool.Pool) *LoginRepository {
	return &LoginRepository{Db: db}
}

func (repo *LoginRepository) Login(email string, password string) (interface{}, error) {
	var result models.LoginResponse
	err := repo.Db.QueryRow(
		context.Background(),
		`select tu.id as id, concat(tp.nama_depan, ' ', tp.nama_belakang) as name, tp.position_id as position_id, tp2.name as position from tm_user tu
join tm_profile tp on tp.user_id = tu.id 
join tm_position tp2 on tp2.id = tp.position_id 
where tu.email = @email and tu.password = @password`,
		pgx.NamedArgs{
			"email":    email,
			"password": password,
		},
	).Scan(&result.Id.Id, &result.Name, &result.PositionId, &result.Position)

	if err != nil {
		log.Error("repositories.LoginRepository.Login QueryErrors:", err)
		return nil, err
	}

	// fmt.Print(rows)

	// for rows.Next() {
	// 	err = rows.Scan(
	// 		&result.Id.Id,
	// 		&result.Name,
	// 		&result.PositionId,
	// 		&result.Position,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &result, nil
}

func (repo *LoginRepository) CheckEmail(email string) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		`select id FROM tm_user tu 
where tu.email = @email`,
		pgx.NamedArgs{
			"email": email,
		},
	).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *LoginRepository) GetUrl() (*string, error) {
	var result string
	err := repo.Db.QueryRow(
		context.Background(),
		`SELECT name FROM settings WHERE parent_id = (SELECT id FROM settings WHERE name = 'BASE_URL')`,
	).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *LoginRepository) InsertToken(userId int, token string) (*int, error) {
	var err error
	var lastInsertedId int
	now := time.Now()
	tomorrow := time.Now().Add(24 * time.Hour)

	request := new(models.TmUserForgotToken)

	request.CreatedAt = &now
	request.ExpiredAt = &tomorrow
	request.UserId = &userId
	request.Token = &token

	err = repo.Db.QueryRow(
		context.Background(),
		queries.TmUserForgotToken,
		request.UserId,
		request.Token,
		request.ExpiredAt,
		request.CreatedAt,
	).Scan(&lastInsertedId)
	if err != nil {
		log.Error("QueryErrors:", err)
		return nil, err
	}

	return &lastInsertedId, nil
}

func (repo *LoginRepository) CheckKode(code string) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		`select user_id FROM tm_user_forgot_token tu 
where tu.token = @code and tu.expired_at >= now() and tu.deleted_at is null`,
		pgx.NamedArgs{
			"code": code,
		},
	).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *LoginRepository) ResetPassword(userid int, password string) (*string, error) {
	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		context.Background(),
		`UPDATE tm_user_forgot_token
		SET deleted_at = now()
		WHERE user_id = @userid`,
		pgx.NamedArgs{
			"userid": userid,
		},
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		context.Background(),
		`UPDATE tm_user
		SET password = @password
		WHERE id = @userid`,
		pgx.NamedArgs{
			"password": password,
			"userid":   userid,
		},
	)
	if err != nil {
		return nil, err
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		return nil, err
	}
	var success *string
	str := "success"
	success = &str

	return success, nil
}

func (repo *LoginRepository) Profile(id int) (interface{}, error) {
	var result models.ProfileResponse
	err := repo.Db.QueryRow(
		context.Background(),
		`SELECT tu.id,
    tu.username,
    tu.email,
    tp.nama_depan,
    tp.nama_belakang,
    tp.jabatan,
	tp.position_id as jabatan_id,
    tp.organisasi,
    tp.bidang_unit_kerja
   FROM prod.tm_user tu
     JOIN prod.tm_profile tp ON tp.user_id = tu.id where tu.id = @id`,
		pgx.NamedArgs{
			"id": id,
		},
	).Scan(&result.Id.Id, &result.Username, &result.Email, &result.NamaDepan, &result.NamaBelakang, &result.Jabatan, &result.JabatanId, &result.Organisasi, &result.BidangUnitKerja)

	if err != nil {
		log.Error("repositories.LoginRepository.Profile QueryErrors:", err)
		return nil, err
	}

	return &result, nil
}

func (repo *LoginRepository) UpdateProfile(userid int, reqData models.UpdateProfileParam) (*string, error) {
	fmt.Println(reqData)
	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		context.Background(),
		`UPDATE tm_user
		SET username = @username, email = @email
		WHERE id = @userid`,
		pgx.NamedArgs{
			"username": reqData.Username,
			"email":    reqData.Email,
			"userid":   userid,
		},
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		context.Background(),
		`UPDATE tm_profile
		SET nama_depan = @namadepan, nama_belakang = @namabelakang, jabatan = @jabatan, organisasi = @organisasi, bidang_unit_kerja = @buk, position_id = @jabatanid
		WHERE user_id = @userid`,
		pgx.NamedArgs{
			"namadepan":    reqData.NamaDepan,
			"namabelakang": reqData.NamaBelakang,
			"jabatan":      reqData.Jabatan,
			"organisasi":   reqData.Organisasi,
			"buk":          reqData.BidangUnitKerja,
			"jabatanid":    reqData.JabatanId,
			"userid":       userid,
		},
	)
	if err != nil {
		return nil, err
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		return nil, err
	}
	var success *string
	str := "success"
	success = &str

	return success, nil
}

func (repo *LoginRepository) ChangePassword(userid int, password string) (*string, error) {
	tx, err := repo.Db.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		context.Background(),
		`UPDATE tm_user
		SET password = @password
		WHERE id = @userid`,
		pgx.NamedArgs{
			"password": password,
			"userid":   userid,
		},
	)
	if err != nil {
		return nil, err
	}
	_ = tx.Commit(context.Background())

	if err != nil {
		return nil, err
	}
	var success *string
	str := "success"
	success = &str

	return success, nil
}

func (repo *LoginRepository) CheckPassword(id int, password string) (*int, error) {
	var result int
	err := repo.Db.QueryRow(
		context.Background(),
		`select count(1) from tm_user tu
where tu.id = @id and tu.password = @password`,
		pgx.NamedArgs{
			"id":       id,
			"password": password,
		},
	).Scan(&result)

	if err != nil {
		log.Error("repositories.LoginRepository.Login QueryErrors:", err)
		return nil, err
	}

	return &result, nil
}
