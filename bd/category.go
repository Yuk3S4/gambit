package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/Yuk3S4/gambit/models"
	"github.com/Yuk3S4/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza Registro de InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CateName + "','" + c.CatePath + "')"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("Insert Category > Ejecución Exitosa")
	return lastInsertId, nil
}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza Registro de UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE category SET "

	if len(c.CateName) > 0 {
		sentencia += " Categ_Name = '" + tools.EscapeString(c.CateName) + "'"
	}

	if len(c.CatePath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += ", "
		}
		sentencia += "Categ_Path = '" + tools.EscapeString(c.CatePath) + "'"
	}

	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecución Exitosa")
	return nil
}

func DeleteCategory(id int) error {
	fmt.Println("Comienza Registro de DeleteCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "DELETE FROM category WHERE Categ_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Ejecución Exitosa")
	return nil
}

func SelectCategories(categId int, slug string) ([]models.Category, error) {
	fmt.Println("Comienza SelectCategories")

	var Categ []models.Category

	err := DbConnect()
	if err != nil {
		return Categ, err
	}
	defer Db.Close()

	sentencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "

	if categId > 0 {
		sentencia += "WHERE Categ_Id = " + strconv.Itoa(categId)
	} else {
		if len(slug) > 0 {
			sentencia += "WHERE Categ_Path LIKE '%" + slug + "%'"
		}
	}

	fmt.Println(sentencia)

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)

	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err = rows.Scan(&categId, &categName, &categPath)
		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CateName = categName.String
		c.CatePath = categPath.String

		Categ = append(Categ, c)
	}

	fmt.Println("Select Category > Ejecución Exitosa")
	return Categ, nil
}
