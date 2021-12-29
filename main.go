package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PAGE_ITEM_MAX_COUNT = 12

var filteredProducts []product

type product struct {
	ProductId     int64   `json:"productId"`
	Title         string  `json:"title"`
	ColorId       int64   `json:"colorId"`
	ColorTitle    string  `json:"colorTitle"`
	BrandId       int64   `json:"brandId"`
	BrandTitle    string  `json:"brandTitle"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"originalPrice"`
	DiscountRate  float64 `json:"discountRate"`
	ImageUrl      string  `json:"imageURL"`
}
type brand struct {
	BrandId int64  `json:"brandId"`
	Title   string `json:"title"`
}
type color struct {
	ColorId int64  `json:"colorId"`
	Title   string `json:"title"`
}
type orderFilter struct {
	OrderFilterId int64  `json:"orderFilterId"`
	Title         string `json:"title"`
}

type brandResult struct {
	Brand brand `json:"brand"`
	Count int64 `json:"count"`
}

type colorResult struct {
	Color color `json:"color"`
	Count int64 `json:"count"`
}

type queryResult struct {
	Products  []product     `json:"products"`
	Brands    []brandResult `json:"brands"`
	Colors    []colorResult `json:"colors"`
	PageIndex int64         `json:"pageIndex"`
	PageCount int           `json:"pageCount"`
}

func getProducts(c *gin.Context) {
	query := c.Query("q")
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	brandId, _ := strconv.ParseInt(c.Query("brandId"), 10, 64)
	colorId, _ := strconv.ParseInt(c.Query("colorId"), 10, 64)
	orderId, _ := strconv.ParseInt(c.Query("orderId"), 10, 64)

	fmt.Println(query)
	fmt.Println(page)
	fmt.Println(brandId)
	fmt.Println(colorId)

	// Filter Section
	linq.From(products).WhereT(func(p product) bool {
		var queryResult = true
		var brandResult = true
		var colorResult = true

		if query != "" {
			queryResult = strings.Contains(strings.ToLower(p.Title), strings.ToLower(query))
		}

		if brandId != 0 {
			brandResult = p.BrandId == brandId
		}

		if colorId != 0 {
			colorResult = p.ColorId == colorId
		}

		return queryResult && brandResult && colorResult
	}).ToSlice(&filteredProducts)

	// Order Operations
	switch orderId {
	// En Düşük Fiyat
	case 1:
		linq.From(filteredProducts).OrderByT(func(p product) interface{} {
			return p.Price
		}).ToSlice(&filteredProducts)
	// En Yüksek Fiyat
	case 2:
		linq.From(filteredProducts).OrderByDescendingT(func(p product) interface{} {
			return p.Price
		}).ToSlice(&filteredProducts)
	// A-Z
	case 3:
		linq.From(filteredProducts).OrderByT(func(p product) interface{} {
			return p.Title
		}).ToSlice(&filteredProducts)
	// Z-A
	case 4:
		linq.From(filteredProducts).OrderByDescendingT(func(p product) interface{} {
			return p.Title
		}).ToSlice(&filteredProducts)
	}

	// @TODO Shared Function
	// BrandId Filter Counts
	var brandResults []brandResult
	linq.From(filteredProducts).GroupByT(
		func(p product) interface{} {
			return p.BrandId
		}, func(p product) interface{} {
			return p.BrandId
		}).Select(func(group interface{}) interface{} {

		innerBrand := linq.From(brands).FirstWith(
			func(b interface{}) bool {
				return b.(brand).BrandId == group.(linq.Group).Key
			},
		)

		return brandResult{Brand: innerBrand.(brand), Count: int64(linq.From(group.(linq.Group).Group).Count())}
	}).ToSlice(&brandResults)

	// ColorId Filter Counts
	var colorResults []colorResult
	linq.From(filteredProducts).GroupByT(
		func(p product) interface{} {
			return p.ColorId
		}, func(p product) interface{} {
			return p.ColorId
		}).Select(func(group interface{}) interface{} {

		innerColor := linq.From(colors).FirstWith(
			func(b interface{}) bool {
				return b.(color).ColorId == group.(linq.Group).Key
			},
		)

		return colorResult{Color: innerColor.(color), Count: int64(linq.From(group.(linq.Group).Group).Count())}
	}).ToSlice(&colorResults)

	var pageCount = 1
	if len(filteredProducts) > PAGE_ITEM_MAX_COUNT {
		pageCount = len(filteredProducts) / PAGE_ITEM_MAX_COUNT
	}
	// Pagination
	linq.From(filteredProducts).Skip(int(page) * PAGE_ITEM_MAX_COUNT).Take(PAGE_ITEM_MAX_COUNT).ToSlice(&filteredProducts)
	for i, _ := range filteredProducts {
		findedBrand := linq.From(brands).FirstWithT(func(b brand) bool {
			return b.BrandId == filteredProducts[i].BrandId
		})
		if findedBrand != nil {
			filteredProducts[i].BrandTitle = findedBrand.(brand).Title
		}

		findedColor := linq.From(colors).FirstWithT(func(c color) bool {
			return c.ColorId == filteredProducts[i].ColorId
		})
		if findedColor != nil {
			filteredProducts[i].ColorTitle = findedColor.(color).Title
		}
	}

	//@TODO Calculate discount price

	var result queryResult
	result.Products = filteredProducts
	result.Brands = brandResults
	result.Colors = colorResults
	result.PageIndex = page
	result.PageCount = pageCount

	c.IndentedJSON(http.StatusOK, result)
}

func getOrderFilters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, orderFilters)
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/products", getProducts)
	router.GET("/order-filters", getOrderFilters)
	err := router.Run()

	if err != nil {
		fmt.Println(err)
	}
}

// Product Data
var products = []product{
	{ProductId: 1, Title: "Apple iPhone 6 64 GB", ColorId: 1, BrandId: 1, Price: 90.85, OriginalPrice: 2000, DiscountRate: 24, ImageUrl: "./{productId}.png"},
	{ProductId: 2, Title: "Apple iPhone 7 64 GB", ColorId: 4, BrandId: 1, Price: 1999.85, OriginalPrice: 2000, DiscountRate: 12, ImageUrl: "./{productId}.png"},
	{ProductId: 3, Title: "Apple iPhone 8 128 GB", ColorId: 3, BrandId: 1, Price: 90.85, OriginalPrice: 124, DiscountRate: 3, ImageUrl: "./{productId}.png"},
	{ProductId: 4, Title: "Apple iPhone 9 64 GB", ColorId: 6, BrandId: 1, Price: 95.85, OriginalPrice: 124, DiscountRate: 4, ImageUrl: "./{productId}.png"},
	{ProductId: 5, Title: "Apple iPhone 10 128 GB (Uzun isimli Türkiye Apple Garantili Telefon)", ColorId: 5, BrandId: 1, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 6, Title: "Apple iPhone 11 256 GB (Uzun isimli Türkiye Apple Garantili Telefon)", ColorId: 1, BrandId: 1, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 7, Title: "Apple iPhone 12 64 GB", ColorId: 1, BrandId: 1, Price: 1750, OriginalPrice: 1850, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 8, Title: "Apple iPhone 13 64 GB (Uzun isimli Türkiye Apple Garantili Telefon)", ColorId: 2, BrandId: 1, Price: 2000, OriginalPrice: 2000, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 9, Title: "Apple iPhone 13 Pro", ColorId: 4, BrandId: 1, Price: 2100, OriginalPrice: 2200, DiscountRate: 24, ImageUrl: "./{productId}.png"},
	{ProductId: 10, Title: "Samsung Galaxy S20 128 GB", ColorId: 2, BrandId: 2, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 11, Title: "Samsung Galaxy A22", ColorId: 3, BrandId: 2, Price: 1475, OriginalPrice: 1800, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 12, Title: "Samsugn Galaxy A52s 64 GB", ColorId: 2, BrandId: 1, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 13, Title: "Samsung Galaxy M32 5G", ColorId: 2, BrandId: 2, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 14, Title: "Samsung Galaxy Z Fold3 5G (Uzun isimli Türkiye Samsung Garantili Telefon)", ColorId: 1, BrandId: 2, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 15, Title: "Samsung Galaxy M12", ColorId: 1, BrandId: 2, Price: 190.85, OriginalPrice: 2150, DiscountRate: 15, ImageUrl: "./{productId}.png"},
	{ProductId: 16, Title: "Huawei P40 Lite 128 GB", ColorId: 2, BrandId: 3, Price: 95.85, OriginalPrice: 110, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 17, Title: "Huawei P Smart 2021 128 GB (Uzun isimli Türkiye Huawei Garantili Telefon)", ColorId: 1, BrandId: 3, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 18, Title: "Huawei Y5p 32 GB", ColorId: 1, BrandId: 3, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 19, Title: "Huawei Mate 20 Lite 64 GB (Uzun isimli Türkiye Huawei Garantili Telefon)", ColorId: 1, BrandId: 3, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 20, Title: "Huawei Y6S 32 GB", ColorId: 6, BrandId: 3, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 21, Title: "Huawei Xs 512 GB", ColorId: 2, BrandId: 3, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 22, Title: "Nokia 1280", ColorId: 1, BrandId: 4, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 23, Title: "Nokia C3", ColorId: 6, BrandId: 4, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 24, Title: "Nokia 101", ColorId: 3, BrandId: 4, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 25, Title: "Nokia C1-01", ColorId: 1, BrandId: 4, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 26, Title: "Nokia 6700 Classic (Uzun isimli Türkiye Nokia Garantili Telefon)", ColorId: 1, BrandId: 4, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 27, Title: "Xiaomi Redmi Note 10S 128 GB", ColorId: 5, BrandId: 5, Price: 590.85, OriginalPrice: 6124, DiscountRate: 80, ImageUrl: "./{productId}.png"},
	{ProductId: 28, Title: "Xiaomi Redmi Note 10 Pro 128 GB (Uzun isimli Türkiye Xiaomi Garantili Telefon)", ColorId: 1, BrandId: 5, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 29, Title: "Xiaomi Redmi Note 9 Pro 128 GB", ColorId: 1, BrandId: 5, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 30, Title: "Xiaomi Redmi 9T 128 GB (Uzun isimli Türkiye Xiaomi Garantili Telefon)", ColorId: 1, BrandId: 5, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 31, Title: "Xiaomi Mi 11T 256 GB", ColorId: 6, BrandId: 5, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 32, Title: "General Mobile Gm 21 Pro 128 GB", ColorId: 2, BrandId: 6, Price: 90.85, OriginalPrice: 1224, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 33, Title: "General Mobile Gm 21 Plus 64 GB", ColorId: 6, BrandId: 6, Price: 90.85, OriginalPrice: 1324, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 34, Title: "General Mobile Gm 21 32 GB (Uzun isimli Türkiye General Mobile Garantili Telefon)", ColorId: 1, BrandId: 6, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 35, Title: "General Mobile GM8 Go", ColorId: 3, BrandId: 6, Price: 90.85, OriginalPrice: 124, DiscountRate: 0, ImageUrl: "./{productId}.png"},
	{ProductId: 36, Title: "LG K61 128 GB", ColorId: 1, BrandId: 7, Price: 190.85, OriginalPrice: 1254, DiscountRate: 25, ImageUrl: "./{productId}.png"},
	{ProductId: 37, Title: "LG K4S1 32 GB", ColorId: 4, BrandId: 7, Price: 90.85, OriginalPrice: 1224, DiscountRate: 45, ImageUrl: "./{productId}.png"},
	{ProductId: 38, Title: "LG K10 2017", ColorId: 5, BrandId: 7, Price: 90.85, OriginalPrice: 1224, DiscountRate: 34, ImageUrl: "./{productId}.png"},
}

// Color Data
var colors = []color{
	{ColorId: 1, Title: "Siyah"},
	{ColorId: 2, Title: "Kırmızı"},
	{ColorId: 3, Title: "Sarı"},
	{ColorId: 4, Title: "Turuncu"},
	{ColorId: 5, Title: "Mor"},
	{ColorId: 6, Title: "Beyaz"},
}

// Brand Data
var brands = []brand{
	{BrandId: 1, Title: "Apple"},
	{BrandId: 2, Title: "Samsung"},
	{BrandId: 3, Title: "Huawei"},
	{BrandId: 4, Title: "Nokia"},
	{BrandId: 5, Title: "Xiaomi"},
	{BrandId: 6, Title: "General Mobile"},
	{BrandId: 7, Title: "LG"},
}

// Order Data
var orderFilters = []orderFilter{
	{OrderFilterId: 1, Title: "En Düşük Fiyat"},
	{OrderFilterId: 2, Title: "En Yüksek Fiyat"},
	{OrderFilterId: 3, Title: "En Yeniler (A>Z)"},
	{OrderFilterId: 4, Title: "En Yeniler (Z<A)"},
}
