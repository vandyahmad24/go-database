package Go_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB{
	db, err := sql.Open("mysql","root:@tcp(localhost:3310)/godb?parseTime=true")
	if err != nil{
		panic(err)
	}
	//minimal koneksi awal
	db.SetMaxIdleConns(10)
	//maksimal koneksi
	db.SetMaxOpenConns(100)
	//waktu begong
	db.SetConnMaxIdleTime(5*time.Minute)
	//koneksi apapun jika sudah 60 diganti koneksi baru
	db.SetConnMaxLifetime(20*time.Minute)
	return db
}
