package main

import (
	"bufio"
	"fmt"
	ip "github.com/xiote/ticketing-app/interpark"
	"os"

	"github.com/tebeka/selenium"
)

func main() {

	DoWork()
	//TestAll()
}

func TestAll() {
	Test1() // 인터파크, 주의메세지, 현장수령, 카드

	Test2() // 인터파크, 캡챠, 영역, 배송, 무통장

	Test4() // 인터파크, 현장수령, 무통장

	Test3() // 예술의전당, 현장수령, 무통장
}

func DoWork() {

	var scanner *bufio.Scanner

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

	loginInfo := ip.NewLoginInfo("chmartha", "ch079577#")
	goodsInfo := ip.NewGoodsInfo("http://ticket.interpark.com/Ticket/Goods/GoodsInfo.asp?GroupCode=20006900")
	playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20200905", "17시 00분")
	seatsInfo := ip.NewSeatsInfo2([]string{"[R석] 1층-C블록10열-8"}, "N", "N", "N", "")
	priceList := []ip.PriceItem{ip.NewPriceItem("R석", "일반", "1")}
	priceInfo := ip.NewPriceInfo(priceList)
	deliveryInfo := ip.NewDeliveryInfo("24000", "771110")
	paymentInfo := ip.NewPaymentInfo2("22004", "", "", "농협(중앙)")

	c := ip.NewController3(wd, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo, deliveryInfo, paymentInfo)

	if err := c.Login(); err != nil {
		panic(err)
	}
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

func Test4() {

	var scanner *bufio.Scanner

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

	loginInfo := ip.NewLoginInfo("xiote12", "gkswlsdn78#")
	goodsInfo := ip.NewGoodsInfo("http://ticket.interpark.com/Ticket/Goods/GoodsInfo.asp?GroupCode=20006380")
	playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20201111", "20시 00분")
	seatsInfo := ip.NewSeatsInfo2([]string{"[R석] 객석1층-22열-42", "[R석] 객석1층-22열-43"}, "N", "N", "N", "")
	priceList := []ip.PriceItem{ip.NewPriceItem("R석", "일반", "2")}
	priceInfo := ip.NewPriceInfo(priceList)
	deliveryInfo := ip.NewDeliveryInfo("24000", "781025")
	paymentInfo := ip.NewPaymentInfo2("22004", "", "", "농협(중앙)")

	c := ip.NewController3(wd, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo, deliveryInfo, paymentInfo)

	if err := c.Login(); err != nil {
		panic(err)
	}
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

func Test3() {

	var scanner *bufio.Scanner

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

	loginInfo := ip.NewLoginInfo2("chmartha", "ch079577#", "sac.or.kr")
	goodsInfo := ip.NewGoodsInfo("http://www.sac.or.kr/SacHome/perform/detail?searchSeq=39139")
	playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20200828", "20시 00분")
	seatsInfo := ip.NewSeatsInfo2([]string{"[R석] 1층-C블록18열-1", "[R석] 1층-C블록18열-3"}, "N", "N", "N", "")
	priceList := []ip.PriceItem{ip.NewPriceItem("R석", "일반", "2")}
	priceInfo := ip.NewPriceInfo(priceList)
	deliveryInfo := ip.NewDeliveryInfo2("24000", "771110", "chmartha@naver.com")
	paymentInfo := ip.NewPaymentInfo2("22004", "", "", "농협(중앙)")

	c := ip.NewController3(wd, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo, deliveryInfo, paymentInfo)

	if err := c.Login(); err != nil {
		panic(err)
	}
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

func Test2() {

	var scanner *bufio.Scanner

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

	loginInfo := ip.NewLoginInfo("xiote12", "gkswlsdn78#")
	goodsInfo := ip.NewGoodsInfo("http://ticket.interpark.com/Ticket/Goods/GoodsInfo.asp?GroupCode=20005097")
	playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20200911", "19시 30분")
	seatsInfo := ip.NewSeatsInfo2([]string{"[B석] 3층-동R구역 7열-50", "[B석] 3층-동R구역 7열-52"}, "N", "Y", "Y", "306")
	priceList := []ip.PriceItem{ip.NewPriceItem("B석", "일반", "2")}
	priceInfo := ip.NewPriceInfo(priceList)
	deliveryInfo := ip.NewDeliveryInfo("24001", "781025")
	paymentInfo := ip.NewPaymentInfo2("22004", "", "", "농협(중앙)")

	c := ip.NewController3(wd, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo, deliveryInfo, paymentInfo)

	if err := c.Login(); err != nil {
		panic(err)
	}
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
func Test1() {
	var scanner *bufio.Scanner

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

	loginInfo := ip.NewLoginInfo("xiote12", "gkswlsdn78#")
	goodsInfo := ip.NewGoodsInfo("http://ticket.interpark.com/Ticket/Goods/GoodsInfo.asp?GroupCode=20003772")
	playDatePlaySeqInfo := ip.NewPlayDatePlaySeqInfo2("20200822", "18시 30분")
	seatsInfo := ip.NewSeatsInfo([]string{"[A석] 2층-C구역9열-31", "[A석] 2층-C구역9열-32"})
	priceList := []ip.PriceItem{ip.NewPriceItem("A석", "일반", "2")}
	priceInfo := ip.NewPriceInfo(priceList)
	deliveryInfo := ip.NewDeliveryInfo("24000", "781025")
	paymentInfo := ip.NewPaymentInfo("22003", "C1", "신한카드")

	c := ip.NewController3(wd, loginInfo, goodsInfo, playDatePlaySeqInfo, seatsInfo, priceInfo, deliveryInfo, paymentInfo)

	if err := c.Login(); err != nil {
		panic(err)
	}

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
