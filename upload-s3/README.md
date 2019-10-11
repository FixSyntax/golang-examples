
File upload script for AWS Bucket S3.

To upload the file, change to function main () the following line:

err = AddFileToS3 (s, "FILE")

*Don't forget the bucket name, region, and having aws cli keys on your computer.