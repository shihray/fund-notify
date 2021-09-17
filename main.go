package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"stocker/enum"
	"stocker/tool"
)

var conf enum.Config

func init() {
	// Load yaml file
	conf.GetConf()
}

func main() {

	tool.InitChromeDP()
	defer tool.Cancel()
	defer tool.CancelCtx()

	for _, info := range conf.Fund {
		value := tool.GetWebKValue(info.Url)
		log.Println(info.Name + ": K9值:" + value)

		kdkDec, _ := decimal.NewFromString(value)
		if kdkDec.LessThan(decimal.NewFromFloat(30)) {
			msg := fmt.Sprintf("%v: K9值:%v", info.Name, value)
			// line Notify
			if err := tool.LineNotify(msg); err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}
