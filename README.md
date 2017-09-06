# Hex

## Purpose
 * Used to convert a hex string to a formatted ip address
 * Provides information regarding request such as:
  * Proper IP address
  * IP address type
  * Upon errors what specifically triggered

## Installation
With git installed simply run


```bash
go get github.com/calebice/hex
```
Following this, then be sure to include in imports
```go
import(
  "github.com/calebice/hex"
)
```
## Features
### Decode(request string) IpMessage
  * Takes in a requested address to be converted
  * Returns an IpMessage object (detailed below)

  ```go
  package main

  import(
    "github.com/calebice/hex"
  )

  func main(){}
    request:="AAAAAAAA"
    convert:=hex.Decode(request)
    IpAddress:=convert.GetIP()  //Return address
    IpType:=convert.GetAddrType() //Return type of address
    ErrorMsg:=convert.GetErrMsg() //Return the error msg
  }
  ```


#### IpMessage
  * **address (string)**
    * The converted formatted IP address
  * **addressType (string)**
    * The type of IP message of the request
  * **errorMessage (error)**
    * Informational error message if the address is invalid

### Using IpMessages
  * **GetIP() string**
    * Returns the address item
  * **GetAddrType() string**
    * Returns the address type
  * **GetErrMsg() error**
    * Returns error associated with conversion failure
