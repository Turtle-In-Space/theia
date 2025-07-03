% theia-init(1) | Perform inital scan and create .env file
% 
% July 2025

NAME
==================================================

**theia init** - Perform inital scan and create .env file

SYNOPSIS
==================================================

**theia init** IPADDR [OPTIONS]

DESCRIPTION
==================================================

Perform inital scan and create .env file


ARGUMENTS
==================================================

IPADDR
--------------------------------------------------

IP address of target

- *Required*

OPTIONS
==================================================

--name, -n NAME
--------------------------------------------------

Name of target

- *Required*

--hostname, -h HOSTNAME
--------------------------------------------------

Hostname of target

- Needs: **--extension**

--extension, -e EXTENSION
--------------------------------------------------

URL extension of target

- Needs: **--hostname**

DEPENDENCIES
==================================================

nmap
--------------------------------------------------


ffuf
--------------------------------------------------


enum4linux-ng
--------------------------------------------------


EXAMPLES
==================================================

~~~
theia init 127.0.0.1 --name anthem --hostname anthemvm --extension thm

theia init 10.13.202.203 -n lazarus

~~~

SEE ALSO
==================================================

**theia**(1)


