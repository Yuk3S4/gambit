package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Yuk3S4/gambit/bd"
	"github.com/Yuk3S4/gambit/models"
	"github.com/aws/aws-lambda-go/events"
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

func SelectOrders(user string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var fechaDesde, fechaHasta string
	var orderId int
	var page int

	if len(request.QueryStringParameters["fechaDesde"]) > 0 {
		fechaDesde = request.QueryStringParameters["fechaDesde"]
	}
	if len(request.QueryStringParameters["fechaHasta"]) > 0 {
		fechaHasta = request.QueryStringParameters["fechaHasta"]
	}
	if len(request.QueryStringParameters["page"]) > 0 {
		page, _ = strconv.Atoi(request.QueryStringParameters["page"])
	}
	if len(request.QueryStringParameters["orderId"]) > 0 {
		orderId, _ = strconv.Atoi(request.QueryStringParameters["orderId"])
	}

	result, err := bd.SelectOrders(user, fechaDesde, fechaHasta, page, orderId)
	if err != nil {
		return 400, "Ocurrió un error al intentar capturar los registros de órdenes del " + fechaDesde + " al " + fechaHasta + " > " + err.Error()
	}

	orders, err := json.Marshal(result)
	if err != nil {
		return 400, "Ocurrió un error al intentar convertir en JSON el registro de Órden"
	}

	return 200, string(orders)
}
