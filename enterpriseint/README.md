# Enterprise Int

Super useful integer type. Try it out.

```
package main

import (
	"fmt"

	ei "github.com/porty/playpit/enterpriseint"
)

func main() {
	ei.EnterpriseInt(12).ForEach(func(i ei.EnterpriseInt) error {
		fmt.Printf("%d is prime? %t\n", i, i.IsPrime())
		return nil
	})

	b := make([]byte, 10)
	ei.EnterpriseInt(3).Read(b)
	fmt.Printf("%#v\n", b)
}
```
