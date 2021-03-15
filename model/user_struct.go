package model

import (
	"fmt"
	"gosample/dbUtils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPhone    string `json:"user_phone"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserSchoolID string `json:"user_school_id"`

	UserType   string `json:"user_type"`
	UserStatus string `json:"user_status"`

	UserDeviceInfo string `json:"user_device_info"`

	UserCreatedAt string `json:"user_created_at"`
	UserUpdatedAt string `json:"user_updated_at"`
}

func GetUserList(c *gin.Context) {
	var classesList = getUsers(c)

	if classesList != nil {
		c.JSON(200, gin.H{"data": classesList, "msg": "data found succes", "status": true})
	} else {
		c.JSON(404, gin.H{"data": nil, "msg": "data not found ", "status": false})
	}
}

func getUsers(c *gin.Context) []User {

	var request Request

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	fmt.Println(request.SchoolID)
	var queryString string
	if request.SchoolID != "" {
		queryString = "Select * from test.user_master where user_status =1 and user_school_id=" + request.SchoolID
	} else {
		queryString = "Select * from test.user_master where user_status =1"
	}
	fmt.Println(queryString)

	res, err := dbUtils.DbConn().Query(queryString)
	var user User
	var usersList []User
	fmt.Printf("%v\n", res)
	fmt.Printf("err %v\n", err)
	if err == nil {
		for res.Next() {
			res.Scan(&user.UserID, &user.UserName, &user.UserPhone, &user.UserEmail,
				&user.UserPassword, &user.UserSchoolID, &user.UserType, &user.UserStatus,
				&user.UserDeviceInfo, &user.UserCreatedAt, &user.UserUpdatedAt)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", user)
			fmt.Printf("%v\n", row)
			usersList = append(usersList, user)
		}
		return usersList
	}
	return nil
}

func RegisterUser(c *gin.Context) {
	var user User
	fmt.Println("func RegisterUser(c ")
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("func RegisterUser(c  shouldbind")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var fetchUser User
	fetchUser = checkIfUserExist(user)
	if fetchUser != (User{}) {
		fmt.Println("func RegisterUser(c fetchUser ", fetchUser)
		c.JSON(200, gin.H{"data": fetchUser, "msg": "User Already exist", "status": false})
	} else {
		fmt.Println("func RegisterUser(c kjas ")
		response, userNew := insertUser(user)
		if response {
			c.JSON(200, gin.H{"data": userNew, "msg": "Registered Su ", "status": true})
		} else {
			c.JSON(200, gin.H{"data": nil, "msg": "Error occur", "status": false})
		}

	}
}

func insertUser(user User) (bool, User) {
	fmt.Println("func insertUser ")
	var queryPreParedStatement = "Insert into test.user_master(user_name,user_phone,user_email,user_password,user_status,user_device_info,user_created_at,user_updated_at) values(?,?,?,?,?,?,?,?)"
	currentTime := time.Now()
	//res, err := utils.DbConn().Query(queryPreParedStatement, user.UserName, user.UserPhone, user.UserEmail, user.UserPassword, "1", user.UserDeviceInfo, currentTime, currentTime)
	insForm, err := dbUtils.DbConn().Prepare(queryPreParedStatement)

	if err != nil {
		panic(err.Error())
	}
	res, err := insForm.Exec(user.UserName, user.UserPhone, user.UserEmail, user.UserPassword, "1", user.UserDeviceInfo, currentTime, currentTime)
	if err != nil {
		return false, User{}
	}
	lid, err := res.LastInsertId()
	defer dbUtils.DbConn().Close()
	if err == nil {
		fmt.Println(lid," last inserted id")
		user.UserID = fmt.Sprintf("%v", lid)
		return true, user

	} else {
		fmt.Println(err)
		return false, User{}
	}
	return false, User{}
}

func checkIfUserExist(user User) User {
	var query string
	query = "SELECT * FROM  user_master where ifnull(user_phone,'') = '" + user.UserPhone + "'  AND IFNULL(user_email,'')='" + user.UserEmail + "' AND  user_status ='1';"
	fmt.Println("func checkIfUserExist query " + query)
	res, err := dbUtils.DbConn().Query(query)
	if err == nil {
		for res.Next() {
			res.Scan(&user.UserID, &user.UserName, &user.UserPhone, &user.UserEmail,
				&user.UserPassword, &user.UserSchoolID, &user.UserType, &user.UserStatus,
				&user.UserDeviceInfo, &user.UserCreatedAt, &user.UserUpdatedAt)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", user)
			fmt.Printf("%v\n", row)
			return user
		}
		return User{}
	}
	return User{}
}

func LoginUser(c *gin.Context) {
	var user User
	fmt.Println("func RegisterUser(c ")
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("func RegisterUser(c  shouldbind")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var fetchUser User
	fetchUser = checkIfUserExist(user)
	if fetchUser != (User{}) {
		fmt.Println("func RegisterUser(c fetchUser ", fetchUser)
		c.JSON(200, gin.H{"data": fetchUser, "msg": "User Logged in successfully ", "status": true})
	} else {

			c.JSON(200, gin.H{"data": nil, "msg": "Error occur", "status": false})

	}
}