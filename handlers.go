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
		CompanyNum:   "eeeeeeee",
		JsonTextArea: "I'm inside the getCompanyProfile",
		OfficerJson:  "To Evaluate 2",
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error[%s] on HTML Template Parse", err)
	}

	a := r.FormValue("CompanyNum")
	pageVars.CompanyNum = r.FormValue("CompanyNum")
	fmt.Printf("-----received--------[%s]\n", a)
	pageVars.JsonTextArea = r.FormValue("JsonTextArea")

	param := TemplateUrl{CompanyNum: pageVars.CompanyNum}
	url, err := getApiUrl(COMPANY, ROA, &param)
	fmt.Printf("-------------[%s]\n", url)
	err = t.Execute(w, pageVars)
	if err != nil {
		log.Printf("Error[%s] on HTML Template Execute [%v]", err, pageVars)
	}
}

//__________________________________________________
func startPage(w http.ResponseWriter, r *http.Request) {
	pageVars := companyInfo{
		CompanyNum:   "10989097",
		JsonTextArea: "I'm inside the startPage",
		OfficerJson:  "To Evaluate 1",
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
