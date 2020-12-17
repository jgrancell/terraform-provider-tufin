terraform {
  required_providers {
    tufin = {
      source = "jgrancell/tufin"
      version = "0.0.1"
    }
  }
}

provider "tufin" {
  securetrack_host = "localhost:8888"
  securechange_host = "localhost:8888"
  user = "example"
  password = "example"
  allow_insecure = true
}

resource "tufin_group_member" "singleton" {
  group_name = "TEST-001"
  ip_address = "1.1.1.1"
}
