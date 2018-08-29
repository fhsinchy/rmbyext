import glob2
import os
import platform
import sys
from tqdm import tqdm

file_count = int()
log = str()
all_files = list()

if len(sys.argv) > 1:
    for ext in sys.argv[1:]:
        all_files = glob2.glob("**/*." + ext)

        if not all_files:
            print("NO {} FILES TO REMOVE.".format(ext.upper()))
            
        else:
            print("Removing: {}".format(ext.upper()))
            for file in tqdm(all_files):
                if os.path.isfile(file):
                    os.remove(file)
                    file_count += 1
                    log +=  "{}\n".format(file)
            print("\n")
    log += "{} FILES DELETED.\n".format(file_count)

    with open("delete_log.log", "a") as log_file:
        log_file.write(log)

    if platform.system() == "Windows":
        os.system("notepad.exe delete_log.log")
    else:
        "LOG: delete_log.log"
