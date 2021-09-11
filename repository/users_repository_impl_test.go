package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	Go_database "github.com/vandyahmad24/go-database-mysql"
	"github.com/vandyahmad24/go-database-mysql/model"
	"testing"
)

func TestUsersInsert(t *testing.T) {
	usersRepository := NewUsersRepository(Go_database.GetConnection())
	ctx := context.Background()
	user := model.Users{
		Username: "Vandy Ahmad Misry",
		Password: "Ahmad misry Vandy",
	}
	result, err:=usersRepository.Insert(ctx,user)
	if err !=nil{
		panic(err)
	}
	fmt.Println(result)
}


func TestUsersByID(t *testing.T) {
	usersRepository := NewUsersRepository(Go_database.GetConnection())
	ctx := context.Background()

	result, err:=usersRepository.FindByID(ctx,3)
	if err !=nil{
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T){
	usersRepository := NewUsersRepository(Go_database.GetConnection())
	ctx := context.Background()

	users, err:=usersRepository.FindAll(ctx)
	if err !=nil{
		panic(err)
	}
	for _,v :=range users {
		fmt.Println("ID ke-",v.Id,"Value Username",v.Username)
	}


}

func TestDeleteByID(t *testing.T) {
	userRepository := NewUsersRepository(Go_database.GetConnection())
	ctx := context.Background()
	id := int32(34)
	err :=userRepository.DeleteByID(ctx, id)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Berhasil menghapus user dengan id-",id)
	}


}