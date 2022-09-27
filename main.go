package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model
// 	Id       int
// 	Name     string
// 	Age      int
// 	Birthday time.Time
// }
// type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
// }

// type User struct {
// 	gorm.Model
// 	Name       string
// 	CreditCard CreditCard
// }
// type Product struct{
// 	gorm.Model
// 	name string
// 	categoryid Category
// 	typeid Type
// 	Storageid []Storage `gorm:"many2many:product_storages;"`
// }
// type Type struct{
// 	gorm.Model
// 	name string
// }
// type Category struct{
// 	gorm.Model
// 	name string
// }
// type Storage struct{
// 	gorm.Model
// 	name string
// 	Addresses []Address `gorm:"many2many:storage_addresses;"`
// }
// type Address struct{
// 	gorm.Model
// 	street string
// }
type Product struct {
	ID     int
	Name   string
	Model  string
	Stores []Store `gorm:"many2many:product_storages;"`
}

type Store struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:storage_addresses;"`
}

type Address struct {
	ID       int
	District string
	Street   string
}

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=productdb port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Address{}, &Store{}, &Product{})
	tx:=db.Begin()
	product := Product{
		Name:  "Sabzi",
		Model: "Kizil",
		Stores: []Store{
			{
				Name: "Farhadskiy",
				Addresses: []Address{
					{
						District: "Yakasaroy",
						Street:   "Shota Rustaveli",
					},
				},
			},
		},
	}
	
	if err:= tx.Create(&product).Error;err!=nil{
		tx.Rollback()
		fmt.Println(err)
	}
	tx.Commit()
	// db.AutoMigrate(&Category{},&Type{},&Address{},&Storage{},&Product{})
	// category:=Category{name: "Food"}
	// producttype:=Type{name: "Fast-food"}
	// address:=Address{street: "Shota rustaveli"}
	// cat:=db.Create(&category)
	// type:=db.Create(&producttype)
	// address:=db.Create(&address)
	// storage:=Storage{name: "Malika",Addresses:[address]}
	// stor:=db.Create(&storage)

	// user := User{Name: "Nodir", Age: 21, Birthday: time.Now()}
	// err = db.AutoMigrate(&User{})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// result := db.Create(&user)
	// fmt.Println(user.Id)
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)
	// db.AutoMigrate(&User{}, &CreditCard{})
	// res := db.Create(&User{
	// 	Name:       "jinzhu",
	// 	CreditCard: CreditCard{Number: "411111111111"},
	// })
	// fmt.Println(res.RowsAffected)

}
