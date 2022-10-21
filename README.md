[![Go](https://github.com/dhcgn/age-fs/actions/workflows/go.yml/badge.svg)](https://github.com/dhcgn/age-fs/actions/workflows/go.yml)

# age-fs

Mount a folder and access it through your os, all files will be transparent encrypted und decrypted with https://age-encryption.org/.

## Goals

1. Leave the file format intact, so that the tooling with age is always working
1. Everything is encrypted (files, folders, meta data)
1. Source can be Filesystem, S3 or other APIs

### Milestones

1. 🟢 Source is Filesystem
1. 🟡 Files are encrypted
1. 🔴 Files and Folders are encrypted
1. 🔴 Source can be S3



## Usage

### Prerequisites

1. Install [age and age-keygen](https://age-encryption.org/)
1. Create a keypair with `age-keygen -o key.txt`

```
# created: 2022-10-21T21:39:35+02:00
# public key: age1gy7cnv24p5ezejztmszdalvw3zach42tzxsp03rmsdxj3n3524us76zc2z
AGE-SECRET-KEY-13CA3YEGDVYC8FAGN8L6X89CQ3SGJUFUKVJQWPRTJY7RAHZPHGJGSXR9T0W
```

### Usage Linux

```
age-fs -key key.txt -mount /tmp/age-fs
```

### Usage Windows
```
age-fs -key key.txt -mount X:\
```
