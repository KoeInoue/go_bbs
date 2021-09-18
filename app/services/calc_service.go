package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_bbs/models"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type CalcService struct{}

func (ser CalcService) CalcConcurrently(date string) (models.CalcResult, error) {
	now := time.Now()
	var wg sync.WaitGroup

	// Get Prime numbers
	primeFun := func() <-chan []int {
		primeNumCh := make(chan []int)
		wg.Add(1)
		go func() {
			primeNumbers := make([]int, 168)
			defer wg.Done()
			defer close(primeNumCh)
			defer fmt.Println("end prime")
			fmt.Println("start prime")
			ser.getPrimeNumbers(&primeNumbers)
			primeNumCh <- primeNumbers
		}()
		return primeNumCh
	}

	type newsResult struct {
		title string
		url   string
		err   error
	}
	newsFun := func(date string) <-chan newsResult {
		newsCh := make(chan newsResult)
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(newsCh)
			defer fmt.Println("end news")
			fmt.Println("start news")
			title, url, err := ser.getNews(date)
			newsCh <- newsResult{title: title, url: url, err: err}
		}()
		return newsCh
	}

	type weatherResult struct {
		content string
		err     error
	}
	weatherFun := func() <-chan weatherResult {
		wheatherCh := make(chan weatherResult)
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(wheatherCh)
			defer fmt.Println("end weather")
			fmt.Println("start weather")
			content, err := ser.getWheather()
			wheatherCh <- weatherResult{content: content, err: err}
		}()
		return wheatherCh
	}

	primeNumCh := primeFun()
	newsCh := newsFun(date)
	weatherCh := weatherFun()

	primeNum := <-primeNumCh
	newsRes := <-newsCh
	weatherRes := <-weatherCh

	wg.Wait()
	spent := time.Since(now).Seconds()

	return models.CalcResult{
		Spent:        spent,
		PrimeNumbers: primeNum,
		NewsTitle:    newsRes.title,
		NewsUrl:      newsRes.url,
		Wheather:     weatherRes.content,
	}, nil
}

func (ser CalcService) CalcInSeries(date string) (models.CalcResult, error) {
	now := time.Now()
	// Get Prime numbers
	fmt.Println("start prime")
	primeNumbers := make([]int, 168)
	ser.getPrimeNumbers(&primeNumbers)
	fmt.Println("end prime")

	// Get today's news
	fmt.Println("start news")
	newsTitle, newsUrl, err := ser.getNews(date)
	if err != nil {
		return models.CalcResult{}, err
	}
	fmt.Println("end news")

	fmt.Println("start wheather")
	// Get what day is this day
	wheather, err := ser.getWheather()
	if err != nil {
		return models.CalcResult{}, err
	}
	fmt.Println("end wheather")

	return models.CalcResult{
		Spent:        time.Since(now).Seconds(),
		PrimeNumbers: primeNumbers,
		NewsTitle:    newsTitle,
		NewsUrl:      newsUrl,
		Wheather:     wheather,
	}, nil
}

func (CalcService) getPrimeNumbers(prime *[]int) {
	ptr := 0

	(*prime)[ptr] = 2 //2は素数
	ptr++
	(*prime)[ptr] = 3 //3も素数
	ptr++

L:
	for n := 5; n < 1001; n += 2 { //対象は5以上の奇数のみ
		i := 1 //素数のスライスの中を走査する添字。prime[1]の3から
		for (*prime)[i]*(*prime)[i] <= n {
			if n%(*prime)[i] == 0 { //割り切れると素数では無い
				continue L //この素数で割るループを抜ける
			}
			i++
		}
		//最後まで割り切れなかったら
		(*prime)[ptr] = n //素数を登録
		ptr++
	}
}

func (CalcService) getNews(date string) (string, string, error) {
	apiKey := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=apple&to=%s&sortBy=publishedAt&apiKey=%s", date, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
		fmt.Println("Error Response:", resp.Status)
		return "", "", err
	}

	var result map[string][]map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	if len(result["articles"]) == 0 {
		return "", "", errors.New("no articles")
	}
	newsTitle := result["articles"][0]["title"]
	newsUrl := result["articles"][0]["url"]

	return newsTitle, newsUrl, nil
}

func (CalcService) getWheather() (string, error) {
	resp, err := http.Get("https://wttr.in/Tokyo?format=3")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
		fmt.Println("Error Response:", resp.Status)
		return "", err
	}

	weather, _ := ioutil.ReadAll(resp.Body)

	return string(weather), nil
}
