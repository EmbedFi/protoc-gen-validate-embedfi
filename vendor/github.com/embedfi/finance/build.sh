#!/bin/bash

buf generate

sed -i scaled_amount.pb.go -e 's/String() string/ProtoString() string/'


