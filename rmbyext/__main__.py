import sys
from pathlib import Path

def main():
    deleted_file_count = int()
    log = str()

    if len(sys.argv) > 1: # checking if user has supplied any file extension as parameters or not.
        p = Path(".")
        for ext in sys.argv[1:]:
            file_list = list(p.rglob("*." + ext)) # creating list of files for each extension supplied by the user.

            if not file_list:
                print(f"NO {ext.upper()} FILES TO REMOVE.")
                
            else:
                print(f"Removing: {ext.upper()}")
                for file in file_list:
                    f = Path(file)
                    if f.is_file():
                        f.unlink() # removing the file.
                        deleted_file_count += 1 # increase deleted files count.
                        print(file) # print deleted filename to the console.
                        log +=  f"{file}\n" # appends deleted filename to the log string.
                print("\n", end="")
        log += f"{deleted_file_count} FILES DELETED.\n"

        if deleted_file_count:
            with open("delete_log.log", "a") as log_file:
                try:
                    log_file.write(log) # write the entire log string to a file.
                except UnicodeEncodeError as err:
                    print(f"Filename contains non-standard character. ({err})")
    else:
        print("NO ARGUMENT PASSED.")

if __name__ == "__main__":
    main()