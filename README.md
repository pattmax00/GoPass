# GoPass

An easy to use command-line password generator that creates secure passwords (Version 1.1 adds proper cryptographic
seeding of math/rand)

# Usage:

`./gopass [number of characters]`<br>
or <br>
`./gopass [number of characters] "[excluded characters]"`

<hr>

eg: `./gopass 32` <br>
output: `E$bGOiiPASS*,ISl{!MJ&<\[COOL0eVw` <br>
eg (with excluded characters): `./gopass 32 "$,!"` <br>
output: `EYbGOiiPASS*2ISl{?MJ&<\[COOL0eVw` <- note the excluded characters are not present in the output

# How to install/use for Windows:

1. Download Windows binary, it may be helpful to simply rename it to "gopass.exe" (make sure to add .exe, this won't
   need to
   be typed to execute the file from the CLI, but it is required for Windows to recognize the binary) to keep
   commands shorter.
2. Move the file to your `C:\Users\[Username]` directory
3. Opening the console *without* Administrator privileges will put you in the directory mentioned above by default
4. Now type `gopass [number of characters]` (Ex. `gopass 64`) to generate a new password!

# How to install/use for Linux:

1. Download Linux binary, it may be helpful to simply rename it to "gopass" to keep commands shorter.
2. Move the file to your home directory.
3. `chmod +x gopass` if necessary
3. Opening your terminal should put you in your home directory by default, if not, just type `cd` to get sent back to
   home.
4. Now type `./gopass [number of characters]` (Ex. `./gopass 64`) to generate a new password!