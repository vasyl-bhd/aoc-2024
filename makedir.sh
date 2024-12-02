DIRNAME=day$1

mkdir -p $DIRNAME/p1 $DIRNAME/p2
touch $DIRNAME/input.txt
echo package main | tee $DIRNAME/p1/p1.go $DIRNAME/p2/p2.go > /dev/null