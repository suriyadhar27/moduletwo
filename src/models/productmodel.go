package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"task/src/config"
	"task/src/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Student, error) {

	db, err := config.GetDB()
	fmt.Println(db)
	log.Println(db.Ping())
	if err != nil {
		return nil, err
	} else {

		rows, err2 := db.Query("select * from student")
		fmt.Println(rows)
		fmt.Println(err2)
		if err2 != nil {
			return nil, err
		} else {

			var students []entities.Student

			for rows.Next() {
				var student entities.Student
				rows.Scan(&student.Id, &student.Title, &student.Description, &student.Duedate, &student.Status)

				fmt.Println(student)
				students = append(students, student)
			}
			fmt.Println(students)
			return students, nil

		}
	}

}

func (*ProductModel) Search() ([]entities.Student, error) {

	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from student where id = ?")
		if err2 != nil {
			return nil, err2
		} else {

			var students []entities.Student
			for rows.Next() {
				var student entities.Student
				rows.Scan(&student.Id, &student.Title, &student.Description, &student.Duedate, &student.Status)

				fmt.Println(student)
				students = append(students, student)
			}
			fmt.Println(students)
			return students, nil

		}
	}

}
func (*ProductModel) Find(id int64) (entities.Student, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Student{}, err
	} else {
		rows, err2 := db.Query("select * from student where id = ?", id)
		if err2 != nil {
			return entities.Student{}, err2
		} else {
			var student entities.Student
			for rows.Next() {
				rows.Scan(&student.Id, &student.Title, &student.Description, &student.Duedate, &student.Status)

			}

			return student, nil

		}
	}

}
func (*ProductModel) Create(student *entities.Student) string {

	db, err := config.GetDB()
	result, err := db.Exec("insert into student(Id, Title, Description, Duedate, status) values(?, ?, ?, ?, ?)", student.Id, student.Title, student.Description, student.Duedate, student.Status)
	if err != nil {
		return "pending"
	}
	lastId, _ := result.LastInsertId()
	student.Id = lastId
	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return "pending"
	}
	affected := rowsAffected > 0

	return strconv.FormatBool(affected) //return rowsAffected > 0
}

func (*ProductModel) Delete(id int64) bool {

	db, err := config.GetDB()
	result, err := db.Exec("delete from student where id = ?", id)
	if err != nil {
		return false
	}
	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return false
	}
	return rowsAffected > 0
}

// func (*ProductModel) Update(student entities.Student) string {

// 	db, err := config.GetDB()
// 	result, err := db.Exec("update student set title = ?, description = ?, duedate = ?, status = ?, where id = ?", student.Description, student.Id, student.Title, student.Duedate, student.Status)
// 	if err != nil {
// 		return "in progress"
// 	}
// 	rowsAffected, err2 := result.RowsAffected()
// 	if err2 != nil {
// 		return "in progress"
// 	}
// 	affected := rowsAffected > 0

// 	return strconv.FormatBool(affected) //return rowsAffected > 0

func (*ProductModel) Update(student entities.Student) (string, error) {
	db, err := config.GetDB()
	if err != nil {
		return "", err
	}

	result, err := db.Exec("UPDATE student SET title = ?, description = ?, duedate = ?, status = ? WHERE id = ?",
		student.Title, student.Description, student.Duedate, student.Status, student.Id)

	if err != nil {
		return "", err
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return "", err2
	}

	if rowsAffected == 0 {
		return "", errors.New("Task not found")
	}

	return "Update successful", nil
}
