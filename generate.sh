#!/bin/bash

base_dir=`pwd`

cd cmd/mg
go install
cd $base_dir
go generate
