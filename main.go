package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("p", "8686", "port number")
	dir := flag.String("d", "./", "directory")

	flag.Parse()

	router := gin.Default()
	router.StaticFS("/", http.Dir(*dir))

	router.Run(fmt.Sprintf(":%s", *port))
}
