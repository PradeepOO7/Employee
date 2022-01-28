package controllers

import (
   "errors"
   "github.com/gin-gonic/gin"
   "Employee/config"
   "Employee/models"
   "gorm.io/gorm"
   "net/http"
   "strconv"
)

type EmployeeRepo struct {
   Db *gorm.DB
}

func New() *EmployeeRepo {
   db := database.InitDb()
   db.AutoMigrate(&models.Employee{})
   return &EmployeeRepo{Db: db}
}

//create employee
func (repository *EmployeeRepo) CreateEmployee(c *gin.Context) {
   var emp models.Employee
   err:=c.ShouldBindJSON(&emp)
   if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
   err = models.CreateEmployee(repository.Db, &emp)
   if err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   c.JSON(http.StatusOK, emp)
}

//get emp by id
func (repository *EmployeeRepo) GetAllEmployee(c *gin.Context) {
   var emp []models.Employee
   err := models.GetAllEmployee(repository.Db, &emp)
   if err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   c.JSON(http.StatusOK, emp)
}


//get emp by id
func (repository *EmployeeRepo) GetEmployee(c *gin.Context) {
	i,_:=c.Params.Get("id")
   id, _ := strconv.Atoi(i)
   var emp models.Employee

   err := models.GetEmployee(repository.Db, &emp, id)
   if err != nil {
      if errors.Is(err, gorm.ErrRecordNotFound) {
         c.AbortWithStatusJSON(http.StatusNotFound,gin.H{"error": "Record Not Found"})
         return
      }
      
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   c.JSON(http.StatusOK, emp)
}

// update emp
func (repository *EmployeeRepo) UpdateEmployee(c *gin.Context) {
   var emp models.Employee
   i,_:=c.Params.Get("id")
   id, _ := strconv.Atoi(i)
   err := models.GetEmployee(repository.Db, &emp, id)
   if err != nil {
      if errors.Is(err, gorm.ErrRecordNotFound) {
         c.AbortWithStatusJSON(http.StatusNotFound,gin.H{"error": "Record Not Found"})
         return
      }
      
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   var employee models.UpdateEmp
   c.ShouldBindJSON(&employee)
   err = models.UpdateEmployee(repository.Db, &employee,id)
   if err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   emp.Address=employee.Address
   emp.Mobile=employee.Mobile
   c.JSON(http.StatusOK, emp)
}

// delete emp
func (repository *EmployeeRepo) DeleteEmployee(c *gin.Context) {
   var emp models.Employee
   i,_:=c.Params.Get("id")
   id, _ := strconv.Atoi(i)
   err := models.DeleteEmployee(repository.Db, &emp, id)
   if err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}