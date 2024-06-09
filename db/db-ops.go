package db

import (
	"empdb/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"empdb/utils"
)

var Log =utils.Log()

type Connection struct{
	Db *gorm.DB
}


func (conn *Connection) InitTestDB(){
	var err error
	conn.Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  Log().Error("failed to connect database")
	}
	
	conn.Migrate()
}



func (conn *Connection) CreateDB(){
	var err error
	conn.Db, err = gorm.Open(sqlite.Open("employee.db"), &gorm.Config{})
	if err != nil {
	  Log().Error("failed to connect database")
	}
	
	if conn.Db.Migrator().HasTable(&models.Employee{}) {
		Log().Info("table exists")
		return
	}else{
		Log().Info("table not exists")
		conn.Migrate()
		return
	}
	

}

//create connection
func (conn *Connection) Migrate(){
	if conn.Db == nil {
		Log().Error("db is nil")
		return
	}
	// Migrate the schema
	conn.Db.AutoMigrate(&models.Employee{})
}

//insert data
func (conn *Connection) Insert(_emp *models.Employee)bool{
	if conn.Db == nil {
		Log().Error("db is nil")
		return false
	}

	err:=conn.Db.Create(_emp).Error
	if err!=nil{
		Log().Error("Error while inserting data",err.Error())
		return false
	}
	
	return true	
}

//get data bu id
func (conn *Connection) GetById(_id int) *models.Emp{
	if conn.Db == nil {
		Log().Error("db is nil")
	}
	var product models.Employee
	err:=conn.Db.First(&product, _id).Error // find product with integer primary key
	
	if err!=nil{
		Log().Error("Error while getting data",err)
		return nil
	}

	return &models.Emp{Name: product.Name, Position: product.Position, Salary: product.Salary}
}

//update data by id
func (conn *Connection) Update(_id int, _emp map[string]interface{}) (bool,*models.Employee){
	if conn.Db == nil {
		Log().Error("db is nil")
		return false,nil
	}

	var employee models.Employee
	err := conn.Db.First(&employee, _id).Error
	if  err != nil {
		return false,nil
	}

	err=conn.Db.Model(&employee).Updates(_emp).Error

	
	if err!=nil{
		Log().Error("Error while updating data",_emp)
		return false,nil
	}
	
	return true,&employee
}

//delete _id int
func (conn *Connection) Delete (_id int) bool {
	if conn.Db == nil {
		Log().Error("db is nil")
		return false
	}
	txn:=conn.Db.Where("id = ?", _id).Delete(&models.Employee{})

	if txn.RowsAffected==0{
		Log().Error("Error while deleting data")
		return false
	}else{
		return true
	}
}

//_offset is used for pagination
func (conn *Connection) GetEmployees (_page, _perpage int) *[]models.Employee{
	if conn.Db == nil {
		Log().Error("db is nil")
		return nil
	}

	if _page==1{_page=0}; //this will fix the first row missing issue

	var employees []models.Employee
	err:=conn.Db.Limit(_perpage).Offset(_perpage*(_page-1)).Find(&employees).Error

	if err!=nil{
		Log().Error("Error while getting data",err)
		return nil
	}

	return &employees
}


