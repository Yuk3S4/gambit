package bd

import (
	"fmt"

	"github.com/Yuk3S4/gambit/models"
)

func InsertAddress(addr models.Address, User string) error {
	fmt.Println("Comienza el registro InsertAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name)"
	sentencia += " VALUES ('" + User + "', '" + addr.AddAdress + "', '" + addr.AddCity + "', '" + addr.AddState + "', '" + addr.AddPostalCode + "', '" + addr.AddPhone + "', '" + addr.AddTitle + "', '" + addr.AddName + "')"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Insert Address > Ejecución Exitosa")
	return nil
}
