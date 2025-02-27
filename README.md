# Slip2Sure JS
SDK สำหรับเชื่อมต่อกับระบบ [Slip2Sure](https://slip2sure.com) Service. (Golang)

หากต้องการอ่าน API Document สามารถอ่านได้ที่
- [API Document (Truemoney)](https://app.slip2sure.com/user/api/docs/truemoney)
- [API Document (Bank slip)](https://app.slip2sure.com/user/api/docs/bankslip)
- [Error code](https://app.slip2sure.com/user/api/docs/errorcode)

## Install
```sh
go get github.com/Slip2Sure/slip2sure-go
```

## Example 
### Truemoney
```go
package main

import (
	"io"
	"log"
	"os"

	slip2surego "github.com/Slip2Sure/slip2sure-go"
)

func main() {
	client := slip2surego.Slip2SureAPI{
		ApiToken: "YOUR_API_KEY",
	}

	// Open file
	f, err := os.Open("FILE_PATH") // Or you can use Echo, Gin or Fiber to recieve file
	if err != nil {
		log.Panicln(err)
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read all bytes
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.ScanTruemoneySlip(fileBytes)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Response data: ")
	log.Println(result)
}

```

### Bank Slip
```go
package main

import (
	"log"

	slip2surego "github.com/Slip2Sure/slip2sure-go"
)

func main() {
	client := slip2surego.Slip2SureAPI{
		ApiToken: "YOUR_API_KEY",
	}

	result, err := client.ScanBankSlipByPayload("YOUR_QR_CODE")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Response data: ")
	log.Println(result)
}

```


## LICENSE
[MIT](./LICENSE)
