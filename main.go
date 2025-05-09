package main

//
// import (
// 	"fmt"
// 	"log"
//
// 	"github.com/bplong33/gonarqube/services"
// 	"github.com/joho/godotenv"
// )
//
// func main() {
// 	// load env
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Unable to load environment variables")
// 	}
//
// 	//
// 	opts := map[string]string{
// 		"visibility": "private",
// 	}
// 	c := services.NewProjectClient()
// 	projects := c.GetProjects(opts)
// 	for _, p := range projects {
// 		fmt.Println(p.Name)
// 	}
// }
