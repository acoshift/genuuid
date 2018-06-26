# genuuid

Random UUID Generator

## Installation

`go get -u github.com/acoshift/genuuid`

## Usage

```bash
$ genuuid
c591ec98-30d2-4466-8a82-75c18cb3741c

$ genuuid -n
0509726a54bf4a199cd3e8d1f3978130

$ genuuid -u
2005F6CD-AE75-4DDD-A83B-62D7F1416843

$ genuuid -n -u
55BE71D72E6A41D88C4CB87D8D3F8B5E
```

## Alternative

```bash
$ uuidgen
EC67FF7D-46CC-4278-B814-A92551D13DC0

$ uuidgen | tr -d -
CC12DAACDA1A4F68A65D8F3409113146

$ uuidgen | awk '{ print tolower($0) }'
8be8f004-9de4-464f-8db6-fc40e59c5019

$ uuidgen | awk '{ print tolower($0) }' | tr -d -
7acf42479ad848299cdd7b4f9696a181
```
