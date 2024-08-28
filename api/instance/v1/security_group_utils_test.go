package instance_test

import (
	"net"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestAPI_UpdateSecurityGroup(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("security-group-test")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := instance.NewAPI(client)

	createResponse, err := instanceAPI.CreateSecurityGroup(&instance.CreateSecurityGroupRequest{
		Name:                  "name",
		Description:           "description",
		Stateful:              true,
		InboundDefaultPolicy:  instance.SecurityGroupPolicyAccept,
		OutboundDefaultPolicy: instance.SecurityGroupPolicyDrop,
	})

	testhelpers.AssertNoError(t, err)

	accept := instance.SecurityGroupPolicyAccept
	drop := instance.SecurityGroupPolicyDrop

	updateResponse, err := instanceAPI.UpdateSecurityGroup(&instance.UpdateSecurityGroupRequest{
		SecurityGroupID:       createResponse.SecurityGroup.ID,
		Name:                  scw.StringPtr("new_name"),
		Description:           scw.StringPtr("new_description"),
		Stateful:              scw.BoolPtr(false),
		InboundDefaultPolicy:  drop,
		OutboundDefaultPolicy: accept,
		// Keep false here, switch it to true is too dangerous for the one who update the test cassette.
		ProjectDefault: scw.BoolPtr(false),
		Tags:           scw.StringsPtr([]string{"foo", "bar"}),
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, "new_name", updateResponse.SecurityGroup.Name)
	testhelpers.Equals(t, "new_description", updateResponse.SecurityGroup.Description)
	testhelpers.Equals(t, instance.SecurityGroupPolicyDrop, updateResponse.SecurityGroup.InboundDefaultPolicy)
	testhelpers.Equals(t, instance.SecurityGroupPolicyAccept, updateResponse.SecurityGroup.OutboundDefaultPolicy)
	testhelpers.Equals(t, false, updateResponse.SecurityGroup.Stateful)
	testhelpers.Equals(t, false, updateResponse.SecurityGroup.ProjectDefault)
	testhelpers.Equals(t, false, *updateResponse.SecurityGroup.OrganizationDefault)
	testhelpers.Equals(t, []string{"foo", "bar"}, updateResponse.SecurityGroup.Tags)

	err = instanceAPI.DeleteSecurityGroup(&instance.DeleteSecurityGroupRequest{
		SecurityGroupID: createResponse.SecurityGroup.ID,
	})
	testhelpers.AssertNoError(t, err)
}

func TestAPI_UpdateSecurityGroupRule(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("security-group-rule-test")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := instance.NewAPI(client)

	bootstrap := func(t *testing.T) (*instance.SecurityGroup, *instance.SecurityGroupRule, func()) {
		t.Helper()
		createSecurityGroupResponse, err := instanceAPI.CreateSecurityGroup(&instance.CreateSecurityGroupRequest{
			Name:                  "name",
			Description:           "description",
			Stateful:              true,
			InboundDefaultPolicy:  instance.SecurityGroupPolicyAccept,
			OutboundDefaultPolicy: instance.SecurityGroupPolicyDrop,
		})

		testhelpers.AssertNoError(t, err)
		_, ipNet, _ := net.ParseCIDR("8.8.8.8/32")
		createRuleResponse, err := instanceAPI.CreateSecurityGroupRule(&instance.CreateSecurityGroupRuleRequest{
			SecurityGroupID: createSecurityGroupResponse.SecurityGroup.ID,
			Direction:       instance.SecurityGroupRuleDirectionInbound,
			Protocol:        instance.SecurityGroupRuleProtocolTCP,
			DestPortFrom:    scw.Uint32Ptr(1),
			DestPortTo:      scw.Uint32Ptr(1024),
			IPRange:         scw.IPNet{IPNet: *ipNet},
			Action:          instance.SecurityGroupRuleActionAccept,
			Position:        1,
		})
		testhelpers.AssertNoError(t, err)

		return createSecurityGroupResponse.SecurityGroup, createRuleResponse.Rule, func() {
			err = instanceAPI.DeleteSecurityGroup(&instance.DeleteSecurityGroupRequest{
				SecurityGroupID: createSecurityGroupResponse.SecurityGroup.ID,
			})
			testhelpers.AssertNoError(t, err)
		}
	}

	t.Run("Simple update", func(t *testing.T) {
		group, rule, cleanUp := bootstrap(t)
		defer cleanUp()

		_, ipNet, _ := net.ParseCIDR("1.1.1.1/32")
		updateResponse, err := instanceAPI.UpdateSecurityGroupRule(&instance.UpdateSecurityGroupRuleRequest{
			SecurityGroupID:     group.ID,
			SecurityGroupRuleID: rule.ID,
			Action:              instance.SecurityGroupRuleActionDrop,
			IPRange:             &scw.IPNet{IPNet: *ipNet},
			DestPortFrom:        scw.Uint32Ptr(1),
			DestPortTo:          scw.Uint32Ptr(2048),
			Protocol:            instance.SecurityGroupRuleProtocolUDP,
			Direction:           instance.SecurityGroupRuleDirectionOutbound,
		})

		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, instance.SecurityGroupRuleActionDrop, updateResponse.Rule.Action)
		testhelpers.Equals(t, scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(0x1, 0x1, 0x1, 0x1), Mask: net.IPMask{0xff, 0xff, 0xff, 0xff}}}, updateResponse.Rule.IPRange)
		testhelpers.Equals(t, scw.Uint32Ptr(1), updateResponse.Rule.DestPortFrom)
		testhelpers.Equals(t, scw.Uint32Ptr(2048), updateResponse.Rule.DestPortTo)
		testhelpers.Equals(t, instance.SecurityGroupRuleProtocolUDP, updateResponse.Rule.Protocol)
		testhelpers.Equals(t, instance.SecurityGroupRuleDirectionOutbound, updateResponse.Rule.Direction)
	})

	t.Run("From a port range to a single port", func(t *testing.T) {
		group, rule, cleanUp := bootstrap(t)
		defer cleanUp()

		_, ipNet, _ := net.ParseCIDR("1.1.1.1/32")
		updateResponse, err := instanceAPI.UpdateSecurityGroupRule(&instance.UpdateSecurityGroupRuleRequest{
			SecurityGroupID:     group.ID,
			SecurityGroupRuleID: rule.ID,
			Action:              instance.SecurityGroupRuleActionDrop,
			IPRange:             &scw.IPNet{IPNet: *ipNet},
			DestPortFrom:        scw.Uint32Ptr(22),
			DestPortTo:          scw.Uint32Ptr(22),
			Protocol:            instance.SecurityGroupRuleProtocolUDP,
			Direction:           instance.SecurityGroupRuleDirectionOutbound,
		})

		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, instance.SecurityGroupRuleActionDrop, updateResponse.Rule.Action)
		testhelpers.Equals(t, scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(0x1, 0x1, 0x1, 0x1), Mask: net.IPMask{0xff, 0xff, 0xff, 0xff}}}, updateResponse.Rule.IPRange)
		testhelpers.Equals(t, uint32(22), *updateResponse.Rule.DestPortFrom)
		testhelpers.Equals(t, (*uint32)(nil), updateResponse.Rule.DestPortTo)
		testhelpers.Equals(t, instance.SecurityGroupRuleProtocolUDP, updateResponse.Rule.Protocol)
		testhelpers.Equals(t, instance.SecurityGroupRuleDirectionOutbound, updateResponse.Rule.Direction)
	})

	t.Run("Switching to ICMP", func(t *testing.T) {
		group, rule, cleanUp := bootstrap(t)
		defer cleanUp()

		updateResponse, err := instanceAPI.UpdateSecurityGroupRule(&instance.UpdateSecurityGroupRuleRequest{
			SecurityGroupID:     group.ID,
			SecurityGroupRuleID: rule.ID,
			Protocol:            instance.SecurityGroupRuleProtocolICMP,
		})

		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, instance.SecurityGroupRuleActionAccept, updateResponse.Rule.Action)
		testhelpers.Equals(t, scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(0x8, 0x8, 0x8, 0x8), Mask: net.IPMask{0xff, 0xff, 0xff, 0xff}}}, updateResponse.Rule.IPRange)
		testhelpers.Equals(t, (*uint32)(nil), updateResponse.Rule.DestPortFrom)
		testhelpers.Equals(t, (*uint32)(nil), updateResponse.Rule.DestPortTo)
		testhelpers.Equals(t, instance.SecurityGroupRuleProtocolICMP, updateResponse.Rule.Protocol)
		testhelpers.Equals(t, instance.SecurityGroupRuleDirectionInbound, updateResponse.Rule.Direction)
	})

	t.Run("Remove ports", func(t *testing.T) {
		group, rule, cleanUp := bootstrap(t)
		defer cleanUp()

		updateResponse, err := instanceAPI.UpdateSecurityGroupRule(&instance.UpdateSecurityGroupRuleRequest{
			SecurityGroupID:     group.ID,
			SecurityGroupRuleID: rule.ID,
			DestPortFrom:        scw.Uint32Ptr(0),
			DestPortTo:          scw.Uint32Ptr(0),
		})

		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, instance.SecurityGroupRuleActionAccept, updateResponse.Rule.Action)
		testhelpers.Equals(t, scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(0x8, 0x8, 0x8, 0x8), Mask: net.IPMask{0xff, 0xff, 0xff, 0xff}}}, updateResponse.Rule.IPRange)
		testhelpers.Equals(t, (*uint32)(nil), updateResponse.Rule.DestPortFrom)
		testhelpers.Equals(t, (*uint32)(nil), updateResponse.Rule.DestPortTo)
		testhelpers.Equals(t, instance.SecurityGroupRuleProtocolTCP, updateResponse.Rule.Protocol)
		testhelpers.Equals(t, instance.SecurityGroupRuleDirectionInbound, updateResponse.Rule.Direction)
	})
}
