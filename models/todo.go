package models

import (
	"database/sql"
	"fmt"
	"main/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Todo struct {
	Id          int    `json: id`
	Title       string `json: title`
	Description string `json: description`
	Status      bool   `jon: status`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func CreateTodo(todo *Todo) *Todo {
	db.Exec("INSERT INTO Todo (title, description, status) VALUES ( ?, ?,?)", todo.Title, todo.Description, todo.Status)
	return todo
}

func GetAllToDo() ([]Todo, error) {
	var todos []Todo
	rows, err := db.Query("Select * from Todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodoByID(ID int) (*Todo, error) {
	var to_do Todo
	row := db.QueryRow("select * from Todo where id = ?", ID)
	if err := row.Scan(&to_do.Id, &to_do.Title, &to_do.Description); err != nil {
		return nil, err
	}
	return &to_do, nil
}

func DeleteTodo(ID int) {
	db.Exec("delete from Todo where id =?", ID)
}

func UpdateTodo(todo *Todo, ID int) {
	stmt, err := db.Exec("update Todo set title = ?, description =?, status =? where Id =?", todo.Title, todo.Description, todo.Status, ID)
	if err != nil {
		return
	}
	fmt.Print(stmt)
}
