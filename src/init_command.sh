print_info() {
  echo "$(blue [*])" "$1"
}

has_match() {
  local pattern="$1"
  local file="$2"

  if  rg -q "$pattern" "$file" ; then
    return 0
  else
    return 1
  fi
}

name=${args[--name]}

print_info "creating $name/..."
mkdir $name

print_info "creating .env file"
# Create .env
cd $name
theia_set_command
cd ..

source $name/$name.env

print_info "starting nmap scan..."
nmap -sC -sV -T4 -p- $ip -oX $name/port-scan.xml -oN $name/port-scan.nmap

# enum web service shares
if has_match "http" $name/port-scan.nmap ; then
  print_info "starting ffuf scan..."
  ffuf -u $url/FUZZ -w /usr/share/seclists/Discovery/Web-Content/big.txt -recursion -recursion-depth 2 -t 100 -of md -o $name/web-dir.md
  firefox $url
else
  print_info "found no web port"
fi

# enum smb shares
if has_match "netbios-ssn" $name/port-scan.nmap ; then
  print_info "startig enum4linux-ng scan..."
  enum4linux-ng -As $ip -oJ smb-enum
else
  print_info "found no smb ports"
fi

print_info "all scans complete."
print_info "add the following to /etc/hosts"
echo $ip $host.$ext
