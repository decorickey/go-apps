package timeutil_test

import (
	"testing"

	"github.com/decorickey/go-apps/lib/timeutil"
)

func TestJST(t *testing.T) {
	jst := timeutil.JST
	if s := jst.String(); s != "Asia/Tokyo" {
		t.Fatalf("expected 'Asia/Tokyo', but actual %s", s)
	}
}
