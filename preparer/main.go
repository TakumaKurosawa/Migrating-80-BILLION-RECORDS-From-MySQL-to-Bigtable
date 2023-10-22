package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	shardCount          = 6
	recordCountPerShard = 5000
)

func main() {
	db := setupMySQL()

	fmt.Println("-------- Start create records --------")
	for i := 1; i <= shardCount; i++ {
		fmt.Printf("shard-%d...", i)
		recoords := genRecords(recordCountPerShard)

		if err := db.Table(fmt.Sprintf("hashdb-%d", i)).Create(&recoords).Error; err != nil {
			log.Printf("failed to create records, err: %#v\n", err)

			return
		}
		fmt.Println("done!")
	}
	fmt.Println("-------- Finish create records --------")
}

func setupMySQL() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:local-pass@tcp(127.0.0.1:3306)/local?charset=utf8mb4&parseTime=True&loc=UTC"), nil)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
