package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type LoginService struct {
	LoginRepository *repositories.LoginRepository
}

func NewLoginService(loginRepository *repositories.LoginRepository) *LoginService {
	return &LoginService{
		LoginRepository: loginRepository,
	}
}

func (r *LoginService) Login(params models.LoginRequestParams) (interface{}, *common.ErrorDomain) {
	hash := sha256.New()

	// Write the data to the hash
	hash.Write([]byte(params.Password))

	// Get the final hash result as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the hash to a hexadecimal string
	password := hex.EncodeToString(hashBytes)

	data, err := r.LoginRepository.Login(params.Email, password)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *LoginService) ForgotPassword(params models.ForgotRequestParams) (*string, *common.ErrorDomain) {
	exist, err := r.LoginRepository.CheckEmail(params.Email)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	if *exist < 1 {
		return nil, &common.ErrorDomain{
			Message: "Email Tidak Terdaftar",
		}
	}

	url, err := r.LoginRepository.GetUrl()
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	unique := common_helper.RandSeq(20)

	insToken, err := r.LoginRepository.InsertToken(*exist, unique)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}
	if *insToken < 1 {
		return nil, &common.ErrorDomain{
			Message: "Gagal Generate Token",
		}
	}
	link := fmt.Sprintf("%s/reset?code=%s", *url, unique)
	if params.IsCms > 0 {
		link = fmt.Sprintf("%s/cms/reset?code=%s", *url, unique)
	}

	// // SMTP Server Configuration
	// smtpHost := "smtp.gmail.com"
	// smtpPort := "587"

	// // Sender's email credentials
	// senderEmail := "yansen@sertikom.id"
	// password := "drlt vjim nbjh bxah" // Use App Password for Gmail

	// SMTP Server Configuration
	smtpHost := "smtp-relay.brevo.com"
	smtpPort := "587"

	// Sender's email credentials
	senderEmail := "88e719001@smtp-brevo.com"
	password := "1CHBKy76F0nkpSqE" // Use App Password for Gmail

	from := "info@neracapangansulsel.id"
	// from := "yansen@sertikom.id"

	// Receiver email address
	to := []string{params.Email}

	// Email Headers and Body
	subject := "Subject: Forgot Password\n"
	fromHeader := "From: \"neraca pangan sulsel\" <" + from + ">\r\n" // ✅ Correct name format
	toHeader := "To: " + params.Email + "\r\n"

	body := fmt.Sprintf("Klik Link Berikut Untuk Mereset Password Anda %s", link)
	// Message body
	message := []byte(fromHeader + toHeader + subject + "\r\n" + body)

	// Authentication
	auth := smtp.PlainAuth("", senderEmail, password, smtpHost)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return &link, nil
}

func (r *LoginService) CheckKode(params models.ForgotTokenParam) (*string, *common.ErrorDomain) {
	exist, err := r.LoginRepository.CheckKode(params.Code)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	if *exist < 1 {
		return nil, &common.ErrorDomain{
			Message: "Kode Tidak Valid atau sudah Kadaluarsa",
		}
	}

	var success *string
	str := "kode valid"
	success = &str

	return success, nil
}

func (r *LoginService) ResetPassword(params models.ResetRequestParams) (*string, *common.ErrorDomain) {
	exist, err := r.LoginRepository.CheckKode(params.Code)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	if *exist < 1 {
		return nil, &common.ErrorDomain{
			Message: "Kode Tidak Valid atau sudah Kadaluarsa",
		}
	}

	hash := sha256.New()

	hash.Write([]byte(params.NewPassword))

	hashBytes := hash.Sum(nil)

	password := hex.EncodeToString(hashBytes)

	_, err = r.LoginRepository.ResetPassword(*exist, password)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	var success *string
	str := "success"
	success = &str

	return success, nil
}

func (r *LoginService) Profile(id int) (interface{}, *common.ErrorDomain) {
	data, err := r.LoginRepository.Profile(id)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *LoginService) UpdateProfile(userid int, params models.UpdateProfileParam) (*string, *common.ErrorDomain) {
	_, err := r.LoginRepository.UpdateProfile(userid, params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	var success *string
	str := "success"
	success = &str

	return success, nil
}

func (r *LoginService) UpdatePassword(userid int, oldpassword string, password string) (*string, *common.ErrorDomain) {
	hash := sha256.New()

	hash.Write([]byte(oldpassword))

	hashBytes := hash.Sum(nil)

	old := hex.EncodeToString(hashBytes)
	exist, err := r.LoginRepository.CheckPassword(userid, old)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	if *exist < 1 {
		return nil, &common.ErrorDomain{
			Message: "Password lama salah",
		}
	}

	hash2 := sha256.New()

	hash2.Write([]byte(password))

	hashBytes2 := hash2.Sum(nil)

	new := hex.EncodeToString(hashBytes2)
	_, err = r.LoginRepository.ChangePassword(userid, new)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	var success *string
	str := "success"
	success = &str

	return success, nil
}
