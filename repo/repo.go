package repo

import (
	"database/sql"
	"fmt"
)

const table = "foos"

func GetAll(db *sql.DB) ([]string, error) {
	var names []string

	query := fmt.Sprintf("select * from %s", table)
	rows, err := db.Query(query)
	if err != nil {
		return names, err
	}

	defer rows.Close()

	for rows.Next() {
		var name string

		if err = rows.Scan(&name); err != nil {
			return names, err
		} else {
			names = append(names, name)
		}
	}

	return names, nil
}
