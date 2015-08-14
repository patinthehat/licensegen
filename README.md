## licensegen ##
---

`licensegen` is a small utility for both linux and windows that generates a LICENSE file for new open source projects.  It [comes with](#available-licenses) templates for several of the most popular OSS licenses.

`licensegen` generates both a LICENSE file containing the text of the selected OSS license and a file containing a comment header for source files.

---
#### Compilation/Installation ####
---

  - `$ git clone https://github.com/patinthehat/licensegen.git`
  - `$ cd licensegen`
  - `$ make`
  - [Configuration](#configuration) 
  - [Usage](#usage) / [Usage Examples](#usage-examples) 

---
#### Configuration ####
---

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
You may leave Email Address and/or Website empty and it will not appear in the generated output.

##### Adding a new license #####
To make a new license template available, you must:

  - Create a `NAME.license` file in the `licenses/` directory, containing the contents of the license with variables such as Year replaced with the appropriate go template variable.  _See existing license templates for available template variables._
  - Create a `NAME.header` file in the `licenses/` directory, also with the appropriate template variables.
  - Add a new entry to `licensegen.json`, under "Licenses": 
```
{ "Name": "NAME", "LicenseFile":"licenses/NAME.license", "HeaderFile":"licenses/NAME.header" }
```

---
#### Usage ####
---

Using licensegen is fairly simple.  After compiling, run 
`$ licensegen --list` 
to list which license templates are available.  Next, run 
`$ licensegen --license [license-name]`.

This will generate a `LICENSE` file and a `LICENSE_FILE_HEADER` file into the current directory.

The `LICENSE` file contains the selected license with name, email, website, copyright year, etc. completed according to your `licensegen.json` configuration.

The `LICENSE_FILE_HEADER` file contains a commented header to insert at the top of each file in your project.

---
#### Usage Examples ####
---

  - `$ licensegen --list`
  - `$ licensegen --license MIT`
  - `$ licensegen MIT`
  - `$ licensegen Apache-2.0`

---
#### Available Licenses ####
---

  - Apache 2.0
  - GPL 2.0
  - GPL 3.0
  - MIT
  - MPL 2.0
  - None (simple copyright)

---
#### @todo ####
---

  - ~~complete the README~~
  - ~~add README entry to explain how to add new license templates~~
  - ~~add more license templates~~
  - add interactive mode when started with -i flag

---
#### Licensing ####
---

`licensegen` is open source software, available under the [MIT license](LICENSE).
