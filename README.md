# jscomp
Javascript compressor.

Generates a JS file containing the compressed input encoded in base64 with a light bootstrap JS code that decompresses and evals the original JS input.

## Why?
We needed to reduce GopherJS generated files to fit into wechat miniprogram restrictions.

## Install
go get github.com/jaracil/jscomp

## Usage
jscomp -i input_file > output_file

