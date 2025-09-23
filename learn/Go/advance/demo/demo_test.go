package demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_GetItemsPtr(t *testing.T) {
	a := assert.New(t)

	jobs := &MyJobList{}
	listPtr, err := GetItemsPtr(jobs)

	a.Nil(err)

}
