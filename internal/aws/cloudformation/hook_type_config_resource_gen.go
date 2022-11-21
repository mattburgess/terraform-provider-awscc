// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_cloudformation_hook_type_config", hookTypeConfigResource)
}

// hookTypeConfigResource returns the Terraform awscc_cloudformation_hook_type_config resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFormation::HookTypeConfig resource.
func hookTypeConfigResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The configuration data for the extension, in this account and region.",
			//	  "pattern": "[\\s\\S]+",
			//	  "type": "string"
			//	}
			Description: "The configuration data for the extension, in this account and region.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("[\\s\\S]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"configuration_alias": {
			// Property: ConfigurationAlias
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "default",
			//	  "description": "An alias by which to refer to this extension configuration data.",
			//	  "enum": [
			//	    "default"
			//	  ],
			//	  "pattern": "^[a-zA-Z0-9]{1,256}$",
			//	  "type": "string"
			//	}
			Description: "An alias by which to refer to this extension configuration data.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9]{1,256}$"), ""),
				validate.StringInSlice([]string{
					"default",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.StringValue("default")),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"configuration_arn": {
			// Property: ConfigurationArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) for the configuration data, in this account and region.",
			//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type(-configuration)?/hook/.+$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) for the configuration data, in this account and region.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"type_arn": {
			// Property: TypeArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the type without version number.",
			//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the type without version number.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"type_name": {
			// Property: TypeName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			//	  "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$",
			//	  "type": "string"
			//	}
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Specifies the configuration data for a registered hook in CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::HookTypeConfig").WithTerraformTypeName("awscc_cloudformation_hook_type_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"configuration":       "Configuration",
		"configuration_alias": "ConfigurationAlias",
		"configuration_arn":   "ConfigurationArn",
		"type_arn":            "TypeArn",
		"type_name":           "TypeName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	opts = opts.WithRequiredAttributesValidators(validate.OneOfRequired(
		validate.Required(
			"type_arn",
			"configuration",
		),
		validate.Required(
			"type_name",
			"configuration",
		),
	),
	)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
