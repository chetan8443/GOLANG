// Write a program for bank details :
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Account struct { // making a struct for individual account
	AccNo    int    `json:"AccNo"`
	Name     string `json:"Name"`
	Balance  int    `json:"Balance"`
	Password string `json:"Password"`
}

type Trax struct { // struct made for collectting transaction data
	Acc      int    `json:"Acc"`
	Password string `json:"Password"`
	Amount   int    `json:"Amount"`
}

var accounts []Account  // all accounts will be stored here
var number int = 230300 // account number will start from this

func main() {
	//calling different functions using different api
	http.HandleFunc("/check_bal", checkBal)
	http.HandleFunc("/add_acc", addAcc)
	http.HandleFunc("/withdraw", withdrawal)
	http.HandleFunc("/deposit", deposit)
	//listening to the port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func checkBal(w http.ResponseWriter, r *http.Request) {
	//checking balance for given account number
	acc, err := strconv.Atoi(r.URL.Query().Get("accNo")) // to get parameter from url and convert into int
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	for _, v := range accounts {
		if acc == v.AccNo {
			bal := map[string]int{"Total balance": v.Balance}
			name := map[string]string{"name": v.Name}
			json.NewEncoder(w).Encode(name)
			json.NewEncoder(w).Encode(bal)
			return
		}
	}
	resp, _ := json.Marshal(map[string]string{"Messege": "Account not found !"})
	w.Write(resp)
}

func addAcc(w http.ResponseWriter, r *http.Request) {
	//adding an account taking arguments from post method
	if r.Method != "POST" {
		http.Error(w, "Use POST method !!", http.StatusBadRequest)
	}
	var acc Account
	a, err := ioutil.ReadAll(r.Body) // reading from post method arguments
	json.Unmarshal(a, &acc)          // converting json data into go format
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	number++
	acc.AccNo = number
	accounts = append(accounts, acc)
	//converting go format data into json format
	resp, _ := json.Marshal(map[string]string{"Messege": "Account added successfully !", "Account Number": strconv.Itoa(number)})
	w.Write(resp)
}

func withdrawal(w http.ResponseWriter, r *http.Request) {
	// withdraw money from given account
	if r.Method != "POST" { // checking for post method
		http.Error(w, "Use POST method !!", http.StatusBadRequest)
		return
	}
	var t Trax
	a, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(a, &t)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	for i, v := range accounts {
		if t.Acc == v.AccNo {
			if t.Password == v.Password {
				accounts[i].Balance -= t.Amount
				resp, _ := json.Marshal(map[string]string{"Messege": "Money has been withdrawn successfully !"})
				w.Write(resp)
				return
			} else {
				resp, _ := json.Marshal(map[string]string{"Messege": "Incorrect password !"})
				w.Write(resp)
				return
			}
		}
	}
	resp, _ := json.Marshal(map[string]string{"Messege": "Invalid account number!"})
	w.Write(resp)
}

func deposit(w http.ResponseWriter, r *http.Request) {
	// to deposite money
	if r.Method != "POST" {
		http.Error(w, "Use POST method !!", http.StatusBadRequest)
		return
	}
	var t Trax
	a, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(a, &t)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	for i, v := range accounts {
		if t.Acc == v.AccNo {
			if t.Password == v.Password {
				accounts[i].Balance += t.Amount
				resp, _ := json.Marshal(map[string]string{"Messege": "Money has been deposited successfully !"})
				w.Write(resp)
				return
			} else {
				resp, _ := json.Marshal(map[string]string{"Messege": "Incorrect password !"})
				w.Write(resp)
				return
			}
		}
	}
	resp, _ := json.Marshal(map[string]string{"Messege": "Invalid account number!"})
	w.Write(resp)
}
