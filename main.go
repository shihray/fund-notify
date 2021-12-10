package main

import (
	"flag"
	"fmt"
	"github.com/shopspring/decimal"
	"stocker/enum"
	"stocker/tool"
	"strings"
)

var (
	conf       enum.Config
	msg        string
	alwaysSend = flag.Bool("s", false, "是否一定要發送")
	allMsg     string
)

func init() {
	// Load yaml file
	conf.GetConf()
	flag.Parse()
}

func main() {

	for _, info := range conf.Fund {
		value := tool.GetWebKValueByRod(info.Name, info.Url)

		strSplit := strings.Split(value, "|")

		fmt.Println(info.Name + ": K9值:" + strSplit[0] + " 價位:" + strSplit[1])

		kdkDec, _ := decimal.NewFromString(strSplit[0])
		_str := fmt.Sprintf("\n%v: K9值:%v, 價位:%v", info.Name, strSplit[0], strSplit[1])
		if kdkDec.LessThan(decimal.NewFromFloat(30)) {
			msg += _str
		}
		allMsg += _str
	}
	// line Notify
	if err := tool.LineNotify(conf.Token, msg); err != nil {
		fmt.Println("line notify err:", err.Error())
	}

	if *alwaysSend {
		if err := tool.LineNotify(conf.Token, allMsg); err != nil {
			fmt.Println("line notify err:", err.Error())
		}
	}
}
