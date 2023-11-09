package it

import (
	"github.com/yangbiny/reason-commons/common"
	"testing"
)

func TestStrEmpty(test *testing.T) {
	str := "   	  "
	blankStr := common.IsBlankStr(&str)
	println(blankStr)
}
