// A package for updating our Backend state
package backend

import (
    "fmt"
    "context"
    "flag"
    "bytes"
    "io"
    "os"
    "time"
    "strings"

    // AWS SDK Go packages:
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/request"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// TODO: Make these configurable
// Our S3 bucket names
const meta_data_bucket = "meta-data-test-bucket-jmdagwystvo2jt4b"
const obao_bucket = "obao-test-bucket-jmdagwystvo2jt4b"

// Where we temporarily store the Obao file
const ObaoTempStore = "/tmp/"

// What data we need to write to S3
type MetaData struct {
    // The CID of the file as a string
    Cid string
    // The name of an obao file
    Obao_name string
    // The blake3 hash of the file as a string
    Hash string
    // The size of the file
    Size int64
    // The endpoint of where the file will be stored
    Endpoint string
    // The port of where the file will be stored
    Port int16
}

// Write our Meta-Data and Obao file to S3 into their respective buckets
func WriteToS3(meta_data MetaData) {
    // Initialize the S3 service client and context for the request
    sess := session.Must(session.NewSession())
    svc := s3.New(sess)
    ctx := context.Background()
    var cancelFn func()
    if timeout > 0 {
        ctx, cancelFn = context.WithTimeout(ctx, timeout)
    }
    // Ensure the context is canceled to prevent leaking.
    if cancelFn != nil {
        defer cancelFn()
    }

    // Write the MetaData to S3
    WriteMetaData(meta_data)
    // Write the Obao file to S3
    WriteObao(meta_data.Obao_name)
}

// Writes the metadata to S3
func write_meta_data(meta_data MetaData, svc *s3.S3, ctx context.Context) (err error) {
    // Return the hash and size of the file
    fmt.Println("Writing MetaData:\n" +
        "Cid: " + meta_data.Cid + "\n" +
        "Obao_name: " + meta_data.Obao_name + "\n" +
        "Hash: " + meta_data.Hash + "\n" +
        fmt.Sprintf("Size: %d", meta_data.Size) + "\n")

    // Upload the MetaData to S3 as a JSON object
    _, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
        Bucket: aws.String(meta_data_bucket),
        Key:    aws.String(meta_data.Cid),
        Body:   aws.ReadSeekCloser(
            strings.NewReader(
                fmt.Sprintf(
                    "{
                        \"cid\":\"%s\",
                        \"hash\":\"%s\",
                        \"obao_name\":\"%s\",
                        \"size\":%d
                    }", meta_data.Cid, meta_data.Obao_name, meta_data.Hash, meta_data.Size)
                )
            ),
    })
    return err
}

// Read in obao file an return the bytes
func read_obao(obao_path string) (obao_bytes []byte, err error) {
    // Read the contents of the obao file into a buffer
    body := &bytes.Buffer{}
    file, err := os.Open(obao_path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    _, err = io.Copy(body, file)
    if err != nil {
        fmt.Println("Error copying file:", err)
        os.Exit(1)
    }

    // Return the bytes
    return body.Bytes(), nil
}

func write_obao(obao_path string, svc *s3.S3, ctx context.Context) (err error) {
    // Read the contents of the obao file into a buffer
    body, err := read_obao(obao_path)
    if err != nil {
        return err
    }

    // Upload the body buffer to S3
    _, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
        Bucket: aws.String(obao_bucket),
        Key:    aws.String(hash),
        Body:   aws.ReadSeekCloser(body),
    })
    return err
}