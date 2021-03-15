package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosample/dbUtils"
)

type Subject struct {
	SubjectID      string `json:"subject_master_id"`
	SubjectName    string `json:"subject_name"`
	SubjectClassID string `json:"subject_class_id"`
	SubjectStatus  string `json:"subject_status"`

	SubjectCreatedAt string `json:"subject_created_at"`
	SubjectUpdatedAt string `json:"subject_updated_at"`
}

func GetSubjectList(c *gin.Context) {
	var subjectList = getSubjectListLocal(c)

	if subjectList != nil {
		c.JSON(200, gin.H{"data": subjectList, "msg": "data found succes", "status": true})
	} else {
		c.JSON(404, gin.H{"data": nil, "msg": "data not found ", "status": false})
	}
}

func getSubjectListLocal(c *gin.Context) []Subject {
	var request Request

	if err := c.ShouldBind(&request); err != nil {
		return nil

	}

	var queryString string

	if request.ClassID != "" {
		queryString = "Select * from test.subject_master where subject_class_id=" + request.ClassID
	} else {
		queryString = "Select * from test.subject_master where subject_status=1"
	}
	fmt.Println(queryString)
	var subject Subject
	var subjectList []Subject
	res, err := dbUtils.DbConn().Query(queryString)
	if err == nil {
		for res.Next() {
			res.Scan(&subject.SubjectID, &subject.SubjectName,
				&subject.SubjectClassID, &subject.SubjectStatus,
				&subject.SubjectCreatedAt, &subject.SubjectUpdatedAt)

			subjectList = append(subjectList, subject)
		}
		return subjectList
	}
	return nil
}
