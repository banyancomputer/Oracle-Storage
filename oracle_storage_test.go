package oracle_storage

import (
    "testing"
)

func TestStore(t *testing.T) {
    // Our Test file
    filename := "test/ethereum.pdf"
    // Get the size of the file
    var filesize int64 = 956232
    // Get the CID of the file
    cid := "bafybeigiysh5xsklm4hailn25bl6ezshkzmtsewo6vbdwjvrpg7lqhz4ae"

    println("Testing Store on file: " + filename)

    // Store the file
    err := Store(filename, filesize, cid)
    if err != nil {
        t.Error(err.Error())
    }
}
