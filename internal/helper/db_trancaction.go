package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx, err *error) {
	errRecover := recover()
	if errRecover != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			*err = errorRollback
			return
		}
		errRecoverValue, ok := errRecover.(error)
		if ok {
			*err = errRecoverValue
			return
		}
		log.Fatal(errRecover)
	} else {
		errorCommit := tx.Commit()
		if errorCommit != nil {
			*err = errorCommit
		}
	}
}
