## Compling

You should to have gccgo complier in your system for building this software.
Install it, if gccgo compiler is missing and run the `make` command in your
favorite command line interpreter

## Usage

Note: This software maintained for education purposes only. Author of this software is not liable for any damages, which you cause within it.

This program must be runned from command line and required tho arguments. First is a url of target, which tested. It must contains the `%d` parameter, which will be replaced by random integer number. Second arguments is count of threads for connection setup. 

### Example

    make && ./httpdos 'http://site.org/viewforum.php?f=%d' 300
