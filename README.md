 dynamodb-conector

Um simples client conector para AWS DynamoDB para realizar operações básicas em uma tabela.

## Operações disponíveis para essa versão

Quais operações esse projeto pode realizar?
* Buscar os dados de um Item (registro) especifico
* Verificar se uma tabela esta acessível e possui dados (Itens)

## Informações de uso

Descrição de uso das operações 

#### GetItem
Retorno: `GetItemOutput | error`  
Parâmetros: `'tableName string', 'key string', 'item string'`

Retorna um conjunto de atributos para do item com a chave informada ou um erro. Todos os parâmetros devem ser informados.

Exemplo:
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
    client := dynamodb.InitDynamodb(config)

	outPut, err := client.GetItem("table", "key", "item")
	if err != nil {
		fmt.Println(err)
	} else {
		var attributes map[string]interface{}
		attributevalue.UnmarshalMap(outPut.Item, &attributes)
		jsonString, _ := json.Marshal(attributes)
		fmt.Println(string(jsonString))
	}

}
```

> Para mais informações sobre "GetItemOutput" consulte a documentação oficial do SKD AWS:
https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#GetItemOutput

#### GetCount 
Retorno: `int32 | error`  
Parâmetros: `'tableName string', 'limit int32'`

Retorna o número de itens da tabela informada observando o limit de registos que possam ser lidos.

Exemplo de código:
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
    client := dynamodb.InitDynamodb(config)

    count, err := client.GetCount("table", 1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("total items ", count)
    }

}
```