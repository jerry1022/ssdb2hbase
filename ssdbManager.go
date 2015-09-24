package main 

import (
	"log"
	"time"
	"strconv"
)

func ReadSSDBData(hname string) map[string]interface{}{
	data, err :=  db.HashRScan(hname, strconv.FormatInt(time.Now().Unix() - (86400 * 40), 10), strconv.FormatInt(time.Now().Unix() - (86400 * 60), 10), 200)
	if err != nil {
		log.Printf("Read SSDB Data Error: %s\n", err)
	}
	return data
}

func DeleteSSDBData() {
//TO-DO
}

