package service

import "testing"

func TestProcess7103(t *testing.T) {
	score := NewProcessingHandler().Process(7, 10, 3).MaxScore
	if score != 370 {
		t.Errorf("score found is: %f while it should be 370", score)
	}
}

func TestProcess8103(t *testing.T) {
	score := NewProcessingHandler().Process(8, 10, 3).MaxScore
	if score != 180 {
		t.Errorf("score found is: %f while it should be 180", score)
	}
}

func TestProcess9104(t *testing.T) {
	score := NewProcessingHandler().Process(9, 10, 4).MaxScore
	if score != 370 {
		t.Errorf("score found is: %f while it should be 370", score)
	}
}

func TestProcess10104(t *testing.T) {
	score := NewProcessingHandler().Process(10, 10, 4).MaxScore
	if score != 180 {
		t.Errorf("score found is: %f while it should be 180", score)
	}
}
