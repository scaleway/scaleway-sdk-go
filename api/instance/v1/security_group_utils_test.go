package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestAPI_UpdateSecurityGroup(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient("security-group-test")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		zone         = scw.ZoneFrPar1
		organization = "b3ba839a-dcf2-4b0a-ac81-fc32370052a0"
	)

	createResponse, err := instanceAPI.CreateSecurityGroup(&CreateSecurityGroupRequest{
		Zone:                  zone,
		Name:                  "name",
		Description:           "description",
		Stateful:              true,
		InboundDefaultPolicy:  SecurityGroupPolicyAccept,
		OutboundDefaultPolicy: SecurityGroupPolicyDrop,
		Organization:          organization,
	})

	testhelpers.Ok(t, err)

	accept := SecurityGroupPolicyAccept
	drop := SecurityGroupPolicyDrop
	f := false

	updateResponse, err := instanceAPI.UpdateSecurityGroup(&UpdateSecurityGroupRequest{
		Zone:                  zone,
		SecurityGroupID:       createResponse.SecurityGroup.ID,
		Name:                  scw.String("new_name"),
		Description:           scw.String("new_description"),
		Stateful:              &f,
		InboundDefaultPolicy:  &drop,
		OutboundDefaultPolicy: &accept,
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, "new_name", updateResponse.SecurityGroup.Name)
	testhelpers.Equals(t, "new_description", updateResponse.SecurityGroup.Description)
	testhelpers.Equals(t, SecurityGroupPolicyDrop, updateResponse.SecurityGroup.InboundDefaultPolicy)
	testhelpers.Equals(t, SecurityGroupPolicyAccept, updateResponse.SecurityGroup.OutboundDefaultPolicy)
	testhelpers.Equals(t, false, updateResponse.SecurityGroup.Stateful)

	err = instanceAPI.DeleteSecurityGroup(&DeleteSecurityGroupRequest{
		Zone:            zone,
		SecurityGroupID: createResponse.SecurityGroup.ID,
	})
	testhelpers.Ok(t, err)
}
