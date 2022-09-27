package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Category string
	Type     string
	Pmodel   string
	Price    int
	Amount   int
	Stores   []Store `gorm:"many2many:product_storages;"`
}

type Store struct {
	gorm.Model
	Name     string
	Adresses []Adress `gorm:"many2many:storage_addresses;"`
}

type Adress struct {
	gorm.Model
	District string
	Street   string
}

type Allinfo struct {
	Name      string
	Category  string
	Type      string
	Pmodel    string
	Price     int
	Amount    int
	Storename string
	dname     string
	Sstreet   string
}

func main() {

	dsn := "host=localhost user=postgres password=sdy12197 dbname=test_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Adress{}, &Store{}, &Product{}, &Allinfo{})
	db.Create(&Product{

		Name:     "model x7",
		Category: "car",
		Type:     "avto",
		Pmodel:   "X7",
		Price:    12000,
		Amount:   32,
		Stores: []Store{
			{
				Name: "SPcarSHop",
				Adresses: []Adress{
					{
						District: "Qoraqamish",
						Street:   "Amir Temur",
					},
				},
			},
			{
				Name: "FAcars",
				Adresses: []Adress{
					{
						District: "yunisobod",
						Street:   "Alisher navoiy 12 3",
					},
					{
						District: "Mirobod",
						Street:   "big street 3 12",
					},
				},
			},
		},
	})

}