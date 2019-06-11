build:
	go build ./cmd/go-http-metrics.go

image:
	docker build --no-cache -t quay.io/charlesakalugwu/go-http-metrics:latest .

push: image
	docker push quay.io/charlesakalugwu/go-http-metrics:latest
