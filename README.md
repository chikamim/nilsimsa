### nilsimsa

Package nilsimsa is a Go implemenation of [Nilsimsa](http://en.wikipedia.org/wiki/Nilsimsa_Hash), ported from
code.google.com/p/py-nilsimsa
but follows the conventions establish by the md5 package in the standard lib

The Java implementaition at
https://github.com/weblyzard/nilsimsa/
blob/master/src/main/java/com/weblyzard/lib/string/nilsimsa/Nilsimsa.java
was also consulted.

There is a discussion about using hash to score string similarities
http://stackoverflow.com/questions/4323977/string-similarity-score-hash


Copyright 2015 Sheng-Te Tsao. All rights reserved.
Use of this source code is governed by the same BSD-style
license that is used by the Go standard library

#### Examples

```
fmt.Println(nilsimsa.HexSum([]byte("Hello nilsimsa!")))
// Output: 436119240183882801210e002e1cb00122a20d11b4268ab8001a51190c08084b

h1 := nilsimsa.HexSum([]byte("Hello nilsimsa!"))
h2 := nilsimsa.HexSum([]byte("Hello world!"))
h3 := nilsimsa.HexSum([]byte("Nobody Expects the Spanish Inquisition"))

fmt.Println(nilsimsa.DiffHexScore(h1, h1))
fmt.Println(nilsimsa.DiffHexScore(h1, h2))
fmt.Println(nilsimsa.DiffHexScore(h1, h3))
// Output:
// 1
// 0.4094488188976378
// 0.14960629921259844
```
