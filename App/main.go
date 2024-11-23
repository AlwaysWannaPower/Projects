package main

import (
	"App/geo"
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {

	city := flag.String("city", "", "Город пользователя")
	// format := flag.Int("format", 1, "Возраст пользователя")

	flag.Parse()

	// fmt.Println(*city)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)

	r := strings.NewReader("Привет! Я поток данных")
	block := make([]byte, 4)
	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}
	}

}
