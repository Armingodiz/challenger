package RandomData

import "testing"

func TestGetData(t *testing.T) {
	bData, cData := GetData()
	for i := 0; i < 200; i++ {
		if bData[i].Ip != cData[i].Ip {
			t.Errorf("Error in setting ip for data in index %d", i)
		}
	}
}
