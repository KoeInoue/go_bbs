package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_bbs/models"
	"net/http"
	"os"
	"time"
)

type CalcService struct{}

func (CalcService) CalcConcurrently(date string) (models.CalcResult, error) {
	return models.CalcResult{}, nil
}

func (ser CalcService) CalcInSeries(date string) (models.CalcResult, error) {
	now := time.Now()
	// Get Prime numbers
	primeNumbers := make([]int, 168)
	ser.getPrimeNumbers(&primeNumbers)

	// Get today's news
	newsTitle, newsUrl, err := ser.getNews(date)
	if err != nil {
		return models.CalcResult{}, err
	}

	// Get what day is this day
	whatDay, err := ser.getWhatDay(date)

	return models.CalcResult{
		Spent:        time.Since(now).Seconds(),
		PrimeNumbers: primeNumbers,
		NewsTitle:    newsTitle,
		NewsUrl:      newsUrl,
		WhatDay:      whatDay,
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
	fmt.Println(newsTitle)

	return newsTitle, newsUrl, nil
}

func (CalcService) getWhatDay(date string) (string, error) {
	return "", errors.New("No")
}
