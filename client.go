package twentythreeandme

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"google.golang.org/appengine/urlfetch"

	"golang.org/x/net/context"
)

var baseURL = `https://api.23andme.com/3/marker/`

func GetTwentyThreeAndMeData(ctx *context.Context, ttam *TwentyThreeAndMe) (*[]GeneMarker, error) {
	var wg sync.WaitGroup

	var accessToken = ttam.Token

	geneMarker := make([]GeneMarker, len(ttam.Scope))
	for i, RSCode := range ttam.Scope {
		// time.Sleep(20 * time.Millisecond)
		wg.Add(1)
		go getGeneMarker(ctx, RSCode, accessToken, &geneMarker[i], &wg)
	}
	fmt.Println("Waiting")
	wg.Wait()

	fmt.Println("Done")

	return &geneMarker, nil
}

func getGeneMarker(ctx *context.Context, RSCode string, accessToken string, geneMarker *GeneMarker, wg *sync.WaitGroup) error {
	defer wg.Done()
	var url = baseURL + RSCode
	attempts := 5

	for i := 0; i < attempts; i++ {
		err := jsonResponse(ctx, "GET", url, accessToken, geneMarker)
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Retrying ", i)
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

	err = json.NewDecoder(response.Body).Decode(geneMarker)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
