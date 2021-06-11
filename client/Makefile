SHELL:=/bin/bash

NG := "./node_modules/@angular/cli/bin/ng"

$(shell cp -n .config.default config.ini)
-include config.ini

local:
	$(NG) serve --port $(port) --host 127.0.0.1 --public-host "$(domain)"

init:
	NG_CLI_ANALYTICS=ci npm install
	npm audit fix

prod:
	./dist/build.sh tank.9farm.com prod
	rsync --partial -vzrtopg -e ssh ./dist/prod/ freya:/www/tank/public/
