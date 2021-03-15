package dbUtils

import (
	"database/sql"
)

/*func DbConn() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}
*/
func DbConn() *sql.DB {
	/*dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")*/

	//db, err := sql.Open("mysql", "sammy:Password@123@tcp(127.0.0.1:3306)/test")
	db, err := sql.Open("mysql", "harishkokcha:Healthians@123@tcp(localhost:3306)/test")
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"dbPort+")/"+dbDatabase)
	if err != nil {
		panic(err.Error())
	}
	return db
}
