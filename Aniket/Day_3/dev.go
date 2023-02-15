package main

import (
	"errors"
	"math"
)

// Code to get area of different entities uisng Interface

type Shape interface {
	area() float32
}

type Circle struct {
	r float32
}

type Rectangle struct {
	w, h float32
}

func (c Circle) area() float32 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float32 {
	return r.h * r.w
}

func getArea(s Shape) float32 {
	return s.area()
}

// "Error Handling" in GoLang

type error interface {
	error() string
}

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
