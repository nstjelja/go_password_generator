package main

import (
	"testing"
	"unicode"
)

func TestGeneratePasswordReturnsPassword(t *testing.T) {
	length := 10
	err, result := generatePassword(length, 0, 0)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(result) < 10 {
		t.Errorf("Password %v has length of %v when it should at least have a length of %v", result, len(result), length)
		t.Fail()
	}
}

func TestGeneratePasswordNoSpecials(t *testing.T) {
	err, result := generatePassword(6, 0, 0)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	for _, letter := range result {
		if !unicode.IsLetter(letter) && !unicode.IsNumber(letter) {
			t.Errorf("%#U is found in the passoword %v while no specials are requested", letter, result)
			break
		}
	}

}

func TestGeneratePasswordNoNumbers(t *testing.T) {
	err, result := generatePassword(6, 0, 0)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	for _, letter := range result {
		if unicode.IsNumber(letter) {
			t.Errorf("%#U is found in the passoword %v while no numbers are requested", letter, result)
			break
		}
	}
}

func TestGeneratePasswordMinimalLegthLessThenRequiredMinimal(t *testing.T) {
	err, _ := generatePassword(3, 0, 0)

	if err != nil {
		return
	}

	t.Error("Error should have been returned")
}

func TestGeneratePasswordNumOfNumbersLessThenZero(t *testing.T) {
	err, _ := generatePassword(6, 0, -1)

	if err == nil {
		t.Errorf("An error was supposed to be thrown if the numOfDigits was -1")
	}
}

func TestGeneratePasswordNumOfSpecialCharactersLessThenZero(t *testing.T) {
	err, _ := generatePassword(6, -1, 0)

	if err == nil {
		t.Errorf("An error was supposed to be thrown if the numOfSpecCharacters was -1")
	}
}

func TestGeneratePasswordLengthCannotExceedLength(t *testing.T) {
	err, _ := generatePassword(6, 6, 0)

	if err != nil {
		t.Error("generatePassword(6, 6, 0) should not have generated an error")
	}

	err, _ = generatePassword(6, 0, 6)

	if err != nil {
		t.Error("generatePassword(6, 0, 6) should not have generated an error")
	}

	err, _ = generatePassword(6, 4, 2)

	if err != nil {
		t.Error("generatePassword(6, 4, 2) should not have generated an error")
	}

	err, _ = generatePassword(6, 0, 7)

	if err == nil {
		t.Error("generatePassword(6, 0, 7) should  have generated an error")
	}

	err, _ = generatePassword(6, 7, 0)

	if err == nil {
		t.Error("generatePassword(6, 7, 0) should  have generated an error")
	}

	err, _ = generatePassword(6, 6, 6)

	if err == nil {
		t.Error("generatePassword(6, 6, 6) should  have generated an error")
	}

}

func TestGeneratePasswordHasSpecialCharacters(t *testing.T) {
	err, result := generatePassword(6, 4, 0)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	count := 0

	for _, letter := range result {
		if !unicode.IsLetter(letter) && !unicode.IsNumber(letter) {
			count++
		}
	}

	if count != 4 {
		t.Errorf("There should have been 4 special characters %v provided", count)
	}
}

func TestGeneratePasswordHasDigits(t *testing.T) {
	err, result := generatePassword(6, 0, 4)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	count := 0

	for _, letter := range result {
		if unicode.IsNumber(letter) {
			count++
		}
	}

	if count != 4 {
		t.Errorf("There should have been 4 digits %v provided in %v", count, result)
	}
}
