"# aerospike" 

To run aerospike : 
docker run -d --name aerospike -p 3000-3002:3000-3002 aerospike:ee-7.2.0.1 


Setup development env 
go get github.com/aerospike/aerospike-client-go/v7@v7.6.0
go run -tags as_proxy main.go