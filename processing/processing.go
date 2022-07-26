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
func getCID(file_path string) (string) {
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

// Get a Deterministic ID for a storage deal
// TODO: Get this to reflect the actual deal-id on the blockchain
func getDealId(cid string, hash string, size int64) (string) {
//     // Generate a Sha256 hash of the CID, Blake3 hash, and size
//     deal_hash := cid + hash + strconv.FormatInt(size, 10)
//     // Generate a Sha246 hash of the deal_hash
//     cmd := exec.Command("sha256sum", "-b", deal_hash)
//     stdout, err := cmd.Output()
//     if err != nil {
//         fmt.Println(err.Error())
//         return ""
//     }
//     return string(stdout)
    return cid
}

// Encodes a file into an obao file
// This file is saved in the system's 'ObaoTempStore' directory, which is 'temp/'
// The deal id of the file is used to name the file
func encodeObao(file_path string, deal_id string) (error) {
    // Determine the path of the obao file
    obao_path := backend.ObaoTempStore + deal_id
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
// - a meta_data struct containing the file's cid, blake3 hash, and size
func ProcessFile(file_path string) (string, backend.MetaData) {
    println("Processing file: " + file_path)

    // Get the CID of the file
    cid := getCID(file_path)
    // Get the size of the file
    size := getSize(file_path)
    // Generate an obao file for the file
    hash := getHash(file_path)

    // Generate a deterministic deal id for the file
    deal_id := getDealId(cid, hash, size)

    // Encode the file into an obao file
    err := encodeObao(file_path, deal_id)
    if err != nil {
        fmt.Println(err.Error())
    }

    // Return the meta_data struct
    return deal_id, backend.MetaData{cid, hash, size}
}
