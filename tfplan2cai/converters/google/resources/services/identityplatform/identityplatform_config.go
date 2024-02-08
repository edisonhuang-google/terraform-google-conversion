// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package identityplatform

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IdentityPlatformConfigAssetType string = "identitytoolkit.googleapis.com/Config"

func ResourceConverterIdentityPlatformConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IdentityPlatformConfigAssetType,
		Convert:   GetIdentityPlatformConfigCaiObject,
	}
}

func GetIdentityPlatformConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/config")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIdentityPlatformConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IdentityPlatformConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "Config",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIdentityPlatformConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	autodeleteAnonymousUsersProp, err := expandIdentityPlatformConfigAutodeleteAnonymousUsers(d.Get("autodelete_anonymous_users"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autodelete_anonymous_users"); !tpgresource.IsEmptyValue(reflect.ValueOf(autodeleteAnonymousUsersProp)) && (ok || !reflect.DeepEqual(v, autodeleteAnonymousUsersProp)) {
		obj["autodeleteAnonymousUsers"] = autodeleteAnonymousUsersProp
	}
	signInProp, err := expandIdentityPlatformConfigSignIn(d.Get("sign_in"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sign_in"); !tpgresource.IsEmptyValue(reflect.ValueOf(signInProp)) && (ok || !reflect.DeepEqual(v, signInProp)) {
		obj["signIn"] = signInProp
	}
	blockingFunctionsProp, err := expandIdentityPlatformConfigBlockingFunctions(d.Get("blocking_functions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("blocking_functions"); !tpgresource.IsEmptyValue(reflect.ValueOf(blockingFunctionsProp)) && (ok || !reflect.DeepEqual(v, blockingFunctionsProp)) {
		obj["blockingFunctions"] = blockingFunctionsProp
	}
	quotaProp, err := expandIdentityPlatformConfigQuota(d.Get("quota"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("quota"); !tpgresource.IsEmptyValue(reflect.ValueOf(quotaProp)) && (ok || !reflect.DeepEqual(v, quotaProp)) {
		obj["quota"] = quotaProp
	}
	authorizedDomainsProp, err := expandIdentityPlatformConfigAuthorizedDomains(d.Get("authorized_domains"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_domains"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizedDomainsProp)) && (ok || !reflect.DeepEqual(v, authorizedDomainsProp)) {
		obj["authorizedDomains"] = authorizedDomainsProp
	}
	smsRegionConfigProp, err := expandIdentityPlatformConfigSmsRegionConfig(d.Get("sms_region_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sms_region_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(smsRegionConfigProp)) && (ok || !reflect.DeepEqual(v, smsRegionConfigProp)) {
		obj["smsRegionConfig"] = smsRegionConfigProp
	}
	clientProp, err := expandIdentityPlatformConfigClient(d.Get("client"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientProp)) && (ok || !reflect.DeepEqual(v, clientProp)) {
		obj["client"] = clientProp
	}
	mfaProp, err := expandIdentityPlatformConfigMfa(d.Get("mfa"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mfa"); !tpgresource.IsEmptyValue(reflect.ValueOf(mfaProp)) && (ok || !reflect.DeepEqual(v, mfaProp)) {
		obj["mfa"] = mfaProp
	}
	multiTenantProp, err := expandIdentityPlatformConfigMultiTenant(d.Get("multi_tenant"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("multi_tenant"); !tpgresource.IsEmptyValue(reflect.ValueOf(multiTenantProp)) && (ok || !reflect.DeepEqual(v, multiTenantProp)) {
		obj["multiTenant"] = multiTenantProp
	}
	monitoringProp, err := expandIdentityPlatformConfigMonitoring(d.Get("monitoring"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("monitoring"); !tpgresource.IsEmptyValue(reflect.ValueOf(monitoringProp)) && (ok || !reflect.DeepEqual(v, monitoringProp)) {
		obj["monitoring"] = monitoringProp
	}

	return obj, nil
}

func expandIdentityPlatformConfigAutodeleteAnonymousUsers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignIn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandIdentityPlatformConfigSignInEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	transformedPhoneNumber, err := expandIdentityPlatformConfigSignInPhoneNumber(original["phone_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPhoneNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["phoneNumber"] = transformedPhoneNumber
	}

	transformedAnonymous, err := expandIdentityPlatformConfigSignInAnonymous(original["anonymous"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnonymous); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["anonymous"] = transformedAnonymous
	}

	transformedAllowDuplicateEmails, err := expandIdentityPlatformConfigSignInAllowDuplicateEmails(original["allow_duplicate_emails"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowDuplicateEmails); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowDuplicateEmails"] = transformedAllowDuplicateEmails
	}

	transformedHashConfig, err := expandIdentityPlatformConfigSignInHashConfig(original["hash_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHashConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hashConfig"] = transformedHashConfig
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformConfigSignInEmailEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedPasswordRequired, err := expandIdentityPlatformConfigSignInEmailPasswordRequired(original["password_required"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPasswordRequired); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["passwordRequired"] = transformedPasswordRequired
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInEmailEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInEmailPasswordRequired(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInPhoneNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformConfigSignInPhoneNumberEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedTestPhoneNumbers, err := expandIdentityPlatformConfigSignInPhoneNumberTestPhoneNumbers(original["test_phone_numbers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTestPhoneNumbers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["testPhoneNumbers"] = transformedTestPhoneNumbers
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInPhoneNumberEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInPhoneNumberTestPhoneNumbers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandIdentityPlatformConfigSignInAnonymous(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformConfigSignInAnonymousEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInAnonymousEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInAllowDuplicateEmails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInHashConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAlgorithm, err := expandIdentityPlatformConfigSignInHashConfigAlgorithm(original["algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["algorithm"] = transformedAlgorithm
	}

	transformedSignerKey, err := expandIdentityPlatformConfigSignInHashConfigSignerKey(original["signer_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignerKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signerKey"] = transformedSignerKey
	}

	transformedSaltSeparator, err := expandIdentityPlatformConfigSignInHashConfigSaltSeparator(original["salt_separator"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSaltSeparator); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["saltSeparator"] = transformedSaltSeparator
	}

	transformedRounds, err := expandIdentityPlatformConfigSignInHashConfigRounds(original["rounds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRounds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rounds"] = transformedRounds
	}

	transformedMemoryCost, err := expandIdentityPlatformConfigSignInHashConfigMemoryCost(original["memory_cost"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemoryCost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memoryCost"] = transformedMemoryCost
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSignInHashConfigAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInHashConfigSignerKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInHashConfigSaltSeparator(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInHashConfigRounds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSignInHashConfigMemoryCost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTriggers, err := expandIdentityPlatformConfigBlockingFunctionsTriggers(original["triggers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTriggers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["triggers"] = transformedTriggers
	}

	transformedForwardInboundCredentials, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials(original["forward_inbound_credentials"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedForwardInboundCredentials); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["forwardInboundCredentials"] = transformedForwardInboundCredentials
	}

	return transformed, nil
}

func expandIdentityPlatformConfigBlockingFunctionsTriggers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFunctionUri, err := expandIdentityPlatformConfigBlockingFunctionsTriggersFunctionUri(original["function_uri"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFunctionUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["functionUri"] = transformedFunctionUri
		}

		transformedUpdateTime, err := expandIdentityPlatformConfigBlockingFunctionsTriggersUpdateTime(original["update_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["updateTime"] = transformedUpdateTime
		}

		transformedEventType, err := tpgresource.ExpandString(original["event_type"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedEventType] = transformed
	}
	return m, nil
}

func expandIdentityPlatformConfigBlockingFunctionsTriggersFunctionUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctionsTriggersUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdToken, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsIdToken(original["id_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idToken"] = transformedIdToken
	}

	transformedAccessToken, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsAccessToken(original["access_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accessToken"] = transformedAccessToken
	}

	transformedRefreshToken, err := expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsRefreshToken(original["refresh_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRefreshToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["refreshToken"] = transformedRefreshToken
	}

	return transformed, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsIdToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsAccessToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsRefreshToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigQuota(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSignUpQuotaConfig, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfig(original["sign_up_quota_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignUpQuotaConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signUpQuotaConfig"] = transformedSignUpQuotaConfig
	}

	return transformed, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedQuota, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuota(original["quota"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQuota); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["quota"] = transformedQuota
	}

	transformedStartTime, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfigStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedQuotaDuration, err := expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuotaDuration(original["quota_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQuotaDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["quotaDuration"] = transformedQuotaDuration
	}

	return transformed, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuota(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfigStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigQuotaSignUpQuotaConfigQuotaDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigAuthorizedDomains(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSmsRegionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowByDefault, err := expandIdentityPlatformConfigSmsRegionConfigAllowByDefault(original["allow_by_default"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowByDefault); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowByDefault"] = transformedAllowByDefault
	}

	transformedAllowlistOnly, err := expandIdentityPlatformConfigSmsRegionConfigAllowlistOnly(original["allowlist_only"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowlistOnly); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowlistOnly"] = transformedAllowlistOnly
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowByDefault(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisallowedRegions, err := expandIdentityPlatformConfigSmsRegionConfigAllowByDefaultDisallowedRegions(original["disallowed_regions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisallowedRegions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disallowedRegions"] = transformedDisallowedRegions
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowByDefaultDisallowedRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowlistOnly(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedRegions, err := expandIdentityPlatformConfigSmsRegionConfigAllowlistOnlyAllowedRegions(original["allowed_regions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedRegions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedRegions"] = transformedAllowedRegions
	}

	return transformed, nil
}

func expandIdentityPlatformConfigSmsRegionConfigAllowlistOnlyAllowedRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigClient(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPermissions, err := expandIdentityPlatformConfigClientPermissions(original["permissions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPermissions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["permissions"] = transformedPermissions
	}

	transformedApiKey, err := expandIdentityPlatformConfigClientApiKey(original["api_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApiKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["apiKey"] = transformedApiKey
	}

	transformedFirebaseSubdomain, err := expandIdentityPlatformConfigClientFirebaseSubdomain(original["firebase_subdomain"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFirebaseSubdomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["firebaseSubdomain"] = transformedFirebaseSubdomain
	}

	return transformed, nil
}

func expandIdentityPlatformConfigClientPermissions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisabledUserSignup, err := expandIdentityPlatformConfigClientPermissionsDisabledUserSignup(original["disabled_user_signup"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisabledUserSignup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disabledUserSignup"] = transformedDisabledUserSignup
	}

	transformedDisabledUserDeletion, err := expandIdentityPlatformConfigClientPermissionsDisabledUserDeletion(original["disabled_user_deletion"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisabledUserDeletion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disabledUserDeletion"] = transformedDisabledUserDeletion
	}

	return transformed, nil
}

func expandIdentityPlatformConfigClientPermissionsDisabledUserSignup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigClientPermissionsDisabledUserDeletion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigClientApiKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigClientFirebaseSubdomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigMfa(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedState, err := expandIdentityPlatformConfigMfaState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedEnabledProviders, err := expandIdentityPlatformConfigMfaEnabledProviders(original["enabled_providers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabledProviders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabledProviders"] = transformedEnabledProviders
	}

	transformedProviderConfigs, err := expandIdentityPlatformConfigMfaProviderConfigs(original["provider_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProviderConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["providerConfigs"] = transformedProviderConfigs
	}

	return transformed, nil
}

func expandIdentityPlatformConfigMfaState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigMfaEnabledProviders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigMfaProviderConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedState, err := expandIdentityPlatformConfigMfaProviderConfigsState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["state"] = transformedState
		}

		transformedTotpProviderConfig, err := expandIdentityPlatformConfigMfaProviderConfigsTotpProviderConfig(original["totp_provider_config"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTotpProviderConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["totpProviderConfig"] = transformedTotpProviderConfig
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIdentityPlatformConfigMfaProviderConfigsState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigMfaProviderConfigsTotpProviderConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdjacentIntervals, err := expandIdentityPlatformConfigMfaProviderConfigsTotpProviderConfigAdjacentIntervals(original["adjacent_intervals"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdjacentIntervals); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["adjacentIntervals"] = transformedAdjacentIntervals
	}

	return transformed, nil
}

func expandIdentityPlatformConfigMfaProviderConfigsTotpProviderConfigAdjacentIntervals(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigMultiTenant(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowTenants, err := expandIdentityPlatformConfigMultiTenantAllowTenants(original["allow_tenants"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowTenants); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowTenants"] = transformedAllowTenants
	}

	transformedDefaultTenantLocation, err := expandIdentityPlatformConfigMultiTenantDefaultTenantLocation(original["default_tenant_location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultTenantLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultTenantLocation"] = transformedDefaultTenantLocation
	}

	return transformed, nil
}

func expandIdentityPlatformConfigMultiTenantAllowTenants(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigMultiTenantDefaultTenantLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformConfigMonitoring(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequestLogging, err := expandIdentityPlatformConfigMonitoringRequestLogging(original["request_logging"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["requestLogging"] = transformedRequestLogging
	}

	return transformed, nil
}

func expandIdentityPlatformConfigMonitoringRequestLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformConfigMonitoringRequestLoggingEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandIdentityPlatformConfigMonitoringRequestLoggingEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
