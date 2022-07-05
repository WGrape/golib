echo -e "------------ The script unit_test.sh is running ------------"
cd ../../ # The base directory

# 1、Test the directory of time
echo "1. test the directory of time"
cd time && go test -v .
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, unit test failed <<<<<<<<<<<<"
  exit 1
fi

echo -e "------------ The script unit_test.sh is stopped ------------"
