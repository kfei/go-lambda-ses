all: build

build: ses zip

clean:
	rm -f go-lambda-ses lambda.zip

ses: *.go
	GOOS=linux GOARCH=amd64 go build

zip: ses
	zip lambda.zip lambda.js go-lambda-ses 
