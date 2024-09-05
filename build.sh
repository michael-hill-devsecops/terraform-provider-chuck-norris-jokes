#!/bin/bash
# build.sh
#
# This script builds the Terraform provider binary and places it in the
# appropriate directory for local Terraform use. It performs the following steps:
# - Downloads Go module dependencies
# - Builds the provider binary
# - Sets executable permissions on the binary
# - Creates the necessary directories in the Terraform plugin cache
# - Copies the binary to the plugin cache directory

go get
go build -o ./artifacts/terraform-provider-chucknorris_v${1}
chmod +x ./artifacts/terraform-provider-chucknorris_v${1}
mkdir -p ~/.terraform.d/plugins/terraform.local/local/chucknorris/1.0.0/linux_amd64/
cp ./artifacts/terraform-provider-chucknorris_v${1} ~/.terraform.d/plugins/terraform.local/local/chucknorris/1.0.0/linux_amd64/