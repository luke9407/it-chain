package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


type testObject struct{
	name string
}

func (t testObject) Serialize() string{
	return t.name
}

func TestHasing(t *testing.T){

	t1 := testObject{name:"asd"}
	t2 := testObject{name:"asd"}
	t3 := testObject{name:"asd2"}

	assert.Equal(t,t1,t2)
	assert.NotEqual(t,t1,t3)
}
