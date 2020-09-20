package main

import (
	"context"

	"time"

	"github.com/go-benchmark/scenario/config"
	"github.com/go-benchmark/scenario/device"
	"github.com/go-benchmark/scenario/user"
	gobench "github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/executor/scenario"
	"github.com/gobench-io/gobench/logger"
)

var opts *config.Options
var log *logger.Log

// export scenarios
func export() scenario.Vus {
	const vu = 10000
	log = logger.NewStdLogger()
	opts, _ = config.ConfigureOptions(vu)
	log.Infow("starting run scenario", "opts", opts)

	return scenario.Vus{
		{
			Nu:   vu,
			Rate: 0.5,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	var err error
	const realtimeInterval = 86400

	// 1. device creation
	d, _ := device.NewDevice(opts, vui, log)

	// 1.1 device calls home
	if err := d.GetMqttAuth(ctx); err != nil {
		log.Errorw("FAILED call home", "vui", vui, "user", nil, "device", d, "error", err)
		return
	}
	// 1.1 device connect MQTT
	if err = d.ConnectMqtt(ctx); err != nil {
		log.Errorw("FAILED connecting to MQTT broker", "vui", vui, "user", nil, "device", d, "error", err)
		return
	}

	// 1.2 run the device, the realtime and history are the background tasks
	if err = d.Run(ctx); err != nil {
		log.Errorw("FAILED starting a device", "vui", vui, "user", nil, "device", d, "error", err)
	}

	// 2. user action
	u := user.NewUser(opts)

	if err = u.GetUserAccount(ctx); err != nil {
		log.Errorw("FAILED get user account", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	if err = u.SignUp(ctx); err != nil {
		log.Errorw("FAILED create a user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	if err = u.Login(ctx); err != nil {
		log.Errorw("FAILED login user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}

	// 2.3 Create a deviceset
	if err = u.CreateDeviceSet(ctx); err != nil {
		log.Errorw("FAILED create a deviceset to a user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	// 2.4 Add devices to deviceset
	if err = u.AddDevicesToDeviceSets(ctx, d); err != nil {
		log.Errorw("FAILED  add devices to devicesets to a user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	// ### 2.5 Add service to a deviceset
	if err = u.AddServicesToDeviceSets(ctx); err != nil {
		log.Errorw("FAILED add services to devicesets to a user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	// ### 2.6. Create 2 zones, a zone has one bot
	if err = u.AddZones(ctx, 2); err != nil {
		log.Errorw("FAILED create 2 zone to a user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	// delay to ensure device is online
	gobench.SleepRateLinear(1 / opts.UC.StartServiceDelay)

	// // ### 2.7. Run services
	if err = u.StartServices(ctx); err != nil {
		log.Errorw("FAILED start services to a user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	//delay to ensure service is running
	gobench.SleepRateLinear(1 / opts.UC.RunServiceDelay)

	//TODO: need a solution to simulate close to reality
	// First scenario: Break user rountine when meet an error in user behavior
	// Second scenario: Continue user rountine when meet errors in user behavior
	// Third scenario: Try to simulate user behavior in case the server is not adapt yet.
	if err = u.CheckConfigs(ctx); err != nil {
		log.Errorw("FAILED check services config to a user", "vui", vui, "user", u, "device", d, "error", err)
		return
	}
	// // loop user : get realtime data of a service, get history of service
	// ### 2.8 Simulate users regular behavior
	// 2.8.1 Get history of a service
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err = u.GetHistoriesByUser(ctx); err != nil {
					log.Errorw("FAILED get histories of services to a user", "vui", vui, "user", u, "device", d, "error", err)
					return
				}
			}
			gobench.SleepRatePoisson(1 / opts.UC.GetDataInterval)
		}
	}(ctx)
	// 2.8.2 create a heartbeat
	go func(ctx context.Context) {
		period := time.Duration(int(opts.UC.RealtimePeriod)) * time.Second
		for {
			select {
			case <-ctx.Done():
				return
			default:
				elapsed := time.Now().Add(period)
				for {
					if time.Now().After(elapsed) {
						break
					}
					if err = u.CreateHeartbeats(ctx); err != nil {
						log.Errorw("FAILED get realtime data of services to a user", "vui", vui, "user", u, "device", d, "error", err)
						return
					}
					gobench.SleepRatePoisson(1 / float64(opts.UC.RealtimeHBInterval))
				}
				// sleep to next in realtime request interval
				gobench.SleepRatePoisson(1 / (realtimeInterval - opts.UC.RealtimePeriod))
			}
		}
	}(ctx)
	// // 2.8.3 Get device status
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err = u.GetDeviceStatus(ctx, d); err != nil {
					log.Errorw("FAILED get device status by a user", "vui", vui, "user", u, "device", d, "error", err)
					return
				}
			}
			gobench.SleepRatePoisson(1 / opts.UC.GetDataInterval)
		}
	}(ctx)
	// 2.8.4 Get service params
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err = u.GetServiceParamsByUser(ctx); err != nil {
					log.Errorw("FAILED get service params of services by a user", "vui", vui, "user", u, "device", d, "error", err)
					return
				}
			}
			gobench.SleepRatePoisson(1 / opts.UC.GetDataInterval)
		}
	}(ctx)
	// wait for x seconds
	gobench.SleepRateLinear(1 / (1 * 60 * 60.0)) // sleep for 1 hours
}
