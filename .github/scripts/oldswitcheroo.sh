#!/bin/bash

mkdir -p themes/messynotes
for i in *; do
	[ "$i" = "themes" ] && continue
	[ "$i" = "exampleSite" ] && continue
	cp -r "$i" themes/messynotes
done
cp -r exampleSite/* ./
sed -i 's!http://example\.org!https://mmessmore.github.io/hugo-messynotes!' config.yaml
