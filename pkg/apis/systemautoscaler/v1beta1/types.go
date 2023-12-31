package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceLevelAgreementList is a list of ServiceLevelAgreement resources
type ServiceLevelAgreementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ServiceLevelAgreement `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceLevelAgreement is a configuration for the autoscaling system.
// It sets a requirement on the services that matches the selector.
type ServiceLevelAgreement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec ServiceLevelAgreementSpec `json:"spec"`
}

// RecommendLogic defines logic used during the recommendation phase
type RecommendLogic string

const (
	FixedGainControl    RecommendLogic = "fixedGainControl"
	AdaptiveGainControl RecommendLogic = "adaptiveGainControl"
)

// ServiceLevelAgreementSpec defines the agreement specifying the
// metric requirement to honor by System Autoscaler, a Selector used
// to match a service with the Service Level Agreement and the
// default resources assigned to pods in case the `requests` field
// is empty in the `PodSpec`.
type ServiceLevelAgreementSpec struct {
	// Specify the metric on which the requirement is set.
	// +kubebuilder:validation:Required
	Metric MetricRequirement `json:"metric"`
	// Specify the logic used during the recommendation phase
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="fixed_gain_control"
	RecommenderLogic RecommendLogic `json:"recommenderLogic"`
	// Specify the default resources assigned to pods in case `requests` field is empty in `PodSpec`.
	// +kubebuilder:validation:Required
	DefaultResources v1.ResourceList `json:"defaultResources,omitempty" protobuf:"bytes,3,rep,name=defaultResources,casttype=ResourceList,castkey=ResourceName"`
	// The lower bound of resources to assign to containers.
	// +kubebuilder:validation:Optional
	MinResources v1.ResourceList `json:"minResources,omitempty" protobuf:"bytes,3,rep,name=minResources,casttype=ResourceList,castkey=ResourceName"`
	// The upper bound of resources to assign to containers.
	// +kubebuilder:validation:Optional
	MaxResources v1.ResourceList `json:"maxResources,omitempty" protobuf:"bytes,3,rep,name=maxResources,casttype=ResourceList,castkey=ResourceName"`
	// The lower bound of replicas for the application.
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=1
	MinReplicas int32 `json:"minReplicas,omitempty" protobuf:"bytes,3,rep,name=minResources,casttype=ResourceList,castkey=ResourceName"`
	// The upper bound of replicas for the application.
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=100
	MaxReplicas int32 `json:"maxReplicas,omitempty" protobuf:"bytes,3,rep,name=maxResources,casttype=ResourceList,castkey=ResourceName"`
	// Identify the Service on which the agreement is defined
	// +kubebuilder:validation:Required
	Service *Service `json:"service"`
}

// Service is used to identify the application to scale by its service Lavels and the container offering the Application service
type Service struct {
	// Specify the selector to match Services and Service Level Agreement
	// +kubebuilder:validation:Required
	Selector *metav1.LabelSelector `json:"selector"`
	// The container to track inside the Pods.
	// +kubebuilder:validation:Required
	Container string `json:"container"`
}

// MetricRequirement specifies a requirement for a metric.
// This means that System Autoscaler will try to honor the
// agreement, making the service metric coherent with it.
// Only one MetricRequirement per ServiceLevelAgreement resource
// must be set to avoid ambiguity.
// Currently supports only ResponseTime.
//
// i.e.: the metric type is the Response Time and the value
// is 4 units of time. This means that the system will try
// to keep the service response time below 4 on average.
type MetricRequirement struct {
	// +kubebuilder:validation:Required
	ResponseTime resource.Quantity `json:"responseTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodScaleList is a list of PodScale resources
type PodScaleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PodScale `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodScale defines the mapping between a `ServiceLevelAgreement` and a
// `Pod` matching the selector. It also keeps track of the resource values
// computed by `Recommender` and adjusted by `Contention Manager`.
type PodScale struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodScaleSpec   `json:"spec"`
	Status PodScaleStatus `json:"status"`
}

// PodScaleSpec is the spec for a PodScale resource
type PodScaleSpec struct {
	Namespace        string          `json:"namespace"`
	SLA              string          `json:"serviceLevelAgreement"`
	Pod              string          `json:"pod"`
	Service          string          `json:"service"`
	Container        string          `json:"container"`
	DesiredResources v1.ResourceList `json:"desired,omitempty" protobuf:"bytes,3,rep,name=desired,casttype=ResourceList,castkey=ResourceName"`
}

// PodScaleStatus contains the resources patched by the
// `Contention Manager` according to the available node resources
// and other pods' SLA
type PodScaleStatus struct {
	CappedResources v1.ResourceList `json:"capped,omitempty" protobuf:"bytes,3,rep,name=actual,casttype=ResourceList,castkey=ResourceName"`
	ActualResources v1.ResourceList `json:"actual,omitempty" protobuf:"bytes,3,rep,name=actual,casttype=ResourceList,castkey=ResourceName"`
}
