resource "vultr_instance" "server" {
  plan = "vc2-2c-4gb"
  region = "fra"
  os_id = "387"
  enable_ipv6 = true
  ssh_key_ids = ["${vultr_ssh_key.root.id}"]
}

output "servers" {
  value = vultr_instance.server.*.main_ip
}
