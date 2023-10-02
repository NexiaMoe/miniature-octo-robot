package function_test

// Test function GetIP
// Path: function/rest_test.go
// Compare this snippet from function/rest_test.go:
// package function_test
//
import (
	"fmt"
	"helloion/database"
	"helloion/filter"
	"helloion/function"
	"helloion/utils"
	"testing"
)

func TestGetIP(t *testing.T) {
	config, err := utils.LoadConfig("..")
	utils.CheckErr(err)

	database.Init(&config)

	f := filter.CLIFilter{
		Json: true,
	}
	_, ip, err := function.GetIP(&f)

	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	if ip == "" {
		t.Error("IP is empty")
	}

	t.Log(ip)
}

// Test function GetIPFromDB

func TestGetIPFromDB(t *testing.T) {
	config, err := utils.LoadConfig("..")
	utils.CheckErr(err)

	database.Init(&config)

	f := filter.CLIFilter{
		Json: true,
	}
	statusCode, ip, err := function.GetIPFromDB(&f)

	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	if statusCode != 200 {
		t.Error("Status code is not 200")
	}

	if len(ip) == 0 {
		t.Error("IP is empty")
	}

	t.Log(ip)
}
