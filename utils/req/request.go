package req

import (
	"advent-of-go/secrets"
	"fmt"
	"io"
	"net/http"
)

// MakeRequest Makes a request to the AoC API and returns the response
func MakeRequest(day int, year int) string {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: secrets.Session})

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
