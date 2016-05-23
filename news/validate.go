package news

//func confirmDelete(form *deleteValidate) *deleteValidate{
//	form.deleteErr = makeErrorMessage("Deleted");
//	return form
//}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}