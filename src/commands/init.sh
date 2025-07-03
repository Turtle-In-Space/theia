has_match() {
  local pattern="$1"
  local file="$2"

  if  rg -q "$pattern" "$file" ; then
    return 0
  else
    return 1
  fi
}

# service name | command
run_script() {
  if has_match "$1" $port_scan ; then
    print_info "starting $1 scan..."
    $3
  else
    print_warn "did not find port for: $1 scan"
  fi
}

# create vars
name=${args[--name]}
env_file=$name.env
port_scan=export/port-scan.xml

# Create dir
print_info "creating $name/..."
mkdir -p $name/export
cd $name

# Create .env
print_info "creating .env file"
theia_set_command
source $env_file

# scan target
print_info "starting nmap scan..."
nmap -sC -sV -T4 -p- $ip -oX $port_scan -oN port-scan.nmap

# script commands
http_scan_command="ffuf -u $url/FUZZ -w /usr/share/seclists/Discovery/Web-Content/big.txt -recursion -recursion-depth 1 -t 100 -of md -o web-dir.md"
smb_scan_command="enum4linux-ng -As $ip -oJ smb-enum"

# run enum scripts
run_script "http" $http_scan_command
run_script "netbios-ssn" $smb_scan_command

print_ok "all scans complete."

if [ -z $host ] ; then
  echo ""
  print_info "add the following to /etc/hosts"
  echo $ip $host.$ext
fi
