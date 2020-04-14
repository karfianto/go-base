package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

const reportTable = `
CREATE TABLE reports (
id serial NOT NULL,
account_id int NOT NULL REFERENCES accounts(id),
updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
complaint text,
PRIMARY KEY (id)
)`

func init() {
	up := []string{
		reportTable,
	}

	down := []string{
		`DROP TABLE reports`,
	}

	migrations.Register(func(db migrations.DB) error {
		fmt.Println("create report table")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("drop report table")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
