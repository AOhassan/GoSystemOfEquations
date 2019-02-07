app: main.go
	# Compile the Go code to a static binary that can be used in a scratch
	# image.
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

docker: app
	docker build -t app .

clean:
	rm -f app

all: docker
