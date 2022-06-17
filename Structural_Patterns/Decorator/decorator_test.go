package decorator

import "testing"

func TestTesla(t *testing.T) {
	model3 := model3{price: 40000}
	model3LongRange := longRange{tesla: &model3}
	model3LongRangeAI := betterAI{tesla: &model3LongRange}

	if model3LongRangeAI.getPrice() != 62000 {
		t.Fatal("The price of the model is wrong, expected 62000 but received ", model3LongRangeAI.getPrice())
	}
}
