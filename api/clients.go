package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"web-application/database"
	"web-application/logger"
)

func ClientsHandler(rw http.ResponseWriter, req *http.Request, db *sql.DB, logger *logger.Logger) {
	var err error
	var res []byte

	switch req.Method {
	case "POST", "PUT":
		res, err = CreateClient(req, db)
	case "PATCH":
		res, err = UpdateClient(req, db)
	case "DELETE":
		res, err = DeleteClient(req, db)
	default:
		res, err = SelectClient(req, db)
	}

	if err != nil {
		data, _ := json.Marshal(map[string]interface{}{"OK": false, "Error": err.Error()})
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(data)
		logger.Info(fmt.Sprintf("[ %s 400 ] %s : %s", req.Method, req.URL, err))
		return
	}

	if res == nil && len(res) == 0 {
		http.NotFound(rw, req)
		logger.Info(fmt.Sprintf("[ %s 404 ] %s", req.Method, req.URL))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
	logger.Info(fmt.Sprintf("[ %s 200 ] %s : %d bytes", req.Method, req.URL, len(res)))
}

func CreateClient(req *http.Request, db *sql.DB) ([]byte, error) {
	var err error
	var data []byte

	if data, err = ioutil.ReadAll(req.Body); err != nil {
		return nil, err
	}
	defer req.Body.Close()

	item := new(food_delivery.Client)
	if err = json.Unmarshal(data, &item); err != nil {
		return nil, err
	}

	err = food_delivery.AddClient(db, item)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]bool{"OK": true})
}

func UpdateClient(req *http.Request, db *sql.DB) ([]byte, error) {
	args := req.URL.Query()

	id, err := getID(&args)
	if err != nil {
		return nil, err
	}

	var data []byte

	if data, err = ioutil.ReadAll(req.Body); err != nil {
		return nil, err
	}
	defer req.Body.Close()

	item := new(food_delivery.ClientInfo)
	if err = json.Unmarshal(data, &item); err != nil {
		return nil, err
	}
	fmt.Print(item)
	err = food_delivery.UpdateClientInfo(db, item.Name, item.Tel, item.Bday, id)
	if err != nil {
		return nil, err
	}

	return json.Marshal(map[string]bool{"OK": true})
}

func DeleteClient(req *http.Request, db *sql.DB) ([]byte, error) {
	args := req.URL.Query()

	id, err := getID(&args)
	if err != nil {
		return nil, err
	}

	err = food_delivery.DeleteClient(db, id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]bool{"OK": true})
}

func SelectClient(req *http.Request, db *sql.DB) ([]byte, error) {
	args := req.URL.Query()

	if id, err := getID(&args); err == nil {
		item, err := food_delivery.GetClient(db, id)

		if err != nil {
			return nil, err
		}
		return json.Marshal(item)
	}

	item, err := food_delivery.GetAllClients(db)
	if err != nil {
		return nil, err
	}
	return json.Marshal(item)
}

func getID(args *url.Values) (int, error) {
	var err error

	if id_str := args.Get("id"); id_str != "" {
		var id uint64
		if id, err = strconv.ParseUint(id_str, 10, 64); err != nil {
			return 0, err
		}
		return int(id), nil
	}
	return 0, fmt.Errorf("Не указан идентификатор")
}
