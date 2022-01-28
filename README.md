# Employee
A simple CRUD API in Golang with Mysql Database

# Prerequiste
Golang v1.16 or greater
Mysql

# Steps to Test on your local Machine
## Step 1:- To clone this repo <br/>
```git clone https://github.com/bossOO7/Employee.git``` or you can download the zip and Extract it.<br/>
## Step 2:- Move to Employee Folder <br/>
'''cd Employee'''<br/>
## Step 3:- Downloads the dependencies <br/>
```go mod Employee ```<br/>
## Step 4:-Start Mysql Server <br/>
Start Your mysql server on port 3306 which is by default and create a database named as employee_details
## Step 5:- Runs the Api on localhost <br/>
```go run . ``` or '''go run main.go'''<br/>

# Testing 
Now you can test the api in postman 

## Api routes 
## 1.To Add new Empolyees:- Method(POST)
http://localhost/api/v1/addemployee
<br/>

## 2.To get All the Employees:- Method(GET)
http://localhost/api/v1/getall
<br/>

## 3.To get a single Employee by id:- Method(GET)
http://localhost/api/v1/getbyid/:id
<br/>

## 4.To Update Existing Employee by id Only Number and Address are Updatable:- Method(PATCH)
http://localhost/api/v1/update/:id
<br/>

## 5.To Delete a Employee by id:- Method(DELETE)
http://localhost/api/v1/delete/:id
