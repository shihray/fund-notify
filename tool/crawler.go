package tool

import (
	"context"
	"encoding/json"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"log"
	"stocker/enum"
	"time"
)

var (
	timeoutCtx context.Context
	CancelCtx  context.CancelFunc
	chromeCtx  context.Context
	Cancel     context.CancelFunc
)

func InitChromeDP() {
	// 設置項，具體參照參考1
	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		chromedp.Flag("headless", false),                       // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"), // 禁用圖片加載
		chromedp.WindowSize(720, 640),                          // 设置浏览器窗口尺寸,
		chromedp.Flag("no-sandbox", true),                      // 启动chrome 不适用沙盒, 性能优先
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	// Chrome初始化代碼如下：
	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	// create context
	chromeCtx, Cancel = chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 執行一個空task, 用提前創建Chrome實例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)
	// 給每個頁面的爬取設置超時時間
	timeoutCtx, CancelCtx = context.WithTimeout(chromeCtx, 30*time.Second)
}

func GetWebKValue(url string) string {
	// run task list
	var res *runtime.RemoteObject

	if err := chromedp.Run(timeoutCtx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Evaluate(`globalData[244].KD_K`, &res, chromedp.EvalAsValue),
	}); err != nil {
		log.Fatal(err)
		return ""
	}

	var resp enum.RespSt
	jsonByte, _ := res.MarshalJSON()

	if err := json.Unmarshal(jsonByte, &resp); err != nil {
		log.Fatal(err)
		return ""
	}
	return resp.Value
}
