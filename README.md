# surferua

![surferua](https://cloud.githubusercontent.com/assets/597902/16172506/9debc136-357a-11e6-90fb-c7c46f50dff0.png)

Logo from [uasurfer](https://github.com/avct/uasurfer).

Surfer User Agent (surferua) is a lightweight Golang package that generate HTTP User-Agent strings with particular attention to device type.

Fake User-Agent in golang, enjoy it.

## Usage

1. Modify the meta data and generate DB code if necessary

```
go run cmd/gen-db.go data.yaml
```

Re-generate 3 go files: `botDB.go`, `browserDB.go` and `platformDB.go`

After this you will get the code with special function what you defined in your input data.

2. Import to your project and use it

```golang
package main

import (
    "github.com/jiusanzhou/surferua"
)

fmt.Println(surferua.New().String())

// functions depends on your generated inputting data.

fmt.Println(surferua.NewBot.String())
fmt.Println(surferua.NewBotGoogle.String())
fmt.Println(surferua.New().Phone().String())
fmt.Println(surferua.New().Android().String())
fmt.Println(surferua.New().Desktop().Firefox().String())
```
