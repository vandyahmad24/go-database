package Go_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql","root:@tcp(localhost:3310)/godb")
	if err != nil{
		panic(err)
	}
//	konfigurasi database pooling


//	close
	defer db.Close()

}