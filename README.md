# Oracle-Storage

> This repository is responsible for managing meta-data storage for our Oracle Ecosystem.
> The described functionality will implement pre-processing that will eventually be moved to providers machines.
> For now it:
> - Manages the deployment of the AWS S3 buckets used for storing the meta-data, endpoints, and obao files.
> - Implements a simple Go Library for preprocessing files into obao files.
> - Writes processed meta-data and files to the S3 buckets.

# Building
First, you need to configure `gobao/gobao.go` to load either a dyanmic or static version of our Rust library.
Edit `gobao/lib/obao/cargo.toml` and to compile to your chosen format.

Build Rust binaries
```bash
$ cd gobao
$ ./lib_build.sh
```
If you compiled the library to a dynamic library, remember to add the `lib` directory to your LD_LIBRARY_PATH.

Build Go binaries
```bash
$ go install .
```

# Usage
Import this library as a git submodule into your project. Be sure to include running the `lib_build.sh` script in your Makefile.

Use Terraform in order to deploy your AWS S3 buckets.