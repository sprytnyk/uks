# UKS

![Example](https://raw.githubusercontent.com/vald-phoenix/uks/master/media/example.svg)

## Overview

This dummy util goes through a list of organisations licensed to sponsor
workers on the Worker and Temporary Worker immigration routes and checks if a
provided company in this list.

It may be useful when you travel abroad and then go back to the UK. Sometimes
companies forget to renew a license and the law says it our requirement to
check it before travelling to the UK, otherwise, it may be considered as an
offence fly and it happens that a company has an expired license, moreover,
nobody wants to stuck at the border.

The list lives here:  
https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers

## Usage

```bash
$ go get -u github.com/vald-phoenix/uks
$ uks -h
Usage of uks:

    $ uks "Acme Ltd."
    Checks if the provided company in the register of licensed sponsors.
    Output:
        Company: Acme Ltd.
        Found: false

    $ uks -o json "Acme Ltd."
    Print output in format: json/text (default "text")
    Output:
        {"company":"Acme Ltd.","found":false}

```

## Goal of this project

The main goal of this project is to make feel of Golang being a Pythonista and
have a handy CLI util because we all love them.

### Things to do

- [ ] Caching
- [ ] Testing
- [ ] Check by a license number and more fancy checks