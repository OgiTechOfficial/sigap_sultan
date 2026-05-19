package common

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/xuri/excelize/v2"
)

func SetWidthColXlsx(sheetDefault string, f *excelize.File) {
	for _, a := range AbjadList {
		if err := f.SetColWidth(sheetDefault, a, a, 20); err != nil {
			log.Error(err)
		}
	}

	if err := f.SetColWidth(sheetDefault, "A", "B", 30); err != nil {
		log.Error(err)
	}
}

func SetHeadersBold(sheetDefault string, f *excelize.File) {
	if rowStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	}); err != nil {
		log.Error(err)
	} else {
		// HEADER INFORMATION
		for i := 1; i <= 5; i++ {
			if err := f.SetRowStyle(sheetDefault, i, i, rowStyle); err != nil {
				log.Error(err)
			}
		}

		// HEADER FOR NERACA / PRICE DATA
		if err := f.SetRowStyle(sheetDefault, 7, 7, rowStyle); err != nil {
			log.Error(err)
		}
	}
}
