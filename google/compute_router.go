// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"reflect"
)

func GetComputeRouterCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	if obj, err := GetComputeRouterApiObject(d, config); err == nil {
		return Asset{
			Name: fmt.Sprintf("//compute.googleapis.com/%s", obj["selfLink"]),
			Type: "google.compute.Router",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Router",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeRouterApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRouterName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRouterDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); ok || !reflect.DeepEqual(v, descriptionProp) {
		obj["description"] = descriptionProp
	}
	networkProp, err := expandComputeRouterNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	bgpProp, err := expandComputeRouterBgp(d.Get("bgp"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bgp"); !isEmptyValue(reflect.ValueOf(bgpProp)) && (ok || !reflect.DeepEqual(v, bgpProp)) {
		obj["bgp"] = bgpProp
	}
	regionProp, err := expandComputeRouterRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRouterName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRouterBgp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAsn, err := expandComputeRouterBgpAsn(original["asn"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAsn); val.IsValid() && !isEmptyValue(val) {
		transformed["asn"] = transformedAsn
	}

	transformedAdvertiseMode, err := expandComputeRouterBgpAdvertiseMode(original["advertise_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdvertiseMode); val.IsValid() && !isEmptyValue(val) {
		transformed["advertiseMode"] = transformedAdvertiseMode
	}

	transformedAdvertisedGroups, err := expandComputeRouterBgpAdvertisedGroups(original["advertised_groups"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["advertisedGroups"] = transformedAdvertisedGroups
	}

	transformedAdvertisedIpRanges, err := expandComputeRouterBgpAdvertisedIpRanges(original["advertised_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["advertisedIpRanges"] = transformedAdvertisedIpRanges
	}

	return transformed, nil
}

func expandComputeRouterBgpAsn(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertiseMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertisedGroups(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertisedIpRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRange, err := expandComputeRouterBgpAdvertisedIpRangesRange(original["range"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["range"] = transformedRange
		}

		transformedDescription, err := expandComputeRouterBgpAdvertisedIpRangesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["description"] = transformedDescription
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRouterBgpAdvertisedIpRangesRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterBgpAdvertisedIpRangesDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouterRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}