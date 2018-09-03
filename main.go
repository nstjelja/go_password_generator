package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type PostGeneratePasswordRequest struct {
	Length                 int
	NumOfSpecialCharacters int
	NumOfDigits            int
}

type PostGeneratePasswordResponse struct {
	Password string
	Message  string
}

type GetGeneratePasswordResponse struct {
	Message string
	Example PostGeneratePasswordRequest
}

func generatePassword(length int, numOfSpecCharacters int, numOfDigits int) (error, string) {

	if length < 6 {
		return fmt.Errorf("Minimal password length is 6, you passed %v", length), ""
	}

	if numOfSpecCharacters < 0 {
		return fmt.Errorf("numOfSpecCharacters can't be less then 0"), ""
	}

	if numOfDigits < 0 {
		return fmt.Errorf("numOfSpecCharacters can't be less then 0"), ""
	}

	if (numOfSpecCharacters + numOfDigits) > length {
		return fmt.Errorf("The sum of special characters and digits cannot be longher then the length of the password"), ""
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	var result []string

	numOfLetters := length - numOfDigits - numOfSpecCharacters

	for i := 0; i < numOfLetters; i++ {
		code := rune(97 + r.Intn(25))

		letter := string(code)

		if r.Intn(100) >= 50 {
			letter = strings.ToUpper(letter)
		}

		result = append(result, letter)
	}

	for i := 0; i < numOfDigits; i++ {
		result = append(result, strconv.Itoa(r.Intn(10)))
	}

	for i := 0; i < numOfSpecCharacters; i++ {
		letter := string(33 + r.Intn(10))
		result = append(result, letter)
	}

	randResult := make([]string, length)
	perm := rand.Perm(length)

	for i, v := range perm {
		randResult[i] = result[v]
	}

	return nil, strings.Join(randResult, "")
}

func GetGeneratePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	postGeneratePasswordRequest := PostGeneratePasswordRequest{
		Length:                 6,
		NumOfDigits:            0,
		NumOfSpecialCharacters: 0,
	}

	getGeneratePasswordResponse := GetGeneratePasswordResponse{
		Message: "Pass the request in the 'Example' format",
		Example: postGeneratePasswordRequest,
	}

	json.NewEncoder(w).Encode(getGeneratePasswordResponse)
}

func PostGeneratePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var postGeneratePasswordRequest PostGeneratePasswordRequest

	err := json.NewDecoder(r.Body).Decode(&postGeneratePasswordRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(PostGeneratePasswordResponse{Message: "Internal error", Password: ""})
		return
	}

	err, password := generatePassword(postGeneratePasswordRequest.Length, postGeneratePasswordRequest.NumOfSpecialCharacters, postGeneratePasswordRequest.NumOfDigits)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(PostGeneratePasswordResponse{Message: err.Error(), Password: ""})
		return
	}

	err = json.NewEncoder(w).Encode(PostGeneratePasswordResponse{Message: "OK", Password: password})
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", PostGeneratePassword).Methods("POST")
	router.HandleFunc("/", GetGeneratePassword).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
