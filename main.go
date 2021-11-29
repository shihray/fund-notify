package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"stocker/enum"
	"stocker/tool"
	"strings"
)

var (
	conf enum.Config
	msg  string
)

func init() {
	// Load yaml file
	conf.GetConf()
}

func main() {

	//tool.InitChromeDP()
	//defer tool.Cancel()
	//defer tool.CancelCtx()

	for _, info := range conf.Fund {
		value := tool.GetWebKValueByRod(info.Name, info.Url)

		strSplit := strings.Split(value, "|")

		fmt.Println(info.Name + ": K9值:" + strSplit[0] + " 價位:" + strSplit[1])

		kdkDec, _ := decimal.NewFromString(strSplit[0])
		if kdkDec.LessThan(decimal.NewFromFloat(30)) {
			_str := fmt.Sprintf("%v: K9值:%v, 價位:%v", info.Name, strSplit[0], strSplit[1])
			if msg == "" {
				msg = _str
			} else {
				msg = fmt.Sprintf("%v\n%v", msg, _str)
			}
		}
	}
	// line Notify
	if err := tool.LineNotify(conf.Token, msg); err != nil {
		fmt.Println("line notify err:", err.Error())
	}
}
