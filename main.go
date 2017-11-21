package main

import (
	"flag"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"os"
	"web-application/database"
	"web-application/logger"
	"web-application/xml_parser"
	"web-application/xml_validator"
)

var (
	configFilename string //String that contains path to config file
	logFilename    string //String that contains path to log file
)

//Reading command line parameters before programm execution
func init() {
	flag.StringVar(&configFilename, "config", "config.xml", "name of config file")
	//Command line parameter for config file. Default is config.xml
	flag.StringVar(&logFilename, "log", "log.txt", "name of log file")
	//Command line parameter for log file. Default is log.txt

	flag.Parse()
}

func check_error(err error, terminate bool, logger *logger.Logger) {
	if err != nil {
		logger.Info(err)
		if terminate {
			os.Exit(1)
		}
	}
	return
}

func main() {
	var logger = logger.Initialize("[WEB APP]", logger.Debug|logger.Info, logger.DateTime|logger.Shortfile)
	defer logger.Close()
	logger.AddOutput(os.Stdout)
	logger.AddOutputFile(logFilename)

	configuration, err := xml_parser.Parse(configFilename)
	check_error(err, true, logger)

	err = xml_validator.Validate(configuration)
	check_error(err, true, logger)

	db, err := food_delivery.Connect(configuration)
	check_error(err, true, logger)
	defer db.Close()

	err = db.Ping()
	check_error(err, true, logger)

	orders, err := food_delivery.GetOrderForClient(db, 1)
	check_error(err, true, logger)

	fmt.Println("\n------------Orders for 1st client-------------")
	for _, element := range orders {
		element.Println()
	}

	products, err := food_delivery.GetAllProducts(db, 1)
	check_error(err, true, logger)

	fmt.Println("\n-----------All avlailable products--------------")
	for _, element := range products {
		element.Println()
	}

	costs, err := food_delivery.GetCostOfAllOrdersForClient(db, 1)
	check_error(err, true, logger)

	fmt.Println("\n--------Total cost of 1st client orders----------")
	for _, element := range costs {
		element.Println()
	}

	clients, err := food_delivery.GetAllClients(db)
	check_error(err, true, logger)

	fmt.Println("\n----------------Clients-----------------")
	for _, element := range clients {
		fmt.Printf("name: %s tel: %s bday: %s\n", element.Name, element.Tel, element.Bday.Format("2 Jan 2006"))
	}

	// err = food_delivery.AddClient(db, "login7", "pass", "88005553535", "1991-01-22", "client8")
	// check_error(err, true, logger)

	// err = food_delivery.ChangeClientTelephone(db, 7, "877454887")
	// check_error(err, true, logger)
	//spew.Dump(configuration)
}
