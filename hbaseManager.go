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
//GetRow(tableName string, row string, attributes map[string]string) (data []*TRowResult, err error) 
//GetRows(tableName string, rows []string, attributes map[string]string) (data []*TRowResult, err error)
//GetRowsWithColumns(tableName string, rows []string, columns []string, attributes map[string]string) (data []*TRowResult, err error) 
}


func UpdateHbaseData(table string, row []byte) {
	mutations := make([]*Hbase.Mutation, 1)
	mutations[0] = goh.NewMutation("cf:c", []byte("value3-mutation"))
	fmt.Println(mutations)
//	err := hclient.MutateRow(tableName string, row []byte, mutations []*Hbase.Mutation, timestamp int64, attributes map[string]string)
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
}

