package main

import (
	"payment/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8000")

}
