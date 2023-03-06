package main

import "fmt"
//to implement an interface in go you need to implement all the 
//methods declared in the interface
type motorvehicle interface{
	mileage() float64
	averagespeed() string
}
type BMW struct{
	distance float64
	fuel float64
	averagespeed string
}
type audi struct{
	distance float64
	fuel float64
}
func (b BMW) mileage()float64{
	return b.distance/b.fuel

}
func (a audi) mileage()float64{
	return a.distance/a.fuel

}
func totalmileage(m []motorvehicle){
	tm:=0.0
	for _,v:=range m{
		tm=tm+v.mileage()
	}
	fmt.Printf("the total mileage per month",tm)
}
func main(){
	b1:=BMW{
		distance:167.9,
		fuel:36,
		averagespeed:"58",
	}
	a1:=audi{
		distance:152.9,
		fuel:30,
	}
    person:=[]motorvehicle{b1,a1}
	totalmileage(person)
}