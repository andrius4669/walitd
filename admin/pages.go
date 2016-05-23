package admin

import (
	//str "strings"
)

type deleteValidate struct {
	deleteErr string;
}

func makeErrorMessage(m string) string{
	return "<p class='error'>" + m + "</p>";
}