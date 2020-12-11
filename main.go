// main.go

package main

import (
	"./models"
	"./utils"
	"github.com/go-session/gin-session"
	"io/ioutil"
	"math/rand"

	//"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()
	router.Use(ginsession.New())
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/**/*")
	router.Static("htmlSupport", "htmlSupport")

	router.GET("/home", func(context *gin.Context) {
		var productList, _ = models.FindAllProduct(context)
		fmt.Println(productList)
		fmt.Println("Harish")
		context.HTML(http.StatusOK, "index.html", gin.H{
			"List":  productList,
			"title": "Home",
		},
		)
	})
	router.GET("/productDetails", func(context *gin.Context) {
		var id = context.Request.URL.Query()["id"]
		fmt.Println(id)
		var product, _ = models.FindSingleProduct(context, id[0])
		context.HTML(http.StatusOK, "single.html", gin.H{
			"Product": product,
			"title":   "Product",
		},
		)
	})
	router.GET("/contact", func(context *gin.Context) {

		context.HTML(http.StatusOK, "contact.html", gin.H{
			"Hello": "world",
			"title": "Contact",
		},
		)
	})
	router.POST("/contact", func(context *gin.Context) {
		var userName = context.PostForm("name")
		if len(userName) > 0 {
			if models.InsertNewContact(context) {
				context.HTML(http.StatusOK, "contact.html", gin.H{
					"Hello": "world",
					"title": "Contact",
				},
				)
			}
		} else {
			context.HTML(http.StatusOK, "contact.html", gin.H{
				"Hello": "world",
				"title": "Contact",
			},
			)
		}

	})

	router.POST("/productInsert", func(context *gin.Context) {
		fmt.Println(context.HandlerNames())
		var userName = context.PostForm("productName")
		var fileName = uploadFile(context)

		if len(userName) > 0 {
			if models.InsertNewProduct(context, fileName) {
				var productList, _ = models.FindAllProduct(context)
				fmt.Println(productList)
				fmt.Println("Harish")
				context.HTML(http.StatusOK, "DkgTable.html", gin.H{
					"List": productList,
				},
				)
			}
		}
	})
	router.GET("/productList", func(context *gin.Context) {
		if checkIsUserLogined(context) {
			var productList, _ = models.FindAllProduct(context)
			fmt.Println(productList)
			fmt.Println("Harish")
			context.HTML(http.StatusOK, "DkgTable.html", gin.H{
				"List": productList,
			},
			)
		}
	})
	router.GET("/contactList", func(context *gin.Context) {
		if checkIsUserLogined(context) {
			var productList, _ = models.FindAllContact(context)
			fmt.Println(productList)
			fmt.Println("Harish")
			context.HTML(http.StatusOK, "DkgContact.html", gin.H{
				"List": productList,
			},
			)
		}
	})
	router.GET("/form", func(context *gin.Context) {
		if checkIsUserLogined(context) {
			context.HTML(http.StatusOK, "DkgForm.html", gin.H{
				"Hello": "world",
			},
			)
		}
	})

	router.GET("/register", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register.html", gin.H{
			"Hello": "world",
		},
		)
	})

	router.POST("/register", func(context *gin.Context) {
		var userName = context.PostForm("username")
		fmt.Println(userName + " harishofdjnsl")
		if len(userName) > 0 {
			if models.CreateNewUser(context) {
				context.HTML(http.StatusOK, "index1.html", gin.H{
					"Hello": "world",
				},
				)
			}
		}
	})
	router.GET("/login", func(context *gin.Context) {

		context.HTML(http.StatusOK, "login.html", gin.H{
			"Hello": "world",
		},
		)
	})

	router.POST("/validate", func(context *gin.Context) {
		var id = models.FindUser(context)
		if id != 0 {
			store := ginsession.FromContext(context)
			store.Set("userId", id)
			err := store.Save()
			if err != nil {
				context.AbortWithError(500, err)
				return
			}
			var productList, _ = models.FindAllProduct(context)
			fmt.Println(productList)
			fmt.Println("Harish")
			context.HTML(http.StatusOK, "DkgTable.html", gin.H{
				"List": productList,
			},
			)
		} else {
			context.HTML(http.StatusOK, "login.html", gin.H{
				"Haell": "elk",
			})
		}

	})
	// Initialize the routes
	//initializeRoutes()

	// Start serving the application
	router.Run()

}

func checkIsUserLogined(context *gin.Context) bool {
	store := ginsession.FromContext(context)
	userId, ok := store.Get("userId")
	fmt.Println(userId)
	fmt.Println(" dsffd ")
	if !ok {
		context.AbortWithStatus(404)
		return false
	}
	return true
}

func uploadFile(context *gin.Context) string {

	fmt.Println("File Upload Endpoint Hit")
	var r = context.Request
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("productImage")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	var random int = rand.Int()
	var fileName = fmt.Sprintf("%s%d%s", "upload", random, "*.png")
	fmt.Println("fileName : " + fileName)
	tempFile, err := ioutil.TempFile("htmlSupport/web/images", fileName)
	if err != nil {
		fmt.Println(err)
	}
	var fileNameSaved = tempFile.Name()
	fmt.Println(tempFile.Name())
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Println("Successfully Uploaded File\n")
	return fileNameSaved
}

func FindById(id int64) {
	db := utils.DbConn()
	defer db.Close()

	query := "SELECT * FROM test.user"

	res, err := db.Query(query)
	//defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	for res.Next() {
		res.Scan(user)
		row, err := res.Columns()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", row)
		fmt.Printf("%v\n", user)

	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
