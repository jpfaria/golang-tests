package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
	"time"
)

type Ext struct {
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {

	salesPrice := doc.Find("p.sales-price").Text()
	salesPriceRegex := regexp.MustCompile("[0-9,]+")
	salesPriceFinal := strings.Join(salesPriceRegex.FindAllString(salesPrice, -1), "")

	productName := doc.Find("h1.product-name").Text()

	productCode := doc.Find("span.product-id").Text()
	productCodeRegex := regexp.MustCompile("[0-9]+")
	productCodeFinal := strings.Join(productCodeRegex.FindAllString(productCode, -1), "")

	if salesPrice != "" {
		fmt.Printf("PRD: %s # %s # %s\n", productName, salesPriceFinal, productCodeFinal)
	}

	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}
	if ctx.URL().Host == "americanas.com" || ctx.URL().Host == "www.americanas.com" || ctx.URL().Host == "www.americanas.com.br" || ctx.URL().Host == "americanas.com.br" {
		return true
	}
	return false
}

func (e *Ext) Visited(ctx *gocrawl.URLContext, harvested interface{}) {

}

func (e *Ext) Enqueued(ctx *gocrawl.URLContext) {

}

func (e *Ext) ComputeDelay(host string, di *gocrawl.DelayInfo, lastFetch *gocrawl.FetchInfo) time.Duration {
	return 0
}

func main() {
	ext := &Ext{&gocrawl.DefaultExtender{}}
	// Set custom options
	opts := gocrawl.NewOptions(ext)
	//opts.CrawlDelay = 1 * time.Second
	opts.CrawlDelay = 0
	opts.LogFlags = gocrawl.LogInfo
	opts.SameHostOnly = false
	opts.WorkerIdleTTL = 0
	//opts.MaxVisits = 100

	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("http://americanas.com")
}