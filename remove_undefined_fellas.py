import os

# Directory containing the files
directory = "/Users/bsoghigian/dev/leetcode_dump/leetcode"

# Iterate through the files in the directory
for filename in os.listdir(directory):
  # Check if the file ends with "undefined"
  if filename.endswith("undefined"):
    # Replace "undefined" with ".go" in the file name
    new_filename = filename.replace("undefined", ".go")
    # Rename the file
    os.rename(os.path.join(directory, filename), os.path.join(directory, new_filename))

