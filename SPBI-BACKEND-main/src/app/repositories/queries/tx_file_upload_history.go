package queries

const (
	PriceTxFileUploadHistoryInsert = `
		INSERT INTO tx_file_upload_history(
			file_name,
			row_total,
			status,
			module_type,
			errors,
			created_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`
)
