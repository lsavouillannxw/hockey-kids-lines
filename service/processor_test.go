package service

import "testing"

func TestProcess7103(t *testing.T) {
	score := NewProcessingHandler(7, 10, 3).Process().MaxScore
	if score != 190 {
		t.Errorf("score found is: %f while it should be 190", score)
	}
}

func TestProcess8103(t *testing.T) {
	score := NewProcessingHandler(8, 10, 3).Process().MaxScore
	if score != 180 {
		t.Errorf("score found is: %f while it should be 180", score)
	}
}

func TestProcess9104(t *testing.T) {
	score := NewProcessingHandler(9, 10, 4).Process().MaxScore
	if score != 220 {
		t.Errorf("score found is: %f while it should be 220", score)
	}
}

func TestProcess10104(t *testing.T) {
	score := NewProcessingHandler(10, 10, 4).Process().MaxScore
	if score != 180 {
		t.Errorf("score found is: %f while it should be 180", score)
	}
}
