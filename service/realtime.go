package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gobench-io/gobench/dis"
)

// RealtimeHandler represent to realtime(heartbeat) handler
func (s *Service) RealtimeHandler(ctx context.Context, realtimeInterval float64, pub func(context.Context, string, byte, []byte) error) (err error) {
	elapsed := time.Now()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.heartbeat:
				period := time.Duration(int(s.OI.RealtimeLength*150/100)) * time.Second
				elapsed = time.Now().Add(period)
				break
			default:
				if elapsed.Before(time.Now()) {
					dis.SleepRateLinear(1.0 / float64(realtimeInterval))
					continue
				}

				var dc decision
				if err = faker.FakeData(&dc); err != nil {
					return
				}

				rd := realtimeDecision{
					Timestamp:  time.Now().Unix(),
					EngineType: s.EngineType,
					ServiceID:  s.ServiceID,
					Decision:   dc,
				}

				rdBytes, err := json.Marshal(rd)
				if err != nil {
					return
				}
				// send heartbeat to user
				if err = pub(ctx, "", 1, rdBytes); err != nil {
					return
				}
				dis.SleepRateLinear(1.0 / float64(realtimeInterval))
				break
			}
		}
	}(ctx)

	return
}

// RealtimeConfig  represent to handle realtime config
func (s *Service) RealtimeConfig(ctx context.Context) (err error) {
	s.heartbeat <- true
	return
}
