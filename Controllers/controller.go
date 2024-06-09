package controllers

import (
	"empdb/db"
	"empdb/models"
	"empdb/utils"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	// "fmt"
)

var Log =utils.Log()
	
func Heart_Beat(c *gin.Context){
	c.String(200,"Working")
}

func AddEmployee(c *gin.Context){

	Log().Info("IP-"+c.Request.RemoteAddr+";Method-"+c.Request.Method+";Path-"+c.Request.URL.Path+";UserAgent-"+c.Request.UserAgent())

	var emp models.Employee

	err:=c.BindJSON(&emp)
	if err!=nil{
		Log().Error("failed to bind json,"+err.Error())
		c.String(400,"Data is not valid")
		return
	}

	conn := new(db.Connection);
	conn.CreateDB()

	if conn.Db ==nil{
		c.String(400,"failed to connect to db")
		return
	}

	ok:=conn.Insert(&emp)

	if !ok{
		c.String(400,"failed")
		return
	}else{
		res,_:=json.Marshal(emp)
		c.String(200,string(res))
		return
	}
}

func GetEmployeeByID(c *gin.Context){
	Log().Info("IP-"+c.Request.RemoteAddr+";Method-"+c.Request.Method+";Path-"+c.Request.URL.Path+";UserAgent-"+c.Request.UserAgent())

	id := c.Param("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        Log().Error("Failed to convert ID: " + err.Error())
        c.String(400, "Valid ID not present")
        return
    }

	conn := new(db.Connection)
	conn.CreateDB()
    if conn == nil {
        Log().Error("Failed to connect to DB")
        c.String(500, "Failed to connect to DB")
        return
    }

	emp:=conn.GetById(idInt)

	if emp==nil{
		c.String(400,"failed")
		return
	}
	
	res,_:=json.Marshal(emp)
	c.String(200,string(res))
	return


	

}

func UpdateEmployee(c *gin.Context){
	Log().Info("IP-"+c.Request.RemoteAddr+";Method-"+c.Request.Method+";Path-"+c.Request.URL.Path+";UserAgent-"+c.Request.UserAgent())

	id := c.Param("id")

    idInt, err := strconv.Atoi(id)
    if err != nil {
        Log().Error("Failed to convert ID: " + err.Error())
        c.String(400, "Valid ID not present")
        return
    }

    var emp map[string]interface{}
    err = c.ShouldBindJSON(&emp)
    if err != nil {
        Log().Error("Failed to bind JSON: " + err.Error())
        c.String(400, "Data is not valid")
        return
    }

    conn := new(db.Connection)
	conn.CreateDB()
    if conn == nil {
        Log().Error("Failed to connect to DB")
        c.String(500, "Failed to connect to DB")
        return
    }
    // defer conn.Close() // Ensure the connection is closed after the operation

    ok,employee:= conn.Update(idInt, emp)

    if !ok {
        c.String(500, "Failed to update employee")
        return
    }
	
    res,err:=json.Marshal(employee)

	if err!=nil{
		Log().Error("Failed to unmarshal")
		c.String(400,"failed")
		return
	}
	c.String(200,string(res))
	return

}

func DeleteEmployee(c *gin.Context){
	Log().Info("IP-"+c.Request.RemoteAddr+";Method-"+c.Request.Method+";Path-"+c.Request.URL.Path+";UserAgent-"+c.Request.UserAgent())

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
    if err != nil {
        Log().Error("Failed to convert ID: " + err.Error())
        c.String(400, "Valid ID not present")
        return
    }

	conn := new(db.Connection)
	conn.CreateDB()
	if conn == nil{
		c.String(400,"failed to connect to db")
		return
	}

	ok:=conn.Delete(idInt)

	if !ok{
		c.String(400,"failed")
		return
	}else{
		c.String(200,"success")
		return
	}

}

func GetEmployees(c *gin.Context){
	Log().Info("IP-"+c.Request.RemoteAddr+";Method-"+c.Request.Method+";Path-"+c.Request.URL.Path+";UserAgent-"+c.Request.UserAgent())

	page := c.Param("page")
	perpage := c.Param("perpage")

    pageInt, err := strconv.Atoi(page);
    if err != nil {
        Log().Error("Failed to convert page: " + err.Error())
        c.String(400, "Valid page not present")
        return
    }

	if pageInt==0{pageInt=1}
		
    perpageInt, err := strconv.Atoi(perpage)
    if err != nil {
        Log().Error("Failed to convert perpage: " + err.Error())
        c.String(400, "Valid perpage not present")
        return
    }

	conn := new(db.Connection)
	conn.CreateDB()
    if conn == nil {
        Log().Error("Failed to connect to DB")
        c.String(500, "Failed to connect to DB")
        return
    }

	emp:=conn.GetEmployees(pageInt,perpageInt)

	if emp==nil{
		c.String(400,"failed")
		return
	}else{
		res,_:=json.Marshal(emp)
		c.String(200,string(res))
		return
	}

	

}