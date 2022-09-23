# dul

Deta drive files management through command-line-interface.

## Usage

- Uploading files

  ```
  dul --files=sample.txt --files=another.txt --files=document.pdf --drive=myfiles
  ```

- Showing the files in drive

  ```
  dul list --drive=myfiles
  ```

- Deleting / removing files

  ```
  dul delete sample.txt another.txt --drive=myfiles
  ```

- Downloading a file
  ```
  dul get --file document.pdf --drive=myfiles
  ```

## Global Flags and Config

The following flags are required in every usage of the commands from above.

- `--project-key` - your Deta project key
- `--drive` - name of the Deta drive within the project

These can be set in a global config for the cli app to use.

    ```
    dul set --projectKey=your-deta-project-key --driveName=your-drive-name
    ```

##

**TheBoringDude &copy; 2022 | [LICENSE](./LICENSE)**
