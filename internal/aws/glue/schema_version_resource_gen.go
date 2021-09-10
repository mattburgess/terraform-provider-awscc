// Code generated by generators/resource/main.go; DO NOT EDIT.

package glue

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_glue_schema_version", schemaVersionResourceType)
}

// schemaVersionResourceType returns the Terraform awscc_glue_schema_version resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Glue::SchemaVersion resource type.
func schemaVersionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"schema": {
			// Property: Schema
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Identifier for the schema where the schema version will be created.",
			//   "properties": {
			//     "RegistryName": {
			//       "description": "Name of the registry to identify where the Schema is located.",
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SchemaArn": {
			//       "description": "Amazon Resource Name for the Schema. This attribute can be used to uniquely represent the Schema.",
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "SchemaName": {
			//       "description": "Name of the schema. This parameter requires RegistryName to be provided.",
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Identifier for the schema where the schema version will be created.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"registry_name": {
						// Property: RegistryName
						Description: "Name of the registry to identify where the Schema is located.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"schema_arn": {
						// Property: SchemaArn
						Description: "Amazon Resource Name for the Schema. This attribute can be used to uniquely represent the Schema.",
						Type:        types.StringType,
						Optional:    true,
					},
					"schema_name": {
						// Property: SchemaName
						Description: "Name of the schema. This parameter requires RegistryName to be provided.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // Schema is a force-new property.
			},
		},
		"schema_definition": {
			// Property: SchemaDefinition
			// CloudFormation resource type schema:
			// {
			//   "description": "Complete definition of the schema in plain-text.",
			//   "maxLength": 170000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Complete definition of the schema in plain-text.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 170000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // SchemaDefinition is a force-new property.
			},
		},
		"version_id": {
			// Property: VersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "Represents the version ID associated with the schema version.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Represents the version ID associated with the schema version.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "This resource represents an individual schema version of a schema defined in Glue Schema Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::SchemaVersion").WithTerraformTypeName("awscc_glue_schema_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"registry_name":     "RegistryName",
		"schema":            "Schema",
		"schema_arn":        "SchemaArn",
		"schema_definition": "SchemaDefinition",
		"schema_name":       "SchemaName",
		"version_id":        "VersionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_glue_schema_version", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
