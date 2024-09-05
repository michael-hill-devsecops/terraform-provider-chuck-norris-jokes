
go get
go build -o terraform-provider-chucknorris
chmod +x terraform-provider-chucknorris
mkdir -p ~/.terraform.d/plugins/terraform.local/local/chucknorris/1.0.0/linux_amd64/
cp terraform-provider-chucknorris ~/.terraform.d/plugins/terraform.local/local/chucknorris/1.0.0/linux_amd64/