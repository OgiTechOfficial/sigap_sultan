package csv_helper

import (
	"encoding/csv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"os"
	"regexp"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/common"
	"strings"
)

func OpenCsvFile(params common.FileStructure) (*csv.Reader, *common.ErrorDomain) {
	log.Info("=> open csv file")

	f, err := os.Open(params.FilePath)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: "Data is not valid",
			Details: err.Error(),
		}
	}

	reader := csv.NewReader(f)
	return reader, nil
}

func CsvFileProcess(params common.FileStructure, moduleType common.ModuleType) ([][]interface{}, *common.ErrorDomain) {
	var csvValues [][]interface{}
	var errDomain *common.ErrorDomain

	csvReader, errDomain := OpenCsvFile(params)
	if errDomain != nil {
		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	if moduleType.Price {
		csvValues, errDomain = ReadCsvFile(csvReader, moduleType)
		if errDomain.Details != nil {
			return nil, errDomain
		}
	} else {
		csvValues, errDomain = ReadCsvFileNeraca(csvReader, moduleType)
		if errDomain.Details != nil {
			return nil, errDomain
		}
	}

	return csvValues, nil
}

func ReadCsvFile(cr *csv.Reader, moduleType common.ModuleType) ([][]interface{}, *common.ErrorDomain) {
	var errorDetails []map[string]any
	errorDomain := &common.ErrorDomain{
		Message: "Data is not valid",
	}

	rowIdx := 0
	finalCsvValues := make([][]interface{}, 0)

	for {
		//records, err := cr.Read()
		records, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		//records = cleanDateField(records)

		colIdx := 0
		rowOrdered := make([]interface{}, 0)
		for _, rowValue := range records {
			switch {
			//case strings.ToLower(rowValue) == "date" || strings.ToLower(rowValue) == "tanggal":
			case strings.Contains(strings.ToLower(rowValue), "date") || strings.Contains(strings.ToLower(rowValue), "tanggal"):
				common.PriceColumns["last_update"] = ""
			case strings.ToLower(rowValue) == "wilayah":
				common.PriceColumns["city"] = ""
			case strings.ToLower(rowValue) == "komoditas":
				common.PriceColumns["commodity_name"] = ""
			case strings.ToLower(rowValue) == "price":
				common.PriceColumns["price"] = ""
			}

			//if (strings.ToLower(rowValue) != "date" && strings.ToLower(rowValue) != "tanggal") && colIdx == 0 {
			if !strings.Contains(strings.ToLower(rowValue), "date") && !strings.Contains(strings.ToLower(rowValue), "tanggal") && colIdx == 0 {
				if moduleType.Price {
					dateRegex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
					if !dateRegex.MatchString(rowValue) {
						errorDetails = append(errorDetails, map[string]any{
							"baris": rowIdx,
							"value": records,
						})
					}
				} else {
					dateRegex := regexp.MustCompile(`\d{4}-\d{2}`)
					if !dateRegex.MatchString(rowValue) {
						errorDetails = append(errorDetails, map[string]any{
							"baris": rowIdx,
							"value": records,
						})
					}
				}
			}

			if strings.ToLower(rowValue) != "wilayah" && colIdx == 1 {
				dateRegex := regexp.MustCompile(`^.*?(kabupaten|kota|nasional|sulawesi).*$`)
				if !dateRegex.MatchString(strings.ToLower(rowValue)) {
					errorDetails = append(errorDetails, map[string]any{
						"baris": rowIdx,
						"value": records,
					})
				}
			}

			//if strings.ToLower(rowValue) != "date" &&
			//	strings.ToLower(rowValue) != "tanggal" &&
			if !strings.Contains(strings.ToLower(rowValue), "date") &&
				!strings.Contains(strings.ToLower(rowValue), "tanggal") &&
				strings.ToLower(rowValue) != "wilayah" &&
				strings.ToLower(rowValue) != "komoditas" &&
				strings.ToLower(rowValue) != "harga" &&
				strings.ToLower(rowValue) != "ketersediaan" &&
				strings.ToLower(rowValue) != "kebutuhan" &&
				strings.ToLower(rowValue) != "neraca" {
				rowOrdered = append(rowOrdered, rowValue)
			}

			if len(errorDetails) == 10 {
				break
			}

			colIdx++
		}

		finalCsvValues = append(finalCsvValues, rowOrdered)
		rowIdx++
	}

	if len(errorDetails) > 0 {
		log.Info("helper.csv_helper.ReadCsvFile.Error:", errorDetails)
		errorDomain.Details = errorDetails
		return nil, errorDomain
	}

	return finalCsvValues[1:], errorDomain
}

func ReadCsvFileNeraca(cr *csv.Reader, moduleType common.ModuleType) ([][]interface{}, *common.ErrorDomain) {
	var errorDetails []map[string]any
	errorDomain := &common.ErrorDomain{
		Message: "Data is not valid",
	}

	rowIdx := 0
	finalCsvValues := make([][]interface{}, 0.0)

	for {
		//records, err := cr.Read()
		records, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		colIdx := 0
		rowOrdered := make([]interface{}, 0.0)
		for _, rowValue := range records {
			switch {
			//case strings.ToLower(rowValue) == "date" || strings.ToLower(rowValue) == "tanggal":
			case strings.Contains(strings.ToLower(rowValue), "date") || strings.Contains(strings.ToLower(rowValue), "tanggal"):
				common.PriceColumns["last_update"] = ""
			case strings.ToLower(rowValue) == "wilayah":
				common.PriceColumns["city"] = ""
			case strings.ToLower(rowValue) == "komoditas":
				common.PriceColumns["commodity_name"] = ""
			case strings.ToLower(rowValue) == "price":
				common.PriceColumns["price"] = ""
			}

			//if (strings.ToLower(rowValue) != "date" && strings.ToLower(rowValue) != "tanggal") && colIdx == 0 {
			if !strings.Contains(strings.ToLower(rowValue), "date") && !strings.Contains(strings.ToLower(rowValue), "tanggal") && colIdx == 0 {
				if moduleType.Price {
					dateRegex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
					if !dateRegex.MatchString(rowValue) {
						errorDetails = append(errorDetails, map[string]any{
							"baris": rowIdx,
							"value": records,
						})
					}
				} else {
					dateRegex := regexp.MustCompile(`\d{4}-\d{2}`)
					if !dateRegex.MatchString(rowValue) {
						errorDetails = append(errorDetails, map[string]any{
							"baris": rowIdx,
							"value": records,
						})
					}
				}
			}

			if strings.ToLower(rowValue) != "wilayah" && colIdx == 1 {
				dateRegex := regexp.MustCompile(`^.*?(kabupaten|kota|nasional|sulawesi).*$`)
				if !dateRegex.MatchString(strings.ToLower(rowValue)) {
					errorDetails = append(errorDetails, map[string]any{
						"baris": rowIdx,
						"value": records,
					})
				}
			}

			//if strings.ToLower(rowValue) != "date" &&
			//	strings.ToLower(rowValue) != "tanggal" &&
			if !strings.Contains(strings.ToLower(rowValue), "date") &&
				!strings.Contains(strings.ToLower(rowValue), "tanggal") &&
				strings.ToLower(rowValue) != "wilayah" &&
				strings.ToLower(rowValue) != "komoditas" &&
				strings.ToLower(rowValue) != "harga" &&
				strings.ToLower(rowValue) != "ketersediaan" &&
				strings.ToLower(rowValue) != "kebutuhan" &&
				strings.ToLower(rowValue) != "neraca" {
				rowOrdered = append(rowOrdered, rowValue)
			}

			colIdx++
		}

		finalCsvValues = append(finalCsvValues, rowOrdered)
		rowIdx++
	}

	if len(errorDetails) > 0 {
		log.Info("helper.csv_helper.ReadCsvFile.Error:", errorDetails)
		errorDomain.Details = errorDetails
		return nil, errorDomain
	}

	return finalCsvValues[1:], errorDomain
}

func DeleteTempDir(c *fiber.Ctx, temp string) {
	err := os.RemoveAll(temp)
	if err != nil {
		response := models.SetError().Details(500, err.Error(), nil)
		c.Status(500).JSON(response)
	}

	log.Info("Deleting temp dir: %s was succeed\n", temp)
}

func cleanDateField(record []string) []string {
	// Assuming DATE is the first column
	if len(record) > 0 {
		record[0] = strings.TrimPrefix(record[0], "@")
	}
	return record
}
