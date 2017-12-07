package main

import (
	"time"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

type Store struct {
	Id                    string `json: "id"`
	Name                  string `json: "name"`
	Url                   string `json: "url"`
	Currency              string `json: "currency"`
	DescriptionQuery      string `json: "descriptionQuery"`
	PriceQuery            string `json: "priceQuery"`
	InternalCodeQuery     string `json: "internalCodeQuery"`
	EANQuery              string `json: "eanQuery"`
	TechnicalContentQuery string `json: "technicalContentQuery"`
	TechnicalKeyQuery     string `json: "technicalKeyQuery"`
	TechnicalValueQuery   string `json: "technicalValueQuery"`
}

type ProductPage struct {
	Id       string    `json: "id"`
	StoreId  string    `json: "storeId"`
	Url      string    `json: "url"`
	Parsed   bool      `json: "parsed"`
	ParsedAt time.Time `json: "parsedAt"`
}

type ProductStore struct {
	Id           string    `json: "id"`
	StoreId      string    `json: "storeId"`
	Description  string    `json: "description"`
	InternalCode int64     `json: "internalCode"`
	CreatedAt    time.Time `json: "createdAt"`
	UpdatedAt    time.Time `json: "updatedAt"`
}

type ProductOfferStore struct {
	Id             string    `json: "id"`
	ProductStoreId string    `json: "storeId"`
	Price          float64   `json: "price"`
	CreatedAt      time.Time `json: "createdAt"`
}

var Format struct {
	TimestampFormat string `default:"2006/01/02 15:04:05.000"`
}

type StoreSet map[string]*Store
type ProductPageSet map[string]*ProductPage
type ProductStoreSet map[string]*ProductStore
type ProductOfferStoreSet map[string]*ProductStore

var (
	storeSet       = make(StoreSet)
	productPageSet = make(ProductPageSet)
)

func main() {

	storeSet["store1"] = &Store{
		Id:                    "store1",
		Name:                  "Americanas",
		Url:                   "http://www.americanas.com.br",
		Currency:              "BRL",
		DescriptionQuery:      "div.product-info-area.col-sm-7.col-xs-12 > section > h1",
		PriceQuery:            "div > div.main-price > p.sales-price",
		InternalCodeQuery:     "div.product-info-area.col-sm-7.col-xs-12 > section > span.product-id",
		TechnicalContentQuery: "div.card-panel.card-info > section:nth-child(3) > div > table > tbody > tr",
		TechnicalKeyQuery:     "td:nth-child(1)",
		TechnicalValueQuery:   "td:nth-child(2)",
	}

	storeSet["store2"] = &Store{
		Id:                "store2",
		Name:              "Ponto Frio",
		Url:               "http://www.pontofrio.com.br",
		Currency:          "BRL",
		DescriptionQuery:  "#ctl00_Conteudo_upMasterProdutoBasico > div.produtoNome > h1 > b",
		PriceQuery:        "#ctl00_Conteudo_ctl01_precoPorValue > i",
		InternalCodeQuery: "#ctl00_Conteudo_upMasterProdutoBasico > div.produtoNome > div > span:nth-child(1)",
		EANQuery:          "#ctl00_Conteudo_upMasterProdutoBasico > div.produtoNome > div > span.productEan",
		TechnicalContentQuery: "#caracteristicas > div > dl",
		TechnicalKeyQuery:     "dt",
		TechnicalValueQuery:   "dd",


	}

	storeSet["store3"] = &Store{
		Id:                "store3",
		Name:              "Ricardo Eletro",
		Url:               "http://www.ricardoeletro.com.br",
		Currency:          "BRL",
		DescriptionQuery:  "#ProdutoDetalhesNomeProduto > h1",
		PriceQuery:        "#ProdutoDetalhesPrecoComprarAgoraPrecoDePreco",
		InternalCodeQuery: "#ProdutoDetalhesCodigoProduto",
		TechnicalContentQuery: "#aba-caracteristicas > div:nth-child(2) > table > tbody > tr",
		TechnicalKeyQuery:     "td:nth-child(1)",
		TechnicalValueQuery:   "td:nth-child(2)",
	}

	storeSet["store4"] = &Store{
		Id:                "store4",
		Name:              "Magazine Luiza",
		Url:               "http://www.magazineluiza.com.br",
		Currency:          "BRL",
		DescriptionQuery:  "div.header-product.js-header-product > h1",
		PriceQuery:        "div.information-values__product-page > div > div > div > span.price-template__text",
		InternalCodeQuery: "div.header-product.js-header-product > small",
	}

	productPageSet["ppage10"] = &ProductPage{
		Id:       "ppage10",
		StoreId:  "store1",
		Parsed:   true,
		Url:      "https://www.americanas.com.br/produto/132246221/iphone-7-128gb-preto-matte-desbloqueado-ios-10-wi-fi-4g-camera-12mp-apple?pfm_carac=iphone%207&pfm_index=2&pfm_page=search&pfm_pos=grid&pfm_type=search_page%20",
		ParsedAt: time.Now(),
	}

	productPageSet["ppage11"] = &ProductPage{
		Id:       "ppage11",
		StoreId:  "store1",
		Parsed:   true,
		Url:      "https://www.americanas.com.br/produto/132118431/smartphone-samsung-galaxy-s8-dual-chip-android-7.0-tela-5.8-octa-core-2.3ghz-64gb-4g-camera-12mp-preto?pfm_carac=galaxy%20s8&pfm_index=0&pfm_page=search&pfm_pos=grid&pfm_type=search_page%20",
		ParsedAt: time.Now(),
	}

	productPageSet["ppage20"] = &ProductPage{
		Id:       "ppage20",
		StoreId:  "store2",
		Parsed:   true,
		Url:      "https://www.pontofrio.com.br/TelefoneseCelulares/Smartphones/iPhone/iphone-7-apple-128gb-tela-retina-hd-de-47-3d-touch-ios-10-touch-id-cam-12mp-resistente-a-agua-e-sistema-de-alto-falantes-estereo-preto-matte-11526482.html?IdProduto=7995624&recsource=btermo&rectype=p1_op_s2",
		ParsedAt: time.Now(),
	}

	productPageSet["ppage30"] = &ProductPage{
		Id:       "ppage30",
		StoreId:  "store3",
		Parsed:   true,
		Url:      "http://www.ricardoeletro.com.br/Produto/Celular-Smartphone-Samsung-Galaxy-S7-Edge-G935F-Dourado-4G-Tela-Curva-55-AMOLED-Camera-12MPFrontal-5MP-Octa-Core-23Ghz-32GB-4GB-RAM-Android-6/44-491-496-585914/?utm_source=Google_Shopping&prc=24522&utm_medium=CPC_Celulares_e_Telefones_Google_Shopping&utm_campaign=Smartphones&utm_content=Samsung",
		ParsedAt: time.Now(),
	}

	productPageSet["ppage40"] = &ProductPage{
		Id:       "ppage40",
		StoreId:  "store4",
		Parsed:   true,
		Url:      "https://www.magazineluiza.com.br/smartphone-galaxy-s8-g955fd-ametista-dual-chip-tela-6.2-cam.-12mp-64gb-android-7.0-samsung/p/6052629/te/tcsp/",
		ParsedAt: time.Now(),
	}

	//productStoreSet := make(ProductStoreSet)

	for _, productPage := range productPageSet {

		doc, _ := goquery.NewDocument(productPage.Url)

		name := storeSet[productPage.StoreId].Name
		fmt.Printf("STORE: %s\n", name)

		description := doc.Find(storeSet[productPage.StoreId].DescriptionQuery).Text()
		fmt.Printf("Description: %s\n", description)

		priceQuery := storeSet[productPage.StoreId].PriceQuery
		curr := storeSet[productPage.StoreId].Currency
		price := doc.Find(priceQuery).Text()
		fmt.Printf("Price: %s\n", strconv.FormatFloat(parseCurrency(price, curr), 'f', -1, 64))

		internalCodeQuery := storeSet[productPage.StoreId].InternalCodeQuery

		if internalCodeQuery != "" {
			internalCode := doc.Find(internalCodeQuery).Text()
			fmt.Printf("InternalCode: %s\n", strconv.Itoa(int(onlyNumbers(internalCode))))
		}

		eanQuery := storeSet[productPage.StoreId].EANQuery

		if eanQuery != "" {
			ean := doc.Find(eanQuery).Text()
			fmt.Printf("EAN: %s\n", strconv.Itoa(int(onlyNumbers(ean))))
		}

		technicalContentQuery := storeSet[productPage.StoreId].TechnicalContentQuery

		if technicalContentQuery != "" {

			fmt.Println("# Technical Info #")

			doc.Find(technicalContentQuery).Each(func(i int, s *goquery.Selection) {
				key := strings.TrimSpace(s.Find(storeSet[productPage.StoreId].TechnicalKeyQuery).Text())
				value := strings.TrimSpace(s.Find(storeSet[productPage.StoreId].TechnicalValueQuery).Text())
				fmt.Printf("%s: %s\n", key, value)
			})

		}

		fmt.Printf("\n\n")

	}

}

func onlyNumbers(value string) int64 {

	reg, _ := regexp.Compile("[0-9]+")
	chars := reg.FindAllString(value, -1)

	r := strings.Join(chars, "")

	i, _ := strconv.ParseInt(r, 0, 64)

	return i
}

func parseCurrency(value string, curr string) float64 {

	reg, _ := regexp.Compile("[0-9,]+")
	chars := reg.FindAllString(value, -1)
	fValue := strings.Join(chars, "")

	/*
	//https://github.com/golang/text/blob/master/currency/format_test.go
	formatter := currency.ISO.Default(currency.MustParseISO(curr))

	//formatter := currency.Symbol

	tag := language.BrazilianPortuguese

	printer := message.NewPrinter(tag)

	v := formatter(fValue)

	f := printer.Sprint(v)

	fmt.Println(f)

	fmt.Println(v)

	return 0.0

	*/

	rr := strings.Replace(fValue, ",", ".", -1)

	f, _ := strconv.ParseFloat(rr, 64)

	return f

}
