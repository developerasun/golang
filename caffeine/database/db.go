package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Declare constant variables to Sqlite queries
const (
	CREATE_TABlE = "CREATE TABlE caffeine_checker(date text not null, name text not null, amount float64 not null)"

	INSERT_DATA = "INSERT INTO caffeine_checker values(?, ?, ?)"

	SELECT_DATA = "SELECT * FROM caffeine_checker"
)

// Create a map to index coffee drinks and their caffeine amount
var ProductList = map[string]float64{
	"LetsBe":  10,
	"Georgia": 20,
	"Maxim":   30,
}

// Create a embodying class : CoffeeTable
type CoffeeTable struct{}

// Declare a function to create a table
func (c CoffeeTable) CreateTable() {
	// Open a sqlite3 and create a database file.
	myDB, _ := sql.Open("sqlite3", "./myTesting.db")

	// Prepare queries to manipulate tables.
	statement, _ := myDB.Prepare(CREATE_TABlE)
	statement.Exec()
}

// Declare a function to insert data into table
func (c CoffeeTable) InsertData() {
	myDB, _ := sql.Open("sqlite3", "./myTesting.db")
	statement, _ := myDB.Prepare(INSERT_DATA)

	// Jonghyun drinks "Georgia"
	statement.Exec("7월 28일", "Jonghyun", 20)
	// Taewoong drinks "LetsBe"
	statement.Exec("7월 29일", "Taewoong", 10)
	// Taewoong drinks "Maxim"
	statement.Exec("7월 30일", "Jeonggi", 30)
}

// Declare a function to select data from the table
func (c CoffeeTable) SelectData() {
	myDB, _ := sql.Open("sqlite3", "./myTesting.db")
	rows, _ := myDB.Query(SELECT_DATA)
	fmt.Println(rows)
}

func main() {

	// Declare a CoffeeTable structure to use the connected methods. d
	var myCaffeineChecker CoffeeTable

	// Manipulate data with query statement
	myCaffeineChecker.CreateTable()
	myCaffeineChecker.InsertData()
	myCaffeineChecker.SelectData()

}

// ===================== Bonus exercise ================= //
// 1. Create a excel file and enter some data
// 2. Insert the data into sqlite3 database

// exploration ---
// Create a product list using excel and put into the Sqlite table
// CoffeeExcel := excelize.NewFile() // Create a new xlsx file

// CoffeeExcel.SetCellValue("CoffeeData","A1", CoffeeProductMap)

// err := CoffeeExcel.Save(./"CoffeeExcel.xlsx")

// statement = myDB.Prepare(InsertCoffeeData, ("./CoffeeExcel.xslx"))
// if err != nil {
//     log.Fatal(err)
// }
