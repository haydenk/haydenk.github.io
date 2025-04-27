.PHONY: clean hugo serve


clean:
	-rm -rv public
	-rm -rv .hugo_build.lock
	-rm -rv hugo_stats.json

serve:
	hugo server -D --watch

build:
	hugo --gc --minify

dev: clean serve

slug := $(shell echo $(title) | tr '[:upper:]' '[:lower:]' | sed 's/ /_/g')

post:
	hugo new content --kind default "posts/$(shell date +%s)_$(slug).md"

today:
	hugo new content --kind today "today/$(shell date +%s)_$(shell date +%Y%m%d).md"