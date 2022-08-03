package processing

import (
    "fmt"
    "os/exec"
    "strings"
)


// Get the Blake3 hash of a file
func get_hash(file_path string) (string) {
    cmd := exec.Command("bao", "hash", file_path)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return ""
    }

    return strings.Split(string(stdout), "\n")[0]
}

// Encodes a file into an obao file
// This file is saved in the system's 'ObaoTempStore' directory, which is 'temp/'
// The  cid of the file is used to name the file
func encode_obao(file_path string, hash string, output_dir string) (error) {
    // Determine the path of the obao file
    obao_path := output_dir + hash
    // Generate an obao file for the file
    cmd := exec.Command("bao", "encode", file_path, "--outboard", obao_path)
    _, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    return nil
}

// Takes a file in, returns the Blake3 hash and writes the obao to temporary storage, indexed by the hash
func ProcessFile(file_path string, output_dir string) (string, error) {
    hash := get_hash(file_path)
    err := encode_obao(file_path, hash, output_dir)
    if err != nil {
        fmt.Println(err.Error())
    }

    return hash, err
}

// Don't need to get the CID anymore, Estuary does that
// Get the CID of a file
// func getCID(file_path string) (string) {
//     cmd := exec.Command("ipfs", "add", file_path, "-q", "--cid-version", "1")
//     stdout, err := cmd.Output()
//
//     if err != nil {
//         fmt.Println(err.Error())
//         return ""
//     }
//
//     return strings.Split(string(stdout), "\n")[0]
// }

// Don't need to het the size anymore, Estuary does that
// // Get the size of a file
// func getSize(file_path string) (int64) {
//     // Determine the size of the file
//     cmd := exec.Command("stat", "-c%s", file_path)
//     stdout, err := cmd.Output()
//
//     if err != nil {
//         fmt.Println(err.Error())
//         return 0
//     }
//
//     val, err := strconv.Atoi(strings.Split(string(stdout), "\n")[0])
//     if err != nil {
//         fmt.Println(err.Error())
//         return 0
//     }
//     return int64(val)
// }