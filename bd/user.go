package bd

import (
	"fmt"

	"github.com/Yuk3S4/gambit/models"
	"github.com/Yuk3S4/gambit/tools"
)

func UpdateUser(u models.User, User string) error {
	fmt.Println("Comienza UpdateUser")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE users SET "

	coma := ""
	if len(u.UserFirstName) > 0 {
		coma = ","
		sentencia += "User_FirstName = '" + u.UserFirstName + "'"
	}

	if len(u.UserLastName) > 0 {
		sentencia += coma + "User_LastName = '" + u.UserLastName + "'"
	}

	sentencia += ", User_DateUpg = '" + tools.FechaMySQL() + "' WHERE User_UUID ='" + User + "'"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update User > Ejecución Exitosa")
	return nil
}
