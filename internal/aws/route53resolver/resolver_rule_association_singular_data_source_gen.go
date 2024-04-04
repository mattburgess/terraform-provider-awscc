// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_resolver_rule_association", resolverRuleAssociationDataSource)
}

// resolverRuleAssociationDataSource returns the Terraform awscc_route53resolver_resolver_rule_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::ResolverRuleAssociation resource.
func resolverRuleAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of an association between a Resolver rule and a VPC.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of an association between a Resolver rule and a VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResolverRuleAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"resolver_rule_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResolverRuleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Resolver rule that you associated with the VPC that is specified by ``VPCId``.",
		//	  "type": "string"
		//	}
		"resolver_rule_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Resolver rule that you associated with the VPC that is specified by ``VPCId``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VPCId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC that you associated the Resolver rule with.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC that you associated the Resolver rule with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53Resolver::ResolverRuleAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverRuleAssociation").WithTerraformTypeName("awscc_route53resolver_resolver_rule_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"name":                         "Name",
		"resolver_rule_association_id": "ResolverRuleAssociationId",
		"resolver_rule_id":             "ResolverRuleId",
		"vpc_id":                       "VPCId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
