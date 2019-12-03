package main

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHaltCondition(c *C) {
	c.Assert(-1, Equals, computeIntCode([]int{99}[:],0))
}

func (s *MySuite) TestMovingForwardCode(c *C) {
	c.Assert(4, Equals, computeIntCode([]int{0,1,2,3}[:],0))
}

func (s *MySuite) TestSimpleAdd(c *C) {
	testArray := []int{1,1,2,2}
	computeIntCode(testArray[:],0)
  c.Assert(3, Equals, testArray[2])
}
