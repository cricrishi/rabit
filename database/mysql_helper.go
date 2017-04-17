package database

import (
	//"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	conf "github.com/rabit/config"
	dbModel "github.com/rabit/database/models"
)

func GetConncetion() (*sql.DB, error) {
	c := conf.ConfigData["config"].Get("mysql")
	mysqlConfig, _ := c.(map[string]interface{})
	db, err := sql.Open("mysql", mysqlConfig["username"].(string)+":"+mysqlConfig["password"].(string)+"@/"+mysqlConfig["database"].(string))
	return db, err
}

func GetUserDetail() (error, []dbModel.User) {
	var users []dbModel.User

	db, e := GetConncetion()
	if e != nil {
		return e, users
	}
	sqlQuery := "SELECT * FROM User"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return err, users
	}
	defer rows.Close() //after function is complete this statment close the rows tosave memory
	for rows.Next() {
		var u dbModel.User
		rows.Scan(&u.FirstName, &u.LastName, &u.Email, &u.Task)
		users = append(users, u)
	}
	return nil, users
}

func GetUserFromEmail(email string) (error,dbModel.Users) {
	var u dbModel.Users
	db, e := GetConncetion()
	if e != nil {
		return e, u
	}
	sqlQuery := "SELECT user_name,first_name,last_name,email,password  FROM Users where email = ? limit 1"
	rows, err := db.Query(sqlQuery,email)
	if err != nil {
		return err, u
	}
	defer rows.Close() //after function is complete this statment close the rows tosave memory
	for rows.Next() {
		rows.Scan(&u.UserName, &u.FirstName, &u.LastName,&u.Email, &u.Password)
		}
		//fmt.Println(u)
	return nil, u	
}
