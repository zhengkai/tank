SHELL:=/bin/bash

NG := "./node_modules/@angular/cli/bin/ng.js"

$(shell cp -n .config.default config.ini)
-include config.ini

local:
	$(NG) serve --port $(port) --host 127.0.0.1 --public-host "$(domain)"

init:
	NG_CLI_ANALYTICS=ci npm install
	npm audit fix

prod:
	./dist/build.sh tank.9farm.com prod
	ssh eirena 'mkdir -p /www/tank/public/'
	rsync --partial -vzrtopg -e ssh ./dist/prod/ eirena:/www/tank/public/
