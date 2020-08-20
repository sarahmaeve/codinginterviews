#!/bin/bash

for BUCKET in $(aws s3api list-buckets | grep \"Name\" | cut -d '"' -f 4); 
  do echo -ne $BUCKET ": ";
    if aws s3api get-bucket-encryption --bucket $BUCKET > /dev/null 2>&1; then
       echo "encrypted"
    else
       echo "unencrypted"
    fi
  done

