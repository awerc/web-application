package food_delivery

import (
	"database/sql"
	"time"
)

type Client struct {
	Id       uint8
	Login    string
	Password string
	Tel      string
	Birthday time.Time
	Fullname string
}

func GetAllClients(db *sql.DB) (list []*struct {
	Name string
	Tel  string
	Bday time.Time
}, err error) {
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT fullname, telephone, birthday
		FROM clients;
    	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list = make([]*struct {
		Name string
		Tel  string
		Bday time.Time
	}, 0)
	for rows.Next() {
		client := new(struct {
			Name string
			Tel  string
			Bday time.Time
		})
		err = rows.Scan(&client.Name, &client.Tel, &client.Bday)
		if err != nil {
			return nil, err
		}
		list = append(list, client)
	}
	return list, nil
}

func AddClient(db *sql.DB, login, password, telephone, birthday, fullname string) (err error) {
	stmt, err := db.Prepare(`
INSERT INTO clients (login, password, telephone, birthday, fullname) 
VALUES ($1, $2, $3, $4, $5);
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(login, password, telephone, birthday, fullname)
	if err != nil {
		return err
	}

	return nil
}

func ChangeClientTelephone(db *sql.DB, client int, telephone string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit() //or defer tx.Rollback()

	_, err = tx.Exec(`
UPDATE clients
SET telephone = $1
WHERE id = $2;
	`, telephone, client)
	if err != nil {
		return err
	}
	return nil
}
