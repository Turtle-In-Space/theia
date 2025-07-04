# theia set

Create .env file

## Usage

```bash
theia set IPADDR [OPTIONS]
```

## Examples

```bash
theia set 127.0.0.1 --name anthem --hostname anthemvm --extension thm
```

```bash
theia set 10.13.202.203 -n lazarus
```

## Arguments

#### *IPADDR*

IP address of target

| Attributes      | &nbsp;
|-----------------|-------------
| Required:       | ✓ Yes

## Options

#### *--name, -n NAME*

Name of target

| Attributes      | &nbsp;
|-----------------|-------------
| Required:       | ✓ Yes

#### *--hostname, -h HOSTNAME*

Hostname of target

| Attributes      | &nbsp;
|-----------------|-------------
| Needs: | *--extension*

#### *--extension, -e EXTENSION*

URL extension of target

| Attributes      | &nbsp;
|-----------------|-------------
| Needs: | *--hostname*


