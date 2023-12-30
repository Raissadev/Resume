terraform {
  required_version = "1.4.2"

  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.token
}

resource "digitalocean_ssh_key" "mk" {
  name       = "example"
  public_key = file(var.pub_key)
}

resource "digitalocean_droplet" "dev" {
  name   = "dev"
  region = "" // hidden ¬¬
  size   = "" // hidden ¬¬
  image  = "" // hidden ¬¬
  ssh_keys = [
    digitalocean_ssh_key.mk.id,
  ]

  connection {
    type        = "ssh"
    user        = var.ssh_user
    private_key = file(var.priv_key)
    timeout     = "2m"
    host        = self.ipv4_address
  }

  provisioner "file" {
    source      = "./down.sh"
    destination = "/down.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /down.sh",
      "/down.sh"
    ]
  }
}