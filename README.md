Lambda of Thrones
=====

[High-Concept](https://en.wikipedia.org/wiki/High-concept) GOT story plot generator written for Apex-Lambda.

# Introduction
Ever wanted to have Game of Thrones story plots being randomly generated by a Lambda set of functions? No? Well, that's too bad... this compelling need actually comes from a little experiment trying to answer 2 questions:

1. Is it possible to have a coherent strategy for logging and monitoring Lambda functions?
1. Is it possible to chain Lambda functions in the Terminal to provide a UNIX-like piping?

Both answers turned out to be yes.

# Requirements
For this project to work you'll need to:
1. have [Apex](http://apex.run/) installed on your local machine
1. have an [AWS account](https://console.aws.amazon.com) with full access to:
  - IAM roles
  - Lambda functions
  - CloudWatch
1. have [GoLang](https://golang.org/doc/install) installed on your local machine
1. have [AWS CLI](https://aws.amazon.com/cli/) tools installed on your local machine with valid `~/.aws/config` and `~/.aws/credentials` files
1. have [json - JSON love for your command line](http://trentm.com/json/#INSTALL-PROJECT-BUGS) installed

## Recommended modules
Optionally, if you want to see your logs in [ELK](https://www.elastic.co/products) :
1. have an AWS account with further full access to:
  - AWS ElasticSearch
1. have [Sense Chrome extension](https://chrome.google.com/webstore/detail/sense-beta/lhjgkmllcaadmopgmanpapmpjgmfcfig?hl=en) installed

# Installation
> Git clone yada yada

From **within** your `GOPATH`:
```bash
# create the
mkdir $GOPATH/src/github.com/edoardo849 && cd "$_"
# clone the repo
git clone https://github.com/edoardo849/lot

# enter the folder
cd lot
```

## Configuration
Initialise Apex
```bash
# set your default profile for AWS and AWS region
source ./config

# run the setup
make setup

# Project name
Project name: lot

# Project description
Project description: Lambda of Thrones
```

I always find useful to setup the aws profile in the project.json file that apex will generate:

```json
file: ./project.json
{
  "profile":"default",
}
```

you may still need to source the ./config file for every new bash session though: that's a known bug in apex.

# Running the functions

As a shorthand I've set up a `Makefile` that takes care of everything. If you want to tweak the configuration have a look at it

```bash
# run the tests
make test

# deploy the code on AWS Lambda
make deploy

# invoke the functions
make plot
```

Something like this should come out: hopefully it will make great sense... but most of the time is just silly.

```txt
_                    _         _                __   _   _
| |    __ _ _ __ ___ | |__   __| | __ _    ___  / _| | |_| |__  _ __ ___  _ __   ___  ___
| |   / _` | '_ ` _ \| '_ \ / _` |/ _` |  / _ \| |_  | __| '_ \| '__/ _ \| '_ \ / _ \/ __|
| |__| (_| | | | | | | |_) | (_| | (_| | | (_) |  _| | |_| | | | | | (_) | | | |  __/\__ \
|_____\__,_|_| |_| |_|_.__/ \__,_|\__,_|  \___/|_|    \__|_| |_|_|  \___/|_| |_|\___||___/


What if Jon Snow drank all the wine of Tywin Lannister after winter came in the Eyre because when dead men come hunting... you think it matters who sits on the iron throne?

```

## Under the hood
TODO
