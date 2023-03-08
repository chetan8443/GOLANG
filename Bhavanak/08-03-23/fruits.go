package main

import "fmt"

func main() {
	var s, t, a, b, m, n, temp, ac, oc int
	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&m, &n)

	for i := 0; i < m; i++ {
		fmt.Scan(&temp)
		if a+temp >= s && a+temp <= t {
			ac++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		if b+temp >= s && b+temp <= t {
			oc++
		}
	}
	fmt.Printf("%d\n%d", ac, oc)
}
