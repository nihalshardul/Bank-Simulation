package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	doc1 := os.Args[1]
	doc2 := os.Args[2]
	doc3 := os.Args[3]

	s1 := strings.Split(doc1, "=")[1]

	s2 := strings.Split(doc2, "=")[1]

	s3 := strings.Split(doc3, "=")[1]

	num_cashier, _ := strconv.Atoi(s1)
	num_customer, _ := strconv.Atoi(s2)
	time_per_customer, _ := strconv.Atoi(s3)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "--> Bank Simulation Started")

	if num_cashier == 0 {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "--> No Cashier is present at the moment")
	} else {
		cust := make(chan int, num_customer)
		results := make(chan int, num_customer)

		for counter := 1; counter <= num_cashier; counter++ {
			go bank_simulation(counter, cust, results, time_per_customer)
		}

		for cc := 1; cc <= num_customer; cc++ {
			cust <- cc
		}
		close(cust)

		for a := 1; a <= num_customer; a++ {
			<-results
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "--> Bank Simulation Ended")
}

func bank_simulation(counter int, cust <-chan int, results chan<- int, ct int) {
	for num := range cust {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "--> Cashier "+strconv.Itoa(counter)+": Customer "+strconv.Itoa(num)+" Started")
		time.Sleep(time.Duration(ct) * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "--> Cashier "+strconv.Itoa(counter)+": Customer "+strconv.Itoa(num)+" Completed")
		results <- 1
	}
}
