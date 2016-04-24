package forum

// sql related stuff
// currently only github.com/lib/pq lib usage

import (
	//"fmt"
	"database/sql"
	"strconv"
)

func queryThread(db *sql.DB, board, thread string) bool {
	// sanity checks first
	if !validBoardName(board) {
		return false
	}

	//var tid uint32
	var err error
	// no casting for multiple return values. no nested funcs too. fuck you golang.
	// I hope this shit will atleast get inlined.
	cst := func(i uint64, e error) (uint32, error) {
		return uint32(i), e
	}
	_, err = cst(strconv.ParseUint(thread, 10, 32))
	if err != nil {
		return false
	}

	// TODO(andrius)
	return false
}
