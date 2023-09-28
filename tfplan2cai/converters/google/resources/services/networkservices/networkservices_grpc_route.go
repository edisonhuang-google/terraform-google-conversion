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

package networkservices

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkServicesGrpcRouteAssetType string = "networkservices.googleapis.com/GrpcRoute"

func ResourceConverterNetworkServicesGrpcRoute() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkServicesGrpcRouteAssetType,
		Convert:   GetNetworkServicesGrpcRouteCaiObject,
	}
}

func GetNetworkServicesGrpcRouteCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkservices.googleapis.com/projects/{{project}}/locations/global/grpcRoutes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkServicesGrpcRouteApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkServicesGrpcRouteAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkservices/v1/rest",
				DiscoveryName:        "GrpcRoute",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkServicesGrpcRouteApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesGrpcRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	hostnamesProp, err := expandNetworkServicesGrpcRouteHostnames(d.Get("hostnames"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("hostnames"); !tpgresource.IsEmptyValue(reflect.ValueOf(hostnamesProp)) && (ok || !reflect.DeepEqual(v, hostnamesProp)) {
		obj["hostnames"] = hostnamesProp
	}
	meshesProp, err := expandNetworkServicesGrpcRouteMeshes(d.Get("meshes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("meshes"); !tpgresource.IsEmptyValue(reflect.ValueOf(meshesProp)) && (ok || !reflect.DeepEqual(v, meshesProp)) {
		obj["meshes"] = meshesProp
	}
	gatewaysProp, err := expandNetworkServicesGrpcRouteGateways(d.Get("gateways"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gateways"); !tpgresource.IsEmptyValue(reflect.ValueOf(gatewaysProp)) && (ok || !reflect.DeepEqual(v, gatewaysProp)) {
		obj["gateways"] = gatewaysProp
	}
	rulesProp, err := expandNetworkServicesGrpcRouteRules(d.Get("rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}
	labelsProp, err := expandNetworkServicesGrpcRouteEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkServicesGrpcRouteDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteHostnames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteMeshes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteGateways(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMatches, err := expandNetworkServicesGrpcRouteRulesMatches(original["matches"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMatches); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["matches"] = transformedMatches
		}

		transformedAction, err := expandNetworkServicesGrpcRouteRulesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesGrpcRouteRulesMatches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedHeaders, err := expandNetworkServicesGrpcRouteRulesMatchesHeaders(original["headers"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["headers"] = transformedHeaders
		}

		transformedMethod, err := expandNetworkServicesGrpcRouteRulesMatchesMethod(original["method"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["method"] = transformedMethod
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandNetworkServicesGrpcRouteRulesMatchesHeadersKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValue, err := expandNetworkServicesGrpcRouteRulesMatchesHeadersValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedType, err := expandNetworkServicesGrpcRouteRulesMatchesHeadersType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesHeadersKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesHeadersValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesHeadersType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGrpcService, err := expandNetworkServicesGrpcRouteRulesMatchesMethodGrpcService(original["grpc_service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrpcService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["grpcService"] = transformedGrpcService
	}

	transformedGrpcMethod, err := expandNetworkServicesGrpcRouteRulesMatchesMethodGrpcMethod(original["grpc_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrpcMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["grpcMethod"] = transformedGrpcMethod
	}

	transformedCaseSensitive, err := expandNetworkServicesGrpcRouteRulesMatchesMethodCaseSensitive(original["case_sensitive"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaseSensitive); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["caseSensitive"] = transformedCaseSensitive
	}

	return transformed, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesMethodGrpcService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesMethodGrpcMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesMatchesMethodCaseSensitive(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDestinations, err := expandNetworkServicesGrpcRouteRulesActionDestinations(original["destinations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestinations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destinations"] = transformedDestinations
	}

	transformedFaultInjectionPolicy, err := expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy(original["fault_injection_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFaultInjectionPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["faultInjectionPolicy"] = transformedFaultInjectionPolicy
	}

	transformedTimeout, err := expandNetworkServicesGrpcRouteRulesActionTimeout(original["timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeout"] = transformedTimeout
	}

	transformedRetryPolicy, err := expandNetworkServicesGrpcRouteRulesActionRetryPolicy(original["retry_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetryPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retryPolicy"] = transformedRetryPolicy
	}

	return transformed, nil
}

func expandNetworkServicesGrpcRouteRulesActionDestinations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedServiceName, err := expandNetworkServicesGrpcRouteRulesActionDestinationsServiceName(original["service_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["serviceName"] = transformedServiceName
		}

		transformedWeight, err := expandNetworkServicesGrpcRouteRulesActionDestinationsWeight(original["weight"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["weight"] = transformedWeight
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesGrpcRouteRulesActionDestinationsServiceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionDestinationsWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDelay, err := expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay(original["delay"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDelay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["delay"] = transformedDelay
	}

	transformedAbort, err := expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort(original["abort"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAbort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["abort"] = transformedAbort
	}

	return transformed, nil
}

func expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixedDelay, err := expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayFixedDelay(original["fixed_delay"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixedDelay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixedDelay"] = transformedFixedDelay
	}

	transformedPercentage, err := expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayPercentage(original["percentage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercentage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["percentage"] = transformedPercentage
	}

	return transformed, nil
}

func expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayFixedDelay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayPercentage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHttpStatus, err := expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortHttpStatus(original["http_status"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHttpStatus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["httpStatus"] = transformedHttpStatus
	}

	transformedPercentage, err := expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortPercentage(original["percentage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercentage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["percentage"] = transformedPercentage
	}

	return transformed, nil
}

func expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortHttpStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortPercentage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionRetryPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRetryConditions, err := expandNetworkServicesGrpcRouteRulesActionRetryPolicyRetryConditions(original["retry_conditions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetryConditions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retryConditions"] = transformedRetryConditions
	}

	transformedNumRetries, err := expandNetworkServicesGrpcRouteRulesActionRetryPolicyNumRetries(original["num_retries"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumRetries); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["numRetries"] = transformedNumRetries
	}

	return transformed, nil
}

func expandNetworkServicesGrpcRouteRulesActionRetryPolicyRetryConditions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteRulesActionRetryPolicyNumRetries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGrpcRouteEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
