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

func UpdateCategory(body, User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CateName) == 0 && len(t.CatePath) == 0 {
		return 400, "Debe de especificar CateName y CatePath para actualizar"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.CategID = id
	err = bd.UpdateCategory(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar hacer el UPDATE de la categoria " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Update OK"
}

func DeleteCategory(body, User string, id int) (int, string) {
	if id == 0 {
		return 400, "Debe especificar el ID de la Categoría a borrar"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteCategory(id)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el DELETE de la Categoría " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Delete OK"
}
