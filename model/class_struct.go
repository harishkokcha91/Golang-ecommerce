package model

import (
	"fmt"
	"gosample/dbUtils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Class struct {
	ClassID        string `json:"class_id"`
	ClassName      string `json:"class_name"`
	ClassSchoolID  string `json:"class_school_id"`
	ClassCreatedAt string `json:"class_created_at"`
	ClassUpdatedAt string `json:"class_updated_at"`
}

func GetClassesForSchool(c *gin.Context) {
	var classesList = getClassList(c)

	if classesList != nil {
		c.JSON(200, gin.H{"data": classesList, "msg": "data found succes", "status": true})
	} else {
		c.JSON(404, gin.H{"data": nil, "msg": "data not found ", "status": false})
	}
}

func getClassList(c *gin.Context) []Class {

	var request Request

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	fmt.Println(request.SchoolID)
	var queryString string
	if request.SchoolID != "" {
		queryString = "Select * from test.class_master where class_school_id=" + request.SchoolID
	} else {
		queryString = "Select * from test.class_master"
	}
	fmt.Println(queryString)

	res, err := dbUtils.DbConn().Query(queryString)
	var class Class
	var classList []Class
	fmt.Printf("%v\n", res)
	fmt.Printf("err %v\n", err)
	if err == nil {
		for res.Next() {
			res.Scan(&class.ClassID, &class.ClassName, &class.ClassSchoolID, &class.ClassCreatedAt, &class.ClassUpdatedAt)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", class)
			fmt.Printf("%v\n", row)
			classList = append(classList, class)
		}
		return classList
	}
	return nil
}
