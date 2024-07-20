package main

import (
    "context"
    "fmt"
    "os"
    "path/filepath"
    "github.com/ipfs/kubo/client/rpc"
    path "github.com/ipfs/boxo/path"
    files "github.com/ipfs/boxo/files"
)

func main() {
    // Connect to local IPFS node
    node, err := rpc.NewLocalApi()
    if err != nil {
        fmt.Printf("Error creating IPFS client: %v\n", err)
        return
    }

    // Pin a given file by its CID
    // create a context
    ctx := context.Background()

    // Get the current working directory
    currentDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current directory:", err)
        return
    }
    // Get the grandparent directory
    grandparentDir := filepath.Dir(filepath.Dir(currentDir))
    // Get the directory of the file
    filePath := filepath.Join(grandparentDir, "ipfs-node1-data", "file_storage", "test_it.txt")
    //Print out
    fmt.Println("Current directory:", currentDir)
    fmt.Println("Grandparent directory:", grandparentDir)
    fmt.Println("Path of file:", filePath)

    // Create a temporary file to add to IPFS
    tempFile, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error creating temporary file: %v\n", err)
        return
    }
    defer tempFile.Close() // Clean up

    // Write some data to the file
    /*fileContent := []byte("Hello, IPFS!")
    if _, err := tempFile.Write(fileContent); err != nil {
        fmt.Printf("Error writing to temporary file: %v\n", err)
        return
    }*/

    // Rewind the file
    if _, err := tempFile.Seek(0, 0); err != nil {
        fmt.Printf("Error seeking to beginning of file: %v\n", err)
        return
    }

    // Create a Unixfs file from the temporary file
    file := files.NewReaderFile(tempFile)

    // Generate the CID of file + Add the file to IPFS
    cidFile, err := node.Unixfs().Add(ctx, file)
    if err != nil {
        fmt.Printf("Error adding file to IPFS: %v\n", err)
        return
    }

    // Print the CID of the added file
    fmt.Printf("File added to IPFS with CID: %s\n", cidFile.RootCid().String())

    // This is the cid of the file we wanna pin
    cid := filepath.Join("/ipfs", cidFile.RootCid().String())
    p, err := path.NewPath(cid)
	if err != nil {
		fmt.Printf("Error creating path: %v]n", err)
	}
    //Pin file
    err = node.Pin().Add(ctx, p)
    if err != nil {
        fmt.Printf("Error pinning file: %v\n", err)
        return
    }

    fmt.Println("File pinned successfully!")
}