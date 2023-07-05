package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Yuk3S4/gambit/models"
	"github.com/Yuk3S4/gambit/tools"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Comineza Registro de InsertProduct")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO products (Prod_Title "

	if len(p.ProdDescription) > 0 {
		sentencia += ", Prod_Description"
	}

	if p.ProdPrice > 0 {
		sentencia += ", Prod_Price"
	}

	if p.ProdCategId > 0 {
		sentencia += ", Prod_CategoryId"
	}

	if p.ProdStock > 0 {
		sentencia += ", Prod_Stock"
	}

	if len(p.ProdPath) > 0 {
		sentencia += ", Prod_Path"
	}

	sentencia += ") VALUES ('" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdDescription) + "'"
	}

	if p.ProdPrice > 0 {
		sentencia += ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}

	if p.ProdCategId > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdCategId)
	}

	if p.ProdStock > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdStock)
	}

	if len(p.ProdPath) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdPath) + "'"
	}

	sentencia += ")"

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

	fmt.Println("Insert Product > Ejecución Exitosa")
	return lastInsertId, nil
}

func UpdateProduct(p models.Product) error {
	fmt.Println("Comienza UpdateProduct")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "UPDATE products SET "

	sentencia = tools.ArmoSentencia(sentencia, "Prod_Title", p.ProdTitle, "S", 0, 0)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Description", p.ProdDescription, "S", 0, 0)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Price", "", "F", 0, p.ProdPrice)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_CategoryId", "", "N", p.ProdCategId, 0)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Stock", "", "N", p.ProdStock, 0)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Path", p.ProdPath, "S", 0, 0)

	sentencia += " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Product > Ejecución Exitosa")
	return nil
}
