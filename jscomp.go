package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
)

var base64Code = `(function(r){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=r()}else if(typeof define==="function"&&define.amd){define([],r)}else{var e;if(typeof window!=="undefined"){e=window}else if(typeof global!=="undefined"){e=global}else if(typeof self!=="undefined"){e=self}else{e=this}e.base64js=r()}})(function(){var r,e,n;return function(){function r(e,n,t){function o(f,i){if(!n[f]){if(!e[f]){var u="function"==typeof require&&require;if(!i&&u)return u(f,!0);if(a)return a(f,!0);var v=new Error("Cannot find module '"+f+"'");throw v.code="MODULE_NOT_FOUND",v}var d=n[f]={exports:{}};e[f][0].call(d.exports,function(r){var n=e[f][1][r];return o(n||r)},d,d.exports,r,e,n,t)}return n[f].exports}for(var a="function"==typeof require&&require,f=0;f<t.length;f++)o(t[f]);return o}return r}()({"/":[function(r,e,n){"use strict";n.byteLength=d;n.toByteArray=h;n.fromByteArray=p;var t=[];var o=[];var a=typeof Uint8Array!=="undefined"?Uint8Array:Array;var f="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";for(var i=0,u=f.length;i<u;++i){t[i]=f[i];o[f.charCodeAt(i)]=i}o["-".charCodeAt(0)]=62;o["_".charCodeAt(0)]=63;function v(r){var e=r.length;if(e%4>0){throw new Error("Invalid string. Length must be a multiple of 4")}var n=r.indexOf("=");if(n===-1)n=e;var t=n===e?0:4-n%4;return[n,t]}function d(r){var e=v(r);var n=e[0];var t=e[1];return(n+t)*3/4-t}function c(r,e,n){return(e+n)*3/4-n}function h(r){var e;var n=v(r);var t=n[0];var f=n[1];var i=new a(c(r,t,f));var u=0;var d=f>0?t-4:t;for(var h=0;h<d;h+=4){e=o[r.charCodeAt(h)]<<18|o[r.charCodeAt(h+1)]<<12|o[r.charCodeAt(h+2)]<<6|o[r.charCodeAt(h+3)];i[u++]=e>>16&255;i[u++]=e>>8&255;i[u++]=e&255}if(f===2){e=o[r.charCodeAt(h)]<<2|o[r.charCodeAt(h+1)]>>4;i[u++]=e&255}if(f===1){e=o[r.charCodeAt(h)]<<10|o[r.charCodeAt(h+1)]<<4|o[r.charCodeAt(h+2)]>>2;i[u++]=e>>8&255;i[u++]=e&255}return i}function s(r){return t[r>>18&63]+t[r>>12&63]+t[r>>6&63]+t[r&63]}function l(r,e,n){var t;var o=[];for(var a=e;a<n;a+=3){t=(r[a]<<16&16711680)+(r[a+1]<<8&65280)+(r[a+2]&255);o.push(s(t))}return o.join("")}function p(r){var e;var n=r.length;var o=n%3;var a=[];var f=16383;for(var i=0,u=n-o;i<u;i+=f){a.push(l(r,i,i+f>u?u:i+f))}if(o===1){e=r[n-1];a.push(t[e>>2]+t[e<<4&63]+"==")}else if(o===2){e=(r[n-2]<<8)+r[n-1];a.push(t[e>>10]+t[e>>4&63]+t[e<<2&63]+"=")}return a.join("")}},{}]},{},[])("/")});`

var inflateCode = `/** @license zlib.js 2012 - imaya [ https://github.com/imaya/zlib.js ] The MIT License */(function() {'use strict';var k=void 0,aa=this;function r(c,d){var a=c.split("."),b=aa;!(a[0]in b)&&b.execScript&&b.execScript("var "+a[0]);for(var e;a.length&&(e=a.shift());)!a.length&&d!==k?b[e]=d:b=b[e]?b[e]:b[e]={}};var t="undefined"!==typeof Uint8Array&&"undefined"!==typeof Uint16Array&&"undefined"!==typeof Uint32Array&&"undefined"!==typeof DataView;function u(c){var d=c.length,a=0,b=Number.POSITIVE_INFINITY,e,f,g,h,l,n,m,p,s,x;for(p=0;p<d;++p)c[p]>a&&(a=c[p]),c[p]<b&&(b=c[p]);e=1<<a;f=new (t?Uint32Array:Array)(e);g=1;h=0;for(l=2;g<=a;){for(p=0;p<d;++p)if(c[p]===g){n=0;m=h;for(s=0;s<g;++s)n=n<<1|m&1,m>>=1;x=g<<16|p;for(s=n;s<e;s+=l)f[s]=x;++h}++g;h<<=1;l<<=1}return[f,a,b]};function w(c,d){this.g=[];this.h=32768;this.c=this.f=this.d=this.k=0;this.input=t?new Uint8Array(c):c;this.l=!1;this.i=y;this.p=!1;if(d||!(d={}))d.index&&(this.d=d.index),d.bufferSize&&(this.h=d.bufferSize),d.bufferType&&(this.i=d.bufferType),d.resize&&(this.p=d.resize);switch(this.i){case A:this.a=32768;this.b=new (t?Uint8Array:Array)(32768+this.h+258);break;case y:this.a=0;this.b=new (t?Uint8Array:Array)(this.h);this.e=this.u;this.m=this.r;this.j=this.s;break;default:throw Error("invalid inflate mode");
}}var A=0,y=1;
w.prototype.t=function(){for(;!this.l;){var c=B(this,3);c&1&&(this.l=!0);c>>>=1;switch(c){case 0:var d=this.input,a=this.d,b=this.b,e=this.a,f=d.length,g=k,h=k,l=b.length,n=k;this.c=this.f=0;if(a+1>=f)throw Error("invalid uncompressed block header: LEN");g=d[a++]|d[a++]<<8;if(a+1>=f)throw Error("invalid uncompressed block header: NLEN");h=d[a++]|d[a++]<<8;if(g===~h)throw Error("invalid uncompressed block header: length verify");if(a+g>d.length)throw Error("input buffer is broken");switch(this.i){case A:for(;e+g>
b.length;){n=l-e;g-=n;if(t)b.set(d.subarray(a,a+n),e),e+=n,a+=n;else for(;n--;)b[e++]=d[a++];this.a=e;b=this.e();e=this.a}break;case y:for(;e+g>b.length;)b=this.e({o:2});break;default:throw Error("invalid inflate mode");}if(t)b.set(d.subarray(a,a+g),e),e+=g,a+=g;else for(;g--;)b[e++]=d[a++];this.d=a;this.a=e;this.b=b;break;case 1:this.j(ba,ca);break;case 2:for(var m=B(this,5)+257,p=B(this,5)+1,s=B(this,4)+4,x=new (t?Uint8Array:Array)(C.length),Q=k,R=k,S=k,v=k,M=k,F=k,z=k,q=k,T=k,q=0;q<s;++q)x[C[q]]=
B(this,3);if(!t){q=s;for(s=x.length;q<s;++q)x[C[q]]=0}Q=u(x);v=new (t?Uint8Array:Array)(m+p);q=0;for(T=m+p;q<T;)switch(M=D(this,Q),M){case 16:for(z=3+B(this,2);z--;)v[q++]=F;break;case 17:for(z=3+B(this,3);z--;)v[q++]=0;F=0;break;case 18:for(z=11+B(this,7);z--;)v[q++]=0;F=0;break;default:F=v[q++]=M}R=t?u(v.subarray(0,m)):u(v.slice(0,m));S=t?u(v.subarray(m)):u(v.slice(m));this.j(R,S);break;default:throw Error("unknown BTYPE: "+c);}}return this.m()};
var E=[16,17,18,0,8,7,9,6,10,5,11,4,12,3,13,2,14,1,15],C=t?new Uint16Array(E):E,G=[3,4,5,6,7,8,9,10,11,13,15,17,19,23,27,31,35,43,51,59,67,83,99,115,131,163,195,227,258,258,258],H=t?new Uint16Array(G):G,I=[0,0,0,0,0,0,0,0,1,1,1,1,2,2,2,2,3,3,3,3,4,4,4,4,5,5,5,5,0,0,0],J=t?new Uint8Array(I):I,K=[1,2,3,4,5,7,9,13,17,25,33,49,65,97,129,193,257,385,513,769,1025,1537,2049,3073,4097,6145,8193,12289,16385,24577],L=t?new Uint16Array(K):K,N=[0,0,0,0,1,1,2,2,3,3,4,4,5,5,6,6,7,7,8,8,9,9,10,10,11,11,12,12,13,
13],O=t?new Uint8Array(N):N,P=new (t?Uint8Array:Array)(288),U,da;U=0;for(da=P.length;U<da;++U)P[U]=143>=U?8:255>=U?9:279>=U?7:8;var ba=u(P),V=new (t?Uint8Array:Array)(30),W,ea;W=0;for(ea=V.length;W<ea;++W)V[W]=5;var ca=u(V);function B(c,d){for(var a=c.f,b=c.c,e=c.input,f=c.d,g=e.length,h;b<d;){if(f>=g)throw Error("input buffer is broken");a|=e[f++]<<b;b+=8}h=a&(1<<d)-1;c.f=a>>>d;c.c=b-d;c.d=f;return h}
function D(c,d){for(var a=c.f,b=c.c,e=c.input,f=c.d,g=e.length,h=d[0],l=d[1],n,m;b<l&&!(f>=g);)a|=e[f++]<<b,b+=8;n=h[a&(1<<l)-1];m=n>>>16;if(m>b)throw Error("invalid code length: "+m);c.f=a>>m;c.c=b-m;c.d=f;return n&65535}
w.prototype.j=function(c,d){var a=this.b,b=this.a;this.n=c;for(var e=a.length-258,f,g,h,l;256!==(f=D(this,c));)if(256>f)b>=e&&(this.a=b,a=this.e(),b=this.a),a[b++]=f;else{g=f-257;l=H[g];0<J[g]&&(l+=B(this,J[g]));f=D(this,d);h=L[f];0<O[f]&&(h+=B(this,O[f]));b>=e&&(this.a=b,a=this.e(),b=this.a);for(;l--;)a[b]=a[b++-h]}for(;8<=this.c;)this.c-=8,this.d--;this.a=b};
w.prototype.s=function(c,d){var a=this.b,b=this.a;this.n=c;for(var e=a.length,f,g,h,l;256!==(f=D(this,c));)if(256>f)b>=e&&(a=this.e(),e=a.length),a[b++]=f;else{g=f-257;l=H[g];0<J[g]&&(l+=B(this,J[g]));f=D(this,d);h=L[f];0<O[f]&&(h+=B(this,O[f]));b+l>e&&(a=this.e(),e=a.length);for(;l--;)a[b]=a[b++-h]}for(;8<=this.c;)this.c-=8,this.d--;this.a=b};
w.prototype.e=function(){var c=new (t?Uint8Array:Array)(this.a-32768),d=this.a-32768,a,b,e=this.b;if(t)c.set(e.subarray(32768,c.length));else{a=0;for(b=c.length;a<b;++a)c[a]=e[a+32768]}this.g.push(c);this.k+=c.length;if(t)e.set(e.subarray(d,d+32768));else for(a=0;32768>a;++a)e[a]=e[d+a];this.a=32768;return e};
w.prototype.u=function(c){var d,a=this.input.length/this.d+1|0,b,e,f,g=this.input,h=this.b;c&&("number"===typeof c.o&&(a=c.o),"number"===typeof c.q&&(a+=c.q));2>a?(b=(g.length-this.d)/this.n[2],f=258*(b/2)|0,e=f<h.length?h.length+f:h.length<<1):e=h.length*a;t?(d=new Uint8Array(e),d.set(h)):d=h;return this.b=d};
w.prototype.m=function(){var c=0,d=this.b,a=this.g,b,e=new (t?Uint8Array:Array)(this.k+(this.a-32768)),f,g,h,l;if(0===a.length)return t?this.b.subarray(32768,this.a):this.b.slice(32768,this.a);f=0;for(g=a.length;f<g;++f){b=a[f];h=0;for(l=b.length;h<l;++h)e[c++]=b[h]}f=32768;for(g=this.a;f<g;++f)e[c++]=d[f];this.g=[];return this.buffer=e};
w.prototype.r=function(){var c,d=this.a;t?this.p?(c=new Uint8Array(d),c.set(this.b.subarray(0,d))):c=this.b.subarray(0,d):(this.b.length>d&&(this.b.length=d),c=this.b);return this.buffer=c};r("Zlib.RawInflate",w);r("Zlib.RawInflate.prototype.decompress",w.prototype.t);var X={ADAPTIVE:y,BLOCK:A},Y,Z,$,fa;if(Object.keys)Y=Object.keys(X);else for(Z in Y=[],$=0,X)Y[$++]=Z;$=0;for(fa=Y.length;$<fa;++$)Z=Y[$],r("Zlib.RawInflate.BufferType."+Z,X[Z]);}).call(this);`

var opts struct {
	InputFile string `short:"i" long:"input" description:"JS input file"`
}

func main() {
	flags.Parse(&opts)
	if opts.InputFile == "" {
		fmt.Fprintf(os.Stderr, "Input file not specified\n")
		return
	}
	f, err := os.Open(opts.InputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %#v\n", err)
		return
	}
	input, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %#v\n", err)
		return
	}

	buf := &bytes.Buffer{}
	def, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating deflate object: %#v\n", err)
		return
	}
	n, err := def.Write(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compressing data: %#v\n", err)
		return
	}
	if len(input) != n {
		fmt.Fprintf(os.Stderr, "buffer size missmatch")
		return
	}
	err = def.Flush()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error flushing data: %#v\n", err)
		return
	}
	def.Close()
	b64str := base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Fprintf(os.Stdout,
		`/* Compressed with jscomp */

%s

%s

var b64str = "%s";
var inflate = new Zlib.RawInflate(base64js.toByteArray(b64str));
var dec = new TextDecoder("utf-8")
eval(dec.decode(inflate.decompress()));
		 `,
		base64Code,
		inflateCode,
		b64str,
	)

}
