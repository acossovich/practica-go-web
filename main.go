package main

import (
	//"context"
	// "fmt"
	// "errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"io"
	"encoding/json"
	"strconv"
)

type Producto struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"code_value"`
	IsPublished bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}

var Productos []Producto

func main(){

	router := gin.Default() //se puede tambien con gin.New() --> es mejor para tener mas control de los middlewares

	archivo,err := os.Open("products.json")

	if err != nil{
		panic(err)
	}
	
	bytesFromArch,err := io.ReadAll(archivo)

	json.Unmarshal(bytesFromArch,&Productos)

	router.GET("/ping",Pong)
	router.GET("/products",func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,Productos)
	})
	router.GET("/products/:id",SearchProductById)

	if err := router.Run(":8080"); err != nil{
		panic(err)
	}

}

func Pong(c *gin.Context){

	c.String(http.StatusOK,"Pong")
	
}

func SearchProductById(c *gin.Context){

	id := c.Param("id")

	//c.String(http.StatusAccepted,"%s",id)

	for _,valor:= range Productos{
		if idValue,err := strconv.Atoi(id);  err != nil {
			c.String(http.StatusBadRequest,"El id es invalido")
			return
		}else{
			if valor.Id == idValue {
				c.JSON(http.StatusOK,valor)
				return
			}
		}
	}

	c.String(http.StatusOK,"Id no encontrado")
	
}