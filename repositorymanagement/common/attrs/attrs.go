package attrs

// CleanupPolicy struct for CleanupPolicy
type CleanupPolicy struct {
	// Components that match any of the applied policies will be deleted
	PolicyNames []string `json:"policyNames,omitempty"`
}

// Replication struct for Replication
type Replication struct {
	PreemptivePullEnabled bool   `json:"preemptivePullEnabled"`
	AssetPathRegex        string `json:"assetPathRegex,omitempty"`
}

// NegativeCache struct for NegativeCache
type NegativeCache struct {
	// Whether to cache responses for content not present in the proxied repository
	Enabled bool `json:"enabled"`
	// How long to cache the fact that a file was not found in the repository (in minutes)
	TimeToLive int32 `json:"timeToLive"`
}

// Storage struct for Storage
type Storage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName,omitempty"`
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`
}

// LocalStorage struct for LocalStorage
type LocalStorage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName,omitempty"`
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`
	// Controls if deployments of and updates to assets are allowed
	WritePolicy string `json:"writePolicy"`
}

// Group struct for Group
type Group struct {
	// Member repositories' names
	MemberNames []string `json:"memberNames,omitempty"`
}

// Proxy struct for Proxy
type Proxy struct {
	// Location of the remote repository being proxied
	RemoteUrl string `json:"remoteUrl,omitempty"`
	// How long to cache artifacts before rechecking the remote repository (in minutes)
	ContentMaxAge int32 `json:"contentMaxAge"`
	// How long to cache metadata before rechecking the remote repository (in minutes)
	MetadataMaxAge int32 `json:"metadataMaxAge"`
}

// HttpClient struct for HttpClient
type HttpClient struct {
	// Whether to block outbound connections on the repository
	Blocked bool `json:"blocked"`
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock      bool                                `json:"autoBlock"`
	Connection     *HttpClientConnection               `json:"connection,omitempty"`
	Authentication *HttpClientConnectionAuthentication `json:"authentication,omitempty"`
}

// HttpClientConnection struct for HttpClientConnection
type HttpClientConnection struct {
	// Total retries if the initial connection attempt suffers a timeout
	Retries int32 `json:"retries,omitempty"`
	// Custom fragment to append to User-Agent header in HTTP requests
	UserAgentSuffix string `json:"userAgentSuffix,omitempty"`
	// Seconds to wait for activity before stopping and retrying the connection
	Timeout int32 `json:"timeout,omitempty"`
	// Whether to enable redirects to the same location (may be required by some servers)
	EnableCircularRedirects bool `json:"enableCircularRedirects,omitempty"`
	// Whether to allow cookies to be stored and used
	EnableCookies bool `json:"enableCookies,omitempty"`
	// Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	UseTrustStore bool `json:"useTrustStore,omitempty"`
}

// HttpClientConnectionAuthentication struct for HttpClientConnectionAuthentication
type HttpClientConnectionAuthentication struct {
	// Authentication type
	Type       string `json:"type,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	NtlmHost   string `json:"ntlmHost,omitempty"`
	NtlmDomain string `json:"ntlmDomain,omitempty"`
}

// HttpClientConnectionAuthenticationWithPreemptive struct for HttpClientConnectionAuthenticationWithPreemptive
type HttpClientConnectionAuthenticationWithPreemptive struct {
	// Authentication type
	Type       string `json:"type,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	NtlmHost   string `json:"ntlmHost,omitempty"`
	NtlmDomain string `json:"ntlmDomain,omitempty"`
	// Whether to use pre-emptive authentication. Use with caution. Defaults to false.
	Preemptive bool `json:"preemptive,omitempty"`
}

// HttpClientWithPreemptiveAuth struct for HttpClientWithPreemptiveAuth
type HttpClientWithPreemptiveAuth struct {
	// Whether to block outbound connections on the repository
	Blocked bool `json:"blocked"`
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock      bool                                              `json:"autoBlock"`
	Connection     *HttpClientConnection                             `json:"connection,omitempty"`
	Authentication *HttpClientConnectionAuthenticationWithPreemptive `json:"authentication,omitempty"`
}
