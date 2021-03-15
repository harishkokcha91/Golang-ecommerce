package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosample/dbUtils"
)

type Chapter struct {
	ChapterID        string `json:"chapter_id"`
	ChapterName      string `json:"chapter_name"`
	ChapterSubjectID string `json:"chapter_subject_id"`

	ChapterCreatedAt string `json:"chapter_created_at"`
	ChapterUpdatedAt string `json:"chapter_updated_at"`
}

func GetChapterList(c *gin.Context) {
	var chapterList = getChapterListLocal(c)

	if chapterList != nil {
		c.JSON(200, gin.H{"data": chapterList, "msg": "data found succes", "status": true})
	} else {
		c.JSON(404, gin.H{"data": nil, "msg": "data not found ", "status": false})
	}
}

func getChapterListLocal(c *gin.Context) []Chapter {
	var request Request

	if err := c.ShouldBind(&request); err != nil {
		return nil

	}

	var queryString string

	if request.SubjectID != "" {
		queryString = "Select * from test.chapter_master where chapter_subject_id=" + request.SubjectID
	} else {
		queryString = "Select * from test.chapter_master"
	}
	fmt.Println(queryString)
	var chapter Chapter
	var chapterList []Chapter
	res, err := dbUtils.DbConn().Query(queryString)
	if err == nil {
		for res.Next() {
			res.Scan(&chapter.ChapterID, &chapter.ChapterName,
				&chapter.ChapterSubjectID, &chapter.ChapterCreatedAt,
				&chapter.ChapterUpdatedAt)

			chapterList = append(chapterList, chapter)
		}
		return chapterList
	}
	return nil
}
