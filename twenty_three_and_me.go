package twentythreeandme

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", HandleDownload)
}
