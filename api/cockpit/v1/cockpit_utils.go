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

// WaitForAlertRequest is used by WaitForAlert method.
type WaitForAlertRequest struct {
	Region        scw.Region
	AlertID       string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForAlert waits for the alert to be in a "terminal state" before returning.
// This function can be used to wait for an alert to be enabled or disabled.
func (s *RegionalAPI) WaitForAlert(req *WaitForAlertRequest, opts ...scw.RequestOption) (*Alert, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[AlertStatus]struct{}{
		AlertStatusEnabled:  {},
		AlertStatusDisabled: {},
	}

	alert, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			// List all alerts and find the one with matching ID
			res, err := s.ListAlerts(&RegionalAPIListAlertsRequest{
				Region: req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			// Find the alert by ID
			for _, alert := range res.Alerts {
				if alert.ID == req.AlertID {
					_, isTerminal := terminalStatus[alert.RuleStatus]
					return alert, isTerminal, nil
				}
			}

			return nil, false, errors.New("alert not found")
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for alert failed")
	}
	return alert.(*Alert), nil
}

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

	// For enabling/disabling, accept transitional states as terminal
	var terminalStatus map[AlertStatus]struct{}
	switch req.TargetStatus {
	case AlertStatusEnabled:
		// Accept both enabled and enabling as terminal states
		terminalStatus = map[AlertStatus]struct{}{
			AlertStatusEnabled:  {},
			AlertStatusEnabling: {},
		}
	case AlertStatusDisabled:
		// Accept both disabled and disabling as terminal states
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
			// List all preconfigured alerts for the project
			res, err := s.ListAlerts(&RegionalAPIListAlertsRequest{
				Region:          req.Region,
				ProjectID:       req.ProjectID,
				IsPreconfigured: scw.BoolPtr(true),
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			// Create a map of rule ID to alert
			alertsByRuleID := make(map[string]*Alert)
			for _, alert := range res.Alerts {
				if alert.PreconfiguredData != nil && alert.PreconfiguredData.PreconfiguredRuleID != "" {
					alertsByRuleID[alert.PreconfiguredData.PreconfiguredRuleID] = alert
				}
			}

			// Check if all requested alerts are in terminal state
			var matchedAlerts []*Alert
			for _, ruleID := range req.PreconfiguredRules {
				alert, found := alertsByRuleID[ruleID]
				if !found {
					return nil, false, errors.New("preconfigured alert with rule ID %s not found", ruleID)
				}

				_, isTerminal := terminalStatus[alert.RuleStatus]
				if !isTerminal {
					// At least one alert is not in terminal state, continue waiting
					return nil, false, nil
				}

				matchedAlerts = append(matchedAlerts, alert)
			}

			// All alerts are in terminal state
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

