// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSI":               schema_pkg_apis_csi_v1_IBMBlockCSI(ref),
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSIControllerSpec": schema_pkg_apis_csi_v1_IBMBlockCSIControllerSpec(ref),
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSINodeSpec":       schema_pkg_apis_csi_v1_IBMBlockCSINodeSpec(ref),
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSISpec":           schema_pkg_apis_csi_v1_IBMBlockCSISpec(ref),
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSIStatus":         schema_pkg_apis_csi_v1_IBMBlockCSIStatus(ref),
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.NodeInfo":                  schema_pkg_apis_csi_v1_NodeInfo(ref),
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.NodeInfoSpec":              schema_pkg_apis_csi_v1_NodeInfoSpec(ref),
		"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.NodeInfoStatus":            schema_pkg_apis_csi_v1_NodeInfoStatus(ref),
	}
}

func schema_pkg_apis_csi_v1_IBMBlockCSI(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IBMBlockCSI is the Schema for the ibmblockcsis API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSISpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSIStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSISpec", "github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSIStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_csi_v1_IBMBlockCSIControllerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IBMBlockCSIControllerSpec defines the desired state of IBMBlockCSIController",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"repository": {
						SchemaProps: spec.SchemaProps{
							Description: "The repository of the controller image",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tag": {
						SchemaProps: spec.SchemaProps{
							Description: "The tag of the controller image",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"repository", "tag"},
			},
		},
	}
}

func schema_pkg_apis_csi_v1_IBMBlockCSINodeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IBMBlockCSINodeSpec defines the desired state of IBMBlockCSINode",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"repository": {
						SchemaProps: spec.SchemaProps{
							Description: "The repository of the node image",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tag": {
						SchemaProps: spec.SchemaProps{
							Description: "The tag of the node image",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"repository", "tag"},
			},
		},
	}
}

func schema_pkg_apis_csi_v1_IBMBlockCSISpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IBMBlockCSISpec defines the desired state of IBMBlockCSI",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"controller": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSIControllerSpec"),
						},
					},
					"node": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSINodeSpec"),
						},
					},
				},
				Required: []string{"controller", "node"},
			},
		},
		Dependencies: []string{
			"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSIControllerSpec", "github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.IBMBlockCSINodeSpec"},
	}
}

func schema_pkg_apis_csi_v1_IBMBlockCSIStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IBMBlockCSIStatus defines the observed state of IBMBlockCSI",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"controllerReady": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"nodeReady": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"ready", "controllerReady", "nodeReady"},
			},
		},
	}
}

func schema_pkg_apis_csi_v1_NodeInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeInfo is the Schema for the nodeinfos API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.NodeInfoSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.NodeInfoStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.NodeInfoSpec", "github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1.NodeInfoStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_csi_v1_NodeInfoSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeInfoSpec defines the desired state of NodeInfo",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_csi_v1_NodeInfoStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeInfoStatus defines the observed state of NodeInfo",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"iqns": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"wwpns": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"definedOnStorages": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"iqns", "wwpns", "definedOnStorages"},
			},
		},
	}
}
