package twentythreeandme

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//HandleDownload gets all the nessesary RS codes from the 23andMe API
func HandleDownload(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

		if req.Body == nil {
			http.Error(wr, "Recieved body empty", http.StatusBadRequest)
			return
		}
		// fmt.Println(req.Body)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}
		var ttam TwentyThreeAndMe
		err = json.Unmarshal(body, &ttam)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		geneMarkers := GetTwentyThreeAndMeData(&ttam)

		wrBody, err := json.Marshal(geneMarkers)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = wr.Write(wrBody)
		if err != nil {
			// Log
		}

	default:
		http.Error(wr, "Method not supported: "+req.Method, http.StatusMethodNotAllowed)
	}
}
