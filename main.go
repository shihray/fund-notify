package main

import (
	"flag"
	"fmt"
	"stocker/enum"
	"stocker/tool"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
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

		kdkDec, _ := decimal.NewFromString(strSplit[0])

		kfloat, _ := strconv.ParseFloat(strSplit[0], 64)
		pricefloat, _ := strconv.ParseFloat(strSplit[1], 64)

		fmt.Printf("%v: K9值:%0.2f, 價位:%.2f\n", info.Name, kfloat, pricefloat)

		_str := fmt.Sprintf("\n%v: K9值:%0.2f, 價位:%.2f", info.Name, kfloat, pricefloat)
		if kdkDec.LessThan(decimal.NewFromFloat(20)) {
			msg += _str
		}
		allMsg += _str
	}

	if *alwaysSend {
		if err := tool.LineNotify(conf.Token, allMsg); err != nil {
			fmt.Println("line notify err:", err.Error())
		}
	} else {
		// line Notify
		if err := tool.LineNotify(conf.Token, msg); err != nil {
			fmt.Println("line notify err:", err.Error())
		}
	}
}
