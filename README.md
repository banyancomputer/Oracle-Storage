# Oracle-Storage

> This repository is responsible for managing meta-data storage for our Oracle Ecosystem.
> This is a temporary repository until we move this storage on chain/onto providers machines.
> It Implements:
> - Managing the deployment of the AWS S3 buckets used for storing the meta-data and obao files.
> - Implements a simple wrapper around a systems command line to generate obao files from Go.
> - Writes processed meta-data and files to the S3 buckets.

## Dependencies
For now, this repository is dependent on the following:
- `bao`: The bao library must be accessible from the command line in order to generate obao files.

# Building
```bash
$ go install .
```

# Usage
Deploy AWS infrastructure with Terraform. Then call this library to generate obao files.

# Testing
```bash
$ go test oracle_storage.go oracle_storage_test.go
```

# Future Work
Right now, all preprocessing is done through command line invocations.
This is not ideal, and should be changed to a contained library.