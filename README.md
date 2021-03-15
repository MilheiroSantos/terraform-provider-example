# terraform-provider-example
A starter template for creating a terraform provider.

# How to run the example
On the `myfile_provider` folder, build the provider:
```shell
$ go build terraform-provider-myfile
```
Copy the executable to terraform's implied local mirror directory. 
The format should be:
```
<TF_DIR>/plugins/myorg.com/custom/myfile/0.1.0/<OS>_<ARCH>/terraform-provider-myfile
```
- For Linux, `OS` is `linux` and `TF_DIR` can be `$HOME/.terraform.d/plugins`
- For macOS X, `OS` is `darwin` and `TF_DIR` can be `$HOME/.terraform.d/plugins`
- For Windows, `OS` is `windows` and `TF_DIR` can be `%APPDATA%/terraform.d/plugins`

`ARCH` is the CPU architecture. Most common values are `amd64` or `arm64`.
 
For a Linux amd64 build: 
```shell
$ mkdir -p ~/.terraform.d/plugins/myorg.com/custom/myfile/0.1.0/linux_amd64/ && \
  mv terraform-provider-myfile ~/.terraform.d/plugins/myorg.com/custom/myfile/0.1.0/linux_amd64/
```

On the `example` folder, apply the terraform configuration:
```shell
$ terraform init && terraform apply -auto-approve
```

You should have an extra file in the `example` folder named `here.txt` with contents `Hello World!`
