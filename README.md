```markdown
# Aerospike with Go

This repository provides an example of how to set up and interact with Aerospike using Go. It includes instructions for running Aerospike with Docker, setting up the Go development environment, handling policies, and working with Aerospike schema data models.

## Running Aerospike with Docker

To run Aerospike using Docker, execute the following command:

```bash
docker run -d --name aerospike -p 3000-3002:3000-3002 aerospike:ee-7.2.0.1
```

This command starts an **Aerospike Enterprise Edition** instance with ports `3000-3002` exposed for communication.

## Setting Up the Development Environment

1. **Install the Aerospike Go client**:

    To install the Aerospike Go client, run the following command:

    ```bash
    go get github.com/aerospike/aerospike-client-go/v7@v7.6.0
    ```

2. **Run your Go application**:

    To run your application with the `as_proxy` build tag, use the following:

    ```bash
    go run -tags as_proxy main.go
    ```

## Changing Policies in Aerospike

### Will Changing Policies Impact Stored Data?

Changing policies (e.g., write policies, read policies) **does not** affect the data already stored in Aerospike. Policies define how operations (e.g., reads, writes) are handled but do not modify existing data. 

### Can Policies Be Adjusted to Reduce Costs?

Yes, you can adjust policies to optimize costs. Some ways to reduce costs include:
- Reducing the number of replicas to save on storage.
- Lowering the write consistency level to reduce load.
- Adjusting the Time-to-Live (TTL) of records to allow data to expire sooner.

However, when adjusting policies, it’s essential to consider the trade-offs between cost and data reliability.

## Key Terms for Aerospike Schema Data Models

Here are some important terms related to Aerospike schema and data models:

### Namespace

A **namespace** is a logical container for sets, records, and policies in Aerospike. You can create multiple namespaces for different use cases or environments.

### Key

A **key** uniquely identifies a record within a namespace and set. It acts as the primary identifier for accessing records in Aerospike.

### Metadata

**Metadata** for a record includes:
- **Generation Counter**: Tracks the version of the record.
- **Time-to-Live (TTL)**: The expiration time for the record after which it is deleted.

### Bin

A **bin** holds the actual data for a record. Each record can have multiple bins, and each bin can store different types of data (e.g., strings, integers, lists, etc.).

### Sets

A **set** is a collection of records within a namespace. If a record is not assigned to a specific set, it will belong to the default `null` set.

## Struct to Bins Conversion

This repository provides a `StructToBins` function that allows you to convert Go structs into a format suitable for Aerospike operations. The function uses Go reflection to map struct fields to Aerospike bin names.

### Key Features:
- **Struct Tagging**: Maps Go struct fields to Aerospike bin names using struct tags (e.g., `as:"bin_name"`).
- **Support for Nested Structs**: Automatically flattens nested structs into their parent bins.
- **Nil and Zero Value Handling**: Supports the `omitempty` struct tag to skip zero or nil values.
- **Error Handling**: Returns errors if unsupported types or invalid values are encountered.

### Important Note:
When using the `GetRecord` method, ensure you pass a **pointer** to the struct to allow proper data mapping.

## Installing Aerospike Tools

To install Aerospike tools (for managing your Aerospike instance), follow these steps:

1. **Download the Aerospike tools package**:

    ```bash
    wget -O aerospike-tools.tgz https://download.aerospike.com/artifacts/aerospike-tools/latest/aerospike-tools_11.1.1_ubuntu20.04_x86_64.tgz
    ```

2. **Extract the downloaded package**:

    ```bash
    tar -xvf aerospike-tools.tgz
    ```

3. **Navigate to the extracted directory**:

    ```bash
    cd aerospike-tools_*
    ```

4. **Install the tools**:

    ```bash
    sudo ./asinstall
    ```

## Accessing the Aerospike CLI

Once the Aerospike tools are installed, you can access the Aerospike Command Line Interface (CLI) using the `aql` command:

```bash
aql -h 127.0.0.1 -p 3000
```

This will start an interactive session where you can run queries and manage your Aerospike database.

---