package google_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAccessApprovalOrganizationServiceAccount_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id": acctest.GetTestOrgFromEnv(t),
	}

	resourceName := "data.google_access_approval_organization_service_account.aa_account"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAccessApprovalOrganizationServiceAccount_basic(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "account_email"),
				),
			},
		},
	})
}

func testAccDataSourceAccessApprovalOrganizationServiceAccount_basic(context map[string]interface{}) string {
	return Nprintf(`
data "google_access_approval_organization_service_account" "aa_account" {
  organization_id = "%{org_id}"
}
`, context)
}