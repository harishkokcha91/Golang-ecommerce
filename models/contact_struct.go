package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosample/dbUtils"
	log "log"
)

type Contact struct {
	Id         int    `json:"id"`
	Name       string `json:"c_name"`
	Email      string `json:"c_email"`
	Message    string `json:"c_desc"`
	ContactCol string `json:"contactcol"`
}

func FindAllContact(c *gin.Context) ([]Contact, error) {
	var product Contact
	var productList []Contact
	res, err := dbUtils.DbConn().Query("select * from test.contact")

	if err == nil {
		for res.Next() {
			res.Scan(&product.Id, &product.Name, &product.Email, &product.Message, &product.ContactCol)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", row)
			fmt.Printf("%v\n", product)
			productList = append(productList, product)
		}
		return productList, err
		//c.JSON(200, user)
	} else {
		//c.JSON(404, gin.H{"error": "user not found"})
	}
	return nil, nil
}
func FindSingleContact(c *gin.Context, id string) (Contact, error) {
	var product Contact
	var query = "Select id,c_name,c_email,c_desc from test.contact where id=" + id
	res, err := dbUtils.DbConn().Query(query)
	fmt.Println(query)
	if err == nil {
		for res.Next() {
			res.Scan(&product.Id, &product.Name, &product.Email, &product.Message)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("kdsjf %v\n", row)
			fmt.Printf("kdjns %v\n", product)
		}
		return product, err
		//c.JSON(200, user)
	} else {
		//c.JSON(404, gin.H{"error": "user not found"})
	}
	return product, nil
}

func InsertNewContact(c *gin.Context) bool {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")
	var querySql = "Insert into test.contact(name,email,desc) values" +
		" ('" + name + "','" + email + "','" + message + "')"
	fmt.Println(querySql + " querySql")

	var completeMsg = "Hello a new customer " + name + " his email is" + email + " message is :" + message + "wants to connect you"

	var emailsList = []string{email}
	sent, error := SendEmail(completeMsg, emailsList)
	fmt.Println("sent")
	fmt.Println(sent)

	fmt.Println(error)

	var queryPreParedStatement = "Insert into test.contact(c_name,c_email,c_desc) values(?,?,?)"
	_, err := dbUtils.DbConn().Query(queryPreParedStatement, name, email, message)
	if err == nil {
		return true
	} else {
		fmt.Print(err)
		return false
	}
	return false
}
