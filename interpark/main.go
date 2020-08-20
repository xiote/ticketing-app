package interpark

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

type PaymentInfo struct {
	PaymentType   string
	PaymentSelect string
	DiscountCard  string
	BankName      string
}

func NewPaymentInfo(paymentType string, paymentSelect string, discountCard string) PaymentInfo {
	return PaymentInfo{paymentType, paymentSelect, discountCard, ""}
}

func NewPaymentInfo2(paymentType string, paymentSelect string, discountCard string, bankName string) PaymentInfo {
	return PaymentInfo{paymentType, paymentSelect, discountCard, bankName}
}

type DeliveryInfo struct {
	DeliveryType string
	YYMMDD       string
	EMail        string
}

func NewDeliveryInfo2(deliveryType string, yymmdd string, email string) DeliveryInfo {
	return DeliveryInfo{deliveryType, yymmdd, email}
}
func NewDeliveryInfo(deliveryType string, yymmdd string) DeliveryInfo {
	return DeliveryInfo{deliveryType, yymmdd, ""}
}

type LoginInfo struct {
	ID       string
	PWD      string
	SiteName string
}

func NewLoginInfo2(id string, pwd string, siteName string) LoginInfo {
	return LoginInfo{id, pwd, siteName}
}
func NewLoginInfo(id string, pwd string) LoginInfo {
	return LoginInfo{id, pwd, "ticket.interpark.com"}
}

type PriceItem struct {
	SeatGradeName  string
	PriceGradeName string
	SeatCount      string
}

func NewPriceItem(seatGradeName string, priceGradeName string, seatCount string) PriceItem {
	return PriceItem{seatGradeName, priceGradeName, seatCount}
}

type PriceInfo struct {
	PriceList []PriceItem
}

func NewPriceInfo(priceList []PriceItem) PriceInfo {
	return PriceInfo{priceList}
}

type SeatsInfo struct {
	Seats           []string
	ClickCloseBtnYN string
	CaptchaYN       string
	AreaYN          string
	AreaName        string
}

func NewSeatsInfo(seats []string) SeatsInfo {
	return SeatsInfo{seats, "Y", "N", "N", ""}
}
func NewSeatsInfo2(seats []string, clickCloseBtnYN string, captchaYN string, areaYN string, areaName string) SeatsInfo {
	return SeatsInfo{seats, clickCloseBtnYN, captchaYN, areaYN, areaName}
}

type PlayDatePlaySeqInfo struct {
	PlayDate    string
	PlaySeq     string
	PlaySeqText string
}

type GoodsInfo struct {
	URL string
}

func NewGoodsInfo(url string) GoodsInfo {
	return GoodsInfo{url}
}

type Controller struct {
	selenium.WebDriver
	LoginInfo
	GoodsInfo
	PlayDatePlaySeqInfo
	SeatsInfo
	PriceInfo
	DeliveryInfo
	PaymentInfo
}

func NewPlayDatePlaySeqInfo(playDate string, playSeq string) PlayDatePlaySeqInfo {
	return PlayDatePlaySeqInfo{playDate, playSeq, ""}
}

func NewPlayDatePlaySeqInfo2(playDate string, playSeqText string) PlayDatePlaySeqInfo {
	return PlayDatePlaySeqInfo{playDate, "", playSeqText}
}

func NewController(webDriver selenium.WebDriver) Controller {
	return Controller{webDriver, LoginInfo{}, GoodsInfo{}, PlayDatePlaySeqInfo{}, SeatsInfo{}, PriceInfo{}, DeliveryInfo{}, PaymentInfo{}}
}

func NewController2(webDriver selenium.WebDriver, loginInfo LoginInfo) Controller {
	return Controller{webDriver, loginInfo, GoodsInfo{}, PlayDatePlaySeqInfo{}, SeatsInfo{}, PriceInfo{}, DeliveryInfo{}, PaymentInfo{}}
}

func NewController3(webDriver selenium.WebDriver, loginInfo LoginInfo, goodsInfo GoodsInfo, playDatePlaySeqInfo PlayDatePlaySeqInfo, seatsInfo SeatsInfo, priceInfo PriceInfo, deliveryInfo DeliveryInfo, paymentInfo PaymentInfo) Controller {
	return Controller{webDriver, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo, deliveryInfo, paymentInfo}
}

func (c *Controller) Navigate(url string) error {
	return c.WebDriver.Get(url)
}

func (c *Controller) GotoGoodsInfoPage() error {
	var err error

	if c.LoginInfo.SiteName == "sac.or.kr" {
		// 티켓오픈시까지 새로고침
		isOpen := false
		for !isOpen {
			if err = c.WebDriver.Get(c.GoodsInfo.URL); err != nil {
				panic(err)
			}
			// <a href="javascript:;" onclick="fn_ticket('10017373');" class="btn btn_red_full">예매하기</a>
			if _, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[contains(text(), '예매하기')]"); err != nil {
				isOpen = false
			} else {
				isOpen = true
			}

		}
	}
	if c.LoginInfo.SiteName == "ticket.interpark.com" {
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
	}
	return nil
}

func (c *Controller) SelectPlayDatePlaySeq() error {
	var condition selenium.Condition
	var webElement selenium.WebElement
	var err error

	if c.LoginInfo.SiteName == "sac.or.kr" {
		// scroll for clicking
		if _, err = c.WebDriver.ExecuteScript("window.scrollBy(0,250)", nil); err != nil {
			panic(err)
		}

		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[contains(@class, 'btn_red_full')]"); err != nil {
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

		var windowHandles []string

		if windowHandles, err = c.WebDriver.WindowHandles(); err != nil {
			panic(err)
		}
		if err = c.WebDriver.SwitchWindow(windowHandles[1]); err != nil {
			panic(err)
		}

		condition = func(wd selenium.WebDriver) (bool, error) {
			if err = c.WebDriver.SwitchFrame(nil); err != nil {
				//panic(err)
				return false, nil
			}
			if webElement, err = c.WebDriver.FindElement(selenium.ByID, "ifrmBookStep"); err != nil {
				//panic(err)
				return false, nil
			}
			if err = c.WebDriver.SwitchFrame(webElement); err != nil {
				//panic(err)
				return false, nil
			}

			// <a id="CellPlayDate" name="CellPlayDate" class="sel1" href="#;" onclick="fnSelectPlayDate(0, '20200828')">28<span class="blind">일 예매 가능</span></a>
			if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[@id='CellPlayDate' and contains(@onclick, '"+c.PlayDatePlaySeqInfo.PlayDate+"')]"); err != nil {
				//panic(err)
				return false, nil
			}
			if err := webElement.Click(); err != nil {
				//panic(err)
				return false, nil
			}

			// <a id="CellPlaySeq" name="CellPlaySeq" class="sel" href="#;" onclick="fnSelectPlaySeq(0, '001', '20200824', '20시 00분 ', 'N', '202008272359')">20시 00분 </a>
			if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//a[@id='CellPlaySeq' and contains(@onclick, '"+c.PlayDatePlaySeqInfo.PlaySeqText+"')]"); err != nil {
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

		// <img src="//ticketimage.interpark.com/TicketImage/onestop/btn_next_on.gif" alt="다음단계" id="LargeNextBtnImage">
		// 주의 : 위의 경우 anchor가 아닌 image를 클릭해야 한다.
		if err = c.WebDriver.SwitchFrame(nil); err != nil {
			panic(err)
		}
		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//img[@id='LargeNextBtnImage']"); err != nil {
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

	if c.LoginInfo.SiteName == "ticket.interpark.com" {
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

		if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//div[contains(@class, 'myValue')]"); err != nil {
			panic(err)
		}
		if err := webElement.Click(); err != nil {
			panic(err)
		}

		if c.PlayDatePlaySeqInfo.PlaySeqText != "" {
			condition = func(wd selenium.WebDriver) (bool, error) {
				if webElement, err = wd.FindElement(selenium.ByXPATH, "//ul[@id='ulPlaySeq']//label[contains(text(),'"+c.PlayDatePlaySeqInfo.PlaySeqText+"')]"); err != nil {
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
		if c.PlayDatePlaySeqInfo.PlaySeq != "" {
			if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//ul[@id='ulPlaySeq']//label[contains(@onclick,'"+c.PlayDatePlaySeqInfo.PlaySeq+"')]"); err != nil {
				panic(err)
			}
			if err := webElement.Click(); err != nil {
				panic(err)
			}
		}

		// <a href="#" onclick="javascript:NoMemPrivacyCertify('','20003772');" class="btn_rev"><span>예매하기</span></a>
		if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//div[@class='tk_dt_btn_TArea']//a[@class='btn_rev']"); err != nil {
			panic(err)
		}
		if err := webElement.Click(); err != nil {
			panic(err)
		}
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

	if c.SeatsInfo.ClickCloseBtnYN == "Y" {
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
	}

	// <iframe id="ifrmSeat" name="ifrmSeat" scrolling="no" width="100%" height="100%" marginwidth="0" marginheight="0" frameborder="no" src="/Book/loading.html"></iframe>
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = wd.FindElement(selenium.ByID, "ifrmSeat"); err != nil {
			//panic(err)
			return false, nil
		}
		if err = c.WebDriver.SwitchFrame(webElement); err != nil {
			//panic(err)
			return false, nil
		}
		return true, nil
	}
	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	if c.SeatsInfo.CaptchaYN == "Y" {
		// <input type="text" id="txtCaptcha" name="txtCaptcha" value="" maxlength="8" onkeydown="IsEnterGo();" style="text-transform:uppercase;ime-mode:inactive;">
		condition = func(wd selenium.WebDriver) (bool, error) {

			if webElement, err = wd.FindElement(selenium.ByXPATH, "//div[contains(@class, 'validationTxt')]"); err != nil {
				//panic(err)
				return false, nil
			}
			if err := webElement.Click(); err != nil {
				//panic(err)
				return false, nil
			}
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//input[@id='txtCaptcha']"); err != nil {
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

		// 6자리 입력 대기
		condition = func(wd selenium.WebDriver) (bool, error) {

			var activeElement selenium.WebElement
			if activeElement, err = wd.ActiveElement(); err != nil {
				//panic(err)
				return false, nil
			}
			var idString string
			if idString, err = activeElement.GetAttribute("id"); err != nil {
				//panic(err)
				return false, nil
			}

			if webElement, err = wd.FindElement(selenium.ByXPATH, "//input[@id='txtCaptcha']"); err != nil {
				//panic(err)
				return false, nil
			}

			if idString != "txtCaptcha" {
				if webElement, err = wd.FindElement(selenium.ByXPATH, "//div[contains(@class, 'validationTxt')]"); err != nil {
					//panic(err)
					return false, nil
				}
				if err := webElement.Click(); err != nil {
					//panic(err)
					return false, nil
				}
				if webElement, err = wd.FindElement(selenium.ByXPATH, "//input[@id='txtCaptcha']"); err != nil {
					//panic(err)
					return false, nil
				}
				if err := webElement.Click(); err != nil {
					//panic(err)
					return false, nil
				}

				return false, nil
			}

			var valueString string
			if valueString, err = webElement.GetAttribute("value"); err != nil {
				//panic(err)
				return false, nil
			}

			if len(valueString) < 6 {
				fmt.Println(valueString)
				return false, nil
			}

			// 입력완료
			// <a href="javascript:;" onclick="fnCheck()">입력완료</a>
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//a[text()='입력완료']"); err != nil {
				//panic(err)
				return false, nil
			}
			if err := webElement.Click(); err != nil {
				//panic(err)
				return false, nil
			}
			// <div id="divRecaptcha" class="capchaLayer" style=""></div>
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//div[@id='divRecaptcha']"); err != nil {
				//panic(err)
				return false, nil
			}
			var isDisplayed bool
			if isDisplayed, err = webElement.IsDisplayed(); err != nil {
				//panic(err)
				return false, nil
			}
			if isDisplayed {

				if webElement, err = wd.FindElement(selenium.ByXPATH, "//input[@id='txtCaptcha']"); err != nil {
					//panic(err)
					return false, nil
				}
				if err := webElement.Click(); err != nil {
					//panic(err)
					return false, nil
				}
				return false, nil
			}
			return true, nil
		}
		if err = c.WebDriver.Wait(condition); err != nil {
			panic(err)
		}

	}

	// <iframe id="ifrmSeatDetail" name="ifrmSeatDetail" scrolling="auto" width="658px" height="619px" marginwidth="0" marginheight="0" frameborder="no" src=""></iframe>
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = wd.FindElement(selenium.ByID, "ifrmSeatDetail"); err != nil {
			//panic(err)
			return false, nil
		}
		if err = c.WebDriver.SwitchFrame(webElement); err != nil {
			//panic(err)
			return false, nil
		}
		return true, nil
	}

	if err = c.WebDriver.Wait(condition); err != nil {
		panic(err)
	}

	if c.SeatsInfo.AreaYN == "Y" {
		// <area shape="rect" coords="446,425,591,447" onfocus="this.blur()" href="javascript:GetBlockSeatList('', '', '306')" onmouseover="javascript:EventBlockOver(this, '306')" onmouseout="javascript:EventBlockOut(this, '306')">

		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//area[contains(@href, '"+c.SeatsInfo.AreaName+"')]"); err != nil {
				//panic(err)
				return false, nil
			}
			if err = webElement.Click(); err != nil {
				//panic(err)
				return false, nil
			}
			return true, nil
		}
		if err = c.WebDriver.Wait(condition); err != nil {
			panic(err)
		}

		// <td bgcolor="#EBEBEB"><img src="//ticketimage.interpark.com/TicketImage/event/061227/dot_03.gif" width="5" height="5" align="absmiddle"> <b><font color="#3300FF">306 영역</font>의 좌석배치도입니다</b></td>

		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//font[@color='#3300FF' and text()='"+c.SeatsInfo.AreaName+" 영역']"); err != nil {
				//panic(err)
				return false, nil
			}
			return true, nil
		}
		if err = c.WebDriver.Wait(condition); err != nil {
			panic(err)
		}

	}

	// <img src="http://ticketimage.interpark.com/TMGSNAS/TMGS/G/1_90.gif" class="stySeat" style="left:335 ;top:241" alt="[VIP석] 1층-B구역14열-23" title="[VIP석] 1층-B구역14열-23" onclick="javascript: SelectSeat('SID49', '1', '1층', 'B구역14열', '23', '002')">
	for _, seat := range c.Seats {
		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//img[@alt='"+seat+"'] | //span[@title='"+seat+"']"); err != nil {
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
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//select[@seatgradename='"+priceItem.SeatGradeName+"' and @pricegradename='"+priceItem.PriceGradeName+"']"); err != nil {
				//panic(err)
				return false, nil
			}
			if err := webElement.Click(); err != nil {
				//panic(err)
				return false, nil
			}
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//select[@seatgradename='"+priceItem.SeatGradeName+"' and @pricegradename='"+priceItem.PriceGradeName+"']//option[@value='"+priceItem.SeatCount+"']"); err != nil {
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

	// <input type="radio" class="chk" id="Delivery" name="Delivery" value="24000" onclick="fnChange()">
	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//input[@id='Delivery' and @value='"+c.DeliveryInfo.DeliveryType+"']"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	// <input type="text" id="YYMMDD" name="YYMMDD" style="width:45px;" class="txt1" maxlength="6" onkeyup="fnMoveFocus(6, 'YYMMDD', 'HpNo1');">
	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//input[@id='YYMMDD'] | //input[@id='SSN1']"); err != nil {
		panic(err)
	}
	if err := webElement.SendKeys(c.DeliveryInfo.YYMMDD); err != nil {
		panic(err)
	}

	if c.LoginInfo.SiteName == "sac.or.kr" {
		// <input type="text" id="Email" value="" style="width:170px;" class="txt1">
		if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//input[@id='Email']"); err != nil {
			panic(err)
		}
		if err := webElement.SendKeys(c.DeliveryInfo.EMail); err != nil {
			panic(err)
		}

	}

	if c.DeliveryInfo.DeliveryType == "24001" {
		// <input type="checkbox" id="chkSyncAddress" onclick="javascript:fnSyncAddress()">
		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//input[@id='chkSyncAddress']"); err != nil {
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

func (c *Controller) SelectPayment() error {
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

	condition = func(wd selenium.WebDriver) (bool, error) {
		// <input type="radio" class="chk" name="Payment" value="22003" kindofsettle="22003" onclick="fnCheckPayment(this);">
		if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//input[@name='Payment' and @value='"+c.PaymentInfo.PaymentType+"']"); err != nil {
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

	if c.PaymentInfo.PaymentType == "22004" {
		// <select id="BankCode" onchange="fnSelectBankCode(this.value)"><option value="">입금하실 은행을 선택하세요.</option><option value="38052">농협(중앙)</option></select>
		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//select[@id='BankCode']//option[text()='"+c.PaymentInfo.BankName+"']"); err != nil {
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

	if c.PaymentInfo.PaymentType == "22003" {
		// <input type="radio" class="chk" name="PaymentSelect" id="PaymentSelect" value="C1" kindofsettle="22003" kindofsettledetail="12001" onclick="fnPaymentSelect(this.value);">
		if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//input[@name='PaymentSelect' and @value='"+c.PaymentInfo.PaymentSelect+"']"); err != nil {
			panic(err)
		}
		if err := webElement.Click(); err != nil {
			panic(err)
		}

		// <select id="DiscountCard" onchange="fnCardSelect(this.value);"><option value="">카드종류를 선택하세요.</option><option value="62">KB국민카드</option></select>

		condition = func(wd selenium.WebDriver) (bool, error) {
			if webElement, err = wd.FindElement(selenium.ByXPATH, "//select[@id='DiscountCard']//option[text()='"+c.PaymentInfo.DiscountCard+"']"); err != nil {
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

	// <a href="javascript:fnNextStep('P');" id="SmallNextBtnLink" onfocus="this.blur();"><img src="http://ticketimage.interpark.com/TicketImage/onestop/btn_next_02.gif" alt="다음단계" id="SmallNextBtnImage"></a>
	// 주의 : 위의 경우 anchor가 아닌 image를 클릭해야 한다.
	if err = c.WebDriver.SwitchFrame(nil); err != nil {
		panic(err)
	}
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = wd.FindElement(selenium.ByXPATH, "//img[@id='SmallNextBtnImage']"); err != nil {
			//panic(err)
			return false, nil
		}
		// 결제오류 때문에 Sleep
		time.Sleep(2 * time.Second)

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

func (c *Controller) DoPay() error {
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

	// <input id="CancelAgree" onclick="fnCheckCancelAgree()" type="checkbox">
	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//input[@id='CancelAgree']"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	// <input id="CancelAgree2" onclick="fnCheckCancelAgree()" type="checkbox">
	if webElement, err = c.WebDriver.FindElement(selenium.ByXPATH, "//input[@id='CancelAgree2']"); err != nil {
		panic(err)
	}
	if err := webElement.Click(); err != nil {
		panic(err)
	}

	// <a href="javascript:fnNextStep('P');" id="LargeNextBtnLink" onfocus="this.blur();"><img src="http://ticketimage.interpark.com/TicketImage/onestop/btn_buy.gif" alt="다음단계" id="LargeNextBtnImage"> </a>
	// 주의 : 위의 경우 anchor가 아닌 image를 클릭해야 한다.
	if err = c.WebDriver.SwitchFrame(nil); err != nil {
		panic(err)
	}
	condition = func(wd selenium.WebDriver) (bool, error) {
		if webElement, err = wd.FindElement(selenium.ByXPATH, "//img[@id='LargeNextBtnImage']"); err != nil {
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

func (c *Controller) Login() error {

	var err error
	var webElement selenium.WebElement
	var title string
	var currentUrl string
	var condition selenium.Condition

	if c.LoginInfo.SiteName == "sac.or.kr" {
		if err = c.WebDriver.Get("http://sac.or.kr/"); err != nil {
			panic(err)
		}

		condition = func(wd selenium.WebDriver) (bool, error) {

			// <a id="login_state" href="http://www.sac.or.kr/SacHome/login/login">로그인</a>
			if webElement, err = c.WebDriver.FindElement(selenium.ByID, "login_state"); err != nil {
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

		// waiting for https://www.sac.or.kr/SacHome/login/login
		condition = func(wd selenium.WebDriver) (bool, error) {
			// <input type="text" id="uid" name="name" value="" placeholder="띄어쓰기 없이 입력">
			if webElement, err = c.WebDriver.FindElement(selenium.ByID, "uid"); err != nil {
				//panic(err)
				return false, nil
			}
			if err = webElement.SendKeys(c.LoginInfo.ID); err != nil {
				//panic(err)
				return false, nil
			}
			// <input type="password" id="upass" name="password" value="" placeholder="띄어쓰기 없이 입력">
			if webElement, err = c.WebDriver.FindElement(selenium.ByID, "upass"); err != nil {
				//panic(err)
				return false, nil
			}
			if err = webElement.SendKeys(c.LoginInfo.PWD); err != nil {
				//panic(err)
				return false, nil
			}
			return true, nil
		}
		if err = c.WebDriver.Wait(condition); err != nil {
			panic(err)
		}
		// <a href="#" class="st1" id="login"><!-- 회원  -->로그인</a>
		if webElement, err = c.WebDriver.FindElement(selenium.ByID, "login"); err != nil {
			panic(err)
		}
		if err = webElement.Click(); err != nil {
			panic(err)
		}

		// waiting for finished
		condition = func(wd selenium.WebDriver) (bool, error) {
			if err = c.WebDriver.AcceptAlert(); err != nil {
				//panic(err)
				return false, nil
			}
			return true, nil
		}
		if err = c.WebDriver.Wait(condition); err != nil {
			panic(err)
		}

		condition = func(wd selenium.WebDriver) (bool, error) {
			if currentUrl, err = wd.CurrentURL(); err != nil {
				//panic(err)
				return false, nil
			}
			if currentUrl != "http://www.sac.or.kr/SacHome/sachome/main" {
				return false, nil
			}
			return true, nil
		}
		if err = c.WebDriver.Wait(condition); err != nil {
			panic(err)
		}
	}
	if c.LoginInfo.SiteName == "ticket.interpark.com" {
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
	}
	return nil
}
