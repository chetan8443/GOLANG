package main

import (
    "database/sql"
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Set up the database connection
    db, err := sql.Open("sqlite3", "../db/mydatabase.db")
    if err != nil {
        log.Fatal("DB error:", err)
    }
    defer db.Close()

    // Create the users table if it does not exist
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
        studentid INTEGER PRIMARY KEY,
        studentname TEXT,
        studentmarks INTEGER
    )`)
    if err != nil {
        log.Fatal(err)
    }

    // Open the CSV file
    file, err := os.Open("StudentMarks.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Parse the CSV file
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Loop through the CSV records and insert them into the database
    for _, record := range records {
        // Convert the values to the appropriate types
        studentid, err := strconv.Atoi(record[0])
        if err != nil {
            log.Fatal(err)
        }

        studentname := record[1]
        
        studentmarks, err := strconv.Atoi(record[2])
        if err != nil {
            log.Fatal(err)
        }

        // Insert the values into the database
        _, err = db.Exec("INSERT OR IGNORE INTO students (studentid, studentname, studentmarks) VALUES (?, ?, ?)", studentid, studentname, studentmarks)
        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println("Data inserted successfully!")
}
