package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c := http.Client{}
		resp, err := c.Get("https://brasilapi.com.br/api/cep/v1/01153000/?01310-100")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		c1 <- string(body)
	}()

	go func() {
		c := http.Client{}
		resp, err := c.Get("http://viacep.com.br/ws/01310-100/json")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		c2 <- string(body)
	}()

	select {
	case body := <-c1:
		fmt.Printf("received from BrazilApi: %s", body)
	case body := <-c2:
		fmt.Printf("received from viaCep: %s", body)
	case <-time.After(time.Second):
		println("timeout")
	}
}
