// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

package tutorial

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go-client/shared"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

var _ = shared.GoUnusedProtection__
const INT32CONSTANT = 9853
var MAPCONSTANT map[string]string

func init() {
MAPCONSTANT = map[string]string{
  "goodnight": "moon",
  "hello": "world",
}

}

