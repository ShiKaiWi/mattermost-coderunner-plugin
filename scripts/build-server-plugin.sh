WORKDIR=$GOPATH/src/github.com/ShiKaiWi/mattermost-coderunner-plugin/
cd $WORKDIR

GOOS=linux GOARCH=amd64 go build -o plugin server/plugin.go
cp server/plugin.yaml ./
tar czf coderunner-plugin.tar.gz plugin plugin.yaml
rm -rf plugin.yaml plugin