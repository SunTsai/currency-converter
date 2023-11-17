## Currency Converter API

## Getting started

### Prerequisites
* Have [docker](https://www.docker.com/products/docker-desktop/) installed in your computer
* Clone the current repo

### Running Currency Converter
* Build the docker image  
```sh
docker build -t currency-converter:latest .
```

* Start the container  
```sh
docker run --name currency-converter -p 8080:8080 -d currency-converter:latest
```

## Checking the result
* Use terminal
```sh
curl 'http://localhost:8080/currencyConversion?source=JPY&target=USD&amount=100'
```

* Use browser, enter the URL  
```
http://localhost:8080/currencyConversion?source=JPY&target=USD&amount=100
```

* Use Postman

## Cleaning up
* Stop the container  
```sh
docker stop currency-converter
```

* Remove the container  
```sh
docker rm currency-converter
```

* Remove the image  
```sh
docker rmi currency-converter:latest
```

## Running unit test
* Simply use the command  
```sh
go test -v -cover ./...
```
