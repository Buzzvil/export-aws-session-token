/*
AWS GO SDK
- session
- sts
*/
package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	log "github.com/sirupsen/logrus"
)

type Output struct {
	AccessKeyId     string `name:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `name:"AWS_SECRET_ACCESS_KEY"`
	SessionToken    string `name:"AWS_SESSION_TOKEN"`
}

func ExportSessionToken(option ExportOption) {
	// Load session
	log.Info("Load session")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: *option.Profile,
	}))
	svc := sts.New(sess)

	// Get session token
	output, err := svc.GetSessionToken(&sts.GetSessionTokenInput{
		DurationSeconds: option.DurationSeconds,
		SerialNumber:    option.SerialNumber,
		TokenCode:       option.TokenCode,
	})
	if err != nil {
		log.WithError(err).Fatal("Error occured when get session token from STS")
		os.Exit(1)
	}

	log.Info("Exporting...")
	o := Output{
		AccessKeyId:     *output.Credentials.AccessKeyId,
		SecretAccessKey: *output.Credentials.SecretAccessKey,
		SessionToken:    *output.Credentials.SessionToken,
	}
	// Export by using tag
	// https://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go
	// https://golang.org/pkg/reflect/#StructTag
	ret := ""
	vo, to := reflect.ValueOf(o), reflect.TypeOf(o)
	for i := 0; i < vo.NumField(); i++ {
		ret = strings.Join([]string{
			fmt.Sprintf("export %s=%s \n", to.Field(i).Tag.Get("name"), vo.Field(i).Interface().(string)),
			ret,
		}, "")
	}
	fmt.Println(ret)
}
