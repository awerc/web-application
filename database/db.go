package food_delivery

import (
	"database/sql"
	"fmt"
	"web-application/xml_parser"

	_ "github.com/lib/pq"
)

func Connect(info xml_parser.Configuration) (db *sql.DB, err error) {
	connection_info := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		info.Connect.Host,
		info.Connect.Port,
		info.User.Login,
		info.User.Password,
		info.Connect.Db)

	db, err = sql.Open("postgres", connection_info)
	if err != nil {
		return nil, err
	}
	return db, nil
}
