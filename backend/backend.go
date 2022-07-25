// A package for updating our Backend state
package backend

import (
    "fmt"
    "context"
    "flag"
    "os"
    "time"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/request"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)


const meta_data_bucket = "meta-data-test-bucket-jmdagwystvo2jt4b"
const obao_bucket = "obao-test-bucket-jmdagwystvo2jt4b"
const ObaoTempStore = "/tmp/"

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

// Writes the metadata to S3
func WriteMetaData(meta_data MetaData) {
    // Return the hash and size of the file
    fmt.Println("Writing MetaData:\n" +
        "Cid: " + meta_data.Cid + "\n" +
        "Obao_name: " + meta_data.Obao_name + "\n" +
        "Hash: " + meta_data.Hash + "\n" +
        fmt.Sprintf("Size: %d", meta_data.Size) + "\n")

    // Source: https://github.com/aws/aws-sdk-go
    // Write the metadata to S3
    var bucket, key string
    var timeout time.Duration

    flag.StringVar(&bucket, "b", "", "Bucket name.")
    flag.StringVar(&key, "k", "", "Object key name.")
    flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
    flag.Parse()

    sess := session.Must(session.NewSession())
    svc := s3.New(sess)
    ctx := context.Background()
    var cancelFn func()
    if timeout > 0 {
        ctx, cancelFn = context.WithTimeout(ctx, timeout)
    }
    // Ensure the context is canceled to prevent leaking.
    // See context package for more information, https://golang.org/pkg/context/
    if cancelFn != nil {
        defer cancelFn()
    }

    // Upload the MetaData to S3 as a JSON object
    _, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
        Bucket: aws.String(meta_data_bucket),
        Key:    aws.String(meta_data.Cid),
        Body:   aws.ReadSeekCloser(strings.NewReader(fmt.Sprintf("{\"Cid\":\"%s\",\"Obao_name\":\"%s\",\"Hash\":\"%s\",\"Size\":%d}", meta_data.Cid, meta_data.Obao_name, meta_data.Hash, meta_data.Size))),
    })
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
            // If the SDK can determine the request or retry delay was canceled
            // by a context the CanceledErrorCode error code will be returned.
            fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
        } else {
            fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
        }
        os.Exit(1)
    }

    fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
}

func WriteObao(obao_name string) {
    obao_path := "/tmp/" + obao_name

    fmt.Println("Writing Obao:\n" +
        "Obao_name: " + obao_path + "\n")
    // TODO - write the obao file to S3
}