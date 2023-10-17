// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ApplicationPropertyAPIOutType - The type of Risk Cloud application
type ApplicationPropertyAPIOutType string

const (
	ApplicationPropertyAPIOutTypeAccessManagement             ApplicationPropertyAPIOutType = "ACCESS_MANAGEMENT"
	ApplicationPropertyAPIOutTypeAmlKyc                       ApplicationPropertyAPIOutType = "AML_KYC"
	ApplicationPropertyAPIOutTypeAssetManagement              ApplicationPropertyAPIOutType = "ASSET_MANAGEMENT"
	ApplicationPropertyAPIOutTypeAuditManagement              ApplicationPropertyAPIOutType = "AUDIT_MANAGEMENT"
	ApplicationPropertyAPIOutTypeBusinessContinuityPlanning   ApplicationPropertyAPIOutType = "BUSINESS_CONTINUITY_PLANNING"
	ApplicationPropertyAPIOutTypeChangeManagement             ApplicationPropertyAPIOutType = "CHANGE_MANAGEMENT"
	ApplicationPropertyAPIOutTypeComplaintManagement          ApplicationPropertyAPIOutType = "COMPLAINT_MANAGEMENT"
	ApplicationPropertyAPIOutTypeComplianceManagement         ApplicationPropertyAPIOutType = "COMPLIANCE_MANAGEMENT"
	ApplicationPropertyAPIOutTypeContractManagement           ApplicationPropertyAPIOutType = "CONTRACT_MANAGEMENT"
	ApplicationPropertyAPIOutTypeControlsCompliance           ApplicationPropertyAPIOutType = "CONTROLS_COMPLIANCE"
	ApplicationPropertyAPIOutTypeControlsManagement           ApplicationPropertyAPIOutType = "CONTROLS_MANAGEMENT"
	ApplicationPropertyAPIOutTypeCrisisManagement             ApplicationPropertyAPIOutType = "CRISIS_MANAGEMENT"
	ApplicationPropertyAPIOutTypeCustom                       ApplicationPropertyAPIOutType = "CUSTOM"
	ApplicationPropertyAPIOutTypeCustomGrcUseCase             ApplicationPropertyAPIOutType = "CUSTOM_GRC_USE_CASE"
	ApplicationPropertyAPIOutTypeCyberRiskManagement          ApplicationPropertyAPIOutType = "CYBER_RISK_MANAGEMENT"
	ApplicationPropertyAPIOutTypeDataPrivacyManagement        ApplicationPropertyAPIOutType = "DATA_PRIVACY_MANAGEMENT"
	ApplicationPropertyAPIOutTypeEmployeeCompliance           ApplicationPropertyAPIOutType = "EMPLOYEE_COMPLIANCE"
	ApplicationPropertyAPIOutTypeEnterpriseRiskManagement     ApplicationPropertyAPIOutType = "ENTERPRISE_RISK_MANAGEMENT"
	ApplicationPropertyAPIOutTypeEsg                          ApplicationPropertyAPIOutType = "ESG"
	ApplicationPropertyAPIOutTypeIncidentManagement           ApplicationPropertyAPIOutType = "INCIDENT_MANAGEMENT"
	ApplicationPropertyAPIOutTypeInternalAuditManagement      ApplicationPropertyAPIOutType = "INTERNAL_AUDIT_MANAGEMENT"
	ApplicationPropertyAPIOutTypeItRiskManagement             ApplicationPropertyAPIOutType = "IT_RISK_MANAGEMENT"
	ApplicationPropertyAPIOutTypeNone                         ApplicationPropertyAPIOutType = "NONE"
	ApplicationPropertyAPIOutTypeOperationalResiliency        ApplicationPropertyAPIOutType = "OPERATIONAL_RESILIENCY"
	ApplicationPropertyAPIOutTypeOther                        ApplicationPropertyAPIOutType = "OTHER"
	ApplicationPropertyAPIOutTypePolicyAndProcedureManagement ApplicationPropertyAPIOutType = "POLICY_AND_PROCEDURE_MANAGEMENT"
	ApplicationPropertyAPIOutTypePolicyManagement             ApplicationPropertyAPIOutType = "POLICY_MANAGEMENT"
	ApplicationPropertyAPIOutTypePrivacyManagement            ApplicationPropertyAPIOutType = "PRIVACY_MANAGEMENT"
	ApplicationPropertyAPIOutTypeQuantify                     ApplicationPropertyAPIOutType = "QUANTIFY"
	ApplicationPropertyAPIOutTypeRegulatoryCompliance         ApplicationPropertyAPIOutType = "REGULATORY_COMPLIANCE"
	ApplicationPropertyAPIOutTypeRepository                   ApplicationPropertyAPIOutType = "REPOSITORY"
	ApplicationPropertyAPIOutTypeRiskQuantification           ApplicationPropertyAPIOutType = "RISK_QUANTIFICATION"
	ApplicationPropertyAPIOutTypeSoxTesting                   ApplicationPropertyAPIOutType = "SOX_TESTING"
	ApplicationPropertyAPIOutTypeStandardsRegulations         ApplicationPropertyAPIOutType = "STANDARDS_REGULATIONS"
	ApplicationPropertyAPIOutTypeSurvey                       ApplicationPropertyAPIOutType = "SURVEY"
	ApplicationPropertyAPIOutTypeThirdPartyRiskManagement     ApplicationPropertyAPIOutType = "THIRD_PARTY_RISK_MANAGEMENT"
)

func (e ApplicationPropertyAPIOutType) ToPointer() *ApplicationPropertyAPIOutType {
	return &e
}

func (e *ApplicationPropertyAPIOutType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACCESS_MANAGEMENT":
		fallthrough
	case "AML_KYC":
		fallthrough
	case "ASSET_MANAGEMENT":
		fallthrough
	case "AUDIT_MANAGEMENT":
		fallthrough
	case "BUSINESS_CONTINUITY_PLANNING":
		fallthrough
	case "CHANGE_MANAGEMENT":
		fallthrough
	case "COMPLAINT_MANAGEMENT":
		fallthrough
	case "COMPLIANCE_MANAGEMENT":
		fallthrough
	case "CONTRACT_MANAGEMENT":
		fallthrough
	case "CONTROLS_COMPLIANCE":
		fallthrough
	case "CONTROLS_MANAGEMENT":
		fallthrough
	case "CRISIS_MANAGEMENT":
		fallthrough
	case "CUSTOM":
		fallthrough
	case "CUSTOM_GRC_USE_CASE":
		fallthrough
	case "CYBER_RISK_MANAGEMENT":
		fallthrough
	case "DATA_PRIVACY_MANAGEMENT":
		fallthrough
	case "EMPLOYEE_COMPLIANCE":
		fallthrough
	case "ENTERPRISE_RISK_MANAGEMENT":
		fallthrough
	case "ESG":
		fallthrough
	case "INCIDENT_MANAGEMENT":
		fallthrough
	case "INTERNAL_AUDIT_MANAGEMENT":
		fallthrough
	case "IT_RISK_MANAGEMENT":
		fallthrough
	case "NONE":
		fallthrough
	case "OPERATIONAL_RESILIENCY":
		fallthrough
	case "OTHER":
		fallthrough
	case "POLICY_AND_PROCEDURE_MANAGEMENT":
		fallthrough
	case "POLICY_MANAGEMENT":
		fallthrough
	case "PRIVACY_MANAGEMENT":
		fallthrough
	case "QUANTIFY":
		fallthrough
	case "REGULATORY_COMPLIANCE":
		fallthrough
	case "REPOSITORY":
		fallthrough
	case "RISK_QUANTIFICATION":
		fallthrough
	case "SOX_TESTING":
		fallthrough
	case "STANDARDS_REGULATIONS":
		fallthrough
	case "SURVEY":
		fallthrough
	case "THIRD_PARTY_RISK_MANAGEMENT":
		*e = ApplicationPropertyAPIOutType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApplicationPropertyAPIOutType: %v", v)
	}
}

// ApplicationPropertyAPIOut - The parent application of the record
type ApplicationPropertyAPIOut struct {
	// The unique ID of this Risk Cloud resource
	ID *string `json:"id,omitempty"`
	// The name of the application
	Name *string `json:"name,omitempty"`
	// Identifies the type of object this data represents
	Object *string `json:"object,omitempty"`
	// The type of Risk Cloud application
	Type *ApplicationPropertyAPIOutType `json:"type,omitempty"`
}

func (o *ApplicationPropertyAPIOut) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ApplicationPropertyAPIOut) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ApplicationPropertyAPIOut) GetObject() *string {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ApplicationPropertyAPIOut) GetType() *ApplicationPropertyAPIOutType {
	if o == nil {
		return nil
	}
	return o.Type
}