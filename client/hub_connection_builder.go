package client

// HubConnectionBuilder represents a hub connection builder contract
type HubConnectionBuilder interface {
	WithUrl(url string) HubConnectionBuilder
	WithTransport(transportType TransportType) HubConnectionBuilder
	Build() HubConnection
}

type hubConnectionBuilder struct {
	url           string
	transportType TransportType
	transport     Transport
}

func (h *hubConnectionBuilder) WithUrl(url string) HubConnectionBuilder {
	h.url = url
	return h
}

func (h *hubConnectionBuilder) WithTransport(
	transportType TransportType,
) HubConnectionBuilder {
	h.transportType = transportType
	return h
}

func (h *hubConnectionBuilder) Build() HubConnection {
	return newHubConnection()
}

// NewHubConnectionBuilder creates a new hub connection builder
func NewHubConnectionBuilder() HubConnectionBuilder {
	return new(hubConnectionBuilder)
}
