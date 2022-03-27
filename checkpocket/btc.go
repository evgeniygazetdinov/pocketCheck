
package main
import(
 "io/ioutil" 
//  "log"
 "net/http"
 "encoding/json"
  "fmt"
)
const BTC = "Bitcoin"
const RUB = "Russian Ruble"
const BitPay = "https://bitpay.com/api/rates"
const myPocket = 0.01535133

type RatesResponse []Rate
type Rate struct {
// type Response struct {
    Code    string `json:"code"`
    Name    string `json:"name"`
    Rate    float32  `json:"rate"`
}

func getRubleForCalculation(dataForManipulations RatesResponse)float32{
  var rubCourse float32
  for _,value := range dataForManipulations{
    if value.Name == RUB{
      rubCourse = value.Rate
    }
  }
 return rubCourse
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

func main() {
  rates := makeResponseFromMarket()
  rubCourse := getRubleForCalculation(rates) 
  fmt.Println("rusian course is", rubCourse)
  fmt.Println("currnet in my pocket is ", calculateInMyPOcket(rubCourse))
}