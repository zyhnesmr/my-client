package resolvers

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMyBuilder_Build(t *testing.T) {

	var s = "subs-rpc-svc:9091"

	split := strings.Split(s, ":")

	i, err := strconv.ParseInt(split[1], 10, 64)
	fmt.Println(i)
	fmt.Println(err)
}
