build:
	go build ./cmd/go-http-metrics.go

image: build
	docker build --no-cache -t quay.io/charlesakalugwu/go-http-metrics:0.0.5 .

push: image
	docker push quay.io/charlesakalugwu/go-http-metrics:0.0.5
