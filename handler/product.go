package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"httpserver/promc"
	"httpserver/types"
	"net/http"
	"time"
)

var products []types.Product

func CreateProduct(c *gin.Context) {
	var product types.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product data",
		})
		return
	}
	go processProduct(product)

	// 存储产品信息
	products = append(products, product)

	c.JSON(http.StatusCreated, product)
}

func processProduct(product types.Product) {
	// 在这里进行故障诊断的逻辑处理
	// 可以根据产品信息进行质量分析、故障检测等操作
	// 这里只是一个示例，可以根据实际需求来实现故障诊断逻辑

	// 模拟故障诊断的结果
	if product.Quantity < 10 {
		fmt.Println("Product quantity is low, potential issue!")
		// 在实际系统中，可以触发报警、记录日志等操作
	}

	// 模拟耗时的故障诊断处理
	time.Sleep(5 * time.Second)

	// 输出故障诊断结果
	fmt.Printf("Product %s processed.\n", product.Name)
}

func GetAllProducts(c *gin.Context) {
	if products == nil {
		products = make([]types.Product, 0)
	}
	c.JSON(http.StatusOK, products)
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Fault Diagnosis System!",
	})
}
func PurgeReqTotal(c *gin.Context) {
	promc.QPSCounter.Reset()
	c.JSON(200, "purged")
}
