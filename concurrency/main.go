package main

import (
	"fmt"
	"io"

	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"http://management.azure.com/",
		"http://weather.sun.ac.za/apiInfo.php",
		"https://learn.microsoft.com/en-us/graph/overview",
		"https://cloud.google.com/?hl=es-419",
		"http://weaher.",
		"https://cloud.google.com/apis?hl=es",
		"https://www.deepl.com/translator",
		"https://si3.bcentral.cl/SieteRestWS/SieteRestWS.ashx?user=123456789&pass=tuPassword&firstdate=YYYY-MM-DD&lastdate=YYYY-MM-DD&timeseries=codigodeserie&function=GetSeries",
	}
	for _, api := range apis {
		go checkAPI(api)
		//checkAPI3()
	}
	elapsed := time.Since(start)
	fmt.Printf("Listo, tomó %v segundos \n ", elapsed.Seconds())
}

func checkAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: ¡%s está caido! \n", api)
	}
	fmt.Printf("SUCCESS: ¡%s está en funcionamiento! \n", api)

}

func checkAPI2() {
	api := "https://si3.bcentral.cl/SieteRestWS/SieteRestWS.ashx?user=123456789&pass=tuPassword&firstdate=YYYY-MM-DD&lastdate=YYYY-MM-DD&timeseries=codigodeserie&function=GetSeries"
	res, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: ¡%s está caido! \n", api)
	}
	fmt.Printf("SUCCESS: ¡%s está en funcionamiento! \n", api)
	fmt.Printf("############################################")
	response := res.StatusCode
	fmt.Printf("HTTP RESPONSE STATUS CODE : %v \n", response)
}
func checkAPI3() {
	api := "https://si3.bcentral.cl/SieteRestWS/SieteRestWS.ashx?user=admin&pass=admin&frequency=frecuenciaelegida&function=SearchSeries"
	res, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: ¡%s está caido! \n", api)
	}
	fmt.Printf("SUCCESS: ¡%s está en funcionamiento! \n", api)
	fmt.Printf("############################################")
	statusCode := res.StatusCode
	body, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	res.Body.Close()
	fmt.Printf("HTTP RESPONSE STATUS CODE : %v \n", statusCode)
	fmt.Printf("HTTP RESPONSE BODY : %v \n", string(body))
}
