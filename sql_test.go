package Go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()
	script:="INSERT INTO users(username, password) VALUES('budy','budy')"
	_, err := db.ExecContext(ctx,script)
	if err != nil{
		panic(err)
	}
	fmt.Println("Skses Insert new customer")
}

func TestQuerySql(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()
	script:="SELECT id, name, email, balance, rating, birthday, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx,script)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id, balance int32
		var name string
		var email sql.NullString
		var rating float32
		var created_at time.Time
		var birthday sql.NullTime
		var married bool
		rows.Scan(&id, &name, &email, &balance, &rating, &birthday, &married, &created_at)
		if err != nil{
			panic(err)
		}


		fmt.Println("ID:",id,"Name:",name, "Email:", email.String, "Balance:", balance, "Birthday:",birthday.Time,"Married:",married,"Created At:",created_at)

	}

	fmt.Println("Sudah Dapat data")
}


func TestSqlInjection(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()
	username:="admin"
	password:="admin"


	script:="SELECT username from users WHERE username = ? AND password = ? LIMIT 1 "
	//fmt.Println(script)
	rows, err := db.QueryContext(ctx,script, username, password)
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err !=nil{
			panic(err)
		}
		fmt.Println("Sukses login",username)
	}else{
		fmt.Println("gagal login")
	}

	fmt.Println("Sudah Dapat data")
}



func TestContextSqlInsert(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()
	script:="INSERT INTO users(username, password) VALUES(?,?)"
	result, err := db.ExecContext(ctx,script,"Ahmad","Ahmad")
	if err != nil{
		panic(err)
	}
	lastId, err:= result.LastInsertId()
	if err != nil{
		panic(err)
	}
	fmt.Println("ID Terakhir",lastId)
	fmt.Println("Skses Insert new customer")
}

func TestPrepareStatment(t *testing.T) {
	db:= GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO users(username, password) VALUES(?,?)"
	statment, err:= db.PrepareContext(ctx,script)
	if err != nil{
		panic(err)
	}
	defer statment.Close()
	for i:=1; i<=10; i++{
		username := "Vandy ke-"+strconv.Itoa(i)
		password := "Vandy Password ke-"+strconv.Itoa(i)
		result, err:=statment.ExecContext(ctx,username,password)
		if err != nil{
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("IDnya",id)
		
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx,err:=db.Begin()
	if err !=nil{
		panic(err)
	}
//	do transaction
	script := "INSERT INTO users(username, password) VALUES(?,?)"
	for i:=1; i<=10; i++{
		username := "Vandy ke-"+strconv.Itoa(i)
		password := "Vandy Password ke-"+strconv.Itoa(i)
		result, err:=tx.ExecContext(ctx,script,username,password)
		if err != nil{
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("IDnya",id)

	}
	err =tx.Commit()
	if err!=nil{
		panic(err)
	}


}