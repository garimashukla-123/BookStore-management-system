pacakge models

import(
	"github.com/jinzhu/gorm"
	"github.com/Garima/go-workspace/go-bookstore/pkg/config"

)

var db *gorm.DBtype Book struct{
	gorm.model
	Name string 'gorm:""json:"name"'
	Author string 'json:"author"'
	Publication string 'json:"publication"'
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}