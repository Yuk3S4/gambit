package handlers

import (
	"fmt"
	"strconv"

	"github.com/Yuk3S4/gambit/auth"
	"github.com/Yuk3S4/gambit/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(path, method, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Voy a procesar ", path, " > ", method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOK, statusCode, user := validoAuthorization(path, method, headers)
	if !isOK {
		return statusCode, user
	}

	switch path[1:5] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)

	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)

	case "stoc":
		return ProcesoStock(body, path, method, user, idn, request)

	case "addr":
		return ProcesoAddress(body, path, method, user, idn, request)

	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)

	case "orde":
		return ProcesoOrder(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid"
}

func validoAuthorization(path, method string, headers map[string]string) (bool, int, string) {
	if (path == "/product" && method == "GET") ||
		(path == "/category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido"
	}

	todoOK, err, msg := auth.ValidoToken(token)
	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token ", err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token ", msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg
}

func ProcesoUsers(body, path, method, user, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	if path == "/user/me" {
		switch method {
		case "PUT":
			return routers.UpdateUser(body, user)
		case "GET":
			return routers.SelectUser(body, user)

		}
	}
	if path == "/users" {
		if method == "GET" {
			return routers.SelectUsers(body, user, request)
		}
	}

	return 400, "Method Invalid"
}

func ProcesoProducts(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		return routers.SelectProduct(request)
	case "POST":
		return routers.InsertProduct(body, user)
	case "PUT":
		return routers.UpdateProduct(body, user, id)
	case "DELETE":
		return routers.DeleteProduct(user, id)
	}

	return 400, "Method Invalid"
}

func ProcesoCategory(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "GET":
		return routers.SelectCategories(body, request)
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(body, user, id)
	}

	return 400, "Method Invalid"
}

func ProcesoStock(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return routers.UpdateStock(body, user, id)
}

func ProcesoAddress(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertAddress(body, user)
	}

	return 400, "Method Invalid"
}

func ProcesoOrder(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}
