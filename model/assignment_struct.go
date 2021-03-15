package model

import (
	"fmt"
	"gosample/dbUtils"
	"log"

	"github.com/gin-gonic/gin"
)

type Assignemt struct {
	AssignemtID        string `json:"assignment_id"`
	AssignemtName      string `json:"assignment_name"`
	AssignemtChapterID string `json:"assignment_chapter_id"`
	AssignemtStatus    string `json:"assignment_status"`
	AssignemtCreatedAt string `json:"assignment_created_at"`
	AssignemtUpdatedAt string `json:"assignment_updated_at"`
}

func GetAllAssignemt(c *gin.Context) {

	var assignmentList = getListOfAssignment(c)

	if assignmentList != nil {
		c.JSON(200, gin.H{"data": assignmentList, "msg": "data found succes", "status": true})
	} else {
		c.JSON(404, gin.H{"data": nil, "msg": "data not found ", "status": false})
	}
}

func getListOfAssignment(c *gin.Context) []Assignemt {
	var queryString = "Select * from test.assignment_master where assignment_status=1"

	var assignmentList []Assignemt
	var assignment Assignemt

	res, err := dbUtils.DbConn().Query(queryString)

	if err == nil {
		for res.Next() {
			res.Scan(&assignment.AssignemtID, &assignment.AssignemtName, &assignment.AssignemtChapterID, &assignment.AssignemtStatus, &assignment.AssignemtCreatedAt, &assignment.AssignemtUpdatedAt)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", assignment)
			fmt.Printf("%v\n", row)
			assignmentList = append(assignmentList, assignment)
		}
		return assignmentList
	}
	return nil
}
