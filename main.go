package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const PORT = ":3030"

type Employee struct{
	ID int
	Name string
	Age int
	Division string
}
var employess=[] Employee{
	{ID : 1, Name: "inri", Age: 23, Division: "Fronted"},
	{ID : 2, Name: "budi", Age: 27, Division: "Marketing"},
	{ID : 3, Name: "udin", Age: 27, Division: "Backend"},
	{ID : 4, Name: "siti", Age: 27, Division: "Marketing"},
}

func main(){
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {

	if r.URL.Path !="/"{
		fmt.Fprint(w, "route not found")
		return
	}
	fmt.Println("hello")
})
http.HandleFunc("/employess", getEmployees)

http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	if r.Method=="GET"{
		json.NewEncoder(w).Encode(employess)
			return		
	}

	if r.Method=="POST"{
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		ageInt, _ := strconv.Atoi(age)

		newEmployee := Employee{
			ID : len(employess)+1,
			Name : name,
			Age : ageInt,
			Division : division,
		}

		employess = append(employess, newEmployee)
		json.NewEncoder(w).Encode(newEmployee)
		return
	}
}