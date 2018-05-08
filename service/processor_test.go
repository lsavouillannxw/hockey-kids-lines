package service

import (
	"testing"
	"fmt"
	"time"
)

func TestProcess7103(t *testing.T) {
	score := NewProcessingHandler(7, 10, 3).Process().MaxScore
	if score != 190 {
		t.Errorf("score found is: %f while it should be 190", score)
	}
}

func TestProcess7163(t *testing.T) {
	score := NewProcessingHandler(7, 16, 3).Process().MaxScore
	if score != 440 {
		t.Errorf("score found is: %f while it should be 440", score)
	}
}

func TestProcess8103(t *testing.T) {
	score := NewProcessingHandler(8, 10, 3).Process().MaxScore
	if score != 180 {
		t.Errorf("score found is: %f while it should be 180", score)
	}
}

func TestProcess8163(t *testing.T) {
	score := NewProcessingHandler(8, 16, 3).Process().MaxScore
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

func TestProcess9164(t *testing.T) {
	score := NewProcessingHandler(9, 16, 4).Process().MaxScore
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

func TestProcess(t *testing.T) {
	for p := 7; p <= 16; p++ {
		for l := 5; l <= 16; l++ {
			for s := 3; s <= 5; s++ {
				if p % s == 0 {
					continue
				}
				c := make(chan string, 1)
				go func() {
					fmt.Printf("for %d players on %d lines of size %d, maxScore is: %f ", p, l, s, NewProcessingHandler(p, l, s).Process().MaxScore)
					c <- "done"
				}()
				select {
				case _ = <-c:
					fmt.Printf("for %d players on %d lines of size %d done", p, l, s)
				case <-time.After(time.Minute):
					fmt.Printf("for %d players on %d lines of size %d: TIMEOUT" , p, l, s)
				}
			}
		}
	}
}
