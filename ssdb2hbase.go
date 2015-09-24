package main

import (
	"github.com/matishsiao/gossdb/ssdb"
	"github.com/sdming/goh"
	"github.com/sdming/goh/Hbase"
	_"github.com/robfig/cron"

	"log"
	"strconv"
	"time"
	"fmt"
	"flag"
	"os"
	"encoding/json"
)

var (
	host string
	db *ssdb.Client
	hclient *goh.HClient
	err error
	version string = "0.0.1"
	configPath string 
        configs Configs
)

func main() {
	var f *os.File
	if !FileExists("/var/log/ssdb2hbase/ssdb2hbase.log") {
		f, _ = os.Create("/var/log/ssdb2hbase/ssdb2hbase.log")
	} else {
		f, _ = os.OpenFile("/var/log/ssdb2hbase/ssdb2hbase.log", os.O_APPEND|os.O_WRONLY, 0755)
	}
	log.SetOutput(f)
	defer f.Close()

	host, _ = os.Hostname()
	flag.StringVar(&configPath,"c", "/etc/ssdb2hbase/config.json", "config filename")
	flag.Parse()
	ok, config := loadConfigs(configPath)
	if !ok {
		log.Println("load config failed")
		os.Exit(0)
	}

	configs = config
	log.Printf("version:%s ssdb:%s:%d hbase thrift:%s:%d",version, 
				configs.SSDB.Address, configs.SSDB.Port,
				configs.HBase.Address, configs.HBase.Port)

	//ssdb connection setting
	db, err = ssdb.Connect(configs.SSDB.Address, configs.SSDB.Port, configs.SSDB.Auth)
	if err != nil {
		log.Printf("ssdb connect failed: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	//hbase client setting
	Address := fmt.Sprintf("%s:%d", configs.HBase.Address, configs.HBase.Port)
	hclient, err = goh.NewTcpClient(Address, goh.TBinaryProtocol, false) 
	if err != nil {
		log.Printf("hbase connect failed: %s\n", err)
		os.Exit(1)
	}
	if err := hclient.Open(); err != nil {
		log.Printf("hbase open connection failed: %s\n", err)
		os.Exit(1)
	}
	defer hclient.Close()

	// set hash for sync all conf purpose 
	db.HashSet("ssdb2hbase-status", host, strconv.FormatInt(time.Now().Unix(), 10))        

	var cf = []string{"proxyagent"}
	
	var proxyCounter ProxyCounter
	var mutations []*Hbase.Mutation
	var rowBatches []*Hbase.BatchMutation

	result := CreateHbaseTable(configs.HBase.Tbl, cf)
	if !result {
		log.Printf("Create Table Error")
	}

	data := ReadSSDBData(configs.SSDB.HT)
	for k, v := range data {
		jsonerr := json.Unmarshal([]byte(v.(string)), &proxyCounter)
		if jsonerr != nil {
			fmt.Println(jsonerr)
		}
		cfq := GetStructFieldsKV(proxyCounter)
		for cfqk, cfqv := range cfq{
			cfqkey := fmt.Sprintf("%s:%s", cf[0], cfqk)
			switch cfqv.(type) {
				case string:
					mutations = append(mutations, goh.NewMutation(cfqkey, []byte(cfqv.(string))))
				case int:
					mutations = append(mutations, goh.NewMutation(cfqkey, []byte(string(strconv.FormatInt(int64(cfqv.(int)), 10)))))
				case float64: 
					mutations = append(mutations, goh.NewMutation(cfqkey, []byte(strconv.FormatFloat(cfqv.(float64),'E',1,64))))
			}
		}
		rowBatches = append(rowBatches, goh.NewBatchMutation([]byte(k), mutations))
	}
	updateResult := UpdateHbaseBatchData(configs.HBase.Tbl, rowBatches, nil)
	if updateResult {
		fmt.Println("Update Successful")
	}


}


