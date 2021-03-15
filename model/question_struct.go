package model

import (
	"fmt"
	"gosample/dbUtils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Question struct {
	QuestionID           string   `json:"question_id"`
	QuestionName         string   `json:"question_name"`
	QuestionAssignmentID string   `json:"question_assignment_id"`
	QuestionMarks        string   `json:"question_marks"`
	QuestionStatus       string   `json:"question_status"`
	QuestionType         string   `json:"question_type"`
	QuestionImage        string   `json:"question_image"`
	QuestionCreatedAt    string   `json:"created_at"`
	QuestionUpdatedAt    string   `json:"updated_at"`
	QuestionOption       []Entity `json:"option"`
}

type Entity struct {
	EntityID         string `json:"entity_id"`
	OptionName       string `json:"option_name"`
	OptionMarks      string `json:"option_marks"`
	OptionType       string `json:"option_type"`
	OptionQuestionID string `json:"option_question_id"`
	OptionImage      string `json:"option_image"`
}

func GetQuestionForAssignment(c *gin.Context) {

	var questionList = getListOfQuestion(c)

	if questionList != nil {
		c.JSON(200, gin.H{"data": questionList, "msg": "data found succes", "status": true})
	} else {
		c.JSON(404, gin.H{"data": nil, "msg": "data not found ", "status": false})
	}
}

func getListOfQuestion(c *gin.Context) []Question {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	fmt.Printf("body is: %s \n", request.AssignmentID)

	var queryString = "Select * from test.question_master where question_assignment_id=" + request.AssignmentID

	fmt.Println("getListOfQuestion : " + queryString)
	var questionList []Question
	var question Question

	res, err := dbUtils.DbConn().Query(queryString)

	if err == nil {
		for res.Next() {
			res.Scan(&question.QuestionID, &question.QuestionName,
				&question.QuestionAssignmentID, &question.QuestionMarks,
				&question.QuestionStatus, &question.QuestionType, &question.QuestionImage,
				&question.QuestionCreatedAt, &question.QuestionUpdatedAt)

			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			question.QuestionOption = fetchOptionForQuestion(question.QuestionID)
			fmt.Printf("%v\n", question)
			fmt.Printf("%v\n", row)
			questionList = append(questionList, question)
		}
		return questionList
	}
	return nil
}

func fetchOptionForQuestion(questionId string) []Entity {

	var queryString = "Select * from test.option_entity where option_question_id=" + questionId

	res, err := dbUtils.DbConn().Query(queryString)

	var optionList []Entity
	var option Entity

	if err == nil {
		for res.Next() {

			res.Scan(&option.EntityID, &option.OptionName,
				&option.OptionMarks, &option.OptionType,
				&option.OptionQuestionID, &option.OptionImage, "", "")

			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", option)
			fmt.Printf("%v\n", row)
			optionList = append(optionList, option)
		}
		return optionList
	}
	return nil
}
