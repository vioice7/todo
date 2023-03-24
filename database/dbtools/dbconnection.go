package dbtools

import (
	"database/sql"
	"fmt"
	"log"
	"todo/database/model"

	_ "github.com/go-sql-driver/mysql"
)

//
// Conection and initialase mysql
//

var driverName string
var dataSourceName string

func DBInitilize(dn, dsn string) {

	driverName = dn
	dataSourceName = dsn

}

func connect() (db *sql.DB) {

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

// -------------------------
// Todo table connection
// -------------------------

func SelectAllTodos() []model.Todo {

	db := connect()

	rows, err := db.Query("select * from todo")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	todos := []model.Todo{}

	for rows.Next() {

		todo := model.Todo{}

		err = rows.Scan(&todo.ID, &todo.Item, &todo.Done)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		todos = append(todos, todo)
	}

	return todos

}

func SelectTodoBasedId(id string) (model.Todo, error) {

	db := connect()

	rows := db.QueryRow("select * from todo where id = ?", id)

	defer db.Close()

	todo := model.Todo{}

	err := rows.Scan(&todo.ID, &todo.Item, &todo.Done)

	return todo, err

}

func SaveTodo(todo model.Todo) int64 {

	db := connect()

	defer db.Close()

	save, err := db.Prepare("insert into todo(id,item,done) values(?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := save.Exec(todo.ID, todo.Item, todo.Done)

	if err != nil {
		log.Fatal(err.Error())
	}

	todoID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	return todoID
}

func UpdateTodo(todo model.Todo) int64 {

	db := connect()

	defer db.Close()

	update, err := db.Prepare("update todo set item=?, done=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(todo.Item, todo.Done, todo.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsEffected

}

func DeleteTodoId(id int) int64 {

	db := connect()

	defer db.Close()

	delete, err := db.Prepare("delete from todo where id = ?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := delete.Exec(id)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}

func DeleteAllTodos() int64 {

	db := connect()

	defer db.Close()

	// count records before delete table

	var count int64

	err := db.QueryRow("select count(*) from todo").Scan(&count)

	if err != nil {
		log.Fatal(err.Error())
	}

	// delete table to delete all records

	deleteAll, err := db.Prepare("delete from todo where id > 0")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = deleteAll.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	alterTable, err := db.Prepare("alter table todo auto_increment = 1")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = alterTable.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	return count

}

func SaveMultipleTodos(todos []model.Todo) int64 {

	db := connect()

	defer db.Close()

	var nrAddedTodos int64

	for _, todo := range todos {

		save, err := db.Prepare("insert into todo(id,item,done) values(?,?,?)")

		if err != nil {
			log.Fatal(err.Error())
		}

		result, err := save.Exec(todo.ID, todo.Item, todo.Done)

		if err != nil {
			log.Fatal(err.Error())
		}

		todoID, err := result.LastInsertId()

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Todo ID ", todoID, " inserted!")

	}

	return nrAddedTodos

}

func UpdateCheck(id int) int64 {

	db := connect()

	defer db.Close()

	// check true or false

	rows := db.QueryRow("select id, done from todo where id = ?", id)

	defer db.Close()

	todoCheck := model.Todo{}

	err := rows.Scan(&todoCheck.ID, &todoCheck.Done)

	if err != nil {
		log.Fatal(err.Error())
	}

	// set true or false

	if todoCheck.Done {
		todoCheck.Done = false
	} else {
		todoCheck.Done = true
	}

	update, err := db.Prepare("update todo set done=? where id=?")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(todoCheck.Done, todoCheck.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsEffected

}
