/*
 * swagger example API
 *
 * API to demo swagger
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 * Hooray! It's out custom header
 *
 */
package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "github.com/grbit/swagger-example/internal/java"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}