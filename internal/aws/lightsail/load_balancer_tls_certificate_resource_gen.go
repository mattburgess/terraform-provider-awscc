// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lightsail_load_balancer_tls_certificate", loadBalancerTlsCertificateResource)
}

// loadBalancerTlsCertificateResource returns the Terraform awscc_lightsail_load_balancer_tls_certificate resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lightsail::LoadBalancerTlsCertificate resource.
func loadBalancerTlsCertificateResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CertificateAlternativeNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"certificate_alternative_names": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
				setplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateDomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
		//	  "type": "string"
		//	}
		"certificate_domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The SSL/TLS certificate name.",
		//	  "type": "string"
		//	}
		"certificate_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The SSL/TLS certificate name.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HttpsRedirectionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.",
		//	  "type": "boolean"
		//	}
		"https_redirection_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IsAttached
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
		//	  "type": "boolean"
		//	}
		"is_attached": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of your load balancer.",
		//	  "pattern": "\\w[\\w\\-]*\\w",
		//	  "type": "string"
		//	}
		"load_balancer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of your load balancer.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("\\w[\\w\\-]*\\w"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerTlsCertificateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"load_balancer_tls_certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The validation status of the SSL/TLS certificate.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The validation status of the SSL/TLS certificate.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::LoadBalancerTlsCertificate").WithTerraformTypeName("awscc_lightsail_load_balancer_tls_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_alternative_names":     "CertificateAlternativeNames",
		"certificate_domain_name":           "CertificateDomainName",
		"certificate_name":                  "CertificateName",
		"https_redirection_enabled":         "HttpsRedirectionEnabled",
		"is_attached":                       "IsAttached",
		"load_balancer_name":                "LoadBalancerName",
		"load_balancer_tls_certificate_arn": "LoadBalancerTlsCertificateArn",
		"status":                            "Status",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
