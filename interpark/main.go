package interpark

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type PlayDatePlaySeq struct {
	PlayDate string
	PlaySeq  string
}
type LoginInfo struct {
	ID  string
	PWD string
}

type Controller struct {
	selenium.WebDriver
	LoginInfo
	GoodsInfoUrl string
	PlayDatePlaySeq
}

func NewController(webDriver selenium.WebDriver) Controller {
	return Controller{webDriver, LoginInfo{}, "", PlayDatePlaySeq{}}
}

func NewController2(webDriver selenium.WebDriver, loginInfo LoginInfo) Controller {
	return Controller{webDriver, loginInfo, "", PlayDatePlaySeq{}}
}

func NewController3(webDriver selenium.WebDriver, loginInfo LoginInfo, goodsInfoUrl string, playDatePlaySeq PlayDatePlaySeq) Controller {
	return Controller{webDriver, loginInfo, goodsInfoUrl, playDatePlaySeq}
}

func (c *Controller) Navigate(url string) error {
	return c.WebDriver.Get(url)
}

func (c *Controller) GotoGoodsInfoPage() error {
	var err error

	// 티켓오픈시까지 새로고침
	isOpen := false
	for !isOpen {
		if err = c.WebDriver.Get(c.GoodsInfoUrl); err != nil {
			panic(err)
		}

		if _, err = c.WebDriver.FindElement(selenium.ByID, "divCarendar"); err != nil {
			isOpen = false
		} else {
			isOpen = true
		}

	}
	return nil
}

func (c *Controller) SelectPlayDayPlaySeq() error {
	var condition selenium.Condition
	var webElement selenium.WebElement
	var err error
	// <a href="javascript:;" onclick="fnSelectPlayDate(5, '20200823')">23</a>
	// waiting for http://ticket.interpark.com/
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = c.WebDriver.FindElement(selenium.ByID, "imgLogin"); err != nil {
			//panic(err)
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "imgLogin"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	// <ul id="ulPlaySeq" class="iList" style="position: absolute;"><li><input name="rdnPS" id="rdnPS0" class="option" type="radio" value="0" checked="checked"><label for="rdnPS0">회차를 선택해주세요.</label></li><li><input name="rdnPS" id="rdnPS1" class="option" type="radio" value="070"><label for="rdnPS1" onclick="fnPlaySeqChange(&quot;1&quot;,&quot;070&quot;,&quot;20시 00분 &quot;,&quot;&quot;);">회차&nbsp;&nbsp;&nbsp;20시 00분 </label></li></ul>

	// <a href="#" onclick="javascript:NoMemPrivacyCertify('','20003772');" class="btn_rev"><span>예매하기</span></a>
	return nil
}

func (c *Controller) SelectSeats() error {
	return nil
}

func (c *Controller) SelectPrice() error {
	return nil
}

func (c *Controller) SelectDelivery() error {
	return nil
}

func (c *Controller) SelectPayment() error {
	return nil
}

func (c *Controller) Login() error {

	var err error
	var webElement selenium.WebElement
	var webElements []selenium.WebElement
	var title string
	var currentUrl string
	var condition selenium.Condition

	if err = c.WebDriver.Get("http://ticket.interpark.com/"); err != nil {
		panic(err)
	}

	// waiting for http://ticket.interpark.com/
	condition = func(wd selenium.WebDriver) (bool, error) {
		if currentUrl, err = wd.CurrentURL(); err != nil {
			//panic(err)
			return false, nil
		}
		if currentUrl != "http://ticket.interpark.com/" {
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "imgLogin"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	// waiting for https://ticket.interpark.com/Gate/TPLogin.asp?CPage=B&MN=Y&tid1=main_gnb&tid2=right_top&tid3=login&tid4=login
	condition = func(wd selenium.WebDriver) (bool, error) {
		if currentUrl, err = wd.CurrentURL(); err != nil {
			//panic(err)
			return false, nil
		}
		if currentUrl != "https://ticket.interpark.com/Gate/TPLogin.asp?CPage=B&MN=Y&tid1=main_gnb&tid2=right_top&tid3=login&tid4=login" {
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	if webElements, err = c.WebDriver.FindElements(selenium.ByTagName, "iframe"); err != nil {
		panic(err)
	}
	//fmt.Println(len(webElements))

	//if title, err = c.WebDriver.Title(); err != nil {
	//	panic(err)
	//}
	//fmt.Println(title)
	//if currentUrl, err = c.WebDriver.CurrentURL(); err != nil {
	//	panic(err)
	//}
	fmt.Println(currentUrl)

	c.WebDriver.SwitchFrame(webElements[0])

	//if title, err = c.WebDriver.Title(); err != nil {
	//	panic(err)
	//}
	fmt.Println(title)
	//if currentUrl, err = c.WebDriver.CurrentURL(); err != nil {
	//	panic(err)
	//}
	//fmt.Println(currentUrl)

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
	condition = func(wd selenium.WebDriver) (bool, error) {
		if currentUrl, err = wd.CurrentURL(); err != nil {
			//panic(err)
			return false, nil
		}
		if currentUrl != "http://ticket.interpark.com/" {
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	return nil
}
