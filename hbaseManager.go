package main

import (
	"fmt"
	"github.com/sdming/goh"
	"github.com/sdming/goh/Hbase"
)

func CreateHbaseTable(table string, cf []string ) bool{
	fmt.Println(table)
	cols := make([]*goh.ColumnDescriptor, len(cf))
	for k, v := range cf {
		cols[k] = goh.NewColumnDescriptorDefault(v)
	}
	exists, err := hclient.CreateTable(table, cols)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if exists {
		fmt.Printf("%s Exists", table)
	} else {
		fmt.Printf("Create %s successful", table)
	}
	return true
}

func ReadHbaseData(table string, row string) {
//TO-DO
}


func UpdateHbaseData(table string, row []byte) {
//TO-DO
}

func UpdateHbaseBatchData(table string, rows []*Hbase.BatchMutation, attr map[string]string) bool{
	err := hclient.MutateRows(table, rows, attr)
	if err != nil {
		return false
	} else {
		return true
	}
}

func DeleteHbaseData() {
//TO-DO
}

