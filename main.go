package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main(){
	requestId := make(chan int)

	concurrency := 500

	for i := 1; i <= concurrency; i++ {
		go worker(requestId, i)
	}

	for i := 0; i < 20000; i++ {
		requestId <- i
	}

}

func worker(requestId chan int, worker int){
	for r := range requestId {
		res, err := http.Get("https://localhost")

		if err != nil {
			log.Fatal("Erro")
		}

		defer res.Body.Close()

		//content, _ := ioutil.ReadAll(res.Body)

		fmt.Printf("Worker %d, RequestId: %d\n", worker, r)
		r := rand.Intn(5)
		time.Sleep(time.Duration(r) * time.Second)
	}
}
