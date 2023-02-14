terraform {
  required_providers {
    libvirt = {
      source = "dmacvicar/libvirt"
      version = "0.7.0"
    }
  }
}

provider "libvirt" {
  uri = "qemu:///system"
}

resource "libvirt_volume" "os" {
  name = "suse_leap"
  source = "https://download.opensuse.org/distribution/leap/15.4/appliances/openSUSE-Leap-15.4-JeOS.x86_64-OpenStack-Cloud.qcow2"
}

resource "libvirt_volume" "default_disk" {
  name = "main.qcow2"
  base_volume_id = libvirt_volume.os.id
  size = 21474836480
}

resource "libvirt_network" "test_nat" {
  name = "testnat"
  mode = "nat"
  addresses = ["10.0.10.0/24"]
  dhcp {
    enabled = true
  }
}

resource "libvirt_cloudinit_disk" "commoninit" {
  name      = "commoninit.iso"
  user_data = <<EOF
#cloud-config
ssh_pwauth: True
chpasswd:
  list: |
     opensuse:linux
  expire: False
EOF
}

resource "libvirt_domain" "default" {
  name = "test8"
  cpu {
    mode = "host-model"
  }
  vcpu = 1
  memory = "1024"
  running = true
  disk {
    volume_id = libvirt_volume.default_disk.id
    scsi = true
  }
  cloudinit = libvirt_cloudinit_disk.commoninit.id
  network_interface {
    network_id = libvirt_network.test_nat.id
    hostname = "test-vm"
    addresses = ["10.0.10.3"]
    wait_for_lease = true
  }
  qemu_agent = true
}