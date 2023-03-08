package main

type User struct {
	Name    string  `json: "name"`
	Age     int     `json: "age"`
	Address Address `json: "address"`
}

type Address struct {
	City    string `json: "city"`
	State   string `json: "state"`
	Country string `json: "country"`
}

func CheckNullValue(err error) {
	if err != nil {
		panic(err)
	}
}
