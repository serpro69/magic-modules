package google_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleSourceRepoRepository_basic(t *testing.T) {
	t.Parallel()

	name := "tf-repository-" + acctest.RandString(t, 10)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:     func() { acctest.TestAccPreCheck(t) },
		Providers:    acctest.TestAccProviders,
		CheckDestroy: testAccCheckSourceRepoRepositoryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleSourceRepoRepositoryConfig(name),
				Check: resource.ComposeTestCheckFunc(
					CheckDataSourceStateMatchesResourceState("data.google_sourcerepo_repository.bar", "google_sourcerepo_repository.foo"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleSourceRepoRepositoryConfig(name string) string {
	return fmt.Sprintf(`
resource "google_sourcerepo_repository" "foo" {
  name               = "%s"
}

data "google_sourcerepo_repository" "bar" {
  name = google_sourcerepo_repository.foo.name
  depends_on = [
    google_sourcerepo_repository.foo,
  ]
}
`, name)
}