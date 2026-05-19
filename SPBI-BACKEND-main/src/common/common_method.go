package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

func GenerateOffset(page int, limit int) int {
	var offset int

	if page > 0 && limit > 0 {
		offset = (page - 1) * limit
	}

	return offset

	//if page == 1 {
	//	page = 0
	//}
	//
	//if page > 0 {
	//	page = page * limit
	//} else {
	//	page = 0 * limit
	//}
	//
	//return page
}

func GetDateNow() string {
	now := time.Now()
	day := now.Day()
	month := now.Month()
	yearDay := now.Year()

	return strconv.Itoa(yearDay) + "-" + MonthWithLeadingZero(int(month)) + "-" + MonthWithLeadingZero(day)
}

func GetDateTimeNow() string {
	currentDate := time.Now().Add(time.Hour * 7)

	day := currentDate.Day()
	month := currentDate.Month()
	yearDay := currentDate.Year()

	hour := currentDate.Hour()
	minute := currentDate.Minute()
	second := currentDate.Second()

	return strconv.Itoa(yearDay) + "-" + MonthWithLeadingZero(int(month)) + "-" + strconv.Itoa(day) + " " + strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)
}

func MonthWithLeadingZero(month int) string {
	if month < 10 {
		return "0" + strconv.Itoa(month)
	}

	return strconv.Itoa(month)
}

func Validate(c *fiber.Ctx, params interface{}) []ErrorFieldValidationResponse {
	var validate = validator.New()

	xValidator := &XValidator{
		Validator: validate,
	}

	if errs := xValidator.Validate(params); len(errs) > 0 && errs[0].Error {
		errorFields := make([]ErrorFieldValidationResponse, 0)

		for _, err := range errs {
			errorFields = append(errorFields, ErrorFieldValidationResponse{
				Field:   err.FailedField,
				Message: "Must be " + err.Tag,
			})
		}

		return errorFields
	}

	return nil
}

func ThousandFormat(n int32) string {
	// Convert integer to string
	s := strconv.Itoa(int(n))
	// Get the length of the string
	length := len(s)
	// Initialize a string builder
	var sb strings.Builder

	// Iterate over the string in reverse
	for i, c := range s {
		if i > 0 && (length-i)%3 == 0 {
			sb.WriteString(".")
		}
		sb.WriteRune(c)
	}

	return sb.String()
}

func RupiahFormat(n int32) string {
	return "Rp " + ThousandFormat(n)
}

func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	firstDayOfNextMonth := StartOfMonth(date).AddDate(0, 1, 0)
	return firstDayOfNextMonth.Add(-time.Second)
}

func TimeToDate(time time.Time) string {
	return strconv.Itoa(time.Year()) + "-" + MonthWithLeadingZero(int(time.Month())) + "-" + MonthWithLeadingZero(time.Day())
}

func NeracaPeriodDateGenerate(startDate *string, endDate *string) error {
	var err error
	var dateFormat = "2006-01-02"
	var periodDateTime time.Time

	periodDateTime, err = time.Parse(dateFormat, *startDate)
	if err != nil {
		return err
	}
	*startDate = TimeToDate(StartOfMonth(periodDateTime))

	periodDateTime, err = time.Parse(dateFormat, *endDate)
	if err != nil {
		return err
	}
	*endDate = TimeToDate(EndOfMonth(periodDateTime))

	return nil
}

func TimeIsBetween(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}
