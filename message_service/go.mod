module github.com/dislinktxws-back/message_service

go 1.18

replace github.com/dislinktxws-back/common => ../common

require (
	github.com/dislinktxws-back/common v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.9.1
	google.golang.org/grpc v1.46.0
)
