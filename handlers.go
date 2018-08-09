package main

//package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//__________________________________________________
func getCompanyProfile(w http.ResponseWriter, r *http.Request) {
	pageVars := companyInfo{
		CompanyNum:  "eeeeeeee",
		CompanyJson: "I'm inside the getCompanyProfile",
		OfficerJson: "To Evaluate 2",
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	fmt.Println("arrived here 1:")

	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Printf("arrived here with CompanyNum: %v %v", t, pageVars)
	companyNum := r.FormValue("CompanyNum")
	fmt.Println("arrived here 2:")
	companyNum2 := "aaaaa"

	fmt.Printf("arrived here with CompanyNum: %s\n", companyNum)
	fmt.Printf("arrived here with CompanyNum: %s\n %v %v", companyNum2, t, pageVars)
	fmt.Fprintln(w, r.FormValue(".CompanyNum"))

	//err = t.Execute(w, pageVars)
	//if err != nil {
	//	log.Print("template executing error: ", err)
	//}
}

//__________________________________________________
func startPage(w http.ResponseWriter, r *http.Request) {
	pageVars := companyInfo{
		CompanyNum:  "10989097",
		CompanyJson: "I'm inside the startPage",
		OfficerJson: "To Evaluate 1",
	}

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	//fmt.Println("arrived here with CompanyNUm:", CompanyNum)

	err = t.Execute(w, pageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
