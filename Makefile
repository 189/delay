
dev:
	go run main.go
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o delay.guzibu.cn

run:
	nohup	delay.guzibu.cn &