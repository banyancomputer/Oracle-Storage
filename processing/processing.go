package processing

import (
    "src/oracle_storage/backend"
    // A package for exposing our Bao processing to go
    "src/oracle_storage/processing/blake3"

    // A package for generating CIDs
)

func getCID(file_path string) string {
    return ""
}

// Takes a file path in, processes it, returning a:
// - a meta_data struct containing the path to a cid, obao file, the file's blake3 hash, and the file's size
//   and other information about the file
func ProcessFile(file_path string) (meta_data backend.MetaData) {
    println("Processing file: " + file_path)

    // Process the file
    obao_name, hash, size := blake3.ObaoData(file_path)

    // Get a CID for the file
    cid := getCID(file_path)

    // Return the meta_data
    // For now just set the endpoint to "localhost" and the port to 5051
    return backend.MetaData{cid, obao_name, hash, size, "localhost", 5051}
}

