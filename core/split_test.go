package core

import "testing"

func TestCalculateSizeSum(t *testing.T) {
	sizeVector := []uint32{1, 2, 3, 4} // Sum = 10
	sum := calculateSizeSum(sizeVector)
	t.Logf("sum: %d\n", sum)

	sizeVector = []uint32{25, 13, 342, 92, 12, 456, 2} // Sum = 942
	sum = calculateSizeSum(sizeVector)
	t.Logf("sum: %d\n", sum)
}

func TestSplitBytes(t *testing.T) {
	bytes := []byte("1234")
	sizes := []uint32{1, 2, 2}

	splitBytes, err := SplitBytes(bytes, sizes)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(splitBytes)
}
