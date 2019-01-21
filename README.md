# AWS token controller

## Why export-aws-session-token?

AWS  is supporting cli MFA by session token. The command is `aws sts get-session-token --serial-number arn:aws:iam::USER-ID:mfa/USER-NAME --token-code XXXXXX`, and the output skeleton is like below. But it is bothering to copy keys, there is needs to do this boring stuff automatically.

```
{
    "Credentials": {
        "SecretAccessKey": "XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
        "SessionToken": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
        "Expiration": "2019-01-06T04:58:13Z",
        "AccessKeyId": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    }
}
``` 


## Installation

Check the [releases](https://github.com/Buzzvil/export-aws-session-token/releases) for the latest version.

## Command 

Command `export-aws-session-token  --serial-number arn:aws:iam::USER-ID:mfa/USER-NAME --token-code XXXXXX` make output like below to export **AWS credential** easily. And you run with copy the result or use backtick `\`` each side of command.

```
AWS_SESSION_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXXXXXX
export AWS_SECRET_ACCESS_KEY=XXXXXXXXXXXXXXXXXXXX
export AWS_ACCESS_KEY_ID=XXXXXXXXXXXXXXXXXXXX
```

## Version
 
### version 0.1

Initialize `aws-session-token` command.
