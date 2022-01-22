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

resource "govia_image" "tf_image01" {
  iso_image   = "/Users/kimjohansson/Downloads/VMware-VMvisor-Installer-7.0U3-18644231.x86_64.iso"
  hash        = "e865b644cea512986a20ec56b313f2d1f167d339b4e37ff826a9ab3821e558a7"
  description = "renameme"
}

resource "govia_pool" "tf_pool01" {
  name          = "tf_test_pool"
  net_address   = "192.168.100.0"
  start_address = "192.168.100.10"
  end_address   = "192.168.100.20"
  netmask       = 24
  gateway       = "192.168.100.1"
}

resource "govia_group" "tf_group01" {
  name        = "test-grp"
  pool_id     = resource.govia_pool.tf_pool01.id
  image_id    = resource.govia_image.tf_image01.id
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

resource "govia_address" "tf_esx01" {
  ip       = "192.168.100.11"
  mac      = "aa:bb:cc:11:22:33"
  hostname = "esx01"
  domain   = "vmlab.se"
  reimage  = true
  pool_id  = resource.govia_group.tf_group01.pool_id
  group_id = resource.govia_group.tf_group01.id
}

resource "govia_address" "tf_esx02" {
  ip       = "192.168.100.13"
  mac      = "aa:bb:cc:11:22:34"
  hostname = "esx02"
  domain   = "vmlab.se"
  reimage  = true
  pool_id  = resource.govia_group.tf_group01.pool_id
  group_id = resource.govia_group.tf_group01.id
}
