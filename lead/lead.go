package lead 

import (

	"mods/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

)

type Lead struct{
	gorm.Model
	Name string `json:"Name"`
	Company string `json:"Company"`
	Email string `json:"Email"`
	Phone int `json:"Phone"`
}


func GetLeads(c *fiber.Ctx){
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)

}
func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db:= database.DBConn
	var lead Lead
	db.Find(&lead,id)
	c.JSON(lead) 
}
func NewLead(c *fiber.Ctx){
	db:= database.DBConn
	lead := new(Lead)
	if err:= c.BodyParser(lead);err!=nil{
		c.Status(401).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
	c.Status(200)
}
func DeleteLead(c *fiber.Ctx){
	id:= c.Params("id")
	db:= database.DBConn
	var lead Lead
	db.First(&lead,id)
	if lead.Name ==""{
		c.Status(400).Send("The id does not exist")
		return
	}
	db.Delete(&lead)
	c.Send("Lead succefully Deleted")

}