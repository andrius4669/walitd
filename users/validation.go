package users

import (
//	"fmt"
//	"net/http"
	"net/mail"
)

func validateRegisterForm(form *userForm) *userForm  {
//	fmt.Printf("%v", form);
	form.ErrorCnt = 0;
	if (len(form.Username) < 5){
		form.p();
		form.UsernameErr = makeErrorMessage("Username is too short");
	}
	if (len(form.FirstName) < 3){
		form.p();
		form.FirstNameErr = makeErrorMessage("First name is too short");
	}
	if (len(form.SecondName) < 3){
		form.p();
		form.SecondNameErr = makeErrorMessage("Last name is too short");
	}
	if (len(form.Pass) < 5){
		form.p();
		form.PassErr = makeErrorMessage("Password is too short");
	}
	//TODO: check if email validation works
	_, err := mail.ParseAddress(form.Email)
	if err != nil {
		form.p();
		form.EmailErr = makeErrorMessage("Bad email");
	}

	return form;
}
func validateLoginForm(form *loginInfo) *loginInfo  {
	//TODO: check login

	return form;
}
func validateProfileForm(form *user) *user{
	form.Err = 0;

	if (len(form.FirstName) < 3){
		form.p();
		form.FirstNameErr = makeErrorMessage("First name is too short");
	}
	if (len(form.SecondName) < 3){
		form.p();
		form.SecondNameErr = makeErrorMessage("Last name is too short");
	}
	if (len(form.Pass) < 5 && len(form.Pass) != 0){
		form.p();
		form.PassErr = makeErrorMessage("Password is  too short");
	}
	//TODO: check if email validation works
	_, err := mail.ParseAddress(form.Email)
	if err != nil {
		form.p();
		form.EmailErr = makeErrorMessage("Bad email");
	}

	return form;
}
