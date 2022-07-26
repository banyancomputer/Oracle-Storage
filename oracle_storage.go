package main

import (
    "fmt"

    "src/oracle_storage/backend"
    "src/oracle_storage/processing"
)

// Takes a file path in, processes it, and returns uploads data to S3.
func oracle_storage(file_path string) {
    // Process the file
    deal_id, meta_data := processing.ProcessFile(file_path)

    // Write the data to S3
    backend.WriteToS3(deal_id, meta_data)

    fmt.Println("Uploaded new deal: ", deal_id)
}

func main() {
    oracle_storage("./test/ethereum.pdf")
}
