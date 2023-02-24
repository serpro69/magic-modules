package google_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceMonitoringService_AppEngine(t *testing.T) {
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMonitoringService_AppEngine(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.google_monitoring_app_engine_service.default", "name"),
					resource.TestCheckResourceAttrSet("data.google_monitoring_app_engine_service.default", "display_name"),
					resource.TestCheckResourceAttr(
						"data.google_monitoring_app_engine_service.default",
						"telemetry.0.resource_name",
						fmt.Sprintf("//appengine.googleapis.com/apps/%s/services/default", acctest.GetTestProjectFromEnv()),
					),
				),
			},
		},
	})
}

// This does not create an app engine service - instead, it uses the
// base App Engine service "default" that cannot be deleted
func testAccDataSourceMonitoringService_AppEngine() string {
	return fmt.Sprintf(`
data "google_monitoring_app_engine_service" "default" {
	module_id = "default"
}`)
}