# surferua

![surferua](https://cloud.githubusercontent.com/assets/597902/16172506/9debc136-357a-11e6-90fb-c7c46f50dff0.png)

Logo from [uasurfer](https://github.com/avct/uasurfer).

Surfer User Agent (surferua) is a lightweight Golang package that generate HTTP User-Agent strings with particular attention to device type.

Fake User-Agent in golang, enjoy it.

## Usage

```golang
package main

import (
    "github.com/jiusanzhou/surferua"
)

fmt.Println(surferua.New().String())

// functions depends on your generated inputting data.

fmt.Println(surferua.New().Phone().String())
fmt.Println(surferua.New().Android().String())
fmt.Println(surferua.New().Desktop().Firefox().String())
```
