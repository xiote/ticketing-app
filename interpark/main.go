package interpark

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type LoginInfo struct {
	ID  string
	PWD string
}

type Controller struct {
	selenium.WebDriver
	LoginInfo
}

func NewController(webDriver selenium.WebDriver) Controller {
	return Controller{webDriver, LoginInfo{}}
}

func NewController2(webDriver selenium.WebDriver, loginInfo LoginInfo) Controller {
	return Controller{webDriver, loginInfo}
}

func (c *Controller) Get(url string) error {
	return c.WebDriver.Get(url)
}

func (c *Controller) Login() error {

	var err error
	var webElement selenium.WebElement
	var webElements []selenium.WebElement
	var title string
	var currentUrl string

	if err = c.WebDriver.Get("http://ticket.interpark.com/"); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "imgLogin"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	if webElements, err = c.WebDriver.FindElements(selenium.ByTagName, "iframe"); err != nil {
		panic(err)
	}
	fmt.Println(len(webElements))

	if title, err = c.WebDriver.Title(); err != nil {
		panic(err)
	}
	fmt.Println(title)
	if currentUrl, err = c.WebDriver.CurrentURL(); err != nil {
		panic(err)
	}
	fmt.Println(currentUrl)

	c.WebDriver.SwitchFrame(webElements[0])

	if title, err = c.WebDriver.Title(); err != nil {
		panic(err)
	}
	fmt.Println(title)
	if currentUrl, err = c.WebDriver.CurrentURL(); err != nil {
		panic(err)
	}
	fmt.Println(currentUrl)

	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "userId"); err != nil {
		panic(err)
	}
	if err = webElement.SendKeys(c.LoginInfo.ID); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "userPwd"); err != nil {
		panic(err)
	}
	if err = webElement.SendKeys(c.LoginInfo.PWD); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "btn_login"); err != nil {
		panic(err)
	}
	if err = webElement.Click(); err != nil {
		panic(err)
	}

	// waiting for http://ticket.interpark.com/
	var condition = func(wd selenium.WebDriver) (bool, error) {
		if currentUrl, err = wd.CurrentURL(); err != nil {
			//panic(err)
			return true, nil
		}

		if currentUrl != "http://ticket.interpark.com" {
			return true, nil
		}

		return false, nil
	}

	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	return nil
}
