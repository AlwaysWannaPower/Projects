package main

import (
	"flag"
	"fmt"
)

func main() {

	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Возраст пользователя")

	flag.Parse()

	fmt.Println(*city, *format)

}
