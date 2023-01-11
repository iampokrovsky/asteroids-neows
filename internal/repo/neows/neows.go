package neows

import (
	"encoding/json"
	"github.com/pokrovsky-io/neows-asteroids/internal/entity"
	"io"
	"net/http"
	"net/url"
	"sort"
	"sync"
)

type Data struct {
	Count int `json:"element_count"`
}

type NeoWsAPI struct {
	url string
	key string
}

func New(url string, key string) *NeoWsAPI {
	return &NeoWsAPI{
		url: url,
		key: key,
	}
}

func (api *NeoWsAPI) makeRequest(wg *sync.WaitGroup, date string, out chan<- entity.AsteroidsReport) {
	defer wg.Done()

	params := url.Values{}
	params.Add("api_key", api.key)
	params.Add("end_date", date)
	params.Add("start_date", date)

	reqUrl := api.url + "?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var data Data
	if err := json.Unmarshal(body, &data); err != nil {
		return
	}

	rep := entity.AsteroidsReport{
		Date:  date,
		Count: data.Count,
	}

	out <- rep
}

func (api *NeoWsAPI) Get(dates []string) ([]entity.AsteroidsReport, error) {
	wg := &sync.WaitGroup{}
	out := make(chan entity.AsteroidsReport, len(dates))

	for _, date := range dates {
		date := date

		wg.Add(1)
		go api.makeRequest(wg, date, out)
	}

	wg.Wait()

	close(out)

	res := make([]entity.AsteroidsReport, 0, len(dates))

	for rep := range out {
		res = append(res, rep)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Date < res[j].Date
	})

	return res, nil
}
