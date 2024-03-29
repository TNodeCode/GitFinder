# ARM and x86
for os in linux windows darwin
do
    for arch in "386" "amd64" "arm" "arm64"
    do
        mkdir -p ./dist/$os/$arch/
        CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -o ./dist/$os/$arch/ .
    done
done