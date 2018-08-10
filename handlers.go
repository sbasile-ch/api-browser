package main

//package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//__________________________________________________
func CallApi(companyNum string) {
	fmt.Printf("-------------%s---%s", companyNum, ApiKey)
}

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

	pageVars.CompanyNum = r.FormValue("CompanyNum")
	CallApi(pageVars.CompanyNum)
	err = t.Execute(w, pageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
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

	err = t.Execute(w, pageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
