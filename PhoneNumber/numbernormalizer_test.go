package main

import "testing"

func TestNumberNormalizer(t *testing.T) {

	testData := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892"}

	resultNormalizer := normalizer(testData)

	if resultNormalizer["1234567890"] != 2 {
		t.Error("1234567890 should have 2")
	}

	if resultNormalizer["1234567891"] != 1 {
		t.Error("1234567891 should have 1")
	}

	if resultNormalizer["1234567892"] != 3 {
		t.Error("1234567892 should have 3")
	}

	if resultNormalizer["1234567893"] != 1 {
		t.Error("1234567893 should have 1")
	}

	if resultNormalizer["1234567894"] != 1 {
		t.Error("1234567894 should have 1")
	}
}
