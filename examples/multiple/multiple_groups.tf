terraform {
  required_providers {
    tufin = {
      source = "jgrancell/tufin"
      version = "0.0.1"
    }
  }
}

locals {
  firewall_groups = [
    "TEST-1",
    "TEST-2",
  ]
}

provider "tufin" {
  securetrack_host = "localhost"
  securechange_host = "localhost"
  user = "example"
  password = "example"
  allow_insecure = true
}

resource "tufin_group_member" "multiples" {
  count = length(local.firewall_groups)
  local.firewall_groups[count.index]
  ip_address = "1.1.1.1"
}
