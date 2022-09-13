package main

import (
	"fmt"

	"github.com/oracle/oci-go-sdk/v49/common"
)

func main() {
	fmt.Println("Start Kafka Stream connect")

	configProvider := common.DefaultConfigProvider()
	fmt.Println(configProvider)

}
