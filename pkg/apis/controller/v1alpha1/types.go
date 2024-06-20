package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Voucher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VoucherSpec   `json:"spec"`
	Status VoucherStatus `json:"status"`
}

// VoucherSpec
type VoucherSpec struct {
	// ResourceCluster is a valid Kubernetes cluster, currently limited to: local.
	ResourceCluster string `json:"resourceCluster"`
	// ResourceNameSpace is a namespace within ResourceCluster.
	//   If the namespace does not existing, the controller will create it.
	// ResourceType is a valid API object, currently limited to pods.
	ResourceType string `json:"resourceType"`
	// BudgetType is a valid metric definition, currently limited to "schedulable"
	// as the number of schedulable pods in a namespace.
	BudgetType string `json:"resourceBudget"`
	// BudgetValue is a valid number related to the BudgetType.
	BudgetValue *int32 `json:"budgetValue"`
}

type VoucherStatus struct {
	// AvailableBudget is BudgetValue minus Allocated
	AvailableBudget int32 `json:"availableBudget"`
	// Currently used budget by ResourceType
	AllocatedBudget int32 `json:"allocatedBudget"`
}

type VoucherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Voucher `json:"items"`
}
