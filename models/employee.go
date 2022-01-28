package models
import (
  
   "gorm.io/gorm"

)

type Employee struct{
    Emp_ID uint `json:"empid" gorm:"primary_key";"auto_increment"`
    Name string `json:"name" binding:"required"`
    Address string `json:"address" binding:"required"`
    Mobile int64  `json:"mobile_number" binding:"required,len=10"`
    Department string  `json:"department" binding:"required"`
    Salary float64  `json:"salary" binding:"required"`
}

type UpdateEmp struct{
   Address string `json:"address"`
   Mobile int64 `json:"mobile_number"`
}

//create a Employee
func CreateEmployee(db *gorm.DB, emp *Employee) (err error) {
    err = db.Create(emp).Error
    if err != nil {
       return err
    }
    return nil
 }
 
 func GetAllEmployee(db *gorm.DB, emp *[]Employee)(err error){
   err = db.Find(emp).Error
   if err != nil {
      return err
   }
   return nil
 }
 //get Employee by id
 func GetEmployee(db *gorm.DB, emp *Employee, id int) (err error) {
    err = db.Where("emp_id = ?", id).First(emp).Error
    if err != nil {
       return err
    }
    return nil
 }
 
 //update Employee
 func UpdateEmployee(db *gorm.DB, emp *UpdateEmp,e *Employee,id int) (err error) {
   employee:=Employee{} 
   switch{
   case emp.Address =="":
     emp.Address=e.Address
   case emp.Mobile==0:
     emp.Mobile=e.Mobile
   }
   db.Model(&employee).Where("emp_id=?",id).Updates(Employee{Address:emp.Address,Mobile:emp.Mobile})
   return nil
}
 
 //delete Employee
 func DeleteEmployee(db *gorm.DB, emp *Employee, id int) (err error) {
    db.Where("emp_id = ?", id).Delete(emp)
    return nil
 }