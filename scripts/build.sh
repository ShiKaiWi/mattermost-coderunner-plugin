WORKDIR=$GOPATH/src/github.com/ShiKaiWi/mattermost-coderunner-plugin/
cd $WORKDIR

echo "building server-coderunner-plugin..."
GOOS=linux GOARCH=amd64 go build -o plugin server/plugin.go
cp server/plugin.yaml ./
tar czf coderunner-plugin.tar.gz plugin plugin.yaml
rm -rf plugin.yaml plugin

echo "building web-coderunner-plugin..."
cd web && npm run build && cd -
WEB_PLUGIN_DIR=com.mattermost.web-coderunner-plugin
mkdir -p $WEB_PLUGIN_DIR
cp web/dist/main.js web/plugin.json $WEB_PLUGIN_DIR
tar czf web-coderunner-plugin.tar.gz $WEB_PLUGIN_DIR
rm -rf $WEB_PLUGIN_DIR
