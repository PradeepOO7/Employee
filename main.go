package main

import (
	"github.com/gin-gonic/gin"
	"Employee/controllers"
)
func main(){
	server:=gin.New()
	server.Use(gin.Logger(),gin.Recovery())
	employee := controllers.New()
	v1:=server.Group("/api/v1")
	{
		v1.GET("/getall",employee.GetAllEmployee)
		v1.POST("/addemployee",employee.CreateEmployee)
		v1.GET("/getbyid/:id",employee.GetEmployee)
		v1.PATCH("/update/:id",employee.UpdateEmployee)
		v1.DELETE("/delete/:id",employee.DeleteEmployee)
	}
	
	server.Run(":9000")
}