package twentythreeandme

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//HandleDownload gets all the nessesary RS codes from the 23andMe API
func HandleDownload(wr http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	switch req.Method {
	case "POST":

		if req.Body == nil {
			log.Errorf(ctx, "Recieved body empty", nil)
			http.Error(wr, "Recieved body empty", http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Errorf(ctx, "Could not read body", err)
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}
		var ttam TwentyThreeAndMe
		err = json.Unmarshal(body, &ttam)
		if err != nil {
			log.Errorf(ctx, "Could not convert JSON body to struct", err)
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}

		geneMarkers, err := GetTwentyThreeAndMeData(&ttam)

		wrBody, err := json.Marshal(geneMarkers)
		if err != nil {
			log.Errorf(ctx, "Could not convert resulting struct to JSON", err)
			http.Error(wr, err.Error(), http.StatusBadRequest)
			return
		}
		wr.Header().Add("Content-Type", "application/json; charset=utf-8")
		_, err = wr.Write(wrBody)
		if err != nil {
			// Log
		}

	default:
		log.Errorf(ctx, "Method not supported: "+req.Method, nil)
		http.Error(wr, "Method not supported: "+req.Method, http.StatusMethodNotAllowed)
	}
}
