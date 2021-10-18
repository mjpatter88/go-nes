[![Go](https://github.com/mjpatter88/go-nes/actions/workflows/go.yml/badge.svg)](https://github.com/mjpatter88/go-nes/actions/workflows/go.yml)

# go-nes
An NES emulator written in go

Following https://bugzmanov.github.io/nes_ebook/ but implementing in go.

https://bugzmanov.github.io/nes_ebook/chapter_3_3.html staes that the Ricoh modificaiton of the 6502 chip,
which was the one used in the NES, removed support for Decimal mode, so it is not necessary to support this
in an emulator.


### Comparison Instructions
I've seen conflicting information, but I believe these insturctions are unsigned comparisons.
The most detailed resource seems to be http://6502.org/tutorials/compare_beyond.html which suggests
that the comparison is unsigned but also describes how to achieve a signed comparison.



### Other Resources
* https://skilldrick.github.io/easy6502/
* http://www.6502.org/tutorials/6502opcodes.html
* http://6502.org/tutorials/compare_beyond.html
