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

type articlesList struct {
	ID int
	Article string
	Name string
	Score int
	Visit_Count int
	Description string
	Category string
	Author       int
	AuthorName string
	UploadDate string
	FullArticleLink string
}