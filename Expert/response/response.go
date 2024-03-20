package response

import (
	"encoding/json"
	"expert/model"
	"net/http"

	"github.com/lib/pq"
)

func MessageShow(code int, errMessage string, w http.ResponseWriter) {
	var message model.Message
	message.Code = code
	message.Message = errMessage
	messageData, err := json.MarshalIndent(message, "", "  ")
	if err != nil {
		message.Code = 400
		message.Message = "Error on marshaling"
	}
	w.WriteHeader(code)
	w.Write(messageData)
}

func DatabaseErrorShow(err error) (string, int) {
	if dbErr, ok := err.(*pq.Error); ok { // For PostgreSQL database driver (pq)
		errCode := dbErr.Code
		switch errCode {
		case "23502":
			// not-null constraint violation
			return "Some required data was left out", 400

		case "23503":
			// foreign key violation
			return "This record can't be changed because another record refers to it", 409

		case "23505":
			// unique constraint violation
			return "This record contains duplicated data that conflicts with what is already in the database", 409

		case "23514":
			// check constraint violation
			return "This record contains inconsistent or out-of-range data", 400

		}
	}
	return "Internal server error", 500
}
