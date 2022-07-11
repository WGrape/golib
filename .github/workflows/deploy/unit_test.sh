echo -e "------------ The script unit_test.sh is running ------------"
if [ -n "$1" ]; then
  noBaseDir="--noBaseDir"
  if [ "$1" != "$noBaseDir" ]; then
    cd $1 # The base directory
  fi
else
  echo "Missed the base directory param."
  exit 1
fi

# 1、Test the directory of time
echo "1. test the directory of time"
cd time && go test -v . && cd ../
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, unit test failed <<<<<<<<<<<<"
  exit 1
fi

# 2、Test the directory of array
echo "2. test the directory of array"
cd array && go test -v . && cd ../
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, unit test failed <<<<<<<<<<<<"
  exit 1
fi

# 3、Test the directory of desensitization
echo "3. test the directory of desensitization"
cd desensitization && go test -v . && cd ../
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, unit test failed <<<<<<<<<<<<"
  exit 1
fi

# 4、Test the directory of permutation
echo "4. test the directory of permutation"
cd permutation && go test -v . && cd ../
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, unit test failed <<<<<<<<<<<<"
  exit 1
fi

echo -e "------------ The script unit_test.sh is stopped ------------"
