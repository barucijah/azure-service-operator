// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "encoding/json"

type Secret_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the Secret to create.
	Properties *SecretProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The JSON object that contains the properties of the Secret to create.
type SecretProperties_STATUS struct {
	DeploymentStatus *SecretProperties_DeploymentStatus_STATUS `json:"deploymentStatus,omitempty"`

	// Parameters: object which contains secret parameters
	Parameters *SecretParameters_STATUS `json:"parameters,omitempty"`

	// ProfileName: The name of the profile which holds the secret.
	ProfileName *string `json:"profileName,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *SecretProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type SecretParameters_STATUS struct {
	// AzureFirstPartyManagedCertificate: Mutually exclusive with all other properties
	AzureFirstPartyManagedCertificate *AzureFirstPartyManagedCertificateParameters_STATUS `json:"azureFirstPartyManagedCertificate,omitempty"`

	// CustomerCertificate: Mutually exclusive with all other properties
	CustomerCertificate *CustomerCertificateParameters_STATUS `json:"customerCertificate,omitempty"`

	// ManagedCertificate: Mutually exclusive with all other properties
	ManagedCertificate *ManagedCertificateParameters_STATUS `json:"managedCertificate,omitempty"`

	// UrlSigningKey: Mutually exclusive with all other properties
	UrlSigningKey *UrlSigningKeyParameters_STATUS `json:"urlSigningKey,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because SecretParameters_STATUS represents a discriminated union (JSON OneOf)
func (parameters SecretParameters_STATUS) MarshalJSON() ([]byte, error) {
	if parameters.AzureFirstPartyManagedCertificate != nil {
		return json.Marshal(parameters.AzureFirstPartyManagedCertificate)
	}

	if parameters.CustomerCertificate != nil {
		return json.Marshal(parameters.CustomerCertificate)
	}

	if parameters.ManagedCertificate != nil {
		return json.Marshal(parameters.ManagedCertificate)
	}

	if parameters.UrlSigningKey != nil {
		return json.Marshal(parameters.UrlSigningKey)
	}

	return nil, nil
}

// UnmarshalJSON unmarshals the SecretParameters_STATUS
func (parameters *SecretParameters_STATUS) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "AzureFirstPartyManagedCertificate" {
		parameters.AzureFirstPartyManagedCertificate = &AzureFirstPartyManagedCertificateParameters_STATUS{}
		return json.Unmarshal(data, parameters.AzureFirstPartyManagedCertificate)
	}
	if discriminator == "CustomerCertificate" {
		parameters.CustomerCertificate = &CustomerCertificateParameters_STATUS{}
		return json.Unmarshal(data, parameters.CustomerCertificate)
	}
	if discriminator == "ManagedCertificate" {
		parameters.ManagedCertificate = &ManagedCertificateParameters_STATUS{}
		return json.Unmarshal(data, parameters.ManagedCertificate)
	}
	if discriminator == "UrlSigningKey" {
		parameters.UrlSigningKey = &UrlSigningKeyParameters_STATUS{}
		return json.Unmarshal(data, parameters.UrlSigningKey)
	}

	// No error
	return nil
}

type SecretProperties_DeploymentStatus_STATUS string

const (
	SecretProperties_DeploymentStatus_STATUS_Failed     = SecretProperties_DeploymentStatus_STATUS("Failed")
	SecretProperties_DeploymentStatus_STATUS_InProgress = SecretProperties_DeploymentStatus_STATUS("InProgress")
	SecretProperties_DeploymentStatus_STATUS_NotStarted = SecretProperties_DeploymentStatus_STATUS("NotStarted")
	SecretProperties_DeploymentStatus_STATUS_Succeeded  = SecretProperties_DeploymentStatus_STATUS("Succeeded")
)

// Mapping from string to SecretProperties_DeploymentStatus_STATUS
var secretProperties_DeploymentStatus_STATUS_Values = map[string]SecretProperties_DeploymentStatus_STATUS{
	"failed":     SecretProperties_DeploymentStatus_STATUS_Failed,
	"inprogress": SecretProperties_DeploymentStatus_STATUS_InProgress,
	"notstarted": SecretProperties_DeploymentStatus_STATUS_NotStarted,
	"succeeded":  SecretProperties_DeploymentStatus_STATUS_Succeeded,
}

type SecretProperties_ProvisioningState_STATUS string

const (
	SecretProperties_ProvisioningState_STATUS_Creating  = SecretProperties_ProvisioningState_STATUS("Creating")
	SecretProperties_ProvisioningState_STATUS_Deleting  = SecretProperties_ProvisioningState_STATUS("Deleting")
	SecretProperties_ProvisioningState_STATUS_Failed    = SecretProperties_ProvisioningState_STATUS("Failed")
	SecretProperties_ProvisioningState_STATUS_Succeeded = SecretProperties_ProvisioningState_STATUS("Succeeded")
	SecretProperties_ProvisioningState_STATUS_Updating  = SecretProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to SecretProperties_ProvisioningState_STATUS
var secretProperties_ProvisioningState_STATUS_Values = map[string]SecretProperties_ProvisioningState_STATUS{
	"creating":  SecretProperties_ProvisioningState_STATUS_Creating,
	"deleting":  SecretProperties_ProvisioningState_STATUS_Deleting,
	"failed":    SecretProperties_ProvisioningState_STATUS_Failed,
	"succeeded": SecretProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  SecretProperties_ProvisioningState_STATUS_Updating,
}

type AzureFirstPartyManagedCertificateParameters_STATUS struct {
	// CertificateAuthority: Certificate issuing authority.
	CertificateAuthority *string `json:"certificateAuthority,omitempty"`

	// ExpirationDate: Certificate expiration date.
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// SecretSource: Resource reference to the Azure Key Vault certificate. Expected to be in format of
	// /subscriptions/{​​​​​​​​​subscriptionId}​​​​​​​​​/resourceGroups/{​​​​​​​​​resourceGroupName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/providers/Microsoft.KeyVault/vaults/{vaultName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/secrets/{certificateName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​
	SecretSource *ResourceReference_STATUS `json:"secretSource,omitempty"`

	// Subject: Subject name in the certificate.
	Subject *string `json:"subject,omitempty"`

	// SubjectAlternativeNames: The list of SANs.
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`

	// Thumbprint: Certificate thumbprint.
	Thumbprint *string                                                 `json:"thumbprint,omitempty"`
	Type       AzureFirstPartyManagedCertificateParameters_Type_STATUS `json:"type,omitempty"`
}

type CustomerCertificateParameters_STATUS struct {
	// CertificateAuthority: Certificate issuing authority.
	CertificateAuthority *string `json:"certificateAuthority,omitempty"`

	// ExpirationDate: Certificate expiration date.
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// SecretSource: Resource reference to the Azure Key Vault certificate. Expected to be in format of
	// /subscriptions/{​​​​​​​​​subscriptionId}​​​​​​​​​/resourceGroups/{​​​​​​​​​resourceGroupName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/providers/Microsoft.KeyVault/vaults/{vaultName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/secrets/{certificateName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​
	SecretSource *ResourceReference_STATUS `json:"secretSource,omitempty"`

	// SecretVersion: Version of the secret to be used
	SecretVersion *string `json:"secretVersion,omitempty"`

	// Subject: Subject name in the certificate.
	Subject *string `json:"subject,omitempty"`

	// SubjectAlternativeNames: The list of SANs.
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`

	// Thumbprint: Certificate thumbprint.
	Thumbprint *string                                   `json:"thumbprint,omitempty"`
	Type       CustomerCertificateParameters_Type_STATUS `json:"type,omitempty"`

	// UseLatestVersion: Whether to use the latest version for the certificate
	UseLatestVersion *bool `json:"useLatestVersion,omitempty"`
}

type ManagedCertificateParameters_STATUS struct {
	// ExpirationDate: Certificate expiration date.
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Subject: Subject name in the certificate.
	Subject *string                                  `json:"subject,omitempty"`
	Type    ManagedCertificateParameters_Type_STATUS `json:"type,omitempty"`
}

type UrlSigningKeyParameters_STATUS struct {
	// KeyId: Defines the customer defined key Id. This id will exist in the incoming request to indicate the key used to form
	// the hash.
	KeyId *string `json:"keyId,omitempty"`

	// SecretSource: Resource reference to the Azure Key Vault secret. Expected to be in format of
	// /subscriptions/{​​​​​​​​​subscriptionId}​​​​​​​​​/resourceGroups/{​​​​​​​​​resourceGroupName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/providers/Microsoft.KeyVault/vaults/{vaultName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/secrets/{secretName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​
	SecretSource *ResourceReference_STATUS `json:"secretSource,omitempty"`

	// SecretVersion: Version of the secret to be used
	SecretVersion *string                             `json:"secretVersion,omitempty"`
	Type          UrlSigningKeyParameters_Type_STATUS `json:"type,omitempty"`
}

type AzureFirstPartyManagedCertificateParameters_Type_STATUS string

const AzureFirstPartyManagedCertificateParameters_Type_STATUS_AzureFirstPartyManagedCertificate = AzureFirstPartyManagedCertificateParameters_Type_STATUS("AzureFirstPartyManagedCertificate")

// Mapping from string to AzureFirstPartyManagedCertificateParameters_Type_STATUS
var azureFirstPartyManagedCertificateParameters_Type_STATUS_Values = map[string]AzureFirstPartyManagedCertificateParameters_Type_STATUS{
	"azurefirstpartymanagedcertificate": AzureFirstPartyManagedCertificateParameters_Type_STATUS_AzureFirstPartyManagedCertificate,
}

type CustomerCertificateParameters_Type_STATUS string

const CustomerCertificateParameters_Type_STATUS_CustomerCertificate = CustomerCertificateParameters_Type_STATUS("CustomerCertificate")

// Mapping from string to CustomerCertificateParameters_Type_STATUS
var customerCertificateParameters_Type_STATUS_Values = map[string]CustomerCertificateParameters_Type_STATUS{
	"customercertificate": CustomerCertificateParameters_Type_STATUS_CustomerCertificate,
}

type ManagedCertificateParameters_Type_STATUS string

const ManagedCertificateParameters_Type_STATUS_ManagedCertificate = ManagedCertificateParameters_Type_STATUS("ManagedCertificate")

// Mapping from string to ManagedCertificateParameters_Type_STATUS
var managedCertificateParameters_Type_STATUS_Values = map[string]ManagedCertificateParameters_Type_STATUS{
	"managedcertificate": ManagedCertificateParameters_Type_STATUS_ManagedCertificate,
}

type UrlSigningKeyParameters_Type_STATUS string

const UrlSigningKeyParameters_Type_STATUS_UrlSigningKey = UrlSigningKeyParameters_Type_STATUS("UrlSigningKey")

// Mapping from string to UrlSigningKeyParameters_Type_STATUS
var urlSigningKeyParameters_Type_STATUS_Values = map[string]UrlSigningKeyParameters_Type_STATUS{
	"urlsigningkey": UrlSigningKeyParameters_Type_STATUS_UrlSigningKey,
}
