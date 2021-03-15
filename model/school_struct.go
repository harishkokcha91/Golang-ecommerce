package model

import (
	"fmt"
	"gosample/dbUtils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type School struct {
	SchoolID         string `json:"school_id"`
	SchoolName       string `json:"school_name"`
	SchoolUniqueCode string `json:"school_unique_code"`
	SchoolAddress    string `json:"school_address"`
	SchoolCity       string `json:"school_city"`
	SchoolState      string `json:"school_state"`

	SchoolStatus    string `json:"school_status"`
	SchoolUpdatedAt string `json:"school_updated_at"`

	SchoolCreatedAt string `json:"school_created_at"`
}

func GetSchool(c *gin.Context) {
	var classesList = getSchool(c)

	if classesList != nil {
		c.JSON(200, gin.H{"data": classesList, "msg": "data found succes", "status": true})
	} else {
		c.JSON(404, gin.H{"data": nil, "msg": "data not found ", "status": false})
	}
}

func getSchool(c *gin.Context) []School {

	var request Request

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	fmt.Println(request.SchoolID)
	var queryString string
	if request.SchoolID != "" {
		queryString = "Select * from test.school_master where school_id=" + request.SchoolID
	} else {
		queryString = "Select * from test.school_master"
	}
	fmt.Println(queryString)

	res, err := dbUtils.DbConn().Query(queryString)
	var school School
	var schoolList []School
	fmt.Printf("%v\n", res)
	fmt.Printf("err %v\n", err)
	if err == nil {
		for res.Next() {
			res.Scan(&school.SchoolID, &school.SchoolName, &school.SchoolUniqueCode,
				&school.SchoolAddress, &school.SchoolCity, &school.SchoolState,
				&school.SchoolStatus, &school.SchoolCreatedAt, &school.SchoolUpdatedAt)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", school)
			fmt.Printf("%v\n", row)
			schoolList = append(schoolList, school)
		}
		return schoolList
	}
	return nil
}
