package main

import (
    "fmt"

    "src/oracle_storage/backend"
    "src/oracle_storage/processing"
)

// Takes a file path in, processes it, and returns uploads data to S3.
func oracle_storage(file_path string) {
    // Process the file
    meta_data := processing.ProcessFile(file_path)

    // Upload the file meta_data to S3
    backend.WriteMetaData(meta_data)
    // Upload the obao file to S3
    backend.WriteObao(meta_data.Obao_name)

    fmt.Println("Done!")
}

func main() {
    oracle_storage("./test/ethereum.pdf")
}
