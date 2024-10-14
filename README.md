#About 

This is a simple git-tracker it will return a table via the terminal of all your tracked repositories. 

#Usage 

#### Setup

You need to add the REPO_FPATH to you zshrc or bashrc file.

`export REPO_FPATH=/path/to/.txt`

This path should lead to a .txt file which contains a list of repos the paths have to be spelt out entirely. 

Example 

```
/home/user/dir/repo
/home/user/dir/repo
```
#### Commands

List status of all repos in the txt 

`helper -list` 

Add Current Working Directory to the txt 

`helper -track`

Open current directory on Github.com

`helper -open`

Update current repo: Add all Files, Commit files with message passed to flag, and pushes to main branch

`helper -sync "commit msg"`

These commands can also be accessed by `helper --help`
