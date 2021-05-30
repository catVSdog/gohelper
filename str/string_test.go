package str

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type StringSuit struct {}

var _ = Suite(&StringSuit{})


// IsUpperCheckInfo 实现自定义checker
type IsUpperCheckInfo struct {
	*CheckerInfo
}

func (u *IsUpperCheckInfo) Check(params []interface{}, names []string) (result bool, error string){
	testString := params[0]
	str ,ok := testString.(string)
	if !ok {
		return false, fmt.Sprintf("%s not a string", testString)
	}
	if IsUpper(str){
		return true, ""
	}
	return false, fmt.Sprintf("%s not a Upper string", testString)
}

var isUpperChecker = &IsUpperCheckInfo{
	&CheckerInfo{Name: "isUpper", Params: []string{"value"}},
}

func (s *StringSuit)TestIsUpperWhenPass(c *C) {
	c.Assert("A",isUpperChecker)
}

func (s *StringSuit)TestIsUpperWhenNotPass(c *C) {
	c.Assert("a", Not(isUpperChecker))
}

func (s *StringSuit)TestIsUpperWhenNotString(c *C) {
	c.Assert(1, Not(isUpperChecker))
}

// IsLowerCheckInfo 实现自定义checker
type IsLowerCheckInfo struct {
	*CheckerInfo
}

func (u *IsLowerCheckInfo) Check(params []interface{}, names []string) (result bool, error string){
	testString := params[0]
	str, ok := testString.(string)
	if !ok {
		return false, fmt.Sprintf("%s not a string", testString)
	}
	if IsLower(str){
		return true, ""
	}
	return false, fmt.Sprintf("%s not a Lower string", testString)
}

var isLowerChecker = &IsLowerCheckInfo{
	&CheckerInfo{Name: "isLower", Params: []string{"value"}},
}

func (s *StringSuit)TestIsLowerWhenPass(c *C) {
	c.Assert("a",isLowerChecker)
}

func (s *StringSuit)TestIsLowerWhenNotPass(c *C) {
	c.Assert("A", Not(isLowerChecker))
}

func (s *StringSuit)TestIsLowerWhenNotString(c *C) {
	c.Assert(true, Not(isLowerChecker))
}