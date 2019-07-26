# Remove by Extension

I often have to remove hundreds of unnecessary files of same extension from my directories. What this script does is, it can take file extensions as command-line argument and remove all files of given extensions recursively.

## Instructions
This script depends on two external modules:

* glob2 [`pip install glob2`]
* tqdm [`pip install tqdm`]

Now, lets say you want to delete all the **PDF** files from **D:\** drive of your computer. You simply put this script on the root of the drive and execute 

`python rbe.py pdf`

The script will go through all your folders and delete all the pdfs from your drive. You can also pass multiple extensions seperated by space

`python rbe.py pdf txt`

This command will delete all **PDF** and **TXT** file.

List of the deleted files are written inside `delete_log.log` file.
