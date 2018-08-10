package main

import (
	//"github.com/sbasile-ch/api-browser/handlers"
	"net/http"
	"os"
)

type companyInfo struct {
	CompanyNum  string
	CompanyJson string
	OfficerJson string
}

var ApiKey string = os.Getenv("MY_CH_API")

//__________________________________________________
func main() {
	http.HandleFunc("/getCompany", getCompanyProfile)
	http.HandleFunc("/index", startPage)
	/*
		http.HandleFunc("/getOfficer", getCompanyOfficer)
		http.HandleFunc("/getPsc", getCompanyPsc)
		http.HandleFunc("/getRegiters", getCompanyRegisters)
		http.HandleFunc("/setAlert", setCompanyAlert)
	*/

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe("127.0.0.1:8080", nil)
}
