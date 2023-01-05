package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateBucket(client *s3.S3, bucketname string) error { //createbucket function passing s3 session and bucketname
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketname),
	})
	return err
}
func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error { //passing s3 session,filenameor path,bucketname
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	return err
}
func main() {
	fmt.Println("Enter your choice")
	fmt.Println("1)File operations 2)S3 bucket operations")
	x := 0
	fmt.Scanln(&x)

	switch x {
	case 1:
		fmt.Println("Enter your choice")
		fmt.Println("1)file creation 2)data insertion in file 3)reading data from file ")
		x := 0
		fmt.Scanln(&x)

		switch x {
		case 1:
			filename := " "
			fmt.Println("enter file name")
			fmt.Scanln(&filename)
			f, errval := os.Create(filename)
			if errval != nil {
				fmt.Println("err", errval)
			}
			defer f.Close()
			fmt.Println(f.Name(), " : Created successfully")

		case 2:
			filename := " "
			fmt.Println("enter file name")
			fmt.Scanln(&filename)
			fmt.Print("Enter your String: \n")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input_str := scanner.Text()
			errval := os.WriteFile(filename, []byte(input_str), 0644)
			if errval != nil {
				fmt.Println("err", errval)
			}
			fmt.Println("Successful write")

		case 3:
			filename := " "
			fmt.Println("enter file name")
			fmt.Scanln(&filename)
			data, _ := ioutil.ReadFile(filename)
			fmt.Printf("%s", data)

		default:
			fmt.Println("please enter valid number")
		}
		fmt.Println("")

	case 2:
		fmt.Println("Enter your choice")
		fmt.Println("1)create s3 bucket 2)insert file into s3 bucket")
		y := 0
		fmt.Scanln(&y)
		switch y {

		case 1:
			sess, err := session.NewSessionWithOptions(session.Options{
				Profile: "default",
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				},
			})
			if err != nil {
				fmt.Println(err)
				return
			}

			s3Client := s3.New(sess)
			a := ""
			fmt.Println("enter unique bucket name")
			fmt.Scanln(&a)
			bucketName := a

			err = CreateBucket(s3Client, bucketName)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Successful bucket creation")

		case 2:
			sess, err := session.NewSessionWithOptions(session.Options{
				Profile: "default",
				Config: aws.Config{
					Region: aws.String("us-west-2"),
				},
			})

			if err != nil {
				fmt.Printf("Failed to initialize new session: %v", err)
				return
			}

			a := ""
			fmt.Println("enter unique bucket name")
			fmt.Scanln(&a)
			bucketName := a
			uploader := s3manager.NewUploader(sess)
			b := ""
			fmt.Println("enter file name")
			fmt.Scanln(&b)
			filename := b

			err = UploadFile(uploader, "test.txt", bucketName, filename)
			if err != nil {
				fmt.Printf("Failed to upload file: %v", err)
			}

			fmt.Println("Successfully uploaded file!")
		}

	}

}
