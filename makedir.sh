base_folder="."

# Find the highest day folder
highest_day=$(ls -d ${base_folder}/day* 2>/dev/null | sed -E 's/.*day([0-9]+)/\1/' | sort -n | tail -n 1)
if [ -z "$highest_day" ]; then
  highest_day=0
fi

# Calculate the next day number
next_day=$((highest_day + 1))

# Create the new directory and subdirectories
DIRNAME="day$next_day"
mkdir -p "$DIRNAME/p1" "$DIRNAME/p2"

# Create an input file and Go files with boilerplate code
touch "$DIRNAME/input.txt"
echo "package main" | tee "$DIRNAME/p1/p1.go" "$DIRNAME/p2/p2.go" > /dev/null
