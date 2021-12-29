# search-api

Deneysel amaçlı olarak Go dilini öğrenmek üzerine geliştirilmiş basit bir search servisidir.

Canlı Demo: https://fierce-citadel-22205.herokuapp.com/products

## Geliştirme Ortamı

Localde geliştirme yapmak için `go run .`

Docker için `docker-compose -f docker-compose.yml up ` komutlarıyla uygulamayı ayağa kaldırabilirsiniz.

localhost:8080 adresinde proje ayağa kalkmış olacaktır.

## Örnek Sorgu


Siyah Apple telefonları sorgulamak için aşağıdaki gibi bir query oluşur.

`https://fierce-citadel-22205.herokuapp.com/products?q=apple&brandId=1&colorId=1&orderId=1`

## In-Memory DataSets

### Sıralama

```javascript
{OrderFilterId: 1, Title: "En Düşük Fiyat"},
{OrderFilterId: 2, Title: "En Yüksek Fiyat"},
{OrderFilterId: 3, Title: "En Yeniler (A>Z)"},
{OrderFilterId: 4, Title: "En Yeniler (Z<A)"},
```
### Firmalar

```javascript
{BrandId: 1, Title: "Apple"},
{BrandId: 2, Title: "Samsung"},
{BrandId: 3, Title: "Huawei"},
{BrandId: 4, Title: "Nokia"},
{BrandId: 5, Title: "Xiaomi"},
{BrandId: 6, Title: "General Mobile"},
{BrandId: 7, Title: "LG"},
```
### Renkler

```javascript
{ColorId: 1, Title: "Siyah"},
{ColorId: 2, Title: "Kırmızı"},
{ColorId: 3, Title: "Sarı"},
{ColorId: 4, Title: "Turuncu"},
{ColorId: 5, Title: "Mor"},
{ColorId: 6, Title: "Beyaz"},
```
### Telefonlar

```javascript
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
