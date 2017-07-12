rm -rf build

function b {
	filename=build/$1/$2/lunch-cli
	
	if [[ "$1" = "windows" ]]
	then
		filename="${filename}.exe"
	fi
	
	GOOS=$1 GOARCH=$2 go build -o $filename src/main.go 
}

echo "Building for Linux"
echo "64-Bit"
b linux amd64
echo "32-Bit"
b linux 386
echo "Building for Windows"
echo "64-Bit"
b windows amd64
echo "32-Bit"
b windows 386
echo "Building for macOS"
echo "64-Bit"
b darwin amd64
echo "32-Bit"
b darwin 386