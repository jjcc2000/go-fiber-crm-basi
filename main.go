package main

import (
	"fmt"
	"mods/database"
	"mods/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)


func setUpRoutes(app *fiber.App)	{
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("api/v1/:id",lead.DeleteLead)
}
func initDatabase(){
	var err error
	database.DBConn , err = gorm.Open("sqlite","leads.db")
	if err != nil {
		panic("Failed to conect database")
	}
	fmt.Println("Coneccion Open to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
func main()  {
	app := fiber.New()	
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000, nil)
	defer database.DBConn.Close()
}
