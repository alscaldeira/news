package service

import (
	"log"
	"time"

	"github.com/alscaldeira/twitter/utils"
	"github.com/fedesog/webdriver"
)

type Web struct {
	chromeDriver *webdriver.ChromeDriver
	session      *webdriver.Session
}

const (
	baseUrlTwitter  = "https://twitter.com/"
	urlTwitterLogin = baseUrlTwitter + "i/flow/login"
	pathDriver      = "driver/chromedriver"
)

func Init(username string, password string) Web {

	chromedriver := webdriver.NewChromeDriver(pathDriver)
	err := chromedriver.Start()
	if err != nil {
		log.Println(err)
	}

	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromedriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}

	session.Url(urlTwitterLogin)
	deadline := time.Now().Add(300 * time.Second)

	for {
		inputLogin, err := session.FindElement(webdriver.FindElementStrategy("xpath"), "//*[@autocomplete='username']")
		if err != nil {
			log.Println(err)
		}

		isAvailable := webdriver.WebElement{} != inputLogin
		if isAvailable {
			inputLogin.SendKeys(username)

			time.Sleep(2 * time.Second)

			inputBtnConfirmLogin, err := session.FindElements(webdriver.FindElementStrategy("xpath"), "//*[@role='button']")
			if err != nil {
				log.Println(err)
			}

			inputBtnConfirmLogin[2].Click()
			time.Sleep(2 * time.Second)

			inputPassword, err := session.FindElement(webdriver.FindElementStrategy("xpath"), "//*[@autocomplete='current-password']")
			if err != nil {
				log.Println(err)
			}

			isAvailable := webdriver.WebElement{} != inputPassword
			if isAvailable {
				utils.Log(password)
				inputPassword.SendKeys(password)
				time.Sleep(2 * time.Second)

				inputBtnConfirmPassword, err := session.FindElement(webdriver.FindElementStrategy("xpath"), "//*[@data-testid='LoginForm_Login_Button']")
				if err != nil {
					log.Println(err)
				}

				inputBtnConfirmPassword.Click()
				time.Sleep(2 * time.Second)
				break
			}

			if time.Now().After(deadline) {
				utils.Logeadline("[Deadline]")
				break
			}
		}
	}

	time.Sleep(5 * time.Second)
	return Web{
		chromeDriver: chromedriver,
		session:      session,
	}
}

func PostContent(content string, web Web) {
	inputPost, err := web.session.FindElement(webdriver.FindElementStrategy("xpath"), "//div[@class='public-DraftStyleDefault-block public-DraftStyleDefault-ltr']//*[@data-text='true']")
	if err != nil {
		log.Println(err)
	}

	inputPost.SendKeys(content)

	btnPost, err := web.session.FindElement(webdriver.FindElementStrategy("xpath"), "//div[@class='css-175oi2r r-sdzlij r-1phboty r-rs99b7 r-lrvibr r-19u6a5r r-2yi16 r-1qi8awa r-ymttw5 r-1loqt21 r-o7ynqc r-6416eg r-1ny4l3l']")
	if err != nil {
		log.Println(err)
	}

	btnPost.Click()
}

func Finish(web Web) {
	web.session.Delete()
	web.chromeDriver.Stop()
}
