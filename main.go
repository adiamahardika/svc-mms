package main

import (
	"fmt"
	"svc-monitoring-maintenance/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func main() {
	dsn := "host=localhost user=postgres search_path='monitoring_maintenance' password=maisenpai dbname=dashboardoa port=5433 sslmode=disable TimeZone=Asia/Jakarta"

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	
	if (error != nil) {
		fmt.Println("Connection to db has been error!")
	} else {
		fmt.Println("Connection to db succeed!")
	}
	
	router.AllRouter(db)
}