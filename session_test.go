package main

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestReflect(t *testing.T) {
	o := Output{
		AccessKeyId:     "foo",
		SecretAccessKey: "boo",
		SessionToken:    "bar",
	}

	answer := "AWS_ACCESS_KEY_ID=foo"
	to := reflect.TypeOf(o)
	vo := reflect.ValueOf(o)
	if test := fmt.Sprintf("%s=%s",
		to.Field(0).Tag.Get("name"),
		vo.Field(0).Interface().(string)); answer != test {

		t.Error(answer, test)
	}
}
