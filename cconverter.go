package currency_converter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func Convert(currency string, value float64) (string, float64) {
	resp, err := http.Get("http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		fmt.Errorf("GET error: %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Status error: %v", resp.StatusCode)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Read body: %v", err)
		os.Exit(1)
	}
	re := regexp.MustCompile("RUB.+rate='([0-9\\.]+)'")
	matches := re.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		fmt.Errorf("No RUB currency")
		os.Exit(1)
	}
	rate, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		fmt.Errorf("Parse error. String to float")
		os.Exit(1)
	}
	if currency == "EUR" {
		return "RUB", value * rate
	}
	return "EUR", value / rate
}
