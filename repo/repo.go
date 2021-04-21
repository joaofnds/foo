package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/opentracing/opentracing-go"
)

const table = "foos"

func GetAll(ctx context.Context, db *sql.DB) ([]string, error) {
	parentSpan := opentracing.SpanFromContext(ctx)
	span := opentracing.StartSpan("repo.GetAll", opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()

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
