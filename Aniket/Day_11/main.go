package main

type employee struct {
	name          string
	designation   string
	department    string
	isPermanant   string
	totalAmount   int
	taxableAmount int
}

type manager struct {
	name            string
	department      string
	totalAmount     int
	taxableAmount   int
	noOfEmployee    int
	projectDuration int
	totalDuration   int
}

func main() {
	v1 := employee{
		name:          "Sanket",
		designation:   "Software Engineer",
		department:    "IT",
		isPermanant:   "Yes",
		totalAmount:   100000,
		taxableAmount: 25000,
	}

	v2 := manager{
		name:            "Tushar",
		department:      "HR",
		noOfEmployee:    25,
		projectDuration: 100,
		totalDuration:   150,
	}

	getInfoD(v1, v1)
	getInfoT(v1, v2)
}
