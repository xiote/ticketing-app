package main

import (
	"bufio"
	"fmt"
	ip "github.com/xiote/ticketing-app/interpark"
	"os"

	"github.com/tebeka/selenium"
)

func main() {
	Example()
}

func Example() {

	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	selenium.SetDebug(true)

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// loginInfo := ip.LoginInfo{"xiote12", "gkswlsdn78#"}
	// goodsInfo := ip.GoodsInfo{"http://ticket.interpark.com/Ticket/Goods/GoodsInfo.asp?GroupCode=20003772"}
	// //playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo("20200822", "075")
	// playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20200822", "18시 30분")
	// seatsInfo := ip.NewSeatsInfo([]string{"[A석] 2층-C구역9열-31", "[A석] 2층-C구역9열-32"})
	// priceList := []ip.PriceItem{ip.PriceItem{"A석", "일반", "2"}}
	// priceInfo := ip.PriceInfo{priceList}
	// deliveryInfo := ip.DeliveryInfo{"24000", "781025"}
	// paymentInfo := ip.NewPaymentInfo("22003", "C1", "신한카드")

	// loginInfo := ip.LoginInfo{"xiote12", "gkswlsdn78#"}
	// goodsInfo := ip.GoodsInfo{"http://ticket.interpark.com/Ticket/Goods/GoodsInfo.asp?GroupCode=20005097"}
	// playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20200911", "19시 30분")
	// seatsInfo := ip.NewSeatsInfo2([]string{"[B석] 3층-동R구역 7열-50", "[B석] 3층-동R구역 7열-52"}, "N", "Y", "Y", "306")
	// priceList := []ip.PriceItem{ip.PriceItem{"B석", "일반", "2"}}
	// priceInfo := ip.PriceInfo{priceList}
	// deliveryInfo := ip.DeliveryInfo{"24001", "781025"}
	// paymentInfo := ip.NewPaymentInfo2("22004", "", "", "농협(중앙)")

	loginInfo := ip.LoginInfo{"xiote12", "gkswlsdn78#"}
	goodsInfo := ip.GoodsInfo{"http://ticket.interpark.com/Ticket/Goods/GoodsInfo.asp?GroupCode=20006380"}
	playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20201121", "19시 00분")
	seatsInfo := ip.NewSeatsInfo2([]string{"[VIP석] 객석1층-20열-21", "[VIP석] 객석1층-20열-22"}, "N", "N", "N", "")
	priceList := []ip.PriceItem{ip.PriceItem{"VIP석", "일반", "2"}}
	priceInfo := ip.PriceInfo{priceList}
	deliveryInfo := ip.DeliveryInfo{"24000", "781025"}
	paymentInfo := ip.NewPaymentInfo2("22004", "", "", "농협(중앙)")

	c := ip.NewController3(wd, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo, deliveryInfo, paymentInfo)

	if err := c.Login(); err != nil {
		panic(err)
	}
	var scanner *bufio.Scanner
	//fmt.Print("Press ENTER to continue")
	//scanner = bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	break
	//}

	if err := c.GotoGoodsInfoPage(); err != nil {
		panic(err)
	}
	if err := c.SelectPlayDatePlaySeq(); err != nil {
		panic(err)
	}
	if err := c.SelectSeats(); err != nil {
		panic(err)
	}
	if err := c.SelectPrice(); err != nil {
		panic(err)
	}
	if err := c.SelectDelivery(); err != nil {
		panic(err)
	}
	if err := c.SelectPayment(); err != nil {
		panic(err)
	}
	if err := c.DoPay(); err != nil {
		panic(err)
	}

	fmt.Print("Press ENTER to continue")
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}

}
