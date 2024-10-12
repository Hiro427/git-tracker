#About 

This is a simple git-tracker it will return a table via the terminal of all your tracked repositories. 

#Usage 

You need to add the REPO_FPATH to you zshrc or bashrc file.

`export REPO_FPATH=/path/to/.txt`

This path should lead to a .txt file which contains a list of repos the paths have to be spelt out entirely. 

Example 

```
/home/user/dir/repo
/home/user/dir/repo
```

You can manually add your repos here, I had a -append-cwd option to add the cwd to that file but a recent update broke this function. I will work on it when I have time.

Run `helper --help` to see options you will see the append option but that isn't working as of 10/12/2024

There is one other option to open the current repo you are in, in the browser on Github.com

