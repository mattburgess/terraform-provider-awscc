// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package transfer_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSTransferProfileDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Transfer::Profile", "awscc_transfer_profile", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSTransferProfileDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Transfer::Profile", "awscc_transfer_profile", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}