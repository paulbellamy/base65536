# Base65536

Human-friendly encoding for binary data.

Uses the most common (and short) 256 nouns and adjectives,
alternating, to encode binary data in a human-friendly format.

## Example

```
$ echo hello, world! | base65536-encode -comma
important internet, popular idea, various sea, eager community, various
story, popular love, empty dust
```

## But, why?

Sometimes it's nice to have a friendly way of transferring, and
remembering small bits of binary data. For example, the bytes
representing `127.0.0.1` encode to: `traditional bird able blossom`

## TODO

* Allow custom dictionaries
* Experiment with adverbs/etc
* Experiment with uneven slicing
  * different number of bytes for nouns vs adjectives
    * e.g. have 256 adjectives, and 1024 nouns
    * or slice over >2 bytes (2 sets of 12? for 4096 of each?)
