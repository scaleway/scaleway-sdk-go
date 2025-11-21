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
	if req == nil {
		return nil, errors.New("WaitForPreconfiguredAlertsRequest cannot be nil")
	}

	if len(req.PreconfiguredRules) == 0 {
		return []*Alert{}, nil
	}

	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	listOpts := append([]scw.RequestOption{scw.WithAllPages()}, opts...)

	result, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.ListAlerts(&RegionalAPIListAlertsRequest{
				Region:          req.Region,
				ProjectID:       req.ProjectID,
				IsPreconfigured: scw.BoolPtr(true),
			}, listOpts...)
			if err != nil {
				return nil, false, err
			}

			alertsByRuleID := make(map[string]*Alert, len(res.Alerts))
			for _, alert := range res.Alerts {
				if alert.PreconfiguredData != nil && alert.PreconfiguredData.PreconfiguredRuleID != "" {
					alertsByRuleID[alert.PreconfiguredData.PreconfiguredRuleID] = alert
				}
			}

			matchedAlerts := make([]*Alert, 0, len(req.PreconfiguredRules))
			for _, ruleID := range req.PreconfiguredRules {
				alert, found := alertsByRuleID[ruleID]
				if !found {
					return nil, false, nil
				}

				if !isAlertInTargetStatus(req.TargetStatus, alert.RuleStatus) {
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

func isAlertInTargetStatus(target, current AlertStatus) bool {
	switch target {
	case AlertStatusEnabled:
		return current == AlertStatusEnabled
	case AlertStatusDisabled:
		return current == AlertStatusDisabled
	default:
		return current == target
	}
}
