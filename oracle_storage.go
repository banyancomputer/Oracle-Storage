package main

import (
    "fmt"

    "github.com/banyancomputer/oracle-storage/backend"
    "github.com/banyancomputer/oracle-storage/processing"
)

// Takes a file path in, processes it, and returns uploads data to S3.
func Store(filename string, filesize int, cid string) (error) {
    // Process the file and get the hash
    hash, err := processing.ProcessFile(filename)
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
   // Declare a MetaData struct to store the file's metadata
    meta_data := backend.MetaData{
        Cid: cid,
        Hash: hash,
        Size: filesize
    }
   // Uplaod our Meta Data to S3
   err = backend.WriteToS3(meta_data)
    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    return nil
}