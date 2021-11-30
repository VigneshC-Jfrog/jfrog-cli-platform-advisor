package inputs

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	commonutils "github.com/jfrog/jfrog-cli-platform-advisor/common"
)

type Config struct {
	XMLName  xml.Name `xml:"config"`
	Security Security `xml:"security"`
}

type Security struct {
	AnonAccess string `xml:"anonAccessEnabled"`
}

func GetConfig() {
	commonutils.NewRtCurlCommand()
	data, _ := ioutil.ReadFile("notes.xml")

	config := &Config{}

	_ = xml.Unmarshal([]byte(data), &config)

	fmt.Println(config.Security.AnonAccess)
}
