VERSION=0.1
GO111MODULE=on

echo "Building project"
mkdir -p builds
gox -output="builds/bitwarden-pinentry_{{.OS}}_{{.Arch}}_$VERSION" -os="linux windows darwin" -arch="amd64 386"
echo "Done"
