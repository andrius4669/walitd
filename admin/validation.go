package admin

func confirmDelete(form *deleteValidate) *deleteValidate{
	form.deleteErr = makeErrorMessage("Deleted");
	form.deleteErr = "deleted";
	return form
}
