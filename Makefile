dep:
	go get github.com/gin-gonic/gin

build:
	(cd cmd/server/ \
	&& go build)

run :
	go run cmd/server/main.go
