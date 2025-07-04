ip=${args[ipaddr]}
name=${args[--name]}
host=${args[--hostname]}
ext=${args[--extension]}
env_file="$name.env"

echo '# Created by Theia' >> $env_file

echo "ip=$ip" >> "$env_file"
if [ -z $host ] ; then
  url="http://$ip"
else
  url="http://$host.$ext"
  echo "host=$host" >> "$env_file"
fi

echo "url=$url" >> "$env_file"

print_ok "Created $env_file"
