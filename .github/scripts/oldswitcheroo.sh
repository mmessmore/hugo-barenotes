#!/bin/bash

mkdir -p themes/messynotes
for i in *; do
	[ "$i" = "themes" ] && continue
	[ "$i" = "exampleSite" ] && continue
	cp -r "$i" themes/messynotes
done
cp -r exampleSite/* ./
sed -i 's/example\.org/mmessmore.github.io\/hugo-messynotes/' config.yaml
