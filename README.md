# Oracle-Storage

> This repository is responsible for managing meta-data storage for our Oracel Ecosystem.
> Therefore it:
> - Manages the deployment of the AWS S3 buckets used for storing the meta-data, endpoints, and obao files.
> - Implements a simple Go Library for preprocessing files.
> - Writes processed meta-data and files to the S3 buckets.

## Dependencies
For now, this repository is dependent on the following:
- `bao`: The bao library must be accessible from the command line.
- `ipfs`: The ipfs CLI must be accessible from the command line.

# Building
```bash
$ go install .
```

# Usage
Right now, the repository processes and uploads the data for the `test/ethereum` file.

```bash
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ oracle_storage
```

# Future Work
Right now, all preprocessing is done through command line invocations.
This is not ideal, and should be changed to a contained library.