package tool

import (
	"github.com/go-rod/rod"
)

func GetWebKValueByRod(name, url string) string {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url)
	page.MustWindowMinimize()
	page.MustWaitLoad()

	// print all screen
	//page.MustScreenshot(name + ".png")

	// define page size
	//img, _ := page.Screenshot(true, &proto.PageCaptureScreenshot{
	//	Format:  proto.PageCaptureScreenshotFormatJpeg,
	//	Quality: 90,
	//	Clip: &proto.PageViewport{
	//		X:      0,
	//		Y:      0,
	//		Width:  700,
	//		Height: 700,
	//		Scale:  1,
	//	},
	//	FromSurface: true,
	//})
	//_ = utils.OutputFile(name+".png", img)

	val := page.MustEval(`globalData[globalData.length - 1].KD_K + "|" + globalData[globalData.length - 1][3]`).Str()

	return val
}
