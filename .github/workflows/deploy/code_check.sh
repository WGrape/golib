# binary will be $(go env GOPATH)/bin/golangci-lint
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

echo -e "------------ The script check_code.sh is running ------------"
if [ -n "$1" ]; then
  noBaseDir="--noBaseDir"
  if [ "$1" != "$noBaseDir" ]; then
    cd $1 # The base directory
  fi
else
  echo "Missed the base directory param."
  exit 1
fi

# 1、Check the golangci-lint is exist
which "golangci-lint" >/dev/null 2>&1
if [ $? -ne 0 ]; then
  echo -e "1. command golangci-lint not exist, installing it ..."
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
else
  echo -e "1. golangci-lint is installed"
fi

# 2、Check the golangci-lint version
# echo "2." $(golangci-lint --version)
echo "2." $($(go env GOPATH)/bin/golangci-lint --version)

# 3、Run the golangci-lint
# Bad case 1 : context loading failed: no go files to analyze, https://github.com/golangci/golangci-lint/issues/825
$(go env GOPATH)/bin/golangci-lint run
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, check code failed <<<<<<<<<<<<"
  exit 1
fi
echo "3. golangci-lint check passed"

echo -e "------------ The script check_code.sh is stopped ------------"