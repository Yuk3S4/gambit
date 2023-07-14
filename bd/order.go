package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Yuk3S4/gambit/models"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza Registro InsertOrder")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES ('"
	sentencia += o.Order_UserUUID + "', " + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + ")"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	for _, od := range o.OrderDetails {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(lastInsertId))
		sentencia += ", " + strconv.Itoa(od.OD_ProdId) + ", " + strconv.Itoa(od.OD_Quantity) + ", " + strconv.FormatFloat(od.OD_Price, 'f', -1, 64) + ")"

		fmt.Println(sentencia)
		_, err = Db.Exec(sentencia)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("InsertOrder > Ejecución exitosa")
	return lastInsertId, nil
}
