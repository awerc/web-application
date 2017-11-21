package food_delivery

import (
	"database/sql"
	"fmt"
	"time"
)

type Orders struct {
	Id         int
	Client     int
	Order_time time.Time
}

type CostForOrders struct {
	Order int
	Cost  int
}

func (this *Orders) Println() {
	fmt.Printf("id: %d client: %d order_time: %s\n", this.Id, this.Client, this.Order_time.Format("2 Jan 2006"))
}
func (this *Products) Println() {
	fmt.Printf("id: %d type: %s cost: %d rubles %d kopec name: %v weight: %v \n",
		this.Id, this.Type, this.Cost/100, this.Cost%100, this.Name, this.Weight)
}
func (this *CostForOrders) Println() {
	fmt.Printf("order: %d cost: %d rubles %d kopec\n",
		this.Order, this.Cost/100, this.Cost%100)
}

func GetOrderForClient(db *sql.DB, client int) (list []*Orders, err error) {
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT id, client, order_time 
		FROM orders  
		WHERE client = $1;
    	`, client)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list = make([]*Orders, 0)
	for rows.Next() {
		order := new(Orders)
		err = rows.Scan(&order.Id, &order.Client, &order.Order_time)
		if err != nil {
			return nil, err
		}
		list = append(list, order)
	}
	return list, nil
}

func GetAllProducts(db *sql.DB, order int) (list []*Products, err error) {
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT id, type, cost, name, weight 
		FROM products  
		WHERE id in (
			SELECT product_id
			FROM products_in_orders
			WHERE order_id = $1
			);		
    	`, order)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list = make([]*Products, 0)
	for rows.Next() {
		product := new(Products)
		err = rows.Scan(&product.Id, &product.Type, &product.Cost, &product.Name, &product.Weight)
		if err != nil {
			fmt.Print(err)
			return nil, err
		}
		list = append(list, product)
	}
	return list, nil
}

func GetCostOfAllOrdersForClient(db *sql.DB, client int) (list []*CostForOrders, err error) {
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
SELECT order_id, sum(cost*count) as cost
FROM (
  SELECT cost, id
  FROM products
  WHERE id = ANY (
    SELECT product_id
    FROM products_in_orders
    WHERE order_id = ANY (
      SELECT id
      FROM orders
      WHERE client = $1
      )
    )
  ) as res1
INNER JOIN
  (SELECT order_id, product_id, count
    FROM products_in_orders
    WHERE order_id = ANY (
      SELECT id
      FROM orders
      WHERE client = $1)
    ) as res2
    ON id = product_id
  GROUP BY order_id; 	
    	`, client)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list = make([]*CostForOrders, 0)
	for rows.Next() {
		result := new(CostForOrders)
		err = rows.Scan(&result.Order, &result.Cost)
		if err != nil {
			fmt.Print(err)
			return nil, err
		}
		list = append(list, result)
	}
	return list, nil
}
