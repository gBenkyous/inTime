package keisan

import (
	_"intimeServer/pkg/testPkg"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
  )
   
  func TestTasizan(t *testing.T) {
	var i int = 1
	i = Tasizan(2, 1)
	fmt.Println(i)
	assert.Equal(t, 3, i)
  }

  func TestHikizan(t *testing.T) {
	var i int = 1
	i = Hikizan(2, 1)
	fmt.Println(i)
	assert.Equal(t, i, 1)
  }

  func TestKakezan(t *testing.T) {
	var i int = 1
	i = Kakezan(2, 1)
	fmt.Println(i)
	assert.Equal(t, i, 2)
  }