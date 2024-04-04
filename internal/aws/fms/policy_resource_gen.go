// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package fms

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_fms_policy", policyResource)
}

// policyResource returns the Terraform awscc_fms_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::FMS::Policy resource.
func policyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource ARN.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "^([^\\s]*)$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource ARN.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeleteAllPolicyResources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"delete_all_policy_resources": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// DeleteAllPolicyResources is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ExcludeMap
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An FMS includeMap or excludeMap.",
		//	  "properties": {
		//	    "ACCOUNT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An AWS account ID.",
		//	        "maxLength": 12,
		//	        "minLength": 12,
		//	        "pattern": "^([0-9]*)$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "ORGUNIT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An Organizational Unit ID.",
		//	        "maxLength": 68,
		//	        "minLength": 16,
		//	        "pattern": "^(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"exclude_map": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ACCOUNT
				"account": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(12, 12),
							stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]*)$"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ORGUNIT
				"orgunit": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(16, 68),
							stringvalidator.RegexMatches(regexp.MustCompile("^(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An FMS includeMap or excludeMap.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExcludeResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"exclude_resource_tags": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[a-z0-9A-Z-]{36}$",
		//	  "type": "string"
		//	}
		"policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IncludeMap
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An FMS includeMap or excludeMap.",
		//	  "properties": {
		//	    "ACCOUNT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An AWS account ID.",
		//	        "maxLength": 12,
		//	        "minLength": 12,
		//	        "pattern": "^([0-9]*)$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "ORGUNIT": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "description": "An Organizational Unit ID.",
		//	        "maxLength": 68,
		//	        "minLength": 16,
		//	        "pattern": "^(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"include_map": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ACCOUNT
				"account": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(12, 12),
							stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]*)$"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ORGUNIT
				"orgunit": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(16, 68),
							stringvalidator.RegexMatches(regexp.MustCompile("^(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An FMS includeMap or excludeMap.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "pattern": "^([a-zA-Z0-9_.:/=+\\-@\\s]+)$",
		//	  "type": "string"
		//	}
		"policy_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("^([a-zA-Z0-9_.:/=+\\-@\\s]+)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "^([a-zA-Z0-9_.:/=+\\-@\\s]+)$",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile("^([a-zA-Z0-9_.:/=+\\-@\\s]+)$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RemediationEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"remediation_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceSetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "description": "A Base62 ID",
		//	    "maxLength": 22,
		//	    "minLength": 22,
		//	    "pattern": "^[a-z0-9A-Z]{22}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"resource_set_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(22, 22),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9A-Z]{22}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A resource tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 8,
		//	  "type": "array"
		//	}
		"resource_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(8),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An AWS resource type",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^([^\\s]*)$",
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An AWS resource type",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^([^\\s]*)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceTypeList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "description": "An AWS resource type",
		//	    "maxLength": 128,
		//	    "minLength": 1,
		//	    "pattern": "^([^\\s]*)$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"resource_type_list": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^([^\\s]*)$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourcesCleanUp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"resources_clean_up": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityServicePolicyData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Firewall security service policy data.",
		//	  "properties": {
		//	    "ManagedServiceData": {
		//	      "description": "Firewall managed service data.",
		//	      "maxLength": 8192,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "PolicyOption": {
		//	      "additionalProperties": false,
		//	      "description": "Firewall policy option.",
		//	      "oneOf": [
		//	        {
		//	          "required": [
		//	            "NetworkFirewallPolicy"
		//	          ]
		//	        },
		//	        {
		//	          "required": [
		//	            "ThirdPartyFirewallPolicy"
		//	          ]
		//	        }
		//	      ],
		//	      "properties": {
		//	        "NetworkFirewallPolicy": {
		//	          "additionalProperties": false,
		//	          "description": "Network firewall policy.",
		//	          "properties": {
		//	            "FirewallDeploymentModel": {
		//	              "description": "Firewall deployment mode.",
		//	              "enum": [
		//	                "DISTRIBUTED",
		//	                "CENTRALIZED"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "FirewallDeploymentModel"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ThirdPartyFirewallPolicy": {
		//	          "additionalProperties": false,
		//	          "description": "Third party firewall policy.",
		//	          "properties": {
		//	            "FirewallDeploymentModel": {
		//	              "description": "Firewall deployment mode.",
		//	              "enum": [
		//	                "DISTRIBUTED",
		//	                "CENTRALIZED"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "FirewallDeploymentModel"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "description": "Firewall policy type.",
		//	      "enum": [
		//	        "WAF",
		//	        "WAFV2",
		//	        "SHIELD_ADVANCED",
		//	        "SECURITY_GROUPS_COMMON",
		//	        "SECURITY_GROUPS_CONTENT_AUDIT",
		//	        "SECURITY_GROUPS_USAGE_AUDIT",
		//	        "NETWORK_FIREWALL",
		//	        "THIRD_PARTY_FIREWALL",
		//	        "DNS_FIREWALL",
		//	        "IMPORT_NETWORK_FIREWALL"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"security_service_policy_data": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ManagedServiceData
				"managed_service_data": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Firewall managed service data.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 8192),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: PolicyOption
				"policy_option": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: NetworkFirewallPolicy
						"network_firewall_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: FirewallDeploymentModel
								"firewall_deployment_model": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Firewall deployment mode.",
									Required:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"DISTRIBUTED",
											"CENTRALIZED",
										),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Network firewall policy.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ThirdPartyFirewallPolicy
						"third_party_firewall_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: FirewallDeploymentModel
								"firewall_deployment_model": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Firewall deployment mode.",
									Required:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"DISTRIBUTED",
											"CENTRALIZED",
										),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Third party firewall policy.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Firewall policy option.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Firewall policy type.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"WAF",
							"WAFV2",
							"SHIELD_ADVANCED",
							"SECURITY_GROUPS_COMMON",
							"SECURITY_GROUPS_CONTENT_AUDIT",
							"SECURITY_GROUPS_USAGE_AUDIT",
							"NETWORK_FIREWALL",
							"THIRD_PARTY_FIREWALL",
							"DNS_FIREWALL",
							"IMPORT_NETWORK_FIREWALL",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Firewall security service policy data.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A policy tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([^\\s]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "pattern": "^([^\\s]*)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("^([^\\s]*)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(256),
							stringvalidator.RegexMatches(regexp.MustCompile("^([^\\s]*)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Creates an AWS Firewall Manager policy.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FMS::Policy").WithTerraformTypeName("awscc_fms_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account":                      "ACCOUNT",
		"arn":                          "Arn",
		"delete_all_policy_resources":  "DeleteAllPolicyResources",
		"exclude_map":                  "ExcludeMap",
		"exclude_resource_tags":        "ExcludeResourceTags",
		"firewall_deployment_model":    "FirewallDeploymentModel",
		"include_map":                  "IncludeMap",
		"key":                          "Key",
		"managed_service_data":         "ManagedServiceData",
		"network_firewall_policy":      "NetworkFirewallPolicy",
		"orgunit":                      "ORGUNIT",
		"policy_description":           "PolicyDescription",
		"policy_id":                    "Id",
		"policy_name":                  "PolicyName",
		"policy_option":                "PolicyOption",
		"remediation_enabled":          "RemediationEnabled",
		"resource_set_ids":             "ResourceSetIds",
		"resource_tags":                "ResourceTags",
		"resource_type":                "ResourceType",
		"resource_type_list":           "ResourceTypeList",
		"resources_clean_up":           "ResourcesCleanUp",
		"security_service_policy_data": "SecurityServicePolicyData",
		"tags":                         "Tags",
		"third_party_firewall_policy":  "ThirdPartyFirewallPolicy",
		"type":                         "Type",
		"value":                        "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DeleteAllPolicyResources",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
