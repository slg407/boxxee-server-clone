package boxee

import (
	"fmt"
	"log"
	"net/http"
)

const loginResponseTemplate = `<?xml version="1.0" encoding="UTF-8" ?>
<object type="user" id="boxeeuser">
<name>Boxee User</name>
<short_name>Boxee User</short_name>
<thumb></thumb>
<thumb_small></thumb_small>
<user_id>boxeeuser</user_id>
<user_display_name>Boxee User</user_display_name>
<user_first_name></user_first_name>
<user_last_name></user_last_name>
<country>US</country>
<show_movie_library>1</show_movie_library>
</object>
`

func Login(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Print("Faking a login")
		_, err := fmt.Fprint(w, loginResponseTemplate)
		if err != nil {
			log.Println(err)
		}
	}
}
