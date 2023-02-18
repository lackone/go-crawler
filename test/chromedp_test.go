package test

import (
	"context"
	"github.com/chromedp/chromedp"
	"testing"
	"time"
)

func TestChromedp(t *testing.T) {
	ctx, fn := chromedp.NewContext(context.Background())
	defer fn()

	ctx, fn = context.WithTimeout(ctx, 15*time.Second)
	defer fn()

	example := ""

	err := chromedp.Run(ctx,
		//指定网址
		chromedp.Navigate("https://pkg.go.dev/time"),
		//等待标签可见，说明dom加载完成
		chromedp.WaitVisible(`body > footer`),
		//模拟点击事件
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		//获取标签数据
		chromedp.Value(`#example-After textarea`, &example),
	)
	if err != nil {
		t.Log(err)
	}
	t.Log(example)
}
