package main

import (

	"backend/driver"
	eHandler "backend/handler/httpHandler"
	"backend/middleware"
	// "backend/migrate"

	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// type Customer struct {
// 	ID   int       `json:"c_id"`
// 	Name string    `json:"c_name"`
// 	DOB  time.Time `json:"c_dob"`
// }
// type Account struct {
// 	AccNo       int           `json:"acc_no"`
// 	Holder      Customer      `json:"acc_holder"`
// 	Transaction []Transaction `json:"acc_transaction"`
// }
// type Transaction struct {
// 	TID    int       `json:"tr_id"`
// 	TTime  time.Time `json:"tr_time"`
// 	Amount int       `json:"tr_amount"`
// 	Cridit bool      `json:"tr_cridit"`
// }
// type Employee1 struct {
// 	EmpID       string `json:"emp_Id"`
// 	FirstName   string `json:"emp_first"`
// 	LastName    string `json:"emp_last"`
// 	Designation string `json:"emp_Designation"`
// 	Salary      string `json:"emp_salary"`
// 	Location    string `json:"emp_location"`
// }

// type Empid struct {
// 	EmpID string `json:"emp_Id"`
// }

// var Employees []Employee
// var Accounts []Account

func main() {

	// handlerequest()
	//    migrate.Migrate()

	connection, err := driver.LoadConnectionForTraining_Db()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	myrout := mux.NewRouter().StrictSlash(true)
	myrout.Use(middleware.LoggingMiddleware)
	employeeRouter := myrout.PathPrefix("/employee").Subrouter()
	registerEmployeeRoutes(employeeRouter, connection)

	corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(corsObj, allowedMethods, allowedHeaders)(myrout)))
}

func registerEmployeeRoutes(m *mux.Router, c *driver.DB) {
	employeeHandler := eHandler.NewEmployeeHandler(c)
	m.HandleFunc("/", employeeHandler.Fetch).Methods("GET")
	m.HandleFunc("/{emp_Id:[0-9]+}", employeeHandler.GetById).Methods("GET")
	m.HandleFunc("/", employeeHandler.Create).Methods("POST")
	m.HandleFunc("/{emp_Id:[0-9]+}", employeeHandler.Delete).Methods("DELETE")
	m.HandleFunc("/{emp_Id:[0-9]+}", employeeHandler.Update).Methods("PUT")
}

// func ConnectDatabase() (*sql.DB, error) {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	var (
// 		server   = os.Getenv("SERVER")
// 		user     = os.Getenv("USER")
// 		password = os.Getenv("PASSWORD")
// 		database = os.Getenv("DATABASE")
// 	)
// 	connectionString := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", server, database, user, password)
// 	db, err := sql.Open("sqlserver", connectionString)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating connection pool: %w", err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		return nil, fmt.Errorf("error connecting: %w", err)
// 	}
// 	fmt.Println("Connected to Azure SQL Managed Instance using AAD!")
// 	return db, nil
// }

// func Retrieve() ([]Employee, error) {
// 	db, err := ConnectDatabase()
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("Connected with database")
// 	defer db.Close()
// 	createTableSql := `SELECT CAST(employee_id AS NVARCHAR) AS ID, first_name + ' ' + last_name AS Name, designation, salary, country FROM anandhu2.employee;`
// 	rows, err := db.Query(createTableSql)
// 	if err != nil {
// 		return nil, fmt.Errorf("error querying employees: %w", err)
// 	}
// 	defer rows.Close()

// 	var employees []Employee
// 	for rows.Next() {
// 		var emp Employee
// 		if err = rows.Scan(&emp.ID, &emp.Name, &emp.Designation, &emp.Salary, &emp.Country); err != nil {
// 			return nil, fmt.Errorf("error scanning row: %w", err)
// 		}
// 		employees = append(employees, emp)

// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return employees, nil
// }

// func accountdetails(w http.ResponseWriter, r *http.Request) {
// 	Accounts = []Account{
// 		{AccNo: 2020,
// 			Holder: Customer{
// 				ID:   567,
// 				Name: "guidehouse",
// 				DOB:  time.Date(2020, time.April, 12, 0, 0, 0, 0, time.Local),
// 			},
// 			Transaction: []Transaction{
// 				{1, time.Now(), 2000, true},
// 				{2, time.Now(), 5000, false},
// 			},
// 		},
// 		{AccNo: 2021,
// 			Holder: Customer{
// 				ID:   600,
// 				Name: "GH",
// 				DOB:  time.Date(2024, time.April, 12, 0, 0, 0, 0, time.Local),
// 			},
// 			Transaction: []Transaction{
// 				{11, time.Now(), 8000, true},
// 				{22, time.Now(), 9000, false},
// 			},
// 		},
// 		{
// 			AccNo: 2022,
// 			Holder: Customer{
// 				ID:   601,
// 				Name: "JK",
// 				DOB:  time.Date(1985, time.January, 25, 0, 0, 0, 0, time.Local),
// 			},
// 			Transaction: []Transaction{
// 				{33, time.Now(), 15000, true},
// 				{44, time.Now(), 5000, true},
// 			},
// 		},
// 		{
// 			AccNo: 2023,
// 			Holder: Customer{
// 				ID:   602,
// 				Name: "LM",
// 				DOB:  time.Date(1992, time.March, 15, 0, 0, 0, 0, time.Local),
// 			},
// 			Transaction: []Transaction{
// 				{55, time.Now(), 12000, false},
// 				{66, time.Now(), 3000, true},
// 			},
// 		},
// 	}
// 	json.NewEncoder(w).Encode(Accounts)
// }

// func employeid(w http.ResponseWriter, r *http.Request) {
// 	employees, err := Retrieve()
// 	if err != nil {
// 		fmt.Printf("Error fetching data")
// 	}

// 	vars := mux.Vars(r)
// 	key := vars["emp_Id"]
// 	for _, Emp := range employees {

// 		if Emp.ID == key {
// 			json.NewEncoder(w).Encode(Emp)
// 		}
// 	}

// }

// func showemp(w http.ResponseWriter, r *http.Request) {
// 	employees, err := Retrieve()

// 	if err != nil {
// 		http.Error(w, "Error fetching data", http.StatusInternalServerError)
// 	}
// 	json.NewEncoder(w).Encode(employees)

// }

// func addemployee(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		db, err := ConnectDatabase()
// 		if err != nil {
// 			fmt.Println(`Connection not successfull`)
// 			return

// 		}
// 		defer db.Close()
// 		var employee Employee1

// 		err = json.NewDecoder(r.Body).Decode(&employee)
// 		fmt.Println(employee)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		insertSql := `INSERT INTO anandhu2.employee (employee_id, first_name, last_name,  designation, salary, country)
//                 VALUES (@EmpID, @FirstName, @LastName, @Designation, CAST(@Salary AS INT), @Country)`

// 		stmt, err := db.Prepare(insertSql)
// 		if err != nil {
// 			log.Fatal("Error preparing statement: " + err.Error())
// 		}
// 		defer stmt.Close()
// 		_, err = stmt.Exec(sql.Named("EmpID", employee.EmpID), sql.Named("FirstName", employee.FirstName), sql.Named("LastName", employee.LastName),
// 			sql.Named("Designation", employee.Designation), sql.Named("Salary", employee.Salary), sql.Named("Country", employee.Location))
// 		if err != nil {
// 			log.Fatal("Error inserting data: " + err.Error())
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Employee added successfully"))

// 	} else {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 	}

// }

// func deleteemp(w http.ResponseWriter, r *http.Request) {

// 	db, err := ConnectDatabase()
// 	if err != nil {
// 		fmt.Println(`Connection not successfull`)
// 		return
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	empIdStr := vars["emp_Id"]
// 	fmt.Println(empIdStr)

// 	empId, err := strconv.Atoi(empIdStr)
// 	if err != nil {
// 		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
// 		return
// 	}
// 	_, err = db.Exec("DELETE FROM anandhu2.Employee WHERE ID = $1", empId)
// 	if err != nil {
// 		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
// 		return
// 	}
// 	employees, err := Retrieve()
// 	if err != nil {
// 		fmt.Printf("Error fetching data")
// 	}
// 	json.NewEncoder(w).Encode(employees)
// }

// func updatemployee(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodPost {
// 		db, err := ConnectDatabase()
// 		if err != nil {
// 			fmt.Println(`Connection not successfull`)
// 			return

// 		}
// 		defer db.Close()
// 		var emp Employee2
// 		err = json.NewDecoder(r.Body).Decode(&emp)
// 		fmt.Println(emp)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		updatesql := `UPDATE anandhu2.Employee SET first_name=@FirstName, last_name=@LastName, designation=@Designation, country=@Country, salary= @Salary WHERE employee_id= @EmpID`
// 		stmt, err := db.Prepare(updatesql)
// 		if err != nil {
// 			log.Fatal("Error updating statement: " + err.Error())
// 		}
// 		defer stmt.Close()
// 		_, err = stmt.Exec(sql.Named("FirstName", emp.FirstName), sql.Named("LastName", emp.LastName),
// 			sql.Named("Designation", emp.Designation), sql.Named("Salary", emp.Salary), sql.Named("Country", emp.Location), sql.Named("EmpID", emp.Id))
// 		if err != nil {
// 			log.Fatal("Error updating data: " + err.Error())
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Employee updated successfully"))

// 	} else {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)

// 	}
// }

// func handlerequest() {
// 	myrout := mux.NewRouter()
// 	myrout.HandleFunc("/", homepage)
// 	myrout.HandleFunc("/accountdetails", accountdetails)
// 	myrout.HandleFunc("/employeid/{emp_Id}", employeid)
// 	myrout.HandleFunc("/employees", showemp)
// 	myrout.HandleFunc("/employees/{emp_Id}", deleteemp)
// 	myrout.HandleFunc("/addemployee", addemployee)
// 	myrout.HandleFunc("/update", updatemployee)

// 	corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"})
// 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
// 	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
// 	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(corsObj, allowedMethods, allowedHeaders)(myrout)))

// }
