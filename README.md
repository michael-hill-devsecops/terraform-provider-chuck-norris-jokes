# Terraform Provider for Chuck Norris Jokes

This is a custom Terraform provider for retrieving Chuck Norris jokes. The provider fetches a random joke from the Chuck Norris API and makes it available as a Terraform resource.

This project is meant to be a fun, simple "Hello, World" style example of a custom Terraform Provider, not a practical application.

## System Requirements
- go >= 1.22.6
- terraform >= 1.0.11

## Project Structure
```bash
.
├── README.md
├── artifacts
│   └── terraform-provider-chucknorris_v1.0.0
├── build.sh
├── go.mod
├── go.sum
├── handlers
│   └── joke_api.go
├── infrastructure
│   ├── .terraform/*
│   ├── .terraform.lock.hcl
│   ├── .terraformrc
│   ├── main.tf
│   ├── terraform.tfstate
│   └── terraform.tfstate.backup
├── main.go
└── resources
    └── resource_joke.go
```

### Directories

- **`handlers/`**: Contains the code for interacting with the Chuck Norris API.
- **`resources/`**: Defines the Terraform resource schema and CRUD operations.
- **`infrastructure/`**: Contains example Terraform configuration files to test the provider.
- **`artifacts/`**: Directory for the compiled Terraform provider binary.
- **`build.sh`**: Script to build and install the Terraform provider.

## Building the Provider
To build the Terraform provider:

1. Run the `build.sh` script, providing the desired version as an argument:
   ```bash
   ./build.sh 1.0.0
   ```

This script will:
- Download dependencies
- Build the provider binary
- Set executable permissions on the binary
- Copy the binary to the appropriate location in your local Terraform plugin directory
- The built provider will be available in the ~/.terraform.d/plugins/terraform.local/local/chucknorris/1.0.0/linux_amd64/ directory.

## Using the Provider
To use this provider in your Terraform configuration, follow these steps:
1. Setup filesystem mirror
    ```bash
    cp ./infrastructure/.terraformrc ~ # Update provider_installation.filesystem_mirror.path with proper user directory 
    ```

2. Create a Terraform configuration file (e.g., main.tf) with the following content:

    ```bash
    terraform {
        required_providers {
            chucknorris = {
            source  = "terraform.local/local/chucknorris"
            version = "1.0.0"
            }
        }
    }

    provider "chucknorris" {}

    resource "chucknorris_joke" "example" {}

    output "joke" {
    value = chucknorris_joke.example.joke
    }
    ```

3. Initialize Terraform to download the provider:
    ```bash
    terraform init
    ```

4. Apply the configuration to get a Chuck Norris joke:
    ```bash
    terraform apply
    ```