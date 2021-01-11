package models

import (
	"gosample/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	UserPass string `json:"user_pass"`
}

func FindAllUser(c *gin.Context) {
	var user User
	res, err := utils.DbConn().Query("select * from test.user")

	if err == nil {
		for res.Next() {
			res.Scan(&user.Id, &user.UserName, &user.UserPass)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", row)
			fmt.Printf("%v\n", user)

		}
		//c.JSON(200, user)
	} else {
		//c.JSON(404, gin.H{"error": "user not found"})
	}
}

func FindUser(c *gin.Context) int {
	name := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Print(name + " " + password)
	var user User
	var querySql = "select * from test.user where user_name='" + name + "' and user_pass='" + password + "'"
	res, err := utils.DbConn().Query(querySql)

	if err == nil {
		for res.Next() {
			res.Scan(&user.Id, &user.UserName, &user.UserPass)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", row)
			fmt.Printf("%v\n", user)
			return user.Id
		}
		//c.JSON(200, user)
	} else {
		fmt.Println("No data found")
		return 0
		//c.JSON(404, gin.H{"error": "user not found"})
	}
	return 0
}

func CreateNewUser(c *gin.Context) bool {
	name := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Print(name + " " + password)

	var querySql = "Insert into test.user(user_name,user_pass) values ('" + name + "','" + password + "')"
	fmt.Println(querySql + "querySql")
	_, err := utils.DbConn().Query(querySql)
	if err == nil {
		return true
	} else {
		fmt.Print(err)
		return false
	}
	return false
}
