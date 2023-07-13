package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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

func AddressExists(User string, id int) (error, bool) {
	fmt.Println("Comienza AddressExists")

	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM addresses WHERE Add_Id = " + strconv.Itoa(id) + " AND Add_UserId = '" + User + "'"

	rows, err := Db.Query(sentencia)
	if err != nil {
		return err, false
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("AddressExists > Ejecución Exitosa - valor devuelto ", valor)
	if valor == "1" {
		return nil, true
	}

	return nil, false
}

func UpdateAddress(addr models.Address) error {
	fmt.Println("Comienza UpdateAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE addresses SET "

	if addr.AddAdress != "" {
		sentencia += "Add_Address = '" + addr.AddAdress + "', "
	}
	if addr.AddCity != "" {
		sentencia += "Add_City = '" + addr.AddCity + "', "
	}
	if addr.AddName != "" {
		sentencia += "Add_Name = '" + addr.AddName + "', "
	}
	if addr.AddPhone != "" {
		sentencia += "Add_Phone = '" + addr.AddPhone + "', "
	}
	if addr.AddPostalCode != "" {
		sentencia += "Add_PostalCode = '" + addr.AddPostalCode + "', "
	}
	if addr.AddState != "" {
		sentencia += "Add_State = '" + addr.AddState + "', "
	}
	if addr.AddTitle != "" {
		sentencia += "Add_Title = '" + addr.AddTitle + "', "
	}

	sentencia, _ = strings.CutSuffix(sentencia, ", ")
	sentencia += " WHERE Add_Id = " + strconv.Itoa(addr.AddId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Address > Ejecución Exitosa")
	return nil
}

func DeleteAddress(id int) error {
	fmt.Println("Comienza Registro de DeleteAddress")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Address > Ejecución Exitosa")
	return nil
}

func SelectAddress(User string) ([]models.Address, error) {
	fmt.Println("Comienza SelectAddress")

	var addrs []models.Address

	err := DbConnect()
	if err != nil {
		return addrs, err
	}
	defer Db.Close()

	sentencia := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name FROM addresses WHERE Add_UserId = '" + User + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		fmt.Println(err)
		return addrs, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Address
		var addId sql.NullInt16
		var addAddress sql.NullString
		var addCity sql.NullString
		var addState sql.NullString
		var addPostalCode sql.NullString
		var addPhone sql.NullString
		var addTitle sql.NullString
		var addName sql.NullString

		err := rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)
		if err != nil {
			return addrs, err
		}

		a.AddId = int(addId.Int16)
		a.AddAdress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String
		addrs = append(addrs, a)
	}

	fmt.Println("Select Addresses > Ejecución Exitosa")
	return addrs, nil
}
