package controller

import (
	convert "challenger/adapter/input/model/converter"
	"challenger/adapter/input/model/request"
	"challenger/adapter/input/model/response"
	"challenger/app/domain"
	"challenger/app/port/input"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/kataras/hcaptcha"
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
	registerForm := template.Must(template.ParseFiles("./form.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	registerForm.Execute(w, map[string]string{
		"SiteKey": siteKey,
	})
}

func (c *Controller) CreateContact(w http.ResponseWriter, r *http.Request) {
	c.log.Println(r.Host, r.RemoteAddr, r.RequestURI)
	hcaptchaResp, ok := hcaptcha.Get(r)
	if !ok {
		c.log.Println(hcaptchaResp.ErrorCodes)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response.ResponseError{TypeErro: "captcha", Title: "the captcha", Detail: []string{"error: captcha is incorrect", "path:" + r.Host + r.RequestURI}})
		return
	}

	var formContact request.ContactRequest

	formContact.Name = r.FormValue("name")
	formContact.Email = r.FormValue("email")
	age := r.FormValue("age")

	num, err := strconv.ParseUint(age, 10, 8)
	if err != nil {
		json.NewEncoder(w).Encode(response.ResponseError{TypeErro: "invalid parameter", Title: "request parameters invalide", Detail: []string{"error: age must be a number between 0 and " + fmt.Sprintf("%d", math.MaxUint8), "path:" + r.Host + r.RequestURI}})
		return
	}
	formContact.Age = uint8(num)
	fmt.Println(num)

	resp, erro := convert.ConvertContactRequestoToResponse(formContact)
	if erro != nil {
		c.log.Println(erro)
	} else {
		c.service.CreateContactServices(domain.ContactDomain{Email: resp.Email, Name: resp.Name, Age: resp.Age})
	}
}
