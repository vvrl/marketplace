package main

import (
	"fmt"
	"marketplace/internal/app"
)

func main() {

	/*
		TODO: тесты сразу на все
			  создание роутинга и хендлера (просто вывод helloworld)
			  echo
			  db
			  миграции
	*/
	market := app.NewApp()

	if err := market.Run(); err != nil {
		fmt.Printf("failed running app: %s", err)
		return
	}
}
