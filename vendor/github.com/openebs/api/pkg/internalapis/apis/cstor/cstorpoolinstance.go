/*
Copyright 2020 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cstor

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=cstorpoolinstance

// CStorPoolInstance describes a cstor pool instance resource.
type CStorPoolInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the specification of the cstorpoolinstance resource.
	Spec CStorPoolInstanceSpec `json:"spec"`
	// Status is the possible statuses of the cstorpoolinstance resource.
	Status CStorPoolInstanceStatus `json:"status"`
	// VersionDetails is the openebs version.
	VersionDetails VersionDetails `json:"versionDetails"`
}

// CStorPoolInstanceSpec is the spec listing fields for a CStorPoolInstance resource.
type CStorPoolInstanceSpec struct {
	// HostName is the name of kubernetes node where the pool
	// should be created.
	HostName string `json:"hostName"`
	// NodeSelector is the labels that will be used to select
	// a node for pool provisioning.
	// Required field
	NodeSelector map[string]string `json:"nodeSelector"`
	// PoolConfig is the default pool config that applies to the
	// pool on node.
	PoolConfig PoolConfig `json:"poolConfig"`
	// DataRaidGroups is the raid group configuration for the given pool.
	DataRaidGroups []RaidGroup `json:"dataRaidGroups"`
	// WriteCacheRaidGroups is the write cache raid group.
	WriteCacheRaidGroups []RaidGroup `json:"writeCacheRaidGroups"`
}

// CStorPoolInstancePhase is the phase for CStorPoolInstance resource.
type CStorPoolInstancePhase string

// Status written onto CStorPool and CStorVolumeReplica objects.
// Resetting state to either Init or CreateFailed need to be done with care,
// as, label clear and pool creation depends on this state.
const (
	// CStorPoolStatusEmpty ensures the create operation is to be done, if import fails.
	CStorPoolStatusEmpty CStorPoolInstancePhase = ""
	// CStorPoolStatusOnline signifies that the pool is online.
	CStorPoolStatusOnline CStorPoolInstancePhase = "Healthy"
	// CStorPoolStatusOffline signifies that the pool is offline.
	CStorPoolStatusOffline CStorPoolInstancePhase = "Offline"
	// CStorPoolStatusDegraded signifies that the pool is degraded.
	CStorPoolStatusDegraded CStorPoolInstancePhase = "Degraded"
	// CStorPoolStatusFaulted signifies that the pool is faulted.
	CStorPoolStatusFaulted CStorPoolInstancePhase = "Faulted"
	// CStorPoolStatusRemoved signifies that the pool is removed.
	CStorPoolStatusRemoved CStorPoolInstancePhase = "Removed"
	// CStorPoolStatusUnavail signifies that the pool is not available.
	CStorPoolStatusUnavail CStorPoolInstancePhase = "Unavail"
	// CStorPoolStatusError signifies that the pool status could not be fetched.
	CStorPoolStatusError CStorPoolInstancePhase = "Error"
	// CStorPoolStatusDeletionFailed ensures the resource deletion has failed.
	CStorPoolStatusDeletionFailed CStorPoolInstancePhase = "DeletionFailed"
	// CStorPoolStatusInvalid ensures invalid resource.
	CStorPoolStatusInvalid CStorPoolInstancePhase = "Invalid"
	// CStorPoolStatusErrorDuplicate ensures error due to duplicate resource.
	CStorPoolStatusErrorDuplicate CStorPoolInstancePhase = "ErrorDuplicate"
	// CStorPoolStatusPending ensures pending task for cstorpool.
	CStorPoolStatusPending CStorPoolInstancePhase = "Pending"
	// CStorPoolStatusInit is initial state of CSP, before pool creation.
	CStorPoolStatusInit CStorPoolInstancePhase = "Init"
	// CStorPoolStatusCreateFailed is state when pool creation failed
	CStorPoolStatusCreateFailed CStorPoolInstancePhase = "PoolCreationFailed"
)

// CStorPoolInstanceStatus is for handling status of pool.
type CStorPoolInstanceStatus struct {
	// Current state of CSPI with details.
	Conditions []CStorPoolInstanceCondition
	//  The phase of a CStorPool is a simple, high-level summary of the pool state on the
	//  node.
	Phase CStorPoolInstancePhase `json:"phase"`
	// Capacity describes the capacity details of a cstor pool
	Capacity CStorPoolInstanceCapacity `json:"capacity"`
	// A human readable message indicating details about why the CSPI is in this
	// condition.
	Message string `json:"message,omitempty"`
}

// CStorPoolCapacityAttr stores the pool capacity related attributes.
type CStorPoolInstanceCapacity struct {
	Total resource.Quantity `json:"total"`
	Free  resource.Quantity `json:"free"`
	Used  resource.Quantity `json:"used"`
}

type CSPIConditionType string

// CSPIConditionType describes the state of a CSPI at a certain point.
type CStorPoolInstanceCondition struct {
	// Type of CSPI condition.
	Type CSPIConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=DeploymentConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,6,opt,name=lastUpdateTime"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,7,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=cstorpoolinstance

// CStorPoolInstanceList is a list of CStorPoolInstance resources
type CStorPoolInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CStorPoolInstance `json:"items"`
}