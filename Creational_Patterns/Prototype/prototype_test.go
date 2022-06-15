package prototype

import "testing"

func TestClone(t *testing.T) {
	robot := NewRobot("android01", "clear water")
	if robot.name != "android01" {
		t.Fatal("Expected name android01 but got instead ", robot.name)
	}
	clone := robot.clone()
	if clone.name != "android01_clone" {
		t.Fatal("Expected name android01_clone but got instead ", robot.name)
	}
}
