package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	//"sync"
	"time"

	//"github.com/google/uuid"
)

var wg sync.WaitGroup

func main() {
	t1 := time.Now()
	//key := uuid.New().String()
	key := "key2"
	fmt.Println("key:", key)
	for i := 0; i < 35; i++ {
		wg.Add(1)
		//time.Sleep(10*time.Millisecond)
		go func() {
			defer wg.Done()
			response, err := http.Post("http://127.0.0.1:8090/v1/check_ratelimit/"+key, "application/json", nil)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				bodyBytes, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)
				if bodyString != "" {
					fmt.Println(bodyString)
				}
				_ = response.Body.Close()
			}
		}()
	}
	wg.Wait()
	/*
			response, err := http.Post("http://127.0.0.1:8090/v1/check_ratelimit/"+key, "application/json", nil)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				bodyBytes, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)
				fmt.Println(bodyString)
				_ = response.Body.Close()
			}
			fmt.Println("total time taken:", time.Since(t1))
		}

	*/
	fmt.Println("total time taken:", time.Since(t1))
}
