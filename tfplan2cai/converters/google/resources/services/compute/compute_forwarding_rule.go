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

package compute

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeForwardingRuleAssetType string = "compute.googleapis.com/ForwardingRule"

func ResourceConverterComputeForwardingRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeForwardingRuleAssetType,
		Convert:   GetComputeForwardingRuleCaiObject,
	}
}

func GetComputeForwardingRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeForwardingRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeForwardingRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "ForwardingRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeForwardingRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	isMirroringCollectorProp, err := expandComputeForwardingRuleIsMirroringCollector(d.Get("is_mirroring_collector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_mirroring_collector"); !tpgresource.IsEmptyValue(reflect.ValueOf(isMirroringCollectorProp)) && (ok || !reflect.DeepEqual(v, isMirroringCollectorProp)) {
		obj["isMirroringCollector"] = isMirroringCollectorProp
	}
	descriptionProp, err := expandComputeForwardingRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPAddressProp, err := expandComputeForwardingRuleIPAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_address"); !tpgresource.IsEmptyValue(reflect.ValueOf(IPAddressProp)) && (ok || !reflect.DeepEqual(v, IPAddressProp)) {
		obj["IPAddress"] = IPAddressProp
	}
	IPProtocolProp, err := expandComputeForwardingRuleIPProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(IPProtocolProp)) && (ok || !reflect.DeepEqual(v, IPProtocolProp)) {
		obj["IPProtocol"] = IPProtocolProp
	}
	backendServiceProp, err := expandComputeForwardingRuleBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backend_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(backendServiceProp)) && (ok || !reflect.DeepEqual(v, backendServiceProp)) {
		obj["backendService"] = backendServiceProp
	}
	loadBalancingSchemeProp, err := expandComputeForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !tpgresource.IsEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	nameProp, err := expandComputeForwardingRuleName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeForwardingRuleNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	portRangeProp, err := expandComputeForwardingRulePortRange(d.Get("port_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(portRangeProp)) && (ok || !reflect.DeepEqual(v, portRangeProp)) {
		obj["portRange"] = portRangeProp
	}
	portsProp, err := expandComputeForwardingRulePorts(d.Get("ports"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ports"); !tpgresource.IsEmptyValue(reflect.ValueOf(portsProp)) && (ok || !reflect.DeepEqual(v, portsProp)) {
		obj["ports"] = portsProp
	}
	subnetworkProp, err := expandComputeForwardingRuleSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	targetProp, err := expandComputeForwardingRuleTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	allowGlobalAccessProp, err := expandComputeForwardingRuleAllowGlobalAccess(d.Get("allow_global_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_global_access"); ok || !reflect.DeepEqual(v, allowGlobalAccessProp) {
		obj["allowGlobalAccess"] = allowGlobalAccessProp
	}
	labelFingerprintProp, err := expandComputeForwardingRuleLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	allPortsProp, err := expandComputeForwardingRuleAllPorts(d.Get("all_ports"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("all_ports"); !tpgresource.IsEmptyValue(reflect.ValueOf(allPortsProp)) && (ok || !reflect.DeepEqual(v, allPortsProp)) {
		obj["allPorts"] = allPortsProp
	}
	networkTierProp, err := expandComputeForwardingRuleNetworkTier(d.Get("network_tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_tier"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkTierProp)) && (ok || !reflect.DeepEqual(v, networkTierProp)) {
		obj["networkTier"] = networkTierProp
	}
	serviceDirectoryRegistrationsProp, err := expandComputeForwardingRuleServiceDirectoryRegistrations(d.Get("service_directory_registrations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_directory_registrations"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceDirectoryRegistrationsProp)) && (ok || !reflect.DeepEqual(v, serviceDirectoryRegistrationsProp)) {
		obj["serviceDirectoryRegistrations"] = serviceDirectoryRegistrationsProp
	}
	serviceLabelProp, err := expandComputeForwardingRuleServiceLabel(d.Get("service_label"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_label"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceLabelProp)) && (ok || !reflect.DeepEqual(v, serviceLabelProp)) {
		obj["serviceLabel"] = serviceLabelProp
	}
	sourceIpRangesProp, err := expandComputeForwardingRuleSourceIpRanges(d.Get("source_ip_ranges"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_ip_ranges"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceIpRangesProp)) && (ok || !reflect.DeepEqual(v, sourceIpRangesProp)) {
		obj["sourceIpRanges"] = sourceIpRangesProp
	}
	allowPscGlobalAccessProp, err := expandComputeForwardingRuleAllowPscGlobalAccess(d.Get("allow_psc_global_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_psc_global_access"); ok || !reflect.DeepEqual(v, allowPscGlobalAccessProp) {
		obj["allowPscGlobalAccess"] = allowPscGlobalAccessProp
	}
	noAutomateDnsZoneProp, err := expandComputeForwardingRuleNoAutomateDnsZone(d.Get("no_automate_dns_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("no_automate_dns_zone"); ok || !reflect.DeepEqual(v, noAutomateDnsZoneProp) {
		obj["noAutomateDnsZone"] = noAutomateDnsZoneProp
	}
	ipVersionProp, err := expandComputeForwardingRuleIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipVersionProp)) && (ok || !reflect.DeepEqual(v, ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	labelsProp, err := expandComputeForwardingRuleEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	regionProp, err := expandComputeForwardingRuleRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeForwardingRuleIsMirroringCollector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleIPAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleIPProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleBackendService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	// This method returns a full self link from a partial self link.
	if v == nil || v.(string) == "" {
		// It does not try to construct anything from empty.
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		// Anything that starts with a URL scheme is assumed to be a self link worth using.
		return v, nil
	} else if strings.HasPrefix(v.(string), "projects/") {
		// If the self link references a project, we'll just stuck the compute prefix on it
		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
		if err != nil {
			return "", err
		}
		return url, nil
	} else if strings.HasPrefix(v.(string), "regions/") || strings.HasPrefix(v.(string), "zones/") {
		// For regional or zonal resources which include their region or zone, just put the project in front.
		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/")
		if err != nil {
			return nil, err
		}
		return url + v.(string), nil
	}
	// Anything else is assumed to be a regional resource, with a partial link that begins with the resource name.
	// This isn't very likely - it's a last-ditch effort to extract something useful here.  We can do a better job
	// as soon as MultiResourceRefs are working since we'll know the types that this field is supposed to point to.
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/")
	if err != nil {
		return nil, err
	}
	return url + v.(string), nil
}

func expandComputeForwardingRuleLoadBalancingScheme(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeForwardingRulePortRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRulePorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v.(*schema.Set).List(), nil
}

func expandComputeForwardingRuleSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeForwardingRuleTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	// This method returns a full self link from a partial self link.
	if v == nil || v.(string) == "" {
		// It does not try to construct anything from empty.
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		// Anything that starts with a URL scheme is assumed to be a self link worth using.
		return v, nil
	} else if strings.HasPrefix(v.(string), "projects/") {
		// If the self link references a project, we'll just stuck the compute prefix on it
		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
		if err != nil {
			return "", err
		}
		return url, nil
	} else if strings.HasPrefix(v.(string), "regions/") || strings.HasPrefix(v.(string), "zones/") {
		// For regional or zonal resources which include their region or zone, just put the project in front.
		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/")
		if err != nil {
			return nil, err
		}
		return url + v.(string), nil
	}
	// Anything else is assumed to be a regional resource, with a partial link that begins with the resource name.
	// This isn't very likely - it's a last-ditch effort to extract something useful here.  We can do a better job
	// as soon as MultiResourceRefs are working since we'll know the types that this field is supposed to point to.
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/")
	if err != nil {
		return nil, err
	}
	return url + v.(string), nil
}

func expandComputeForwardingRuleAllowGlobalAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleAllPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleNetworkTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleServiceDirectoryRegistrations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamespace, err := expandComputeForwardingRuleServiceDirectoryRegistrationsNamespace(original["namespace"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["namespace"] = transformedNamespace
		}

		transformedService, err := expandComputeForwardingRuleServiceDirectoryRegistrationsService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeForwardingRuleServiceDirectoryRegistrationsNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleServiceDirectoryRegistrationsService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleServiceLabel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleSourceIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleAllowPscGlobalAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleNoAutomateDnsZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleIpVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeForwardingRuleEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeForwardingRuleRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
