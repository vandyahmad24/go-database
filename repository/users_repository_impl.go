package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vandyahmad24/go-database-mysql/model"
	"strconv"
)

type usersRepositoryImpl struct {
	DB *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository{
	return &usersRepositoryImpl{DB:db}
}


func (u *usersRepositoryImpl) DeleteByID(ctx context.Context, id int32) (error) {
	script := "DELETE FROM users WHERE id=?"
	rows, err := u.DB.QueryContext(ctx,script,id)
	if err!=nil{
		return err
	}
	defer rows.Close()
	return nil


}


func (u *usersRepositoryImpl) Insert(ctx context.Context, users model.Users) (model.Users, error) {
	script:="INSERT INTO users(username, password) value (?,?)"
	result, err := u.DB.ExecContext(ctx,script,users.Username, users.Password)
	if err!=nil{
		return users, err
	}
	id,err := result.LastInsertId()
	if err!=nil{
		return users, err
	}
	users.Id = int32(id)
	return users,nil
}

func (u *usersRepositoryImpl) FindByID(ctx context.Context, id int32) (model.Users, error) {
	script := "SELECT id, username,password FROM users WHERE id=? LIMIT 1"
	rows, err := u.DB.QueryContext(ctx,script,id)
	users:= model.Users{}
	if err!=nil{
		return users, err
	}
	defer rows.Close()
	if(rows.Next()){
	//	ada
		rows.Scan(&users.Id,&users.Username,&users.Password)
		return users,nil
	}else{
	//	tidaka da
		return users,errors.New("Id "+strconv.Itoa(int(id)) +" not found")
	}


}

func (u *usersRepositoryImpl) FindAll(ctx context.Context) ([]model.Users, error) {
	script := "SELECT id, username,password FROM users "
	rows, err := u.DB.QueryContext(ctx,script)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()
	var users []model.Users
	for rows.Next(){
		user := model.Users{}
		rows.Scan(&user.Id,&user.Username, &user.Password)
		users = append(users,user)
	}
	return users,nil
}

