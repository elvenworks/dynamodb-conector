# Simple AWS Dynamodb Connector 

 [![SDK Documentation](https://img.shields.io/badge/SDK-Documentation-blue)](https://aws.github.io/aws-sdk-go-v2/docs/) [![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://pkg.go.dev/mod/github.com/aws/aws-sdk-go-v2) 


`dynamodb-conector` is a simple client AWS DynamoDB to perform basic operations on a table. 
This connector use the v2 SDK to make an API request using the SDK's Amazon DynamoDB client.

This version requires a minimum version of `Go 1.15`.

## Getting started
To get started working with the conector setup your project for Go modules, and retrieve the `dynamodb-conector` dependencies with `go get`.

###### Initialize Project
```sh
$ mkdir hello-dynamodb-conector
$ cd hello-dynamodb-conector
$ go mod init hello-dynamodb-conector
```
###### Add dynamodb-conector Dependencies

```sh
$ go get github.com/elvenworks/dynamodb-conector
```

## Operations available for this version

What operations can this connector perform?
* Check if a table is accessible and has data
* Fetch the data of a specific Item

### Usage - GetCount 
Return: `int32 | error`  
Parameters: `'tableName string', 'limit int32'`

The example uses the given table name, and return number of items in a dynamoDB's table table noting the limit of readable records.
All parameters must be informed.

###### Write Code
In your preferred editor add the following content to main.go

```go
package main

import (
	"fmt"

	"github.com/elvenworks/dynamodb-conector"
)

func main() {

    config := dynamodb.InitConfig{
        AccessKeyID:     "",
        SecretAccessKey: "",
        Region:          "", 
    }
    // Using the Config value, create the DynamoDB client
    client := dynamodb.InitDynamodb(config)

    // Build the request with its parameters
    count, err := client.GetCount("table", 1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("total items: ", count)
    }

}
```
###### Compile and Execute
```sh
$ go run .
total items: 1
```

### Usage - GetItem
Return: `GetItemOutput | error`  
Parameters: `'tableName string', 'key string', 'item string'`

Returns a set of attributes for the item with the given key or an error. 
All parameters must be informed.

###### Write Code

Add dynamodb/attributevalue Dependencies

```sh
$ go get github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue
```

In your preferred editor add the following content to `main.go`
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/elvenworks/dynamodb-conector"
)

func main() {

    config := dynamodb.InitConfig{
        AccessKeyID:     "",
        SecretAccessKey: "",
        Region:          "", 
    }
    // Using the Config value, create the DynamoDB client
    client := dynamodb.InitDynamodb(config)

    // Build the request with its parameters
	outPut, err := client.GetItem("table", "key", "item")
	if err != nil {
		fmt.Println(err)
	} else {
		var attributes map[string]interface{}
		attributevalue.UnmarshalMap(outPut.Item, &attributes)
		jsonString, _ := json.Marshal(attributes)
		fmt.Println("Success: ", string(jsonString))
	}

}
```
###### Compile and Execute
```sh
$ go run .
Success: {"key":"item", ...}
```

## Resources

[SDK API Reference Documentation - GetItemOutput](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#GetItemOutput) - Use this document to look up about "GetItemOutput" output.