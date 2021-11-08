<!--
ignore these words in spell check for this file
// cSpell:ignore ostringstream Downey
-->

[Main](README.md)

Text

## A Crash Course in Unicode for C++ Developers - Steve Downey

<details>
<summary>
Different types of encodings, encoder/decoder (compose / decompose) algorithms. unicode in c++.
</summary>

[A Crash Course in Unicode for C++ Developers](https://youtu.be/iQWtiYNK3kQ)
[unicode](http://unicode.org)
[utf8 encoding](https://en.wikipedia.org/wiki/UTF-8)

std::u8string

code units, code point, graphemes, abstract characters.

> code units
>
> - char
> - wchar_t
> - octet
> - Word

code points and scalar values
grapheme clusters, extended grapheme clusters.

> utf-8 is good
>
> - C string safe
> - No aliasing
> - Self syncing
> - Single errors lost one character
> - ASCII compatible
> - Start is easy to find

a table about how we encode different ranges of values into different bytes.
some 'well formed' byte sequences.

utf-16 and utf-32. if the value fits inside 16 bit, then put it in one, otherwise, split it into two code points (surrogate pairs).
ucs-2, ucs-4. wtf-8 (wobbly transformation format), wtf-16

### Encoding and Decoding

> Encoders take text and output octets.
> Decoders take octets and output text.
> Text is this context is scalar values.

in utf8 the order is set.
in utf-16 there are byte order marks, for big and little endian.

legacy encoding from before unicode

> - Windows 1252, 125x
> - ISO-8859-x
>   ...others

multi-byte encodings

transcoding, from one character set to another.

### Normalization

combined text might have more than one representation, like special forms, canonical equivalence and compatible equivalence

canonical equivalence:

three difference ways to produce the same symbol.  
latin capital with a rings &#x00c5; is like angstrom sigh &#x212b; or combining the letter with the symbol &#x0041;&#x030A;

compatible equivalence:
not the same symbol, losing some data, but meaning is preserved (mostly)

decomposed and composed text. there are some forms of charters that are already predefined, but can also be created by composing different code points together,we have

NFD,NFD,NFKD,NFKC are forms for different usages, like search (human), identifiers in linkage (strong equality), decomposing ignored diacritics.

nfc is the least risky in terms of information lost.

quick_check of code points to test if it's normalized: yes, no and maybe.

stream safe text format, a way to avoid some problems that can occur in full normalization, so there's a stream-safe format.

the unicode character database. UCD files (txt files and xml files) - all sorts of data.

theres an issue with emojis.

### Algorithms

there are still many problems in the standard solutions, different text directions. word wrapping (line breaks): positions where is's possible to break between lines, there are many ways to get this wrong. text segmentation: from data into user perceived characters, words and sentences.

unicode and regular expressions. matching words on word boundaries. sentences embedded within sentences.

### The Future For C++

what might be in the future, c++23 and c++26 are what they hope to achieve.

| version | features                   |
| ------- | -------------------------- |
| C++20   | char8_t                    |
| C++23   | literal Encoding           |
|         | Portable Source Code       |
|         | Encoding / Decoding Ranges |
| C++26   | Algorithms with ranges     |

</details>

[Main](README.md)
