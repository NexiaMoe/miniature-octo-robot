package function

import (
	"encoding/json"
	"helloion/database"
	"helloion/entity"
	"helloion/filter"
	"helloion/utils"
	"log"

	"github.com/go-resty/resty/v2"
)

// GetIP is a function to get IP address from httpbin.org
func GetIP(f *filter.CLIFilter) (statusCode int, res string, err error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}

	ip := new(entity.IP)
	err = json.Unmarshal(resp.Body(), &ip)

	if err != nil {
		log.Fatal(err)
	}

	// Save to Log if status code is 200
	ip.StatusCode = resp.StatusCode()
	err = database.SaveLog(ip)
	utils.CheckErr(err)

	if resp.StatusCode() != 200 {
		return resp.StatusCode(), `Error : ` + resp.Status(), nil
	}

	if f.Json {
		return resp.StatusCode(), resp.String(), nil
	}

	return resp.StatusCode(), `Your IP address is : ` + ip.Origin, nil
}

// GetIPFromDB is a function to get IP address from database
func GetIPFromDB(f *filter.CLIFilter) (statusCode int, res string, err error) {
	ip, err := database.GetLog()
	utils.CheckErr(err)

	b, err := json.MarshalIndent(ip, "", "  ")

	utils.CheckErr(err)

	// Print JSON format

	return 200, string(b), nil

}
