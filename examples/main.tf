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

resource "govia_pool" "pool01" {
  name          = "tf_test_pool"
  net_address   = "192.168.100.0"
  start_address = "192.168.10.10"
  end_address   = "192.168.10.20"
  netmask       = 24
  gateway       = "192.168.10.1"
}

resource "govia_group" "group01" {
  name        = "test-grp"
  pool_id     = 1
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

output "group01" {
  value = resource.govia_group.group01
}

resource "govia_address" "esx01" {
  ip       = "172.16.100.11"
  mac      = "aa:bb:cc:11:22:33"
  hostname = "esx01"
  domain   = "vmlab.se"
  reimage  = true
  pool_id  = resource.govia_group.group01.pool_id
  group_id = resource.govia_group.group01.id
}

resource "govia_address" "esx02" {
  ip       = "172.16.100.13"
  mac      = "aa:bb:cc:11:22:34"
  hostname = "esx02"
  domain   = "vmlab.se"
  reimage  = true
  pool_id  = resource.govia_group.group01.pool_id
  group_id = resource.govia_group.group01.id
}

output "esx01" {
  value = resource.govia_address.esx01
}
output "esx02" {
  value = resource.govia_address.esx02
}

