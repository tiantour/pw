# pw
client for pullword with golang

### how to use

```
package main

import (
	"fmt"

	"github.com/tiantour/pw"
)

func main() {
	args := pw.Options{
		Threshold: 0, // 1开启0关闭
		Debug:     1, // 1开启0关闭
	}
	body, err := pw.Cmd(
		"post", // get or post
		"南京市长江大桥", // resorce
		args, // options
	)
	fmt.Println(string(body), err)
}
```
