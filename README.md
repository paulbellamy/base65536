# Base65536

Human-memorable encoding for binary data.

Uses the most common (and short) 256 nouns and adjectives,
alternating, to encode binary data in a human-memorable format.

## Example

```
$ echo hello, world! | base65536-encode -comma
important internet, popular idea, various sea, eager community, various
story, popular love, empty dust
```

## TODO

* Allow custom dictionaries
* Decoding
* Experiment with adverbs/etc
* Experiment with uneven slicing
  * different number of bytes for nouns vs adjectives
    * e.g. have 256 adjectives, and 1024 nouns
    * or slice over >2 bytes (2 sets of 12? for 4096 of each?)
