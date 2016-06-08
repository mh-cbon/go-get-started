package main

import (
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/mh-cbon/a/sub"
	"github.com/mh-cbon/b"
)

func main() {
	lib.Hello()
	b.Hello()
	c, _ := semver.NewConstraint("<= 1.2.3, >= 1.4")
	fmt.Println(c)
}
