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

package dataplex

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataplexTaskAssetType string = "dataplex.googleapis.com/Task"

func ResourceConverterDataplexTask() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataplexTaskAssetType,
		Convert:   GetDataplexTaskCaiObject,
	}
}

func GetDataplexTaskCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataplexTaskApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataplexTaskAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataplex/v1/rest",
				DiscoveryName:        "Task",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataplexTaskApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDataplexTaskDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataplexTaskDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	triggerSpecProp, err := expandDataplexTaskTriggerSpec(d.Get("trigger_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("trigger_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(triggerSpecProp)) && (ok || !reflect.DeepEqual(v, triggerSpecProp)) {
		obj["triggerSpec"] = triggerSpecProp
	}
	executionSpecProp, err := expandDataplexTaskExecutionSpec(d.Get("execution_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("execution_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(executionSpecProp)) && (ok || !reflect.DeepEqual(v, executionSpecProp)) {
		obj["executionSpec"] = executionSpecProp
	}
	sparkProp, err := expandDataplexTaskSpark(d.Get("spark"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spark"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkProp)) && (ok || !reflect.DeepEqual(v, sparkProp)) {
		obj["spark"] = sparkProp
	}
	notebookProp, err := expandDataplexTaskNotebook(d.Get("notebook"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notebook"); !tpgresource.IsEmptyValue(reflect.ValueOf(notebookProp)) && (ok || !reflect.DeepEqual(v, notebookProp)) {
		obj["notebook"] = notebookProp
	}
	labelsProp, err := expandDataplexTaskEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDataplexTaskDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskTriggerSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandDataplexTaskTriggerSpecType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedStartTime, err := expandDataplexTaskTriggerSpecStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedDisabled, err := expandDataplexTaskTriggerSpecDisabled(original["disabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disabled"] = transformedDisabled
	}

	transformedMaxRetries, err := expandDataplexTaskTriggerSpecMaxRetries(original["max_retries"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetries); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxRetries"] = transformedMaxRetries
	}

	transformedSchedule, err := expandDataplexTaskTriggerSpecSchedule(original["schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schedule"] = transformedSchedule
	}

	return transformed, nil
}

func expandDataplexTaskTriggerSpecType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskTriggerSpecStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskTriggerSpecDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskTriggerSpecMaxRetries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskTriggerSpecSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskExecutionSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedArgs, err := expandDataplexTaskExecutionSpecArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedServiceAccount, err := expandDataplexTaskExecutionSpecServiceAccount(original["service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccount"] = transformedServiceAccount
	}

	transformedProject, err := expandDataplexTaskExecutionSpecProject(original["project"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProject); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["project"] = transformedProject
	}

	transformedMaxJobExecutionLifetime, err := expandDataplexTaskExecutionSpecMaxJobExecutionLifetime(original["max_job_execution_lifetime"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxJobExecutionLifetime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxJobExecutionLifetime"] = transformedMaxJobExecutionLifetime
	}

	transformedKmsKey, err := expandDataplexTaskExecutionSpecKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	return transformed, nil
}

func expandDataplexTaskExecutionSpecArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataplexTaskExecutionSpecServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskExecutionSpecProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskExecutionSpecMaxJobExecutionLifetime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskExecutionSpecKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSpark(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFileUris, err := expandDataplexTaskSparkFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataplexTaskSparkArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	transformedInfrastructureSpec, err := expandDataplexTaskSparkInfrastructureSpec(original["infrastructure_spec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInfrastructureSpec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["infrastructureSpec"] = transformedInfrastructureSpec
	}

	transformedMainJarFileUri, err := expandDataplexTaskSparkMainJarFileUri(original["main_jar_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainJarFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainJarFileUri"] = transformedMainJarFileUri
	}

	transformedMainClass, err := expandDataplexTaskSparkMainClass(original["main_class"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainClass); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainClass"] = transformedMainClass
	}

	transformedPythonScriptFile, err := expandDataplexTaskSparkPythonScriptFile(original["python_script_file"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPythonScriptFile); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pythonScriptFile"] = transformedPythonScriptFile
	}

	transformedSqlScriptFile, err := expandDataplexTaskSparkSqlScriptFile(original["sql_script_file"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSqlScriptFile); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sqlScriptFile"] = transformedSqlScriptFile
	}

	transformedSqlScript, err := expandDataplexTaskSparkSqlScript(original["sql_script"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSqlScript); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sqlScript"] = transformedSqlScript
	}

	return transformed, nil
}

func expandDataplexTaskSparkFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBatch, err := expandDataplexTaskSparkInfrastructureSpecBatch(original["batch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBatch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["batch"] = transformedBatch
	}

	transformedContainerImage, err := expandDataplexTaskSparkInfrastructureSpecContainerImage(original["container_image"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContainerImage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["containerImage"] = transformedContainerImage
	}

	transformedVpcNetwork, err := expandDataplexTaskSparkInfrastructureSpecVpcNetwork(original["vpc_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpcNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpcNetwork"] = transformedVpcNetwork
	}

	return transformed, nil
}

func expandDataplexTaskSparkInfrastructureSpecBatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExecutorsCount, err := expandDataplexTaskSparkInfrastructureSpecBatchExecutorsCount(original["executors_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExecutorsCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["executorsCount"] = transformedExecutorsCount
	}

	transformedMaxExecutorsCount, err := expandDataplexTaskSparkInfrastructureSpecBatchMaxExecutorsCount(original["max_executors_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxExecutorsCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxExecutorsCount"] = transformedMaxExecutorsCount
	}

	return transformed, nil
}

func expandDataplexTaskSparkInfrastructureSpecBatchExecutorsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpecBatchMaxExecutorsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpecContainerImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedImage, err := expandDataplexTaskSparkInfrastructureSpecContainerImageImage(original["image"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["image"] = transformedImage
	}

	transformedJavaJars, err := expandDataplexTaskSparkInfrastructureSpecContainerImageJavaJars(original["java_jars"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJavaJars); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["javaJars"] = transformedJavaJars
	}

	transformedPythonPackages, err := expandDataplexTaskSparkInfrastructureSpecContainerImagePythonPackages(original["python_packages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPythonPackages); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pythonPackages"] = transformedPythonPackages
	}

	transformedProperties, err := expandDataplexTaskSparkInfrastructureSpecContainerImageProperties(original["properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["properties"] = transformedProperties
	}

	return transformed, nil
}

func expandDataplexTaskSparkInfrastructureSpecContainerImageImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpecContainerImageJavaJars(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpecContainerImagePythonPackages(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpecContainerImageProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataplexTaskSparkInfrastructureSpecVpcNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkTags, err := expandDataplexTaskSparkInfrastructureSpecVpcNetworkNetworkTags(original["network_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkTags"] = transformedNetworkTags
	}

	transformedNetwork, err := expandDataplexTaskSparkInfrastructureSpecVpcNetworkNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedSubNetwork, err := expandDataplexTaskSparkInfrastructureSpecVpcNetworkSubNetwork(original["sub_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subNetwork"] = transformedSubNetwork
	}

	return transformed, nil
}

func expandDataplexTaskSparkInfrastructureSpecVpcNetworkNetworkTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpecVpcNetworkNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkInfrastructureSpecVpcNetworkSubNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkMainJarFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkMainClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkPythonScriptFile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkSqlScriptFile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskSparkSqlScript(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebook(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNotebook, err := expandDataplexTaskNotebookNotebook(original["notebook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNotebook); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["notebook"] = transformedNotebook
	}

	transformedInfrastructureSpec, err := expandDataplexTaskNotebookInfrastructureSpec(original["infrastructure_spec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInfrastructureSpec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["infrastructureSpec"] = transformedInfrastructureSpec
	}

	transformedFileUris, err := expandDataplexTaskNotebookFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataplexTaskNotebookArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	return transformed, nil
}

func expandDataplexTaskNotebookNotebook(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBatch, err := expandDataplexTaskNotebookInfrastructureSpecBatch(original["batch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBatch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["batch"] = transformedBatch
	}

	transformedContainerImage, err := expandDataplexTaskNotebookInfrastructureSpecContainerImage(original["container_image"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContainerImage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["containerImage"] = transformedContainerImage
	}

	transformedVpcNetwork, err := expandDataplexTaskNotebookInfrastructureSpecVpcNetwork(original["vpc_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpcNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpcNetwork"] = transformedVpcNetwork
	}

	return transformed, nil
}

func expandDataplexTaskNotebookInfrastructureSpecBatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExecutorsCount, err := expandDataplexTaskNotebookInfrastructureSpecBatchExecutorsCount(original["executors_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExecutorsCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["executorsCount"] = transformedExecutorsCount
	}

	transformedMaxExecutorsCount, err := expandDataplexTaskNotebookInfrastructureSpecBatchMaxExecutorsCount(original["max_executors_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxExecutorsCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxExecutorsCount"] = transformedMaxExecutorsCount
	}

	return transformed, nil
}

func expandDataplexTaskNotebookInfrastructureSpecBatchExecutorsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpecBatchMaxExecutorsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpecContainerImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedImage, err := expandDataplexTaskNotebookInfrastructureSpecContainerImageImage(original["image"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["image"] = transformedImage
	}

	transformedJavaJars, err := expandDataplexTaskNotebookInfrastructureSpecContainerImageJavaJars(original["java_jars"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJavaJars); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["javaJars"] = transformedJavaJars
	}

	transformedPythonPackages, err := expandDataplexTaskNotebookInfrastructureSpecContainerImagePythonPackages(original["python_packages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPythonPackages); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pythonPackages"] = transformedPythonPackages
	}

	transformedProperties, err := expandDataplexTaskNotebookInfrastructureSpecContainerImageProperties(original["properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["properties"] = transformedProperties
	}

	return transformed, nil
}

func expandDataplexTaskNotebookInfrastructureSpecContainerImageImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpecContainerImageJavaJars(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpecContainerImagePythonPackages(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpecContainerImageProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataplexTaskNotebookInfrastructureSpecVpcNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkTags, err := expandDataplexTaskNotebookInfrastructureSpecVpcNetworkNetworkTags(original["network_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkTags"] = transformedNetworkTags
	}

	transformedNetwork, err := expandDataplexTaskNotebookInfrastructureSpecVpcNetworkNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedSubNetwork, err := expandDataplexTaskNotebookInfrastructureSpecVpcNetworkSubNetwork(original["sub_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subNetwork"] = transformedSubNetwork
	}

	return transformed, nil
}

func expandDataplexTaskNotebookInfrastructureSpecVpcNetworkNetworkTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpecVpcNetworkNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookInfrastructureSpecVpcNetworkSubNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskNotebookArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexTaskEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
