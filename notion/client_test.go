package notion

import (
	"fmt"
	"net/http"
)

func getErrorJSON(status int) string {
	return fmt.Sprintf(`{
	"status": %d,
	"message": "error",
	"type": "error"
}`, status)
}

func addHeader(w http.ResponseWriter) {
	w.Header().Add(rateLimitRemainingHeader, "99")
	w.Header().Add(rateLimitLimitHeader, "1000")
	w.Header().Add(rateLimitResetHeader, "1598795193")
}
