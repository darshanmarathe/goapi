package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_CONNECTIONSTRING"))

	port := ":" + os.Getenv("APP_PORT")

	a.Run(port)
}
