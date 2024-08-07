package controller

import (
	convert "challenger/adapter/input/model/converter"
	"challenger/adapter/input/model/response"
	"challenger/app/port/input"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
)

func NewController(service input.ContactDomainService) ControllerInterface {
	file, err := os.OpenFile("requestLog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "request: ", log.LstdFlags)
	return &Controller{
		service: service,
		log:     logger,
	}
}

type ControllerInterface interface {
	Form(w http.ResponseWriter, r *http.Request)
	CreateContact(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	service input.ContactDomainService
	log     *log.Logger
}

func (c *Controller) Form(w http.ResponseWriter, r *http.Request) {
	c.log.Println(r.Host, r.RemoteAddr, r.RequestURI)

	siteKey := os.Getenv("HCAPTCHA-SITE-KEY")
	registerForm := template.Must(template.ParseFiles("./template/form.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	registerForm.Execute(w, map[string]string{
		"SiteKey": siteKey,
	})
}

func (c *Controller) CreateContact(w http.ResponseWriter, r *http.Request) {
	c.log.Println(r.Host, r.RemoteAddr, r.RequestURI)

	if !verifyRecaptcha(r.FormValue("g-recaptcha-response")) {
		json.NewEncoder(w).Encode(response.ResponseError{TypeErro: "captcha", Title: "the captcha", Detail: []string{"error: captcha is incorrect", "path:" + r.Host + r.RequestURI}})
		return
	}

	formContact, err := convert.ConvertHttpRequestToRequestConect(r)

	if err != nil {
		json.NewEncoder(w).Encode(response.ResponseError{TypeErro: "invalid parameter", Title: "request parameters invalide", Detail: []string{"error: age must be a number between 0 and " + fmt.Sprintf("%d", math.MaxUint8), "path:" + r.Host + r.RequestURI}})
		return
	}

	resp, erro := convert.ConvertContactRequestoToResponse(formContact)
	if erro != nil {
		json.NewEncoder(w).Encode(response.ResponseError{TypeErro: "invalid parameter", Title: "request parameters invalide", Detail: erro})
		c.log.Println(erro)
		return
	} else {
		json.NewEncoder(w).Encode(resp)
		c.service.CreateContactServices(resp)
		return
	}
}

func verifyRecaptcha(token string) bool {

	if token == "" {
		return false
	}

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{"secret": {os.Getenv("HCAPTCHA-SECRET-KEY")}, "response": {token}})
	if err != nil {
		log.Printf("error: %v", err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error: %v", err)
		return false
	}

	var rr response.RecaptchaResponse
	err = json.Unmarshal(body, &rr)
	if err != nil {
		log.Printf("error: %v", err)
		return false
	}
	return rr.Success
}
