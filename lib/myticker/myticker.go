package myticker

import (
	"gs/lib/mylog"
	"sync"
	"time"
)

type MyTicker struct {
	tickers []*time.Ticker
}

var myTicker *MyTicker
var once sync.Once

func GetInstance() *MyTicker {
	once.Do(func() {
		myTicker = new(MyTicker)
		myTicker.init()
	})
	return myTicker
}

func (this *MyTicker) init() {

}

func (this *MyTicker) AddTicker(d time.Duration, fn func()) {
	if fn == nil {
		mylog.Warning("fn is nil")
		return
	}
	ticker := time.NewTicker(d)
	this.tickers = append(this.tickers, ticker)
	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			}
		}
	}()
}

func (this *MyTicker) Destory() {
	for _, v := range this.tickers {
		v.Stop()
	}
	this.tickers = this.tickers[:0]
}
