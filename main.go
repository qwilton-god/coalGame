package main

import (
	"coalGame/company"
	myhttp "coalGame/http"
)

func main() {
	myGame := company.RunGame()
	myhttp.StartServer(myGame)
}
