package cookieservice

import (
	"net/http"
	"time"
)

//ReadCookie Use this for getting value of a cookie.
func ReadCookie(r *http.Request, cookieName string) string {
	var cookie, err = r.Cookie(cookieName)
	if cookie == nil || err != nil {
		return ""
	}
	return cookie.Value
}

//NewCookie Use this for creating a new cookie.
func NewCookie(w http.ResponseWriter, cookieName, val string, hour int) {
	expiration := time.Now().Add(time.Duration(hour) * time.Hour)
	cookie := http.Cookie{Name: cookieName, Value: val, Path: "/", Expires: expiration}
	http.SetCookie(w, &cookie)
}

//DeleteCookie Use this for deleting a cookie.
func DeleteCookie(w http.ResponseWriter, name string) {
	expiration := time.Now().Add(-1)
	cookie := http.Cookie{Name: name, Value: "", Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
}




