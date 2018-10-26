WORKDIR=$GOPATH/src/github.com/ShiKaiWi/mattermost-coderunner-plugin/
cd $WORKDIR
GOOS=linux GOARCH=amd64 go build -o coderunner-plugin server/plugin.go
cp server/plugin.yaml coderunner-plugin.yaml
tar czf coderunner-plugin.tar.gz coderunner-plugin coderunner-plugin.yaml
rm -rf coderunner-plugin.yaml coderunner-plugin
