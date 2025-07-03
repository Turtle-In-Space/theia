print_ok() {
  echo "$(green [+])" "$1"
}

print_info() {
  echo "$(blue [*])" "$1"
}

print_warn() {
  echo "$(orange [-])" "$1"
}

print_err() {
  echo "$(red [!])" "$1"
}
