module github.com/yueqingkong/Okex

go 1.12

replace (
	github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
	golang.org/x/text => github.com/golang/text v0.3.2
)

require (
	github.com/go-resty/resty v1.12.0
	gopkg.in/ini.v1 v1.46.0 // indirect
)
