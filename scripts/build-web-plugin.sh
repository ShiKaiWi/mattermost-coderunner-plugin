cd $GOPATH/src/github.com/ShiKaiWi/mattermost-coderunner-plugin/web

npm install
npm run build

cd ..
WEB_PLUGIN_DIR=com.mattermost.web-coderunner-plugin
mkdir -p $WEB_PLUGIN_DIR

cp web/dist/main.js web/plugin.json $WEB_PLUGIN_DIR
tar czf web-coderunner-plugin.tar.gz $WEB_PLUGIN_DIR
rm -rf $WEB_PLUGIN_DIR
