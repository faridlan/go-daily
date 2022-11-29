package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Address struct {
	Street     string `json:"street,omitempty"`
	No         int    `json:"no,omitempty"`
	PostalCode int    `json:"postal_code,omitempty"`
	City       string `json:"city,omitempty"`
}

type User struct {
	UserName string    `json:"user_name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Address  []Address `json:"address,omitempty"`
}

func GetUser() User {

	return User{
		UserName: "nulhakim",
		Email:    "nul@mail.com",
		Password: "123abc",
		Address: []Address{
			{
				Street:     "Jl Mitra Batik",
				No:         20,
				PostalCode: 46717,
				City:       "Tasikmalaya",
			},
			{
				Street:     "Jl Leuwidahu",
				No:         69,
				PostalCode: 46151,
				City:       "Tasikmalaya",
			},
		},
	}

}

func CreateJson(x any) {

	writer, _ := os.Create("sample.json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(&x)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success create json file")
}

func TestReadJson(t *testing.T) {

	user := User{}

	example, err := os.Open("example.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(example)
	err = decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func TestJson(t *testing.T) {
	user := GetUser()
	CreateJson(user)
}

func TestDaily(t *testing.T) {
	user := GetUser()

	fmt.Println(user)
	for _, v := range user.Address {
		fmt.Println(v)
	}
}

func BizzBuzz(number int) {

	for i := 0; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Bizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Bizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}

func TestFactorial(t *testing.T) {
	val := 1

	for i := 10; i > 0; i-- {
		val *= i
	}

	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	assert.Equal(t, val, result)

}
