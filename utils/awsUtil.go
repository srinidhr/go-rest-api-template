package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

func GetAwsSesSession(awsRegion string) (sesiface.SESAPI, error) {
	// create a new AWS session client and pick up the token from .aws/credentials configured
	// by aws duo-sso. In case if token is expired, retry after signing duo-sso locally
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)

	//Creates a new SES session client for sending email
	return ses.New(awsSession), err
}
