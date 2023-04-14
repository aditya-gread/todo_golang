package models

import (
	"database/sql"
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

// initializes db
func init() {
	config.Connect()
	db = config.GetDB()
}

// add data to db
func CreateTodo(todo *Todo) *Todo {
	db.Exec("INSERT INTO Todo (title, description, status) VALUES ( ?, ?,?)", todo.Title, todo.Description, todo.Status)
	return todo
}

// retrieve all data from db
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
	return restructure(todos), nil
}

// retrieves single row fom db
func GetTodoByID(ID int) (*Todo, error) {
	var to_do Todo
	row := db.QueryRow("select * from Todo where id = ?", ID)
	if err := row.Scan(&to_do.Id, &to_do.Title, &to_do.Description); err != nil {
		return nil, err
	}
	return &to_do, nil
}

// deletes data by id from db
func DeleteTodo(ID int) {
	db.Exec("delete from Todo where id =?", ID)
}

// updates data by id in db
func UpdateTodo(todo *Todo, ID int) {
	_, err := db.Exec("update Todo set title = ?, description =?, status =? where Id =?", todo.Title, todo.Description, todo.Status, ID)
	if err != nil {
		return
	}
}

// searches data by id in db
func SearchTodo(word string) []Todo {
	var to_do []Todo
	rows, err := db.Query("select * from Todo where title LIKE ? or description LIKE ?", "%"+word+"%", "%"+word+"%")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status); err != nil {
			return nil
		}
		to_do = append(to_do, todo)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return restructure(to_do)
}

func restructure(todos []Todo) []Todo {
	i := 0
	var to_dos []Todo
	for k := 0; k < len(todos); k++ {
		if todos[k].Status {
			to_dos = append(to_dos, todos[k])
		} else {
			to_dos = append(to_dos[:i+1], to_dos[i:]...)
			to_dos[i] = todos[k]
			i++
		}
	}
	return to_dos
}
