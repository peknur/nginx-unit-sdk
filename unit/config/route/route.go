package route

type Config struct {
	Match  *Match  `json:"match,omitempty"`
	Action *Action `json:"action,omitempty"`
}

type Match struct {
	URI       string            `json:"uri,omitempty"`
	Host      []string          `json:"host,omitempty"`
	Source    []string          `json:"source,omitempty"`
	Scheme    string            `json:"scheme,omitempty"`
	Arguments map[string]string `json:"arguments,omitempty"`
	Cookies   map[string]string `json:"cookies,omitempty"`
	Query     []string          `json:"query,omitempty"`
}

type Action struct {
	Pass           string            `json:"pass,omitempty"`
	Share          []string          `json:"share,omitempty"`
	Chroot         string            `json:"chroot,omitempty"`
	TraverseMounts *bool             `json:"traverse_mounts,omitempty"`
	FollowSymlinks *bool             `json:"follow_symlinks,omitempty"`
	Types          []string          `json:"types,omitempty"`
	Fallback       map[string]string `json:"fallback,omitempty"`
	Return         int               `json:"return,omitempty"`
	Proxy          string            `json:"proxy,omitempty"`
	Location       string            `json:"location,omitempty"`
}
