AlphaNum
========

Simply Go library and command line app to convert 1s based indexes (like you'd use in a program or script for CSV processing) to spreadsheet column letters. Or from spreadsheet type column letters to a 1s based index.


Thoughts on the Future
----------------------

Features I'm considering for the next version

* [ ] Allow lowercase letter workflow
* [ ] Add options for JSON output
* [ ] Add options for POSIX style tabbed output


Prior Art
---------

The code was partially based on code by Robin Houston's Python [Gist](https://gist.github.com/robinhouston/99746cac543e6b3ea61f1d245e9b19cc). The gist was the only example of converting column letters to indexes and back again that was _not_ directly tied to a spreadsheet in Excel or Google Sheets. Go does not have a reduce function as they tend to be very slow and can typically be replaced with a thoughtful for loop. So I figured out my for loop based on Robin's code implemented in my Python example in the `extras` directory.
