// There might have been erors here and there since I coded this in subway by smarthone
// What I treid to do was to a) create an excel file and enter some data b) connect the excel file with sqlite3 db
package main

import "fmt"

func main() {
// Open Sqlite3 and create a DB file. 
myDB := sql.open("sqlite3", /myMac/Caffeine/"Coffee.db")

// Create a table that lists all the coffee drinks 
CreateCoffeeTable := '''CREATE Table Drinks( product text not null, caffeine integer not null )'''
)

InsertCoffeeData := '''INSERT INTO Drinks values(?, ?,?)'''

// Create a product list using excel and put into the Sqlite table
CoffeeExcel := excelize.NewFile() // Create a new xlsx file

Var CoffeeProductMap : map[string]int {
    "LetsBe" : 10
    "Georgia" : 20
    "Maxim" : 30 
}

CoffeeExcel.SetCellValue("CoffeeData","A1", CoffeeProductMap)

err := CoffeeExcel.Save(./"CoffeeExcel.xlsx")

if err != nil {
    log.Fatal(err)
}

statement, _ := myDB.Prepare(CreateCoffeeTable)

statememt.Exec()

statement = myDB.Prepare(InsertCoffeeData, ("./CoffeeExcel.xslx"))

statement.Exec()

 // Grab your coffee from the table and return caffeine amount of the coffee 
func ChooseCoffee() (caffeine float64){
      
}
 
