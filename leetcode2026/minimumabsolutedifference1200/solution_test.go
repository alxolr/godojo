package minimumabsolutedifference1200

import (
	"reflect"
	"testing"
)

func TestMinimumAbsDifference(t *testing.T) {
	want := [][]int{{1, 2}, {2, 3}, {3, 4}}
	got := minimumAbsDifference([]int{4, 2, 1, 3})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
