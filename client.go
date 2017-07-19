package twentythreeandme

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	"golang.org/x/net/context"
)

var baseURL = `https://api.23andme.com/3/profile/`

// GetTwentyThreeAndMeData This function downloads the given SNP's concurrently from 23andMe
func GetTwentyThreeAndMeData(ctx *context.Context, ttam *TwentyThreeAndMe) (*[]GeneMarker, error) {
	var wg sync.WaitGroup

	var accessToken = ttam.Token

	geneMarker := make([]GeneMarker, len(ttam.Scope))

	baseURL := baseURL + ttam.ProfileId + "/marker/"

	for i, RSCode := range ttam.Scope {
		url := baseURL + RSCode + "/"
		// time.Sleep(20 * time.Millisecond)
		wg.Add(1)
		go getGeneMarker(ctx, url, accessToken, &geneMarker[i], &wg)
	}
	wg.Wait()
	return &geneMarker, nil
}

func getGeneMarker(ctx *context.Context, url string, accessToken string, geneMarker *GeneMarker, wg *sync.WaitGroup) error {
	defer wg.Done()
	attempts := 5

	for i := 0; i < attempts; i++ {
		err := jsonResponse(ctx, "GET", url, accessToken, geneMarker)
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
		log.Debugf(*ctx, "Retrying")
	}

	return nil
}

func jsonResponse(ctx *context.Context, httpMethod string, url string, accessToken string, geneMarker *GeneMarker) error {

	client := urlfetch.Client(*ctx)
	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	// requestDump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// log.Debugf(*ctx, "Send", string(requestDump))

	// responseDump, err := httputil.DumpResponse(response, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// log.Debugf(*ctx, "Received", string(responseDump))

	err = json.NewDecoder(response.Body).Decode(geneMarker)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
