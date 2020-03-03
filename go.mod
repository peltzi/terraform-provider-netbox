module github.com/h0x91b-wix/terraform-provider-netbox

replace github.com/netbox-community/go-netbox v0.0.0 => /home/peltzi/go/src/github.com/netbox-community/go-netbox

go 1.13

require (
	github.com/go-openapi/runtime v0.19.11
	github.com/go-openapi/strfmt v0.19.4
	github.com/hashicorp/terraform-plugin-sdk v1.6.0
	github.com/netbox-community/go-netbox v0.0.0
	//github.com/netbox-community/go-netbox v0.0.0-20200211112514-6bc857cb89fd
	github.com/sirupsen/logrus v1.4.2
)
