package main

import (
	// "errors"
	"fmt"
	"time"
	// "log"
	// "net/http"
	// a "hello/arraysSlices"
	s "hello/switchback"
	// "os"

	// d "hello/dependencyInjection"
	// ints "hello/integers"
	// i "hello/iteration"

	// m "hello/mocking"
	// p "hello/pointersErrors"
	// s "hello/structsMethodsInterfaces"
)

// const spanish = "Spanish"
// const french = "French"
// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "

// // Hello returns a personalized greeting in a given langage
// func Hello(name, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	return greetingPrefix(language) + name
// }

// func greetingPrefix(language string) (prefix string) {
// 	switch language {
// 	case french:
// 		prefix = frenchHelloPrefix
// 	case spanish:
// 		prefix = spanishHelloPrefix
// 	default:
// 		prefix = englishHelloPrefix
// 	}
// 	return
// }

func main() {
	// fmt.Println(Hello("Shannon", "French"))
	// fmt.Println(ints.Add(2, 2))
	// fmt.Println(i.Repeat("a"))
	// fmt.Println(i.CharRepeat("a", 12))

	// numbers := []int{1, 2, 3, 4, 5}
	// moreNumbers := []int{3, 5, 7}
	// fmt.Println(a.Sum(numbers))
	// fmt.Println(a.SumAll(numbers, moreNumbers))
	// fmt.Println(a.SumAllTails(numbers, moreNumbers))
	// // fmt.Println(s.Perimeter(10.0, 10.0))


	// w := p.Wallet{}
	
	// fmt.Printf("Balance is %v \n", w.Balance())
	// w.Deposit(50)
	// fmt.Printf("Balance is %v \n", w.Balance())
	// w.Withdraw(60)
	// fmt.Printf("Balance is %v \n", w.Balance())
	
	// d.Greet(os.Stdout, "Elodie")
	//log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(d.MyGreeterHandler)))

	// sleeper := &m.ConfigurableSleeper{1 * time.Second, time.Sleep}
	// m.Countdown(os.Stdout, sleeper)

	sb := s.SwitchBack{T: time.Now()}
	fmt.Println(sb.EvenOrOdd())
	
}
