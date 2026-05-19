package queries

const (
	MapJabatanMenuInsert = `
		INSERT INTO map_jabatan_menu(
			client_id,
			position_id,
			menu_id,
			create_status,
			read_status,
			update_status,
			delete_status,
			created_at
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8
		)
	`
)
