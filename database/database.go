package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:qwerty@tcp(127.0.0.1:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}