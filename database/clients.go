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

type ClientInfo struct {
	Name string
	Tel  string
	Bday time.Time
}

func GetClient(db *sql.DB, id int) (client *ClientInfo, err error) {
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT fullname, telephone, birthday
		FROM clients
		WHERE id = $1;
    	`, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	client = new(ClientInfo)
	rows.Next()
	err = rows.Scan(&client.Name, &client.Tel, &client.Bday)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetAllClients(db *sql.DB) (list []*ClientInfo, err error) {
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

	list = make([]*ClientInfo, 0)
	for rows.Next() {
		client := new(ClientInfo)
		err = rows.Scan(&client.Name, &client.Tel, &client.Bday)
		if err != nil {
			return nil, err
		}
		list = append(list, client)
	}
	return list, nil
}

func AddClient(db *sql.DB, client *Client) (err error) {
	stmt, err := db.Prepare(`
INSERT INTO clients (login, password, telephone, birthday, fullname) 
VALUES ($1, $2, $3, $4, $5);
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.Login, client.Password, client.Tel, client.Birthday, client.Fullname)
	if err != nil {
		return err
	}

	return nil
}

func DeleteClient(db *sql.DB, id int) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
DELETE FROM clients
WHERE id = $1;
	`, id)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func UpdateClientInfo(db *sql.DB, name string, tel string, bday time.Time, id int) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if name != "" {
		_, err = tx.Exec(`
UPDATE clients
SET fullname = $1
WHERE id = $2;
		`, name, id)
		if err != nil {
			return err
		}
	}

	if tel != "" {
		_, err = tx.Exec(`
UPDATE clients
SET telephone = $1
WHERE id = $2;
		`, tel, id)
		if err != nil {
			return err
		}
	}

	if !bday.IsZero() {
		_, err = tx.Exec(`
UPDATE clients
SET birthday = $1
WHERE id = $2;
		`, bday, id)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}
