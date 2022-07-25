// A package for updating our Backend state
package backend

import (
    "fmt"
)

type MetaData struct {
    // The CID of the file as a string
    Cid string
    // The name of an obao file
    Obao_name string
    // The blake3 hash of the file as a string
    Hash string
    // The size of the file
    Size int64
    // The endpoint of where the file will be stored
    Endpoint string
    // The port of where the file will be stored
    Port int16
}

// Writes the metadata to S3
func WriteMetaData(meta_data MetaData) {
    // Return the hash and size of the file
    fmt.Println("Writing MetaData:\n" +
        "Cid: " + meta_data.Cid + "\n" +
        "Obao_name: " + meta_data.Obao_name + "\n" +
        "Hash: " + meta_data.Hash + "\n" +
        fmt.Sprintf("Size: %d", meta_data.Size) + "\n")

    // TODO - write the metadata to S3
}

func WriteObao(obao_name string) {
    obao_path := "/tmp/" + obao_name

    fmt.Println("Writing Obao:\n" +
        "Obao_name: " + obao_path + "\n")
    // TODO - write the obao file to S3
}