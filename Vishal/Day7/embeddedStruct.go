
package main

func main()  {
	nested()
}
func nested()  {
	type contact struct{
		email ,address string
		phone int
	}

	type employee struct{
		name string
		age int
		contactInfo contact
	}

	john:=employee{
		name: "john",
		age: 22,
		contactInfo:contact{
			email:"abc@gmail.com",
            address:"banglore",
			phone:90282,
		},
	}
	fmt.Printf("%+v",john)
}
