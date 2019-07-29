# Remove by Extension

I often have to remove hundreds of unnecessary files of same extension from my directories. What this script does is, it can take file extensions as command-line argument and remove all files of given extensions recursively.

## Instructions
The generic command for executing this script is

`python rbe.py <file extension>`

or in case of multiple extensions

`python rbe.py <file extension> <file extension>`

Lets say you want to delete all the **PDF** files from a directory. You simply put this script on that directory and execute 

`python rbe.py pdf`

The script will look for all **PDF** files recursively and delete all of them. As stated above you can also pass multiple extensions seperated by space

`python rbe.py pdf txt`

This command will delete all **PDF** and **TXT** files.

List of the deleted files will be shown on the console as well as written inside `delete_log.log` file in the same directory as the script.
