base_folder="."

highest_day=$(ls -d ${base_folder}/day* 2>/dev/null | grep -oP 'day\K[0-9]+' | sort -n | tail -n 1)
if [ -z "$highest_day" ]; then
  highest_day=0
fi
next_day=$((highest_day + 1))

DIRNAME=day$next_day

mkdir -p $DIRNAME/p1 $DIRNAME/p2
touch $DIRNAME/input.txt
echo package main | tee $DIRNAME/p1/p1.go $DIRNAME/p2/p2.go > /dev/null