# binary will be $(go env GOPATH)/bin/golangci-lint
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

echo -e "------------ The script check_code.sh is running ------------"
cd ../../ # The base directory

# 1、Check the golangci-lint is exist
which "golangci-lint" >/dev/null 2>&1
if [ $? -ne 0 ]; then
  echo -e "1. command golangci-lint not exist, installing it ..."
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
else
  echo -e "1. golangci-lint is installed"
fi

# 2、Check the golangci-lint version
echo "2." $(golangci-lint --version)

# 3、Run the golangci-lint
golangci-lint run
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, check code failed <<<<<<<<<<<<"
  exit 1
fi
echo "3. golangci-lint check passed"

echo -e "------------ The script check_code.sh is stopped ------------"