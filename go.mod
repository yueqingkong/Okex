module github.com/yueqingkong/Okex

go 1.12

replace golang.org/x/text => github.com/golang/text v0.3.2

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

replace golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297

require (
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000
	gopkg.in/ini.v1 v1.46.0 // indirect
	gopkg.in/resty.v1 v1.12.0 // indirect
)
