package main

import "github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http"


func main(){
	r := http.SetupRouter()

	r.Run(":8080")
	
}