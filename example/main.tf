terraform {
  required_providers {
    myfile = {
      version = "=0.1.0"
      source = "myorg.com/custom/myfile"
    }
  }
}

provider "myfile" {
  encoding = "utf8"
}

resource "myfile_file" "this" {
  path = "${path.module}/here.txt"
  contents = "Hello Word!"
}

output "file_path" { value = myfile_file.this.path }
output "file_contents" { value = myfile_file.this.contents }
output "file_owner" { value = myfile_file.this.owner }
