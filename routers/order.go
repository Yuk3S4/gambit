package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Yuk3S4/gambit/bd"
	"github.com/Yuk3S4/gambit/models"
)

func InsertOrder(body, User string) (int, string) {
	var o models.Orders

	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	o.Order_UserUUID = User

	ok, message := validOrder(o)
	if !ok {
		return 400, message
	}

	result, err := bd.InsertOrder(o)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la órden " + err.Error()
	}

	return 200, "{ OrderID:" + strconv.Itoa(int(result)) + "}"
}

func validOrder(o models.Orders) (bool, string) {
	if o.Order_Total == 0 {
		return false, "Debe indicar el total de la órden"
	}

	count := 0
	for _, od := range o.OrderDetails {
		if od.OD_ProdId == 0 {
			return false, "Debe indicar el ID del producto en el detalle de la Órden"
		}
		if od.OD_Quantity == 0 {
			return false, "Debe indicar la cantidad del producto en el detalle de la Órden"
		}
		count++
	}

	if count == 0 {
		return false, "Debe indicar items en la órden"
	}

	return true, ""
}
