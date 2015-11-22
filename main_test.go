

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

	// if len(result.Accounts) <0 {
	// 	t.Errorf("Accounts count doesn't match. Expected %s, got %s", expectedCount,
	// 		len(result.Accounts))
	// }
}

//Benchmark Test
func BenchmarkGeocode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmn("flipkart.com")
	}
}
