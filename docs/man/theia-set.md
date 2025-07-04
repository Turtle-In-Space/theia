% theia-set(1) | Create .env file
% 
% July 2025

NAME
==================================================

**theia set** - Create .env file

SYNOPSIS
==================================================

**theia set** IPADDR [OPTIONS]

DESCRIPTION
==================================================

Create .env file


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

EXAMPLES
==================================================

~~~
theia set 127.0.0.1 --name anthem --hostname anthemvm --extension thm

theia set 10.13.202.203 -n lazarus

~~~

SEE ALSO
==================================================

**theia**(1)


