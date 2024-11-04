package sentry

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var client sentry.ClientOptions
var sentryFiber fiber.Handler

func initClientOpt() {
	client = sentry.ClientOptions{
		Dsn:              utils.AddToStr(configs.Instance.SentryURL),
		TracesSampleRate: 0.5,
		Debug:            false, //!configs.Instance.IsProduction(),
		AttachStacktrace: true,
	}
}

func Middleware() fiber.Handler {
	return sentryFiber
}

func ClientOpt() sentry.ClientOptions {
	return client
}

func InitializeSentry(l *logrus.Logger) (*sentrylogrus.Hook, error) {
	initClientOpt()

	if err := sentry.Init(client); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return nil, err
	}

	// Send only ERROR and higher level logs to Sentry
	sentryLevels := []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}

	// Initialize Sentry
	sentryHook, err := sentrylogrus.New(sentryLevels, client)
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return nil, err
	}
	log.AddHook(sentryHook)

	logrus.RegisterExitHandler(func() { sentryHook.Flush(5 * time.Second) })

	sentryFiber = sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: false,
	})

	return sentryHook, nil
}
