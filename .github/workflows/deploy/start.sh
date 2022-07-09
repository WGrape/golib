echo -e "------------ The script start.sh is running ------------"
if [ -n "$1" ]; then
  noBaseDir="--noBaseDir"
  if [ "$1" != "$noBaseDir" ]; then
    cd $1 # The base directory
  fi
else
  echo "Missed the base directory param."
  exit 1
fi

echo $(go env)

echo -e "------------ The script start.sh is stopped ------------"