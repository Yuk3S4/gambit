package bd

import (
	"database/sql"
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

func SelectUser(UserId string) (models.User, error) {
	fmt.Println("Comienza SelectUser")
	user := models.User{}

	err := DbConnect()
	if err != nil {
		return user, err
	}
	defer Db.Close()

	sentencia := "SELECT * FROM users WHERE User_UUID = '" + UserId + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime

	rows.Scan(&user.UserUUID, &user.UserEmail, &firstName, &lastName, &user.UserStatus, &user.UserDateAdd, &dateUpg)

	user.UserFirstName = firstName.String
	user.UserLastName = lastName.String
	user.UserDateUpd = dateUpg.Time.String()

	fmt.Println("Select User > Ejecución Exitosa")

	return user, nil
}
