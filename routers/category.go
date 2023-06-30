package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Yuk3S4/gambit/bd"
	"github.com/Yuk3S4/gambit/models"
)

func InsertCategory(body, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CateName) == 0 {
		return 400, "Debe de especificar el Nombre (Title) de la Categoría"
	}
	if len(t.CatePath) == 0 {
		return 400, "Debe de especificar el Path (Ruta) de la Categoría"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err := bd.InsertCategory(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoría " + t.CateName + " > " + err.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}
