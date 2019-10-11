package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
   	"github.com/aws/aws-sdk-go/aws/awserr" 
)
func HandleLambdaEvent () {

svc := ec2.New(session.New())
input := &ec2.StartInstancesInput{
    InstanceIds: []*string{
        aws.String("i-XPTO"),
    },
}
result, err := svc.StartInstances(input)
if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        fmt.Println(err.Error())
    }
    return
    }
    fmt.Println(result)
}
    
func main() {
	lambda.Start(HandleLambdaEvent)
}
