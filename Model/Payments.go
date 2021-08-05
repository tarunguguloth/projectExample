package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"sync"
)
var Mutex sync.Mutex

func AddAmount(payment *Payments, id string, amount string)(err error){
	Mutex.Lock()
	Amount, err := strconv.Atoi(amount)
	if(err!=nil){
		return err
	}else{
	payment.CurrentBalance += Amount
		return nil
	}
	Mutex.Unlock()

	return nil
}
/*
func AddAmount(payment *Payments, id string, amount int)(err error){

	Mutex.Lock()
		response, err := http.Get("https://api.razorpay.com/v1/payment_links")
		if err != nil {
			return err
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	    Mutex.Unlock()
}
 */

func WithdrawAmount(payment *Payments, id string )(err error){
	Mutex.Lock()
	Amount, err := strconv.Atoi(amount)
	if(err!=nil){
		return err
	}else {
		if (payment.CurrentBalance < Amount) {
			return err
		} else {
			payment.CurrentBalance -= Amount
			return nil
		}
	}
	Mutex.Unlock()

	return nil
}



/*function which defines how to connect to the another api
func main() {
	fmt.Println("Starting the application...")
	response, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("Terminating the application...")
}

*/

