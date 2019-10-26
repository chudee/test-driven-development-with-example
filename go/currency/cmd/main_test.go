package cmd

import (
	d "github.com/chudee/currency-go-example/go/currency/pkg/dollar"
	f "github.com/chudee/currency-go-example/go/currency/pkg/franc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := d.NewDollar(5)
	assert.Equal(t, d.NewDollar(10), five.Times(2))
	assert.Equal(t, d.NewDollar(15), five.Times(3))
}

func TestFrancMultiplication(t *testing.T) {
	five := f.NewFranc(5)
	assert.Equal(t, f.NewFranc(10), five.Times(2))
	assert.Equal(t, f.NewFranc(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	five := d.NewDollar(5)
	assert.True(t, five.Equals(d.NewDollar(5)))
	assert.False(t, five.Equals(d.NewDollar(6)))
}
