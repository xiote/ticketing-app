package interpark

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type LoginInfo struct {
	ID  string
	PWD string
}

type PriceItem struct {
	SeatGradeName  string
	PriceGradeName string
	SeatCount      string
}

type PriceInfo struct {
	PriceList []PriceItem
}

type SeatsInfo struct {
	Seats []string
}

type PlayDatePlaySeqInfo struct {
	PlayDate string
	PlaySeq  string
}

type GoodsInfo struct {
	URL string
}

type Controller struct {
	selenium.WebDriver
	LoginInfo
	GoodsInfo
	PlayDatePlaySeqInfo
	SeatsInfo
	PriceInfo
}

func NewController(webDriver selenium.WebDriver) Controller {
	return Controller{webDriver, LoginInfo{}, GoodsInfo{}, PlayDatePlaySeqInfo{}, SeatsInfo{}, PriceInfo{}}
}

func NewController2(webDriver selenium.WebDriver, loginInfo LoginInfo) Controller {
	return Controller{webDriver, loginInfo, GoodsInfo{}, PlayDatePlaySeqInfo{}, SeatsInfo{}, PriceInfo{}}
}

func NewController3(webDriver selenium.WebDriver, loginInfo LoginInfo, goodsInfo GoodsInfo, playDatePlaySeqInfo PlayDatePlaySeqInfo, seatsInfo SeatsInfo, priceInfo PriceInfo) Controller {
	return Controller{webDriver, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo}
}

func (c *Controller) Navigate(url string) error {
	return c.WebDriver.Get(url)
}

func (c *Controller) GotoGoodsInfoPage() error {
	var err error

	// 티켓오픈시까지 새로고침
	isOpen := false
	for !isOpen {
		if err = c.WebDriver.Get(c.GoodsInfo.URL); err != nil {
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

func (c *Controller) SelectPlayDatePlaySeq() error {
	var condition selenium.Condition
	var webElement selenium.WebElement
	var err error

	//<a href="javascript:fnPlayDateTab(2);" id="aPlayDateTab" class="btn_view_calendar">달력</a>
	//	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[text()='달력']"); err != nil {
	//		panic(err)
	//	}
	//	if err := webElement.Click(); err != nil {
	//		panic(err)
	//	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "ifrCalendar"); err != nil {
		panic(err)
	}
	c.WebDriver.SwitchFrame(webElement)

	// <a href="javascript:;" onclick="fnSelectPlayDate(5, '20200823')">23</a>
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[contains(@onclick,'"+c.PlayDate+"')]"); err != nil {
			//panic(err)
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[contains(@onclick,'"+c.PlayDatePlaySeqInfo.PlayDate+"')]"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	// <ul id="ulPlaySeq" class="iList" style="position: absolute;"><li><input name="rdnPS" id="rdnPS0" class="option" type="radio" value="0" checked="checked"><label for="rdnPS0">회차를 선택해주세요.</label></li><li><input name="rdnPS" id="rdnPS1" class="option" type="radio" value="070"><label for="rdnPS1" onclick="fnPlaySeqChange(&quot;1&quot;,&quot;070&quot;,&quot;20시 00분 &quot;,&quot;&quot;);">회차&nbsp;&nbsp;&nbsp;20시 00분 </label></li></ul>

	c.WebDriver.SwitchFrame(nil)

	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//div[@class='myValue']"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//ul[@id='ulPlaySeq']//label[contains(@onclick,'"+c.PlayDatePlaySeqInfo.PlaySeq+"')]"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}
	// <a href="#" onclick="javascript:NoMemPrivacyCertify('','20003772');" class="btn_rev"><span>예매하기</span></a>
	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//div[@class='tk_dt_btn_TArea']//a[@class='btn_rev']"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}
	return nil
}

func (c *Controller) SelectSeats() error {
	var err error
	var webElement selenium.WebElement
	var windowHandles []string
	var condition selenium.Condition

	if windowHandles, err = c.WebDriver.WindowHandles(); err != nil {
		panic(err)
	}
	if err = c.WebDriver.SwitchWindow(windowHandles[1]); err != nil {
		panic(err)
	}

	// <img src="//ticketimage.interpark.com/TicketImage/onestop/cost_close.gif" alt="닫기">
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = wd.FindElement(selenium.ByXPATH, "//a[@class='closeBtn']"); err != nil {
			//panic(err)
			return false, nil
		}
		if err := webElement.Click(); err != nil {
			//panic(err)
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	// <iframe id="ifrmSeat" name="ifrmSeat" scrolling="no" width="100%" height="100%" marginwidth="0" marginheight="0" frameborder="no" src="/Book/loading.html"></iframe>
	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "ifrmSeat"); err != nil {
		panic(err)
	}
	if err = c.WebDriver.SwitchFrame(webElement); err != nil {
		panic(err)
	}
	// <iframe id="ifrmSeatDetail" name="ifrmSeatDetail" scrolling="auto" width="658px" height="619px" marginwidth="0" marginheight="0" frameborder="no" src=""></iframe>
	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "ifrmSeatDetail"); err != nil {
		panic(err)
	}
	if err = c.WebDriver.SwitchFrame(webElement); err != nil {
		panic(err)
	}

	//var pageSource string
	//pageSource, err = c.WebDriver.PageSource()
	//fmt.Printf("%s\n", pageSource)

	// <img src="http://ticketimage.interpark.com/TMGSNAS/TMGS/G/1_90.gif" class="stySeat" style="left:335 ;top:241" alt="[VIP석] 1층-B구역14열-23" title="[VIP석] 1층-B구역14열-23" onclick="javascript: SelectSeat('SID49', '1', '1층', 'B구역14열', '23', '002')">
	for _, seat := range c.Seats {
		if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//img[@title='"+seat+"']"); err != nil {
			panic(err)
		}
		if err := webElement.Click(); err != nil {
			panic(err)
		}
	}
	// <a href="javascript:;" onclick="fnSelect();"><img id="NextStepImage" src="http://ticketimage.interpark.com/TicketImage/onestop/btn_seat_confirm_on.gif" alt="좌석선택완료"></a>
	if err = c.WebDriver.SwitchFrame(nil); err != nil {
		panic(err)
	}

	// <iframe id="ifrmSeat" name="ifrmSeat" scrolling="no" width="100%" height="100%" marginwidth="0" marginheight="0" frameborder="no" src="/Book/loading.html"></iframe>
	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "ifrmSeat"); err != nil {
		panic(err)
	}
	if err = c.WebDriver.SwitchFrame(webElement); err != nil {
		panic(err)
	}

	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[@onclick='fnSelect();']"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	return nil
}

func (c *Controller) SelectPrice() error {
	var err error
	var webElement selenium.WebElement
	var condition selenium.Condition

	if err = c.WebDriver.SwitchFrame(nil); err != nil {
		panic(err)
	}
	if webElement, err = c.WebDriver.FindElement(selenium.ByID, "ifrmBookStep"); err != nil {
		panic(err)
	}
	if err = c.WebDriver.SwitchFrame(webElement); err != nil {
		panic(err)
	}

	for _, priceItem := range c.PriceInfo.PriceList {

		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//select[@seatgradename='"+priceItem.SeatGradeName+"']"); err != nil {
				//panic(err)
				return false, nil
			}
			if err := webElement.Click(); err != nil {
				//panic(err)
				return false, nil
			}
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//select[@seatgradename='"+priceItem.SeatGradeName+"']//option[@value='"+priceItem.SeatCount+"']"); err != nil {
				//panic(err)
				return false, nil
			}
			if err := webElement.Click(); err != nil {
				//panic(err)
				return false, nil
			}
			return true, nil
		}
		if err = c.WebDriver.Wait(condition); err != nil {
			panic(err)
		}

	}

	// <a href="javascript:fnNextStep('P');" id="SmallNextBtnLink" onfocus="this.blur();"><img src="http://ticketimage.interpark.com/TicketImage/onestop/btn_next_02_on.gif" alt="다음단계" id="SmallNextBtnImage"></a>
	// 주의 : 위의 경우 anchor가 아닌 image를 클릭해야 한다.
	if err = c.WebDriver.SwitchFrame(nil); err != nil {
		panic(err)
	}
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = wd.FindElement(selenium.ByXPATH, "//img[@id='SmallNextBtnImage']"); err != nil {
			//panic(err)
			return false, nil
		}
		if err := webElement.Click(); err != nil {
			//panic(err)
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

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

	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//div[@class='leftLoginBox']//iframe"); err != nil {
		panic(err)
	}
	c.WebDriver.SwitchFrame(webElement)

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
