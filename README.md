# Solidgate API

[![Go project version](https://badge.fury.io/go/github.com%2solidgate-tech%2go-sdk.svg)](https://badge.fury.io/go/github.com%2solidgate-tech%2go-sdk)

GO SDK provides API options for integrating Solidgate’s payment orchestrator into your Go applications.

Check our
* <a href="https://docs.solidgate.com/" target="_blank">Payment guide</a> to understand business value better
* <a href="https://api-docs.solidgate.com/" target="_blank">API Reference</a> to find more examples of usage

## Structure

<table style="width: 100%; background: transparent;">
  <colgroup>
    <col style="width: 50%;">
    <col style="width: 50%;">
  </colgroup>
  <tr>
    <th>SDK for GO contains</th>
    <th>Table of contents</th>
  </tr>
  <tr>
    <td>
      <code>api.go/solidgate/</code> – main file for API integration<br>
      <code>encryption.go</code> – library for encryption-related operations<br>
      <code>entries.go</code> – contains form-related methods (e.g., form resign)<br>
      <code>go.mod</code> – dependency file for managing module imports
    </td>
    <td>
      <a href="https://github.com/solidgate-tech/go-sdk?tab=readme-ov-file#requirements">Requirements</a><br>
      <a href="https://github.com/solidgate-tech/go-sdk?tab=readme-ov-file#installation">Installation</a><br>
      <a href="https://github.com/solidgate-tech/go-sdk?tab=readme-ov-file#usage">Usage</a><br>
      <a href="https://github.com/solidgate-tech/go-sdk?tab=readme-ov-file#errors">Errors</a>
    </td>
  </tr>
</table>

## Requirements

* **GO**: 1.13 or later
* **Solidgate account**: Merchant ID and secret key (request via <a href="mailto:sales@solidgate.com">sales@solidgate.com</a>)

<br>

## Installation

To install the Go SDK:

1. Ensure you have your merchant ID and secret key.
2. Run:
   ```bash
   go get github.com/solidgate-tech/go-sdk
   ```
3. Import the library into your Go application:
   ```go
   import solidgate "github.com/solidgate-tech/go-sdk"
   ```
4. Use test credentials to validate your integration before deploying to production.

<br>

## Usage

### Charge a payment

Returns a raw JSON response.

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

### Payment form

Returns a `FormInitDTO` struct in JSON.

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

### Update payment form

Returns a `FormUpdateDTO` struct in JSON.

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

### Resign payment form

Returns a `FormResignDTO` struct in JSON.

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

<br>

## Errors

Handle <a href="https://docs.solidgate.com/payments/payments-insights/error-codes/" target="_blank">errors</a>.

```js
if err != nil {
   fmt.Println(err)
}
```
