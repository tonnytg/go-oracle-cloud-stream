package main

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/identity"
)

func main() {
	fmt.Println("Start Kafka Stream connect")

	configProvider := common.DefaultConfigProvider()
	fmt.Println(configProvider)

	client, err := identity.NewIdentityClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Println("ClientError:", err)
		return
	}

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	if err != nil {
		fmt.Println("TanancyError:", err)
		return
	}

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := client.ListAvailabilityDomains(context.Background(), request)
	if err != nil {
		fmt.Println("ListDomainsError:", err)
		return
	}

	fmt.Printf("List of available domains: %v", r.Items)
}
