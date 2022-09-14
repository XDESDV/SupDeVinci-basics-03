package main

import (
	"reflect"
	"testing"
)

func TestSpirale(t *testing.T) {
	var output3 = [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}

	if spirale := createSpiral(3); !reflect.DeepEqual(spirale, output3) {
		t.Errorf("Spirale N=3 incorrecte, obtenu: %d, attendu: %d.", spirale, output3)
	}

	var output4 = [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}

	if spirale := createSpiral(4); !reflect.DeepEqual(spirale, output4) {
		t.Errorf("Spirale N=4 incorrecte, obtenu: %d, attendu: %d.", spirale, output4)
	}

	var output5 = [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}}

	if spirale := createSpiral(5); !reflect.DeepEqual(spirale, output5) {
		t.Errorf("Spirale N=5 incorrecte, obtenu: %d, attendu: %d.", spirale, output5)
	}

}
