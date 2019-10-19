# Remove by Extension

I often have to remove hundreds of unnecessary files of same extension from my directories. What this program does is, it can take file extensions as command-line argument and remove all files of given extensions recursively.

## Instructions
If you have `$GOPATH` setup on your computer then just clone the repo in your `src` folder. Build the code with `go build` or to run the tests execute `go test`. You can use the binary to remove files of given extension. The generic command would be

`rbe <file extension>`

or in case of multiple extensions

`rbe <file extension> <file extension>`

Lets say you want to delete all the **PDF** files from a directory. You simply put the binary on that directory and execute 

`rbe pdf`

The program will look for all **PDF** files recursively and delete all of them. As stated above you can also pass multiple extensions seperated by space

`rbe pdf txt`

This command will delete all **PDF** and **TXT** files.

List of the deleted files will be shown on the console as well as written inside `delete_log.log` file in the same directory as the binary.
