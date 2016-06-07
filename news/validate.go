package news


func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}