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

`helper -list-repos` 

Add Current Working Directory to the txt 

`helper -append-cwd`

Open current directory on Github.com

`helper -open-repo`
