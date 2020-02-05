# SolidGate API


This library provides basic API options of SolidGate payment gateway.

## Installation


```
$ go get bitbucket.org/solidgate/go-sdk
```

## Usage

```go
package main

func main() {
    //.....
    someRequestStruct = SomeRequestStruct{}
    someStructJson := json.Marshal(someRequestStruct)
    
    api := NewSolidGateApi("YourMerchantId", "YourPrivateKey", nil(for default) or "base url")
    
    response, err := api.Charge(someStructJson)
    
    if err != nil {
        fmt.Print(err) // handle error
    }
    
    someResponeStruct = SomeResponeStruct{}
    err := json.Unmarshal(response, &someResponeStruct)
    
    if err != nil {
        fmt.Print(err) // handle error
    }
    //.....
}
```
