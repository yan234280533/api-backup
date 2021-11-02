package v1alpha1

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// HighestGauranteed is the highest level for guaranteed pod.
	HighestGauranteed = int32(100000)

	// HighestDefinableServiceLevelForGuaranteedPod is the lowest level for guaranteed pod.
	LowestGauranteed = int32(90000)

	// HighestBurstable is the highest level for burstable pod.
	HighestBurstable = int32(89999)

	// LowestBurstable is the lowest level for burstable pod.
	LowestBurstable = int32(10000)

	// HighestDefinableServiceLevelForBestEffortPod is the highest level for bestEffort pod.
	HighestBestEffort = int32(9999)

	// LowestBestEffort is the lowest level for bestEffort pod.
	LowestBestEffort = int32(0)
)

type Operator string

const (
	OperatorNotEqual Operator = "!="
	OperatorGreater  Operator = ">"
	OperatorSmaller  Operator = ">"
)

// URIScheme identifies the scheme used for connection to a host for Get actions
type URIScheme string

const (
	// URISchemeHTTP means that the scheme used will be http://
	URISchemeHTTP URIScheme = "HTTP"
	// URISchemeHTTPS means that the scheme used will be https://
	URISchemeHTTPS URIScheme = "HTTPS"
)

// ServiceLevel defines the service level for pods
// Configure by platform developers

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
type ServiceLevelClass struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServiceLevelSpec `json:"spec,omitempty"`

	Status ServiceLevelStatus `json:"spec,omitempty"`
}

type ServiceLevelSpec struct {
	// The value of pods level. This is the actual level that pods
	// receive when they have the name of this class in their pod evasion policy.
	// Integer value range （0～100000), the highest level is 100000.
	// 100000>1000>1>0
	Value int32 `json:"value"`

	// Description is an arbitrary string that usually provides guidelines on
	// when this qos level class should be used.
	// +optional
	Description string `json:"description,omitempty"`
}

type ServiceLevelStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceLevelClassList contains a list of ServiceLevelClass
type ServiceLevelClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceLevelClass `json:"items"`
}

// EvasionActionClass defines Evasion action
// Configure by platform developers

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
type EvasionActionClass struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EvasionActionClassSpec   `json:"spec"`
	Status EvasionActionClassStatus `json:"status"`
}

type EvasionActionClassSpec struct {
	// how long it should wait between blocking/unblocking scheduling
	SchedulingCoolDown time.Duration `json:"schedulingCoolDown,omitempty"`

	//Action to restrain resource
	// +optional
	Restrain *RestrainAction `json:"restrain,omitempty"`

	//Action to evict low level pods
	// +optional
	Eviction *EvictionAction `json:"eviction,omitempty"`

	// Description is an arbitrary string that usually provides guidelines on
	// when this action should be used.
	// +optional
	Description string `json:"description,omitempty"`
}

type EvasionActionClassStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EvasionActionClassList contains a list of EvasionActionClass
type EvasionActionClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EvasionActionClass `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodEvasionPolicy is the Schema for the podevasionpolicies API
type PodEvasionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodEvasionPolicySpec   `json:"spec,omitempty"`
	Status PodEvasionPolicyStatus `json:"status,omitempty"`
}

// PodEvasionPolicySpec defines the desired state of PodEvasionPolicy
// Configure by businesses developers
type PodEvasionPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// service level for pods
	ServiceLevelName string `json:"serviceLevelName"`

	// select pod used labels
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`

	//pod quality probe
	QualityProbe QualityProbe `json:"qualityProbe,omitempty"`

	//pod objective ensurance check and action
	ObjectiveEnsurance []ObjectiveEnsurance `json:"objectiveEnsurance,omitempty"`
}

// PodEvasionPolicyStatus defines the observed state of PodEvasionPolicy
type PodEvasionPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodEvasionPolicyList contains a list of PodEvasionPolicy
type PodEvasionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodEvasionPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeEvasionPolicy is the Schema for the nodeevasionpolicies API
type NodeEvasionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeEvasionPolicySpec   `json:"spec,omitempty"`
	Status NodeEvasionPolicyStatus `json:"status,omitempty"`
}

// NodeEvasionPolicySpec defines the desired state of NodeEvasionPolicy
//  Configure  by platform developers
type NodeEvasionPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//select nodes use labels
	LabelSelector metav1.LabelSelector `json:"labelSelector"`

	//node quality probe
	NodeQualityProbe NodeQualityProbe `json:"nodeQualityProbe,omitempty"`

	//node objective ensurance check and action
	ObjectiveEnsurance []ObjectiveEnsurance `json:"objectiveEnsurance,omitempty"`
}

// NodeEvasionPolicyStatus defines the observed state of NodeEvasionPolicy
type NodeEvasionPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeEvasionPolicyList contains a list of NodeEvasionPolicy
type NodeEvasionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeEvasionPolicy `json:"items"`
}

type RestrainAction struct {
	// +optional
	CPURestrain CPURestrain `json:"cpuRestrain,omitempty"`

	// +optional
	MemoryRestrain MemoryRestrain `json:"memoryRestrain,omitempty"`
}

type CPURestrain struct {
	// how long it waits for each compress step
	// +optional
	Interval time.Duration `json:"interval,omitempty"`

	//the min of cpu share ratio for low level pods
	//example: the pod share is 4096, ratio is 10, the min is 409
	// +optional
	MinCPUShareRatio int32 `json:"minCPUShareRatio,omitempty"`

	//the min of cpu limit ratio for low level pods
	//example: the pod limit is 4096, ratio is 10, the min is 409
	// +optional
	MinCPULimitRatio uint64 `json:"minCPULimitRatio,omitempty"`

	//the step of cpu share and limit for once down-size (1-100)
	// +optional
	StepCPURatio uint64 `json:"stepCPURatio,omitempty"`
}

type MemoryRestrain struct {
	// how long it waits for each compress step
	// +optional
	Interval time.Duration `json:"interval,omitempty"`

	// to force gc the page cache of low level pods
	// +optional
	ForceGC bool `json:"forceGC,omitempty"`
}

type EvictionAction struct {
	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
	// Value must be non-negative integer. The value zero indicates delete immediately.
	// +optional
	DeletionGracePeriodSeconds *int32 `json:"deletionGracePeriodSeconds,omitempty"`
}

type QualityProbe struct {
	Handler             `json:",inline"`
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32 `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32 `json:"periodSeconds,omitempty"`
}

// Handler defines a specific action that should be taken
type Handler struct {
	HTTPGet *HTTPGet `json:"httpGet,omitempty"`
}

type HTTPGet struct {
	// Path to access on the HTTP server.
	// +optional
	Path string `json:"path,omitempty"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +optional
	Host string `json:"host,omitempty"`
	// Scheme to use for connecting to the host.
	// Defaults to HTTP.
	// +optional
	Scheme URIScheme `json:"scheme,omitempty"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`
}

// HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeader struct {
	// The header field name
	Name string `json:"name"`
	// The header field value
	Value string `json:"value"`
}

type ObjectiveEnsurance struct {
	// MetricName is the name of the given metric
	MetricName string `json:"metricName"`

	// Selector is the selector for the given metric
	// it is the string-encoded form of a standard kubernetes label selector
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// TargetThreshold is the target threshold of the metric (as a quantity).
	TargetThreshold *resource.Quantity `json:"targetThreshold"`

	// Target opterator such as (">","<","!="), default (">")
	TargetOperator Operator `json:"targetOpterator,omitempty"`

	// How many times the QualityObjective is reach, to trigger action
	ReachedThreshold int32 `json:"reachedThreshold,omitempty"`

	// Evasion action when be triggered
	EvasionAction []string `json:"actions,omitempty"`
}

type NodeQualityProbe struct {
	Handler NodeHandler `json:",inline"`
	// +optional
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
	// +optional
	PeriodSeconds int32 `json:"periodSeconds,omitempty"`
}

type NodeHandler struct {
	// currently supported
	// CPU usage, CPU load, Memory Usage, DiskIO
	// +optional
	HTTPGet *HTTPGet `json:"httpGet,omitempty"`

	// Get node metric from local
	// +optional
	NodeLocalGet *NodeLocalGet `json:"nodeLocalGet,omitempty"`
}

type NodeLocalGet struct {
	// +optional
	LocalCacheTTL time.Duration `json:"localCacheTTL,omitempty"`
	// +optional
	MaxHousekeepingInterval time.Duration `json:"maxHousekeepingInterval,omitempty"`
}
