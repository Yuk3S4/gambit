package routers

import (
	"encoding/json"

	"github.com/Yuk3S4/gambit/bd"
	"github.com/Yuk3S4/gambit/models"
)

func UpdateUser(body string, User string) (int, string) {
	var t models.User

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.UserFirstName) == 0 && len(t.UserLastName) == 0 {
		return 400, "Debe especificar el Nombre (FirstName) y (LastName) del usuario"
	}

	_, encontrado := bd.UserExists(User)
	if !encontrado {
		return 404, "No existe un usuario con ese UUID '" + User + "'"
	}

	err = bd.UpdateUser(t, User)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar la actualización del usuario " + User + " > " + err.Error()
	}

	return 200, "Update User OK"
}
