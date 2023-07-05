package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Yuk3S4/gambit/bd"
	"github.com/Yuk3S4/gambit/models"
)

func InsertProduct(body, User string) (int, string) {
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el título del producto"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err := bd.InsertProduct(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro del producto " + t.ProdTitle + " > " + err.Error()
	}

	return 200, "{ ProductID: " + strconv.Itoa(int(result)) + "}"
}
