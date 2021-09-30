package shapes

import (
	"math"
	"testing"
)

// go test ./... -v
// ?       advanced-tests  [no test files]
// === RUN   TestArea
// === RUN   TestArea/Testing_Rectangle
// === RUN   TestArea/Testing_Circle
// --- PASS: TestArea (0.00s)
//     --- PASS: TestArea/Testing_Rectangle (0.00s)
//     --- PASS: TestArea/Testing_Circle (0.00s)
// PASS
// ok      advanced-tests/shapes   0.006s
func TestArea(t *testing.T) {
	t.Run("Testing Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12}
		expectedArea := float64(120)
		actualArea := rect.Area()

		if expectedArea != actualArea {
			// Fatal stops here, instead of Error which continues
			t.Fatalf("Actual area is [%f] which is different from expected area [%f]", actualArea, expectedArea)
		}
	})

	t.Run("Testing Circle", func(t *testing.T) {
		circ := Circle{10}
		expectedArea := float64(math.Pi * 10 * 10)
		actualArea := circ.Area()

		if expectedArea != actualArea {
			t.Fatalf("Actual area is [%f] which is different from expected area [%f]", actualArea, expectedArea)
		}
	})
}
