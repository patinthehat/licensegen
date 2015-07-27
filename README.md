### licensegen ###
----

`licensegen` is a small utility for both linux and windows that generates a LICENSE file for new projects.  

It outputs both a LICENSE file and a file containing the header for new files.

----
#### Compilation

  - `git clone <url>`
  - `go build licensegen.go`
  - This generates the `licensegen` or `licensegen.exe` file.

----

#### Configuration

After cloning the repository, you must create a configuration file for `licensegen` to read settings from, i.e. you name or email address.

Here is a default `licensegen.json`:

```
{ 
   "Author": {
    "FirstName":"firstname",
    "LastName":"lastname",
    "EmailAddress":"user@example.com",
    "Website":"http://www.example.com"
   },

  -- snip --

```

Modify the settings to match your personal details.
You may leave Email Address and/or Website empty and will not appear in the generated output.

----

#### Usage

Using licensegen is fairly simple.  After compiling, move the generated binary to somewhere in your $PATH directories.
After moving it, run `licensegen` in the terminal window.

_example usage:_
  `licensegen --license MIT`
	_- OR -_
  `licensegen MIT`

  This will generate a LICENSE file and a LICENSE_FILE_HEADER.txt to the current directory.

  Place the contents of the LICENSE_FILE_HEADER.txt file as a commented header for each file in your project.


----

#### @todo

	[ ] complete `README.md`
	[ ] add README entry to explain how to add new license templates
	[ ] add more license templates
	[ ] add interactive mode when started with -i flag

----
