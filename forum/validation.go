package forum

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func validBoardName(board string) bool {
	slen := len(board)
	if slen == 0 || slen > 64 {
		// current limit is 64
		return false
	}
	for i := 0; i < slen; i++ {
		if (board[i] >= 'a' && board[i] <= 'z') || (board[i] >= '0' && board[i] <= '9') {
			continue
		}
		if board[i] == '.' && i > 0 && board[i-1] != '.' && i < len(board) {
			continue
		}
		return false
	}
	return true
}

func reservedBoardName(name string) bool {
	return name == "mod" || name == "static"
}
