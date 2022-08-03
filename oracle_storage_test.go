package main

import ("testing", "os")

func TestStore(t *testing.T) {
    // Our Test file
    filename := "test/ethereum.pdf"
    // Get the size of the file
    filesize := 956232
    // Get the CID of the file
    cid = "bafybeigiysh5xsklm4hailn25bl6ezshkzmtsewo6vbdwjvrpg7lqhz4ae"

    // Store the file
    err := Store(filename, filesize, cid)
    if err != nil {
        t.Error(err.Error())
    }
}
