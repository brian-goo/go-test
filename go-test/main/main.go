package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const s = "wake up"

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("test")
		break
	}
}

func condition() {
	if m := 7; m > 0 {
		fmt.Println(m, "is positive")
	} else {
		fmt.Println(m, "is negative")
	}
}

func switchTest() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its noon")
	default:
		fmt.Println("its not noon")
	}
}

func arr() {
	var a [4]int
	fmt.Println(a)
	a[3] = 4
	fmt.Println(a)
	fmt.Println(len(a))

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	c := [3]string{"a", "b", "c"}
	fmt.Println(c)

	var k [3]string
	fmt.Println(k)

	var d [3][3]float64
	fmt.Println(d)
}

func sliceT() {
	// 10 is len and capacity in this case
	a := make([]int, 10)
	fmt.Println(a)

	b := make([]int, 4, 10)
	fmt.Println(b)
	b = append(b, 5)
	fmt.Println(b)
	b = append(b, 1, 2)
	fmt.Println(b)

	// see p can take another value while keeping its type
	p := b[4:]
	fmt.Println(p)
	p = b[:4]
	fmt.Println(p)

	// looks like can init empty slice wo make
	n := []int{}
	fmt.Println(n, len(n))
	n = append(n, 1)
	fmt.Println(n)

	target := make([]int, len(n))
	copy(target, n)
	fmt.Println(target, n, &target, &n)
	n = append(n, 2)
	fmt.Println(target, n)

	test := []int{}
	fmt.Printf("%p\n", test)

	// ref := n
	// fmt.Println(ref, n)
	// n = append(n, 3)
	// fmt.Println(ref, n)
}

func passbywhat() {
	// looks liks array struct etc is passed by value
	// slice, map are passed by pointer
	i := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &i)

	// assigning like this looks like passed by value, different than python
	j := i
	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", &j)
	fmt.Println(i == j)
}

func spitTwo(i int) (x, y int) {
	x = i / 2
	y = i * 2
	return x, y
}

var (
	x int = y
	y int = 2
)

var t [10]int

func init() {
	for i := 0; i < 10; i++ {
		t[i] = i
	}
}

type T struct {
	Name string
	Px   int
}

func getNext() func() float64 {
	v := 1.0
	// b := []float64{-0.1, 0.1}
	return func() float64 {
		rand.Seed(time.Now().UnixNano())
		r := math.Max(rand.NormFloat64(), -0.1)
		t := math.Min(r, 0.1)
		v = v * (1 + t)
		// v = v * b[rand.Intn(len(b))]
		// if v < 0.2 {
		// 	v = 0.2
		// }
		return math.Round(v*100) / 100
	}
}

func getRandom() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.NormFloat64()
}

func getR(ch, quit chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	quit <- 1
}

func main() {
	// fmt.Println(x, y, t, T{})
	// fmt.Printf("%v\n", T{})

	// fmt.Println(time.Now().Format(time.RFC3339))
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Println(time.Now().Add(time.Second).Format("15:04:05"))

	// i := 1
	// fmt.Println(float64(i)*0.01 + 1)

	// t := new(T)
	// var v T
	// fmt.Println(t, v)
	// next := getNext()
	// fmt.Println(next())
	// fmt.Println(next())

	ch := make(chan int)
	quit := make(chan int)
	go getR(ch, quit)
	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
