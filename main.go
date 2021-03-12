package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	times := time.Now()
	name := "タイムスタンプ"
	channel := "times_tanasho"

	time_text := times.String()

	time_text += "\n"

	input_text := ""
	fmt.Print("やったことを入力してね:")
	fmt.Scan(&input_text)
	time_text += input_text

	jsonStr := `{"channel":"` + channel + `", "username":"` + name + `", "text":"` + time_text + `"}`

	req, err := http.NewRequest(
		"POST",
		"XXXXXXX",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	defer resp.Body.Close()
}
