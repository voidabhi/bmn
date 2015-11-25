

package main

import (
	"testing"
)

func TestBmn(t *testing.T) {

	query := "flipkart.com"

	result,err := bmn(query)
	if err != nil {
		t.Error(err)
	}

	if len(result.Accounts) <0 {
		t.Errorf("Error: Unexpected Count!")
	}
}

//Benchmark Test
func BenchmarkGeocode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmn("flipkart.com")
	}
}
