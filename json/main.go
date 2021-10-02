package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type product struct {
	Name string
	Price int
	Published bool
}

type x struct {
	A string
	B int
	C bool
}

func main() {

	p := product{
		Name: "MacBook Pro",
		Price: 1500,
		Published: true,
	}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	jsonDataUnMarchal := `{"Name": "MacBook Air", "Price": 900, "Published": true}`
	product := product{}
	if err := json.Unmarshal([]byte(jsonDataUnMarchal), &product); err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)

	z := x{
		A: "Hola Mundo",
		B: 5,
		C: true,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(z)

	fmt.Println(encoder)

	jsonData01 := `{
				"A":"Hola Mundo",
				"B": 5,
				"C": true}`

	y := x{}
	buff := bytes.NewBuffer([]byte(jsonData01))
	decoder := json.NewDecoder(buff)
	decoder.Decode(&y)
	fmt.Println(y)

}
