package controllers

import "testing"

func TestAddition(t *testing.T) {
	result := addition(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d instead", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := subtraction(5, 2)
	if result != 3 {
		t.Errorf("Expected 3, but got %d instead", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := multiplication(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d instead", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := division(6, 2)
	if err != nil {
		t.Error("Expected nil, but got an error")
	}
	if result != 3 {
		t.Errorf("Expected 3, but got %d instead", result)
	}

	_, err = division(6, 0)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
