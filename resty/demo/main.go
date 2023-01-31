package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	SourceUrl      = "https://cloudso.org"
	LoginUrl       = SourceUrl + "/auth/login"
	GetUserProfile = SourceUrl + "/user/profile"
	CheckInUrl     = SourceUrl + "/user/checkin"
)

func main() {
	accounts := []Account{}
	for {
		for _, account := range accounts {
			cookies, err := login(account.Email, account.Passwd)
			if err != nil {
				log.Println(err)
			}
			//if err := getUserProfile(cookies); err != nil {
			//	log.Println(err)
			//}
			err = checkIn(cookies)
			if err != nil {
				log.Printf("自动签到失败，邮箱：%s，时间：%s，失败原因：%s", account.Email, time.Now(), err)
			} else {
				log.Printf("自动签到成功，邮箱：%s，时间：%s", account.Email, time.Now())
			}
		}
		time.Sleep(time.Hour * 24)
		//dur := int64(rand.Intn(60))
		//time.Sleep(time.Duration(dur) * time.Minute)
	}

	//ticker := time.NewTicker(24 * time.Hour)
	//defer ticker.Stop()
	//for range ticker.C {
	//	for _, account := range accounts {
	//		cookies, err := login(account.Email, account.Passwd)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//		//if err := getUserProfile(cookies); err != nil {
	//		//	log.Println(err)
	//		//}
	//		err = checkIn(cookies)
	//		if err != nil {
	//			log.Printf("自动签到失败，邮箱：%s，时间：%s，失败原因：%s", account.Email, time.Now(), err)
	//		} else {
	//			log.Printf("自动签到成功，邮箱：%s，时间：%s", account.Email, time.Now())
	//		}
	//	}
	//	dur := int64(rand.Intn(60))
	//	time.Sleep(time.Duration(dur) * time.Minute)
	//}
}

type Account struct {
	Email  string
	Passwd string
}

func login(email, passwd string) ([]*http.Cookie, error) {
	client := resty.New().SetTimeout(3 * time.Second).SetDisableWarn(true).SetRetryCount(0).SetRetryWaitTime(500 * time.Millisecond)
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"email":      email,
			"passwd":     passwd,
			"code":       "",
			"remeber_me": "on",
		}).Post(LoginUrl)

	if err != nil {
		return nil, err
	}

	return resp.Cookies(), nil
}

func getUserProfile(cookies []*http.Cookie) error {
	client := resty.New().SetTimeout(3 * time.Second).SetDisableWarn(true).SetRetryCount(0).SetRetryWaitTime(500 * time.Millisecond)
	resp, err := client.R().SetCookies(cookies).Get(GetUserProfile)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("登录失败")
	}
	return nil
}

func checkIn(cookies []*http.Cookie) error {
	client := resty.New().SetTimeout(3 * time.Second).SetDisableWarn(true).SetRetryCount(0).SetRetryWaitTime(500 * time.Millisecond)
	resp, err := client.R().SetCookies(cookies).Post(CheckInUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("签到失败")
	}
	return nil
}
