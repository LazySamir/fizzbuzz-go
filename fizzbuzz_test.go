package main

import "testing"

func TestNumber(t *testing.T) {
	if FizzBuzz(1) != "1" {
		t.Error("Failed test: did not return '1'")
	}
	if FizzBuzz(2) != "2" {
		t.Error("Failed test: did not return '2'")
	}
}

func TestFizz(t *testing.T) {

	if FizzBuzz(3) != "Fizz" {
		t.Error("Failed test: did not return 'Fizz'")
	}
	if FizzBuzz(6) != "Fizz" {
		t.Error("Failed test: did not return 'Fizz'")
	}
}

func TestBuzz(t *testing.T) {
	if FizzBuzz(5) != "Buzz" {
		t.Error("Failed test: did not return 'Buzz'")
	}
	if FizzBuzz(10) != "Buzz" {
		t.Error("Failed test: did not return 'Buzz'")
	}
}

func TestFizzBuzz(t *testing.T) {
	if FizzBuzz(15) != "FizzBuzz" {
		t.Error("Did not return 'FizzBuzz'")
	}
	if FizzBuzz(30) != "FizzBuzz" {
		t.Error("Did not return 'FizzBuzz'")
	}
}
