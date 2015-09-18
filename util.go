package main

import (
	"syscall"
	"os"
	"log"
	"time"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"reflect"
)

func SetUlimit(number uint64) {
	var rLimit syscall.Rlimit
    err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        log.Println("[Error]: Getting Rlimit ", err)
    }    
    rLimit.Max = number
    rLimit.Cur = number
    err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        log.Println("[Error]: Setting Rlimit ", err)
    }
    err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        log.Println("[Error]: Getting Rlimit ", err)
    }
}

func WriteToLogFile(remote string, msg string) {
	logMsg := "[" + remote + "]:" + msg
	log.Println(logMsg)
	t := time.Now()
	fileName := t.Format("20060102") + ".log"
	fmt.Println(logMsg, fileName)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()
	log.SetOutput(f)
	log.Println(logMsg)
}

func SaveFile(dir string,fileName string,data []byte) bool {
	os.Mkdir(dir,0777)
	fo, err := os.Create(dir +"/"+ fileName)
	if err != nil {
		log.Printf("File create error:%v\n",err)			
		return false
	}
	// close fo on exit and check for its returned error
		
	if _, err := fo.Write(data); err != nil {
		log.Printf("File Write Error:%v\n",err)
		return false
	}
	
	if err := fo.Close(); err != nil {
	   log.Printf("File close error:%v\n",err)
	   return false
	}
	
	return true
}

func DeleteFile(dir,fileNmae string) bool {
    fname := dir + "/" + fileNmae
    if err := os.Remove(fname); err != nil {
        log.Printf("File delete fail: %v\n",err)
        return false
    }
    return true
}

func LoadFile(fileName string) []byte {
	contents,err := ioutil.ReadFile(fileName)	
	if err != nil {
		fmt.Printf("Import Error:%s\n", err)
		return nil
	}
	return contents
}

func loadConfigs(configName string) (bool,Configs) {
	fmt.Printf("load config file:%s\n", configName)
	file, e := ioutil.ReadFile(configName)
	if e != nil {
		fmt.Printf("Load config error: %v\n", e)
		os.Exit(1)
	}
	
	var config Configs
	err := json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("Config load error:%v \n",err)
		return false,config
	}
	return true,config
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetStructFieldsKV(v interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	val := reflect.Indirect(reflect.ValueOf(v))
        for i := 0; i < val.NumField(); i++ {
		fields[val.Type().Field(i).Name] = val.Field(i).String()
	}
	return fields
}
