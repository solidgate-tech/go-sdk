[![Go project version](https://badge.fury.io/go/bitbucket.org%2Fsolidgate%2Fgo-sdk.svg)](https://badge.fury.io/go/bitbucket.org%2Fsolidgate%2Fgo-sdk)

# SolidGate API


This library provides basic API options of SolidGate payment gateway.

## Installation


```
$ go get bitbucket.org/solidgate/go-sdk
```

## Usage for h2h

```go
package main

func main() {
    //.....
    someRequestStruct = SomeRequestStruct{}
    someStructJson, err := json.Marshal(someRequestStruct)

    if err != nil {
        fmt.Print(err)
    }

    
    api := NewSolidGateApi("YourMerchantId", "YourPrivateKey", nil(for default) or "base url")
    
    response, err := api.Charge(someStructJson)
    
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

## Usage for payment form data

```go
package main

import (
    solidgate "bitbucket.org/solidgate/go-sdk"
    "github.com/AlekSi/pointer"
)

func main() {

    
    api := NewSolidGateApi("YourMerchantId", "YourPrivateKey")
    someRequestStruct = SomeRequestStruct{}
    someStructJson, err := json.Marshal(someRequestStruct)

    if err != nil {
        fmt.Print(err)
    }

    
    mercahntData, err := api.FormMerchantData(someStructJson)
    
    if err != nil {
        fmt.Print(err)
    }
    
 
    // ...
}
```
