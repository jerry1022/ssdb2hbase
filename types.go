package main
type Configs struct {
	SSDB struct {
		Cron string `json:"Cron"`
		Debug bool `json:"Debug"`
		Address string `json:"Address"`
		Port int `json:"Port"`
		Auth string `json:"Auth"`
		HT string `json:"Hashtable"`
	} `json:"SSDB"`
	HBase struct {
		Cron string `json:"Cron"`
		Debug bool `json:"Debug"`
		Address string `json:"Address"`
		Port int `json:"Port"`
		Tbl string `json:"Table"`
	} `json:"HBASE"`
}

type ProxyCounter struct {
	Idx int `json:"idx"`
	Vip string `json:"vip"`
	Bip string `json:"bip"`
	Type string `json:"type"`
	Status string `json:"status"`
	Busy float64 `json:"busy"`
	Rise float64 `json:"rise"`
	Fall float64 `json:"fall"`
	Access float64 `json:"access"`
	BytesRead float64 `json:"bytes.read"`
	BytesWrite float64 `json:"bytes.write"`
}

