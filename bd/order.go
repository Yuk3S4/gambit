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

func SelectOrders(user, fechaDesde, fechaHasta string, page, orderId int) ([]models.Orders, error) {
	fmt.Println("Comienza SelectOrders")

	var orders []models.Orders

	sentencia := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total, FROM orders "

	if orderId > 0 {
		sentencia += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if page == 0 {
			page = 1
		}
		if page > 1 {
			offset = 10 * (page - 1)
		}

		if len(fechaHasta) == 10 {
			fechaHasta += " 23:59:59"
		}

		var where string
		var whereUser string = " Order_UserUUID = '" + user + "'"
		if len(fechaDesde) > 0 && len(fechaHasta) > 0 {
			where += " WHERE Order_Date BETWEEN '" + fechaDesde + "' AND '" + fechaHasta
		}
		if len(where) > 0 {
			where += " AND " + whereUser
		} else {
			where += " WHERE " + whereUser
		}

		limit := " LIMIT 10 "
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset)
		}

		sentencia += where + limit
	}

	fmt.Println(sentencia)

	err := DbConnect()
	if err != nil {
		return orders, err
	}
	defer Db.Close()

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var Order models.Orders
		var orderAddId sql.NullInt32
		err := rows.Scan(&Order.Order_Id, &Order.Order_UserUUID, &orderAddId, &Order.Order_Date, &Order.Order_Total)
		if err != nil {
			return orders, err
		}
		Order.Order_AddId = int(orderAddId.Int32)

		var rowsD *sql.Rows
		sentenciaD := "SELECT OD_Id, OD_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderId = " + strconv.Itoa(Order.Order_Id)
		rowsD, err = Db.Query(sentenciaD)
		if err != nil {
			return orders, err
		}

		for rowsD.Next() {
			var OD_Id int64
			var OD_ProdId int64
			var OD_Quantity int64
			var OD_Price float64

			err = rowsD.Scan(&OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)
			if err != nil {
				return orders, err
			}

			var od models.OrdersDetail
			od.OD_Id = int(OD_Id)
			od.OD_ProdId = int(OD_ProdId)
			od.OD_Quantity = int(OD_Quantity)
			od.OD_Price = OD_Price

			Order.OrderDetails = append(Order.OrderDetails, od)
		}

		orders = append(orders, Order)

		rowsD.Close()
	}

	fmt.Println("SelectOrders > Ejecución Exitosa")
	return orders, nil
}
