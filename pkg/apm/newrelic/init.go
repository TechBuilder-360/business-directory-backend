package newrelic

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func InitialiseNewRelic(l *logrus.Logger) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(utils.AddToStr(configs.Instance.NewRelicAppName)),
		newrelic.ConfigLicense(utils.AddToStr(configs.Instance.NewRelicLicense)),
		newrelic.ConfigAppLogForwardingEnabled(false),
		//nrlogrus.ConfigLogger(l),
		newrelic.ConfigAppLogEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigAppLogDecoratingEnabled(true),
		newrelic.ConfigInfoLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		log.Error(err.Error())
	}

	err = app.WaitForConnection(10 * time.Second)
	if err != nil {
		log.Error(err.Error())
	}

}
