package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Product struct {
	Name
	Manufacturers []Manufacturer
}

type Name struct {
	Name        string
	Description string
	Active      bool
}

type Manufacturer struct {
	Name     string
	Location string
	RegNo    int64
}

func main() {
	product := Product{
		Name: Name{
			Name:        "macbook",
			Description: "professional laptop",
			Active:      true,
		},
		Manufacturers: []Manufacturer{
			{
				Name:     "apple",
				Location: "USA",
				RegNo:    56,
			},
			{
				Name:     "microsoft",
				Location: "SA",
				RegNo:    100,
			},
		},
	}

	err := saveGob("product.gob", product)
	if err != nil {
		log.Fatalln(err)
	}
}

func saveGob(fileName string, key interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(key)
}
