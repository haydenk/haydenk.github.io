.PHONY: clean hugo serve


clean:
	-rm -rfv public
	-rm -rfv .hugo_build.lock
	-rm -rfv hugo_stats.json

serve:
	hugo server -D --watch --disableFastRender

build:
	hugo --gc --minify

dev: clean serve
