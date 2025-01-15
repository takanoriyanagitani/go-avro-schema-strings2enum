#!/bin/sh

export ENV_SCHEMA_NAME=sample_schema_name
export ENV_SCHEMA_NAME_SPACE=

enum_strings(){
	echo UNSPECFIED
	echo OK
	echo NG
}

enum_strings |
  ./avsc-strings2enum |
  jq .
