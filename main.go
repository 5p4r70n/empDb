package main

import (
	"empdb/Controllers"

	"github.com/gin-gonic/gin"
	"os"
)

func main(){

	os.Mkdir("log",0777);


	router :=gin.Default()


	
	// for the heart beat, automate health check of containers
	router.HEAD("/", controllers.Heart_Beat)

	router.POST("/addEmployee",controllers.AddEmployee)

	router.PATCH("/updateEmployee/:id",controllers.UpdateEmployee)

	router.DELETE("/deleteEmployee/:id",controllers.DeleteEmployee)

	router.GET("/getEmployee/:id",controllers.GetEmployeeByID)

	router.GET("/getEmployees/:page/:perpage",controllers.GetEmployees)



	router.Run(":2255")
}	