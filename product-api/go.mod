module github.com/myk4040okothogodo/Grpc2/product-api

go 1.17

replace github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos => ../currency/protos/currency/protos

replace github.com/myk4040okothogodo/Grpc2/product-api/data => ./data

replace github.com/myk4040okothogodo/Grpc2/product-api/handlers => ./handlers

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/Grpc2/product-api/data v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/Grpc2/product-api/handlers v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.44.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/myk4040okothogodo/env => ../../env
