// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"time"

	"github.com/twilio/twilio-go/client"
	AccountsV1 "github.com/twilio/twilio-go/rest/accounts/v1"
	ApiV2010 "github.com/twilio/twilio-go/rest/api/v2010"
	AutopilotV1 "github.com/twilio/twilio-go/rest/autopilot/v1"
	BulkexportsV1 "github.com/twilio/twilio-go/rest/bulkexports/v1"
	ChatV1 "github.com/twilio/twilio-go/rest/chat/v1"
	ChatV2 "github.com/twilio/twilio-go/rest/chat/v2"
	ConversationsV1 "github.com/twilio/twilio-go/rest/conversations/v1"
	EventsV1 "github.com/twilio/twilio-go/rest/events/v1"
	FaxV1 "github.com/twilio/twilio-go/rest/fax/v1"
	FlexV1 "github.com/twilio/twilio-go/rest/flex/v1"
	InsightsV1 "github.com/twilio/twilio-go/rest/insights/v1"
	IpMessagingV1 "github.com/twilio/twilio-go/rest/ip_messaging/v1"
	IpMessagingV2 "github.com/twilio/twilio-go/rest/ip_messaging/v2"
	LookupsV1 "github.com/twilio/twilio-go/rest/lookups/v1"
	MessagingV1 "github.com/twilio/twilio-go/rest/messaging/v1"
	MonitorV1 "github.com/twilio/twilio-go/rest/monitor/v1"
	NotifyV1 "github.com/twilio/twilio-go/rest/notify/v1"
	NumbersV2 "github.com/twilio/twilio-go/rest/numbers/v2"
	PricingV1 "github.com/twilio/twilio-go/rest/pricing/v1"
	PricingV2 "github.com/twilio/twilio-go/rest/pricing/v2"
	ProxyV1 "github.com/twilio/twilio-go/rest/proxy/v1"
	ServerlessV1 "github.com/twilio/twilio-go/rest/serverless/v1"
	StudioV1 "github.com/twilio/twilio-go/rest/studio/v1"
	StudioV2 "github.com/twilio/twilio-go/rest/studio/v2"
	SupersimV1 "github.com/twilio/twilio-go/rest/supersim/v1"
	SyncV1 "github.com/twilio/twilio-go/rest/sync/v1"
	TaskrouterV1 "github.com/twilio/twilio-go/rest/taskrouter/v1"
	TrunkingV1 "github.com/twilio/twilio-go/rest/trunking/v1"
	TrusthubV1 "github.com/twilio/twilio-go/rest/trusthub/v1"
	VerifyV2 "github.com/twilio/twilio-go/rest/verify/v2"
	VideoV1 "github.com/twilio/twilio-go/rest/video/v1"
	VoiceV1 "github.com/twilio/twilio-go/rest/voice/v1"
	WirelessV1 "github.com/twilio/twilio-go/rest/wireless/v1"
)

// RestClient provides access to Twilio services.
type RestClient struct {
	*client.RequestHandler
	AccountsV1      *AccountsV1.DefaultApiService
	ApiV2010        *ApiV2010.DefaultApiService
	AutopilotV1     *AutopilotV1.DefaultApiService
	BulkexportsV1   *BulkexportsV1.DefaultApiService
	ChatV1          *ChatV1.DefaultApiService
	ChatV2          *ChatV2.DefaultApiService
	ConversationsV1 *ConversationsV1.DefaultApiService
	EventsV1        *EventsV1.DefaultApiService
	FaxV1           *FaxV1.DefaultApiService
	FlexV1          *FlexV1.DefaultApiService
	InsightsV1      *InsightsV1.DefaultApiService
	IpMessagingV1   *IpMessagingV1.DefaultApiService
	IpMessagingV2   *IpMessagingV2.DefaultApiService
	LookupsV1       *LookupsV1.DefaultApiService
	MessagingV1     *MessagingV1.DefaultApiService
	MonitorV1       *MonitorV1.DefaultApiService
	NotifyV1        *NotifyV1.DefaultApiService
	NumbersV2       *NumbersV2.DefaultApiService
	PricingV1       *PricingV1.DefaultApiService
	PricingV2       *PricingV2.DefaultApiService
	ProxyV1         *ProxyV1.DefaultApiService
	ServerlessV1    *ServerlessV1.DefaultApiService
	StudioV1        *StudioV1.DefaultApiService
	StudioV2        *StudioV2.DefaultApiService
	SupersimV1      *SupersimV1.DefaultApiService
	SyncV1          *SyncV1.DefaultApiService
	TaskrouterV1    *TaskrouterV1.DefaultApiService
	TrunkingV1      *TrunkingV1.DefaultApiService
	TrusthubV1      *TrusthubV1.DefaultApiService
	VerifyV2        *VerifyV2.DefaultApiService
	VideoV1         *VideoV1.DefaultApiService
	VoiceV1         *VoiceV1.DefaultApiService
	WirelessV1      *WirelessV1.DefaultApiService
}

// Meta holds relevant pagination resources.
type Meta struct {
	FirstPageURL    *string `json:"first_page_url"`
	Key             *string `json:"key"`
	LastPageURL     *string `json:"last_page_url,omitempty"`
	NextPageURL     *string `json:"next_page_url"`
	Page            *int    `json:"page"`
	PageSize        *int    `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url"`
	URL             *string `json:"url"`
}

type RestClientParams struct {
	AccountSid string
	Client     client.BaseClient
}

// NewRestClientWithParams provides an initialized Twilio RestClient with params.
func NewRestClientWithParams(username string, password string, params RestClientParams) *RestClient {
	requestHandler := client.NewRequestHandler(params.Client)

	if params.Client == nil {
		defaultClient := &client.Client{
			Credentials: client.NewCredentials(username, password),
		}

		if params.AccountSid != "" {
			defaultClient.SetAccountSid(params.AccountSid)
		} else {
			defaultClient.SetAccountSid(username)
		}
		requestHandler = client.NewRequestHandler(defaultClient)
	}

	c := &RestClient{
		RequestHandler: requestHandler,
	}

	c.AccountsV1 = AccountsV1.NewDefaultApiService(c.RequestHandler)
	c.ApiV2010 = ApiV2010.NewDefaultApiService(c.RequestHandler)
	c.AutopilotV1 = AutopilotV1.NewDefaultApiService(c.RequestHandler)
	c.BulkexportsV1 = BulkexportsV1.NewDefaultApiService(c.RequestHandler)
	c.ChatV1 = ChatV1.NewDefaultApiService(c.RequestHandler)
	c.ChatV2 = ChatV2.NewDefaultApiService(c.RequestHandler)
	c.ConversationsV1 = ConversationsV1.NewDefaultApiService(c.RequestHandler)
	c.EventsV1 = EventsV1.NewDefaultApiService(c.RequestHandler)
	c.FaxV1 = FaxV1.NewDefaultApiService(c.RequestHandler)
	c.FlexV1 = FlexV1.NewDefaultApiService(c.RequestHandler)
	c.InsightsV1 = InsightsV1.NewDefaultApiService(c.RequestHandler)
	c.IpMessagingV1 = IpMessagingV1.NewDefaultApiService(c.RequestHandler)
	c.IpMessagingV2 = IpMessagingV2.NewDefaultApiService(c.RequestHandler)
	c.LookupsV1 = LookupsV1.NewDefaultApiService(c.RequestHandler)
	c.MessagingV1 = MessagingV1.NewDefaultApiService(c.RequestHandler)
	c.MonitorV1 = MonitorV1.NewDefaultApiService(c.RequestHandler)
	c.NotifyV1 = NotifyV1.NewDefaultApiService(c.RequestHandler)
	c.NumbersV2 = NumbersV2.NewDefaultApiService(c.RequestHandler)
	c.PricingV1 = PricingV1.NewDefaultApiService(c.RequestHandler)
	c.PricingV2 = PricingV2.NewDefaultApiService(c.RequestHandler)
	c.ProxyV1 = ProxyV1.NewDefaultApiService(c.RequestHandler)
	c.ServerlessV1 = ServerlessV1.NewDefaultApiService(c.RequestHandler)
	c.StudioV1 = StudioV1.NewDefaultApiService(c.RequestHandler)
	c.StudioV2 = StudioV2.NewDefaultApiService(c.RequestHandler)
	c.SupersimV1 = SupersimV1.NewDefaultApiService(c.RequestHandler)
	c.SyncV1 = SyncV1.NewDefaultApiService(c.RequestHandler)
	c.TaskrouterV1 = TaskrouterV1.NewDefaultApiService(c.RequestHandler)
	c.TrunkingV1 = TrunkingV1.NewDefaultApiService(c.RequestHandler)
	c.TrusthubV1 = TrusthubV1.NewDefaultApiService(c.RequestHandler)
	c.VerifyV2 = VerifyV2.NewDefaultApiService(c.RequestHandler)
	c.VideoV1 = VideoV1.NewDefaultApiService(c.RequestHandler)
	c.VoiceV1 = VoiceV1.NewDefaultApiService(c.RequestHandler)
	c.WirelessV1 = WirelessV1.NewDefaultApiService(c.RequestHandler)

	return c
}

// NewRestClient provides an initialized Twilio RestClient.
func NewRestClient(username string, password string) *RestClient {
	return NewRestClientWithParams(username, password, RestClientParams{
		AccountSid: username,
	})
}

// SetTimeout sets the Timeout for Twilio HTTP requests.
func (c *RestClient) SetTimeout(timeout time.Duration) {
	c.RequestHandler.Client.SetTimeout(timeout)
}

// SetEdge sets the Edge for the Twilio request.
func (c *RestClient) SetEdge(edge string) {
	c.RequestHandler.Edge = edge
}

// SetRegion sets the Region for the Twilio request. Defaults to "us1" if an edge is provided.
func (c *RestClient) SetRegion(region string) {
	c.RequestHandler.Region = region
}
