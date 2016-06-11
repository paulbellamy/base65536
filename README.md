# Base65536

Human-memorable encoding for binary data.

Uses the most common (and short) 256 nouns and adjectives,
alternating, to encode binary data in a human-memorable format.

## Example

```
$ echo hello, world! | base65536-encode
visible user exciting communication realistic area significant location
realistic nation exciting marriage successful system
```

## TODO

* Use shorter words! NOTE: This will break backwards compatibility!
* Allow custom dictionaries
* Decoding
* Experiment with adverbs/etc
