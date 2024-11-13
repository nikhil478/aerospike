# Aerospike

## Running Aerospike

To start Aerospike using Docker, run the following command:

```bash
docker run -d --name aerospike -p 3000-3002:3000-3002 aerospike:ee-7.2.0.1
```

## Setting Up the Development Environment

To set up your development environment, execute:

```bash
go get github.com/aerospike/aerospike-client-go/v7@v7.6.0
go run -tags as_proxy main.go
```

## Changing Policies

If you change the policy, will it impact the data stored? Are you allowed to adjust policies to reduce costs? 

## Key Terms for Aerospike Schema Data Models

### Namespace
This is where all policies, sets, and record data reside, along with the last update time. You can create multiple namespaces based on your requirements.

### Key
The unique identifier for each record.

### Metadata
Includes the generation counter and time to live (TTL) for records.

### Bin
This is where the actual data is stored. A bin can accept any type of data.

### Sets
A collection of records. If you don’t define a set for a record, it will belong to the null set within the namespace.

This package offers a function, StructToBins, which takes a struct and converts it into a BinMap suitable for Aerospike operations. The function uses reflection to iterate through the struct fields, retrieve the tags, and handle the conversion to the appropriate Aerospike data types.

Key Features:
Struct Tagging: Uses Go struct tags (as:"bin_name") to map struct fields to Aerospike bin names.
Support for Nested Structs: Automatically handles nested structs and flattens them into the parent bin.
Nil and Zero Value Handling: Supports the omitempty tag to skip zero or nil values.
Error Handling: Returns an error if unsupported types or nil values for required fields are encountered.