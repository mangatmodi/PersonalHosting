resource "digitalocean_droplet" "mangatmodi_personal" {
  image  = "ubuntu-18-04-x64"
  name   = "mangatmodi.digitalocean.com"
  region = "FRA1"
  size   = "s-2vcpu-2gb"
  ipv6   = "true"
  monitoring = "true"
  private_networking = "true"
  ssh_keys = [
      "f6:0d:c0:5d:e9:6d:74:8d:4b:90:56:96:10:fe:9f:94"
  ]
  tags = [
	"Personal"
  ]
  user_data = "sudo apt-get update && sudo apt-get -y install docker"
}
