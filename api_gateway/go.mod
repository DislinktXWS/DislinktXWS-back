module github.com/dislinktxws-back/api_gateway

go 1.18

replace github.com/dislinktxws-back/common => ../common

require (
	github.com/dislinktxws-back/common v0.0.0-00010101000000-000000000000
	github.com/gorilla/handlers v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	google.golang.org/grpc v1.46.0
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
