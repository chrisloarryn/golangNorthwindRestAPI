package main

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"golangNorthwindRestAPI/database"

	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"

	_ "github.com/go-sql-driver/mysql"
)

var databaseConnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func catch (err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	//Logica

	r := chi.NewRouter()
	
	r.Get("/products", AllProducts)

	http.ListenAndServe(":5000", r)

	//fmt.Println(databaseConnection)
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id,product_code,COALESCE(description,'')
				 FROM products`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, products)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
