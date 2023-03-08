package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
)

type Employee struct {
    ID	int    
    Name string 

}

var employees = []Employee{
    {Name: "kavita", ID: 100 },
    {Name: "Alka", ID: 101},
}

func main() {
    http.HandleFunc("/employee/", updateEmployeeByID)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func updateEmployeeByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/employee/"):])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var updatedEmployee Employee
    err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, employee := range employees {
        if employee.ID == id {
            employees[i] = updatedEmployee
            return
        }
    }

    http.NotFound(w, r)
}
