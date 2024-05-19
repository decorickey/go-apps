package main

import (
	"fmt"
	"time"

	"github.com/decorickey/go-apps/lib/timeutil"
)

func main() {
	now := timeutil.NowInJST()
	fmt.Println(now)
	fmt.Println(now.UTC())

	month := time.Month(1)
	fmt.Println(month.String())
}
