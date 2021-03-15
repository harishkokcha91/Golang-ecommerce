package model

import (
	"fmt"
	"gosample/dbUtils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Teacher struct {
	ID        string `json:"teacher_id"`
	Name      string `json:"teacher_name"`
	Email     string `json:"teacher_email"`
	Password  string `json:"teacher_pass"`
	Phone     string `json:"teacher_phone"`
	SubjectID string `json:"teacher_subject_id"`
	Status    string `json:"teacher_status"`
	SchoolID  string `json:"teacher_school_id"`
}

func LoginTeacher(c *gin.Context) {
	var teacher Teacher
	var id = TeacherLogin(c, teacher)
	fmt.Println("dkjs ", id)
	if len(id) > 0 {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"List": "dslj",
		},
		)
	} else {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"List": "ouyoery",
		},
		)
	}
}

func TeacherLogin(c *gin.Context, teacher Teacher) string {
	name := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Print(name + " " + password)

	var querySql = "select * from test.teacher_master where teacher_name='" + name + "' and teacher_pass='" + password + "'"
	res, err := dbUtils.DbConn().Query(querySql)
	fmt.Println(res, " res")
	if err == nil {
		for res.Next() {
			res.Scan(&teacher.ID, &teacher.Name)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", row)
			fmt.Printf("%v\n", teacher)
			return teacher.ID
		}
		//c.JSON(200, user)
	} else {
		fmt.Println("No data found")
		return ""
		//c.JSON(404, gin.H{"error": "user not found"})
	}
	return ""
}
