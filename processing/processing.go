package processing

import (
    "fmt"
    "os/exec"
    "strconv"
    "strings"

    "src/oracle_storage/backend"
)

/* TODO: Transition from calling these as command lines args */

// Get the CID of a file
func getCID(file_path string) string {
    cmd := exec.Command("ipfs", "add", file_path, "-q", "--cid-version", "1")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return ""
    }

    return string(stdout)
}

// Get the Blake3 hash of a file
func getHash(file_path string) (string) {
    cmd := exec.Command("bao", "hash", file_path)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return ""
    }

    return string(stdout)
}

// Get the size of a file
func getSize(file_path string) (int64) {
    // Determine the size of the file
    cmd := exec.Command("stat", "-c%s", file_path)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return 0
    }

    val, err := strconv.Atoi(strings.Split(string(stdout), "\n")[0])
    if err != nil {
        fmt.Println(err.Error())
        return 0
    }
    return int64(val)
}

// Encodes a file into an obao file
func encodeObao(file_path string, hash string) (error) {
    // Determine the path of the obao file
    obao_path := backend.ObaoTempStore + hash
    // Generate an obao file for the file
    cmd := exec.Command("bao", "encode", file_path, "--outboard", obao_path)
    _, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    return nil
}

// Takes a file path in, processes it, returning a:
// - a meta_data struct containing the path to a cid, obao file, the file's blake3 hash, and the file's size
//   and other information about the file
func ProcessFile(file_path string) (meta_data backend.MetaData) {
    println("Processing file: " + file_path)

    // Get the CID of the file
    cid := getCID(file_path)
    // Get the hash of the file
    hash := getHash(file_path)
    // Get the size of the file
    size := getSize(file_path)
    // Generate an obao file for the file
    err := encodeObao(file_path, hash)
    if err != nil {
        fmt.Println(err.Error())
    }

    // Return the meta_data
    // TODO: Be able to account for different endpoints
    // But also keep in mind we might be getting that from our Smart Contract
    // For now just set the endpoint to "localhost" and the port to 5051
    return backend.MetaData{cid, hash, hash, size, "localhost", 5051}
}
