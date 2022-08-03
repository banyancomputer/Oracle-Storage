// A package for updating our Backend state
package backend

import (
    "fmt"
    "context"
    "bytes"
    "io"
    "os"
    "time"
    "strings"

    // AWS SDK Go packages:
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// What data we need to write to S3
type MetaData struct {
    // The CID of the file as a string
    Cid string
    // The blake3 hash of the file as a string
    Hash string
    // The size of the file
    Size int64
}

// // What fields define an endpoint
// type Endpoint struct {
//     Host string
//     Port int16
//
//     // Maybe we add more things here later, for backups and whatnot
// }

// Where we temporarily store the Obao file
const ObaoTempStore = "/tmp/"

// TODO: Make these configurable
// Our S3 bucket names
const meta_data_bucket = "meta-data-bucket-dev-9lz7kptz8kihj7qx"
const obao_bucket = "obao-file-bucket-dev-9lz7kptz8kihj7qx"
// const endpoint_bucket = "endpoint-bucket-dev-9lz7kptz8kihj7qx"
// Our AWS region
const aws_region = "us-east-2"
// The timeout for our requests
const timeout = time.Duration(10) * time.Second

// Write our Meta-Data, Obao file, and Endpoint to S3 into their respective buckets
// All data is indexed by the deal ID
func WriteToS3(meta_data MetaData) {
    // Initialize the S3 service client and context for the request
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(aws_region)},
    ))
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
    err := write_meta_data(meta_data, svc, ctx)
    if err != nil {
        fmt.Println("Error writing meta-data:", err)
    }
    // Write the Obao file to S3. The file is named by its Blake3 hash and is stored in the ObaoTempStore
    err = write_obao(meta_data.Hash, svc, ctx)
    if err != nil {
        fmt.Println("Error writing obao file:", err)
    }
//     // Write the endpoint to S3. TODO: This eventually needs to be updated to be a real endpoint
//     default_endpoint := Endpoint{Host: "localhost", Port: 5001}
//     err = write_endpoint(deal_id, default_endpoint, svc, ctx)
//     if err != nil {
//         fmt.Println("Error writing endpoint:", err)
//     }
}

// Writes the metadata to S3, indexes MetaData by the files CID
func write_meta_data(cid string, meta_data MetaData, svc *s3.S3, ctx context.Context) (error) {
    // Upload the MetaData to S3 as a JSON object
    _, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
        Bucket: aws.String(meta_data_bucket),
        Key:    aws.String(meta_data.Cid),
        Body:   aws.ReadSeekCloser(
            strings.NewReader(
                fmt.Sprintf(
                    `{"cid": "%s","hash":"%s","size":%d}`,
                    meta_data.Cid, meta_data.Hash, meta_data.Size,
                ),
            ),
        ),
    })
    return err
}

// TODO: Take out the endpoint stuff
// // Write an object that indexes an endpoint for a file by its deal_id
// func write_endpoint(deal_id string, endpoint Endpoint, svc *s3.S3, ctx context.Context) (error) {
//     // Upload the endpoint to S3
//     _, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
//         Bucket: aws.String(endpoint_bucket),
//         Key:    aws.String(deal_id),
//         Body:   aws.ReadSeekCloser(
//             strings.NewReader(
//                 fmt.Sprintf(
//                     `{"host": "%s","port": %d}`,
//                     endpoint.Host, endpoint.Port,
//                 ),
//             ),
//         ),
//     })
//     return err
// }

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
        return nil, err
    }

    // Return the bytes
    return body.Bytes(), nil
}

// Write the contents of an obao file to S3, indexed by the file's Blake3 hash
func write_obao(hash string, svc *s3.S3, ctx context.Context) (error) {
    obao_path := ObaoTempStore + hash
    // Read the contents of the obao file into a buffer
    body, err := read_obao(obao_path)
    if err != nil {
        return err
    }

    // Upload the body buffer to S3
    _, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
        Bucket: aws.String(obao_bucket),
        Key:    aws.String(cid),
        Body:   aws.ReadSeekCloser(bytes.NewReader(body)),
    })
    return err
}