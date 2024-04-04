// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_iotwireless_fuota_task", fuotaTaskResource)
}

// fuotaTaskResource returns the Terraform awscc_iotwireless_fuota_task resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTWireless::FuotaTask resource.
func fuotaTaskResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FUOTA task arn. Returned after successful create.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FUOTA task arn. Returned after successful create.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssociateMulticastGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Multicast group to associate. Only for update request.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"associate_multicast_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Multicast group to associate. Only for update request.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssociateWirelessDevice
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device to associate. Only for update request.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"associate_wireless_device": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device to associate. Only for update request.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FUOTA task description",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FUOTA task description",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DisassociateMulticastGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Multicast group to disassociate. Only for update request.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"disassociate_multicast_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Multicast group to disassociate. Only for update request.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DisassociateWirelessDevice
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device to disassociate. Only for update request.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"disassociate_wireless_device": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device to disassociate. Only for update request.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FirmwareUpdateImage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FUOTA task firmware update image binary S3 link",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"firmware_update_image": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FUOTA task firmware update image binary S3 link",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: FirmwareUpdateRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FUOTA task firmware IAM role for reading S3",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"firmware_update_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FUOTA task firmware IAM role for reading S3",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: FuotaTaskStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FUOTA task status. Returned after successful read.",
		//	  "type": "string"
		//	}
		"fuota_task_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FUOTA task status. Returned after successful read.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FUOTA task id. Returned after successful create.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"fuota_task_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FUOTA task id. Returned after successful create.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoRaWAN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "FUOTA task LoRaWAN",
		//	  "properties": {
		//	    "RfRegion": {
		//	      "description": "FUOTA task LoRaWAN RF region",
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "StartTime": {
		//	      "description": "FUOTA task LoRaWAN start time",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "RfRegion"
		//	  ],
		//	  "type": "object"
		//	}
		"lo_ra_wan": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RfRegion
				"rf_region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "FUOTA task LoRaWAN RF region",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 64),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: StartTime
				"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "FUOTA task LoRaWAN start time",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "FUOTA task LoRaWAN",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of FUOTA task",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of FUOTA task",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the FUOTA task.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the FUOTA task.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Create and manage FUOTA tasks.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::FuotaTask").WithTerraformTypeName("awscc_iotwireless_fuota_task")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"associate_multicast_group":    "AssociateMulticastGroup",
		"associate_wireless_device":    "AssociateWirelessDevice",
		"description":                  "Description",
		"disassociate_multicast_group": "DisassociateMulticastGroup",
		"disassociate_wireless_device": "DisassociateWirelessDevice",
		"firmware_update_image":        "FirmwareUpdateImage",
		"firmware_update_role":         "FirmwareUpdateRole",
		"fuota_task_id":                "Id",
		"fuota_task_status":            "FuotaTaskStatus",
		"key":                          "Key",
		"lo_ra_wan":                    "LoRaWAN",
		"name":                         "Name",
		"rf_region":                    "RfRegion",
		"start_time":                   "StartTime",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
