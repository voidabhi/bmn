

package main

import (
	"testing"
)

func TestBmn(t *testing.T) {

	query := "flipkart.com"
	expectedCount := 23

	result,err := bmn(query)
	if err != nil {
		t.Error(err)
	}

	if expectedCount != len(result.Accounts) {
		t.Errorf("Accounts count doesn't match. Expected %s, got %s", expectedCount,
			len(result.Accounts))
	}
}


func BenchmarkGeocode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bmn("flipkart.com")
	}
}