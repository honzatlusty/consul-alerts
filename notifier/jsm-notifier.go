package notifier

import (
	"fmt"
	log "github.com/honzatlusty/consul-alerts/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

type JSMNotifier struct {
	Enabled     bool
	ClusterName string `json:"cluster-name"`
	ApiKey      string `json:"api-key"`
}

// NotifierName provides name for notifier selection
func (jsm *JSMNotifier) NotifierName() string {
	return "jsm"
}

func (jsm *JSMNotifier) Copy() Notifier {
	notifier := *jsm
	return &notifier
}

//Notify sends messages to the endpoint notifier
func (jsm *JSMNotifier) Notify(messages Messages) bool {

	overallStatus, pass, warn, fail := messages.Summary()

	for _, message := range messages {
		title := fmt.Sprintf("\n%s:%s:%s is %s.", message.Node, message.Service, message.Check, message.Status)
		alias := jsm.createAlias(message)
		content := fmt.Sprintf(header, jsm.ClusterName, overallStatus, fail, warn, pass)
		content += fmt.Sprintf("\n%s:%s:%s is %s.", message.Node, message.Service, message.Check, message.Status)
		content += fmt.Sprintf("\n%s", message.Output)

		// create the alert
		switch {
		case message.IsCritical():
		case message.IsWarning():
		case message.IsPassing():
			ok = jsm.sendAlertRequest(title, message.Status, content, alias) && ok
		default:
			ok = false
			log.Warn("Message was not either IsCritical, IsWarning or IsPasssing. No notification was sent for ", alias)
		}
	}
	return ok
}

func (jsm JSMNotifier) createAlias(message Message) string {
	incidentKey := message.Node
	if message.ServiceId != "" {
		incidentKey += ":" + message.ServiceId
	}

	return incidentKey
}

func (jsm *JSMNotifier) sendAlertRequest(title string, status string, content string, alias string) bool {
        return true
}
