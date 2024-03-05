[![Go project version](https://badge.fury.io/go/github.com%2solidgate-tech%2go-sdk.svg)](https://badge.fury.io/go/github.com%2solidgate-tech%2go-sdk)

# SolidGate API

This library provides basic API options of SolidGate payment gateway.

## Installation

```
$ go get github.com/solidgate-tech/go-sdk
```

## Usage for h2h

```go
package main

import (
	"encoding/json"
	"fmt"

	solidgate "github.com/solidgate-tech/go-sdk"
)

func main() {
	//.....
	someRequestStruct := SomeRequestStruct{}
	someStructJson, err := json.Marshal(someRequestStruct)

	if err != nil {
		fmt.Print(err)
	}

	solidgateSdk := solidgate.NewSolidGateApi("YourPublicKey", "YourSecretKey")

	response, err := solidgateSdk.Charge(someStructJson)

	if err != nil {
		fmt.Print(err)
	}

	someResponeStruct = SomeResponeStruct{}
	err := json.Unmarshal(response, &someResponeStruct)

	if err != nil {
		fmt.Print(err)
	}
	//.....
}

```

## Usage for init payment form data

```go
package main

import (
	"encoding/json"
	"fmt"

	solidgate "github.com/solidgate-tech/go-sdk"
)

func main() {

	solidgateSdk := solidgate.NewSolidGateApi("YourPublicKey", "YourSecretKey")
	someRequestStruct := SomeRequestStruct{}
	someStructJson, err := json.Marshal(someRequestStruct)

	if err != nil {
		fmt.Print(err)
	}

	formInitDto, err := solidgateSdk.FormMerchantData(someStructJson)

	if err != nil {
		fmt.Print(err)
	}

	// ...
}

```

## Usage for update form data

```go
package main

import (
	"encoding/json"
	"fmt"

	solidgate "github.com/solidgate-tech/go-sdk"
)

type UpdateParams struct {
	...
}

func main() {
	solidgateSdk := solidgate.NewSolidGateApi("YourPublicKey", "YourSecretKey")
	someRequestStruct := UpdateParams{}
	someStructJson, err := json.Marshal(someRequestStruct)

	if err != nil {
		fmt.Print(err)
	}

	formUpdateDto, err := solidgateSdk.FormUpdate(someStructJson)

	if err != nil {
		fmt.Print(err)
	}

	// ...
}

```

## Usage for resign form data

```go
package main

import (
	"encoding/json"
	"fmt"

	solidgate "github.com/solidgate-tech/go-sdk"
)

func main() {

	solidgateSdk := solidgate.NewSolidGateApi("YourPublicKey", "YourSecretKey")
	someRequestStruct := SomeRequestStruct{}
	someStructJson, err := json.Marshal(someRequestStruct)

	if err != nil {
		fmt.Print(err)
	}

	formResignDto, err := solidgateSdk.FormResign(someStructJson)

	if err != nil {
		fmt.Print(err)
	}

	// ...
}

```