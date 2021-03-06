# goxirr
[![GoDoc](https://godoc.org/github.com/maksim77/goxirr?status.svg)](https://godoc.org/github.com/maksim77/goxirr)

Goxirr is a simple implementation of a function for calculating the Internal Rate of Return for irregular cash flow (XIRR).

## Links
- [Wikipedia](https://en.wikipedia.org/wiki/Internal_rate_of_return)
- [Excel support](https://support.office.com/en-us/article/XIRR-function-DE1242EC-6477-445B-B11B-A303AD9ADC9D)

## Example

```go
import "time"

import "github.com/maksim77/goxirr"

firstDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
t1 := goxirr.Transaction{
	Date: firstDate,
	Cash: -100,
}
t2 := goxirr.Transaction{
    Date: firstDate.Add(time.Hour * 24 * 365),
    Cash: 112,
}

tas := goxirr.Transactions{t1, t2}
fmt.Println(goxirr.Xirr(tas))
```