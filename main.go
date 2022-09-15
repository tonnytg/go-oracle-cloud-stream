package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/streaming"
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

func putMsgInStream(streamEndpoint string, streamOcid string) {
	fmt.Println("Stream endpoint for put msg api is: " + streamEndpoint)

	provider, err := common.ConfigurationProviderFromFileWithProfile(ociConfigFilePath, ociProfileName, "")
	helpers.FatalIfError(err)

	streamClient, err := streaming.NewStreamClientWithConfigurationProvider(provider, streamEndpoint)
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).
	for i := 0; i < 5; i++ {
		putMsgReq := streaming.PutMessagesRequest{StreamId: common.String(streamOcid),
			PutMessagesDetails: streaming.PutMessagesDetails{
				// we are batching 2 messages for each Put Request
				Messages: []streaming.PutMessagesDetailsEntry{
					{Key: []byte("key dummy-0-" + strconv.Itoa(i)),
						Value: []byte("value dummy-" + strconv.Itoa(i))},
					{Key: []byte("key dummy-1-" + strconv.Itoa(i)),
						Value: []byte("value dummy-" + strconv.Itoa(i))}}},
		}

		// Send the request using the service client
		putMsgResp, err := streamClient.PutMessages(context.Background(), putMsgReq)
		helpers.FatalIfError(err)

		// Retrieve value from the response.
		fmt.Println(putMsgResp)
	}

}
