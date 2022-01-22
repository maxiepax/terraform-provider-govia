Terraform provider for go-via, the custom deployment tool for VMware ESXi Hypervisor
=========================================

Credits
-------

Massive credits go to one of my best friends, and mentor [Jonathan "stamp" G](https://www.github.com/stamp) for all the help, coaching and lessons during this project. Without your support this project would never have been a reality.


# Building the terraform provider

Run the following command to build the provider

```shell
go build -o terraform-provider-govia
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```

## In wait for documentation, please find syntax in example/main.tf
