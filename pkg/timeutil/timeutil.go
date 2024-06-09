package timeutil

import "time"

var (
	JST *time.Location
)

func init() {
	if jst, err := time.LoadLocation("Asia/Tokyo"); err == nil {
		JST = jst
	} else {
		JST = time.FixedZone("JST", 9*60*60)
	}
}

func NowInJST() time.Time {
	return time.Now().In(JST)
}
