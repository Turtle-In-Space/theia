# theia init

Perform inital scan and create .env file

## Usage

```bash
theia init IPADDR [OPTIONS]
```

## Examples

```bash
theia init 127.0.0.1 --name anthem --hostname anthemvm --extension thm
```

```bash
theia init 10.13.202.203 -n lazarus
```

## Dependencies

#### *nmap*



#### *ffuf*



#### *enum4linux-ng*



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


