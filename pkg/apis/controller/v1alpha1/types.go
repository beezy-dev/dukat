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

type VoucherSpec struct {
}

type VoucherStatus struct {
}

type VoucherList struct {
}
