package status

type Status struct {
	Connections  Connections  `json:"connections"`
	Requests     Requests     `json:"requests"`
	Applications Applications `json:"applications"`
}

type Connections struct {
	Accepted int `json:"accepted"`
	Active   int `json:"active"`
	Idle     int `json:"idle"`
	Closed   int `json:"closed"`
}

type Requests struct {
	Total int `json:"total"`
}

type Application struct {
	Processes ApplicationProcesses `json:"processes"`
	Requests  ApplicationRequests  `json:"requests"`
}

type ApplicationProcesses struct {
	Running  int `json:"running"`
	Starting int `json:"starting"`
	Idle     int `json:"idle"`
}

type ApplicationRequests struct {
	Active int `json:"active"`
}

type Applications map[string]Application
