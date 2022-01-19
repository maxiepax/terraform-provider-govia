terraform {
  required_providers {
    govia = {
      source = "hashicorp.com/edu/govia"
    }
  }
}

provider "govia" {
  username = "admin"
  password = "VMware1!"
  url      = "https://localhost:8443"
}

resource "govia_group" "createdgroup" {
  name        = "lets-test"
  pool_id     = 2
  image_id    = 1
  password    = "VMware1!"
  bootdisk    = "mpx.vmhba1:C0:T0:L0"
  vlan        = "1166"
  callbackurl = "http://stamp.se/callback"
  syslog      = "tcp://172.16.100.4:514,udp://172.16.100.4:514"
  dns         = "172.16.100.4,172.16.100.5"
  ntp         = "172.16.100.4,172.16.100.5"
  options = {
    ssh            = true
    erasedisks     = true
    allowlegacycpu = true
    certificate    = true
    createvmfs     = true
  }
}

output "letstest" {
  value = resource.govia_group.createdgroup
}
