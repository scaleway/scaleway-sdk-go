package cockpit

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 5 * time.Second
	defaultTimeout       = 5 * time.Minute
)

// WaitForPreconfiguredAlertsRequest is used by WaitForPreconfiguredAlerts method.
type WaitForPreconfiguredAlertsRequest struct {
	Region             scw.Region
	ProjectID          string
	PreconfiguredRules []string
	TargetStatus       AlertStatus
	Timeout            *time.Duration
	RetryInterval      *time.Duration
}

// WaitForPreconfiguredAlerts waits for multiple preconfigured alerts to reach a target status.
// This function can be used to wait for alerts to be enabled or disabled after calling
// EnableAlertRules or DisableAlertRules.
func (s *RegionalAPI) WaitForPreconfiguredAlerts(req *WaitForPreconfiguredAlertsRequest, opts ...scw.RequestOption) ([]*Alert, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	var terminalStatus map[AlertStatus]struct{}
	switch req.TargetStatus {
	case AlertStatusEnabled:
		terminalStatus = map[AlertStatus]struct{}{
			AlertStatusEnabled:  {},
			AlertStatusEnabling: {},
		}
	case AlertStatusDisabled:
		terminalStatus = map[AlertStatus]struct{}{
			AlertStatusDisabled:  {},
			AlertStatusDisabling: {},
		}
	default:
		terminalStatus = map[AlertStatus]struct{}{
			req.TargetStatus: {},
		}
	}

	result, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.ListAlerts(&RegionalAPIListAlertsRequest{
				Region:          req.Region,
				ProjectID:       req.ProjectID,
				IsPreconfigured: scw.BoolPtr(true),
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			alertsByRuleID := make(map[string]*Alert)
			for _, alert := range res.Alerts {
				if alert.PreconfiguredData != nil && alert.PreconfiguredData.PreconfiguredRuleID != "" {
					alertsByRuleID[alert.PreconfiguredData.PreconfiguredRuleID] = alert
				}
			}

			var matchedAlerts []*Alert
			for _, ruleID := range req.PreconfiguredRules {
				alert, found := alertsByRuleID[ruleID]
				if !found {
					return nil, false, errors.New("preconfigured alert with rule ID %s not found", ruleID)
				}

				_, isTerminal := terminalStatus[alert.RuleStatus]
				if !isTerminal {
					return nil, false, nil
				}

				matchedAlerts = append(matchedAlerts, alert)
			}

			return matchedAlerts, true, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for preconfigured alerts failed")
	}
	return result.([]*Alert), nil
}
