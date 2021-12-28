// Package Caas provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package caas

import (
	"time"
)

// Account defines model for Account.
type Account struct {
	CreationDate *time.Time `json:"creationDate,omitempty"`
	Id           *int       `json:"id,omitempty"`
	Tags         *[]string  `json:"tags,omitempty"`
	Text         *string    `json:"text,omitempty"`
}

// Cluster defines model for Cluster.
type Cluster struct {
	AccountId    *string    `json:"accountId,omitempty"`
	CreationDate *time.Time `json:"creationDate,omitempty"`
	Id           *int       `json:"id,omitempty"`
	Kubeconfig   *string    `json:"kubeconfig,omitempty"`
	Tags         *[]string  `json:"tags,omitempty"`
	Text         *string    `json:"text,omitempty"`
}

// PostAccountJSONBody defines parameters for PostAccount.
type PostAccountJSONBody struct {
	Due  *time.Time `json:"due,omitempty"`
	Tags *[]string  `json:"tags,omitempty"`
	Text *string    `json:"text,omitempty"`
}

// PostClusterJSONBody defines parameters for PostCluster.
type PostClusterJSONBody struct {
	Due  *time.Time `json:"due,omitempty"`
	Tags *[]string  `json:"tags,omitempty"`
	Text *string    `json:"text,omitempty"`
}

// GetClusterFindByAccountIdParams defines parameters for GetClusterFindByAccountId.
type GetClusterFindByAccountIdParams struct {
	AccountId string `json:"accountId"`
}

// PostAccountJSONRequestBody sdefines body for PostAccount for application/json ContentType.
type PostAccountJSONRequestBody PostAccountJSONBody

// PostClusterJSONRequestBody defines body for PostCluster for application/json ContentType.
type PostClusterJSONRequestBody PostClusterJSONBody

