
package main
import(
 "io/ioutil" 
//  "log"
 "net/http"
 "encoding/json"
  "fmt"
)
const BTC = "Bitcoin"
const RUB = "RUB"
const USD = "USD"
const BitPay = "https://bitpay.com/api/rates"
const myPocket = 0.01535133
const russianRUBtoUSD = "https://free.currconv.com/api/v7/convert?q=USD_RUB&compact=ultra&apiKey=9f066153493cf78a8560"

type RatesResponse []Rate

type Rate struct {
// type Response struct {
    Code    string `json:"code"`
    Name    string `json:"name"`
    Rate    float32  `json:"rate"`
}

type RUBtoUSD struct{
  value float32 "json: USD_RUB"
}

func getRubleForCalculation(dataForManipulations RatesResponse, Cur string)float32{
  var rates float32
  for _,value := range dataForManipulations{
    if value.Code == Cur{
      rates = value.Rate
      break
    }
  }
  return rates
}

func calculateInMyPOcket(rubCourse float32 )int{
  return int(rubCourse * myPocket)
}

func makeResponseFromMarket() RatesResponse{
  var curRate RatesResponse
  // TODO add coinmarketcap for referense
  resp, err := http.Get(BitPay)
  if err != nil {
      fmt.Println("No response from request")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body) // response body is []byte
  errs := json.Unmarshal([]byte(body), &curRate)
  if errs != nil {
      panic(errs)
  }
  return curRate
}

func rubleMarket() float32{
  var curRUB float32
  // TODO add coinmarketcap for referense
  resp, err := http.Get(russianRUBtoUSD)
  if err != nil {
      fmt.Println("No response from request")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body) // response body is []byte
  errs := json.Unmarshal([]byte(body), &curRUB)
  if errs != nil {
      panic(errs)
  }
  return curRUB
}

func main() {
  rates:= makeResponseFromMarket()
  rubCourse := getRubleForCalculation(rates, RUB)
  usdCourse :=  getRubleForCalculation(rates, USD)
  rub := rubleMarket() 
  fmt.Println("RUB to USD", rub)
  fmt.Println("russian course is", rubCourse)
  fmt.Println("usd course is", usdCourse)
  fmt.Println("current in my pocket is", RUB , "=", calculateInMyPOcket(rubCourse), "or", USD ,"=",
   calculateInMyPOcket(usdCourse))
}