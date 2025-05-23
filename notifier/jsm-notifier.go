package notifier

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

	return true
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
