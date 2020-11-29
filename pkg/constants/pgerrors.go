package constants

import "github.com/lib/pq"

/* http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html */

// IsErrUniqueViolation checks if error is unique constraint violation
func IsErrUniqueViolation(err error) bool {
	pqErr, ok := err.(*pq.Error)
	return ok && pqErr.Code == "23505"
}
