package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

// Options
type ListOption struct {
	Verbose *bool
}

type ExportOption struct {
	Verbose         *bool
	DurationSeconds *int64
	SerialNumber    *string
	TokenCode       *string
}

// Global variables
var (
	helpMessage string

	listCommand *flag.FlagSet

	exportCommand *flag.FlagSet
	exportOption  ExportOption
)

// Define flagset
func init() {
	listCommand = flag.NewFlagSet("list", flag.ExitOnError)
	exportCommand = flag.NewFlagSet("export", flag.ExitOnError)

	helpMessage = `Subcommand list
	list: list up profiles in credential file 
	export: export session-token`
}

// Define flag of flagset
func init() {
	exportOption = ExportOption{
		Verbose: exportCommand.Bool("verbose", false, "Display log"),
		DurationSeconds: exportCommand.Int64(
			"duration-seconds", 43200,
			`The  duration, in seconds, that the credentials should remain valid.
Acceptable durations for IAM user sessions range  from  900  seconds
(15  minutes)  to  129600 seconds (36 hours), with 43200 seconds (12
hours)  as  the  default.  Sessions  for  AWS  account  owners   are
restricted  to a maximum of 3600 seconds (one hour). If the duration
is longer than one hour, the session for AWS account owners defaults
to one hour.`),
		SerialNumber: exportCommand.String(
			"serial-number", "",
			`The  identification number of the MFA device that is associated with
the IAM user who is making the GetSessionToken  call.  Specify  this
value if the IAM user has a policy that requires MFA authentication.
The value is either the serial number for a hardware device (such as
GAHT12345678 ) or an Amazon Resource Name (ARN) for a virtual device
(such as arn:aws:iam::123456789012:mfa/user  ).  You  can  find  the
device  for  an  IAM user by going to the AWS Management Console and
viewing the user's security credentials.`),
		TokenCode: exportCommand.String(
			"token-code", "",
			`The  value  provided  by  the MFA device, if MFA is required. If any
policy requires the IAM user to submit an  MFA  code,  specify  this
value. If MFA authentication is required, and the user does not pro-
vide a code when requesting a set of temporary security credentials,
the  user  will  receive an "access denied" response when requesting
resources that require MFA authentication.`),
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
	case "export":
		exportCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *exportOption.Verbose == true {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.FatalLevel)
	}

	if listCommand.Parsed() {
		ListCredentials()
	} else if exportCommand.Parsed() {
		if !exportOption.Valid() {
			log.Fatal(`serial-number, access-token options are required.
export-aws-session-token -h command helps you`)
			os.Exit(1)
		}

		ExportSessionToken(exportOption)
	}
}

func (option ExportOption) Valid() bool {
	if *option.SerialNumber != "" && *option.TokenCode != "" {
		return true
	} else {
		return false
	}
}
