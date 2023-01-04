// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediaconnect_flow_entitlement", flowEntitlementResource)
}

// flowEntitlementResource returns the Terraform awscc_mediaconnect_flow_entitlement resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaConnect::FlowEntitlement resource.
func flowEntitlementResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DataTransferSubscriberFeePercent
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "description": "Percentage from 0-100 of the data transfer cost to be billed to the subscriber.",
		//	  "type": "integer"
		//	}
		"data_transfer_subscriber_fee_percent": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Percentage from 0-100 of the data transfer cost to be billed to the subscriber.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				Int64DefaultValue(0),
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the entitlement.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the entitlement.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Encryption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The type of encryption that will be used on the output that is associated with this entitlement.",
		//	  "properties": {
		//	    "Algorithm": {
		//	      "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
		//	      "enum": [
		//	        "aes128",
		//	        "aes192",
		//	        "aes256"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "ConstantInitializationVector": {
		//	      "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "DeviceId": {
		//	      "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "KeyType": {
		//	      "default": "static-key",
		//	      "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
		//	      "enum": [
		//	        "speke",
		//	        "static-key"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Region": {
		//	      "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "ResourceId": {
		//	      "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
		//	      "type": "string"
		//	    },
		//	    "SecretArn": {
		//	      "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
		//	      "type": "string"
		//	    },
		//	    "Url": {
		//	      "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Algorithm",
		//	    "RoleArn"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Algorithm
				"algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"aes128",
							"aes192",
							"aes256",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: ConstantInitializationVector
				"constant_initialization_vector": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DeviceId
				"device_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KeyType
				"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"speke",
							"static-key",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						StringDefaultValue("static-key"),
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Region
				"region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ResourceId
				"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecretArn
				"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Url
				"url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The type of encryption that will be used on the output that is associated with this entitlement.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EntitlementArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the entitlement.",
		//	  "type": "string"
		//	}
		"entitlement_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the entitlement.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EntitlementStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": " An indication of whether the entitlement is enabled.",
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"entitlement_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: " An indication of whether the entitlement is enabled.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ENABLED",
					"DISABLED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FlowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the flow.",
		//	  "type": "string"
		//	}
		"flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the flow.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the entitlement.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the entitlement.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Subscribers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"subscribers": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.",
			Required:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaConnect::FlowEntitlement",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowEntitlement").WithTerraformTypeName("awscc_mediaconnect_flow_entitlement")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm":                            "Algorithm",
		"constant_initialization_vector":       "ConstantInitializationVector",
		"data_transfer_subscriber_fee_percent": "DataTransferSubscriberFeePercent",
		"description":                          "Description",
		"device_id":                            "DeviceId",
		"encryption":                           "Encryption",
		"entitlement_arn":                      "EntitlementArn",
		"entitlement_status":                   "EntitlementStatus",
		"flow_arn":                             "FlowArn",
		"key_type":                             "KeyType",
		"name":                                 "Name",
		"region":                               "Region",
		"resource_id":                          "ResourceId",
		"role_arn":                             "RoleArn",
		"secret_arn":                           "SecretArn",
		"subscribers":                          "Subscribers",
		"url":                                  "Url",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
