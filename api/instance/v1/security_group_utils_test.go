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

func TestAPI_UpdateSecurityGroupRule(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient("security-group-rule-test")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		zone         = scw.ZoneFrPar1
		organization = "b3ba839a-dcf2-4b0a-ac81-fc32370052a0"
	)

	bootstrap := func(t *testing.T) (*SecurityGroup, *SecurityGroupRule, func()) {

		createSecurityGroupResponse, err := instanceAPI.CreateSecurityGroup(&CreateSecurityGroupRequest{
			Zone:                  zone,
			Name:                  "name",
			Description:           "description",
			Stateful:              true,
			InboundDefaultPolicy:  SecurityGroupPolicyAccept,
			OutboundDefaultPolicy: SecurityGroupPolicyDrop,
			Organization:          organization,
		})

		testhelpers.Ok(t, err)

		createRuleResponse, err := instanceAPI.CreateSecurityGroupRule(&CreateSecurityGroupRuleRequest{
			Zone:            zone,
			SecurityGroupID: createSecurityGroupResponse.SecurityGroup.ID,
			Direction:       SecurityGroupRuleDirectionInbound,
			Protocol:        SecurityGroupRuleProtocolTCP,
			DestPortFrom:    scw.Uint32(1),
			DestPortTo:      scw.Uint32(1024),
			IPRange:         "8.8.8.8/32",
			Action:          SecurityGroupRuleActionAccept,
			Position:        1,
		})
		testhelpers.Ok(t, err)

		return createSecurityGroupResponse.SecurityGroup, createRuleResponse.Rule, func() {
			err = instanceAPI.DeleteSecurityGroup(&DeleteSecurityGroupRequest{
				Zone:            zone,
				SecurityGroupID: createSecurityGroupResponse.SecurityGroup.ID,
			})
			testhelpers.Ok(t, err)
		}
	}

	t.Run("Simple update", func(t *testing.T) {
		group, rule, cleanUp := bootstrap(t)
		defer cleanUp()

		action := SecurityGroupRuleActionDrop
		protocol := SecurityGroupRuleProtocolUDP
		direction := SecurityGroupRuleDirectionOutbound

		updateResponse, err := instanceAPI.UpdateSecurityGroupRule(&UpdateSecurityGroupRuleRequest{
			Zone:                zone,
			SecurityGroupID:     group.ID,
			SecurityGroupRuleID: rule.ID,
			Action:              &action,
			IPRange:             scw.String("1.1.1.1/32"),
			DestPortFrom:        scw.Uint32(1),
			DestPortTo:          scw.Uint32(2048),
			Protocol:            &protocol,
			Direction:           &direction,
		})

		testhelpers.Ok(t, err)
		testhelpers.Equals(t, SecurityGroupRuleActionDrop, updateResponse.Rule.Action)
		testhelpers.Equals(t, "1.1.1.1", updateResponse.Rule.IPRange)
		testhelpers.Equals(t, scw.Uint32(1), updateResponse.Rule.DestPortFrom)
		testhelpers.Equals(t, scw.Uint32(2048), updateResponse.Rule.DestPortTo)
		testhelpers.Equals(t, SecurityGroupRuleProtocolUDP, updateResponse.Rule.Protocol)
		testhelpers.Equals(t, SecurityGroupRuleDirectionOutbound, updateResponse.Rule.Direction)
	})

	t.Run("From a port range to a single port", func(t *testing.T) {
		//t.Skip("Wait for instance team to fix this scenario in the API")
		group, rule, cleanUp := bootstrap(t)
		defer cleanUp()

		action := SecurityGroupRuleActionDrop
		protocol := SecurityGroupRuleProtocolUDP
		direction := SecurityGroupRuleDirectionOutbound

		updateResponse, err := instanceAPI.UpdateSecurityGroupRule(&UpdateSecurityGroupRuleRequest{
			Zone:                zone,
			SecurityGroupID:     group.ID,
			SecurityGroupRuleID: rule.ID,
			Action:              &action,
			IPRange:             scw.String("1.1.1.1/32"),
			DestPortFrom:        scw.Uint32(22),
			DestPortTo:          scw.Uint32(22),
			Protocol:            &protocol,
			Direction:           &direction,
		})

		testhelpers.Ok(t, err)
		testhelpers.Equals(t, SecurityGroupRuleActionDrop, updateResponse.Rule.Action)
		testhelpers.Equals(t, "1.1.1.1", updateResponse.Rule.IPRange)
		testhelpers.Equals(t, uint32(22), *updateResponse.Rule.DestPortFrom)
		testhelpers.Equals(t, (*uint32)(nil), updateResponse.Rule.DestPortTo)
		testhelpers.Equals(t, SecurityGroupRuleProtocolUDP, updateResponse.Rule.Protocol)
		testhelpers.Equals(t, SecurityGroupRuleDirectionOutbound, updateResponse.Rule.Direction)
	})

}
