package models

import (
	"gosample/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type Product struct {
	Id                     int    `json:"product_id"`
	ProductName            string `json:"product_name"`
	ProductDesc            string `json:"product_desc"`
	ProductImage           string `json:"product_image"`
	ProductPrice           string `json:"product_price"`
	ProductDiscountedPrice string `json:"product_discounted_price"`
	ProductLink            string `json:"product_buy_link"`
}

func FindAllProduct(c *gin.Context) ([]Product, error) {
	var product Product
	var productList []Product
	res, err := utils.DbConn().Query("select * from test.product")

	if err == nil {
		for res.Next() {
			res.Scan(&product.Id, &product.ProductName, &product.ProductDesc, &product.ProductImage, &product.ProductPrice, &product.ProductDiscountedPrice, &product.ProductLink)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v\n", row)
			var str1 string = "../../"
			var str2 string = strings.Replace(product.ProductImage, "\\", "/", -2)
			var result string = fmt.Sprintf("%s%s", str1, str2)
			product.ProductImage = result
			fmt.Printf("%v\n", product.ProductImage)
			productList = append(productList, product)
		}
		return productList, err
		//c.JSON(200, user)
	} else {
		//c.JSON(404, gin.H{"error": "user not found"})
	}
	return nil, nil
}
func FindSingleProduct(c *gin.Context, id string) (Product, error) {
	var product Product
	var query = "Select id,product_name,product_desc,product_image,product_price,product_discounted_price,product_buy_link from test.product where id=" + id
	res, err := utils.DbConn().Query(query)
	fmt.Println(query)
	if err == nil {
		for res.Next() {
			res.Scan(&product.Id, &product.ProductName, &product.ProductDesc, &product.ProductImage, &product.ProductPrice, &product.ProductDiscountedPrice, &product.ProductLink)
			row, err := res.Columns()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", row)
			var str1 string = "../../"
			var str2 string = strings.Replace(product.ProductImage, "\\", "/", -2)
			var result string = fmt.Sprintf("%s%s", str1, str2)
			product.ProductImage = result
			fmt.Printf("%v\n", product.ProductImage)
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

func InsertNewProduct(c *gin.Context, fileName string) bool {
	name := c.PostForm("productName")
	price := c.PostForm("productPrice")
	desc := c.PostForm("productDesc")
	link := c.PostForm("productLink")
	productDiscountPrice := c.PostForm("productDiscountPrice")
	var querySql = "Insert into test.product(product_name,product_desc,product_price,product_discounted_price,product_image) values" +
		" ('" + name + "','" + desc + "','" + price + "','" + productDiscountPrice + "','" + fileName + "')"
	fmt.Println(querySql + " querySql")

	var queryPreParedStatement = "Insert into test.product(product_name,product_desc,product_price,product_discounted_price,product_image,product_buy_link) values(?,?,?,?,?,?)"
	_, err := utils.DbConn().Query(queryPreParedStatement, name, desc, price, productDiscountPrice, fileName, link)
	if err == nil {
		return true
	} else {
		fmt.Print(err)
		return false
	}
	return false
}
