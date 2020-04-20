package jikan

// MetaStatus struct for the /meta/status endpoint
type MetaStatus struct {
	CachedRequests           int    `json:"cached_requests"`
	RequestsToday            int    `json:"requests_today"`
	RequestsThisWeek         int    `json:"requests_this_week"`
	RequestsThisMonth        int    `json:"requests_this_month"`
	ConnectedClients         string `json:"connected_clients"`
	TotalConnectionsReceived string `json:"total_connections_received"`
}

// GetMetaStatus returns meta status
func GetMetaStatus() (*MetaStatus, error) {
	res := &MetaStatus{}

	err := urlToStruct("/meta/status", res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
