<!--
// cSpell:ignore Schwarzmüller mdkir
 -->

# Learn Git & GitHub and master working with commits, branches, the stash, cherry picking, rebasing, pull requests & more!

udemy course [Learn Git & GitHub and master working with commits, branches, the stash, cherry picking, rebasing, pull requests & more!](https://www.udemy.com/course/git-github-practical-guide/) by *Manuel Lorenz* and *Maximilian Schwarzmüller*. 


1. [Introduction](#section-1-introduction)
1. [Mac Terminal & Windows Prompt Introduction](#section-2---mac-terminal--windows-prompt-introduction)
1. [Version Management with Git - The Basics](#section-3---version-management-with-git---the-basics)
1. [Diving Deeper Into Git](#section-4---diving-deeper-into-git)
1. [From Local To Remote - Understanding GitHub](#section-5---from-local-to-remote---understanding-github)
1. [GitHub Deep Dive - Collaboration & Contribution](#section-6---github-deep-dive---collaboration--contribution)
1. [Real Project Example: Git & GitHub Applied](#section-7---real-project-example-git--github-applied)

## Section 1: Introduction
<details>
<summary>
Course Introduction
</summary>

Git is a version management tool, github is way to use git on the cloud.

### What is Git?
in the official [git website](git-scm.com) we can see the declaration
> Git is a free and open source distributed version control system designed to handle everything from small to very large projects with speed and efficiency.

managing different version of code/documents.

a naive approach is having multiple files with suffixes like "first_draft","another_draft", "_final", "_final2", and "final_for_real!", and in websites and codes, we can't use this approach easily, because files need to reference each other.

version management tools allow us to have a single file, and still be able to go back in time to earlier versions.

**Control and tracking of code changes over time**

git is a local tool, it can be used on a single machine, and if we use it only on our machine, then it's still suspectable to computer crashes and it isn't the solution for managing project with multiple contributors.

### What is GitHub?

[GitHub website](github.com), a Git Repository Hosting Service.

The largest development platform in the world, a cloud hosting and collabaration tools, it's free for basic usage, and provides paid features for large companies.

it allows different users access to a shared codebase. with all the benefits of version control, and with added features.

### How to Get the Most out of this Course!
### Course Slides

</details>


## Section 2 - Mac Terminal & Windows Prompt Introduction
<details>
<summary>
Optional module on using CLI - shell and command prompts
</summary>


### The Command Line - What & Why?

today, we mostly use the graphic user interface, we interact with the mouse, menus, etc. it's very user friendly and easy to understand. it's also easy to explore.

the alternative is using text-based tools (command line interface), such as bash shell or windows command prompt.

it gives the user more power, like starting servers, install tools, downloading files, run code. and in our case, to interact with git.

### Comparing the Mac & Windows Command Line

Mac terminology - the console is called 'the terminal',the shell is the cli itself, its a software that works with text input. in Mac we have the Bash shell, or the z-shell.

in windows, we have the command prompt, as the basic initial windows shell. there is also a newer version, powershell. we can also use the git-bash cli emulator, which allows us to have a unix-like experience.


### Mac Terminal
<details>
<summary>
Using the z-shell terminal in Mac.
</summary>

in mac computer, we access the **z-shell** terminal (abbreviated as **zsh**). The terminal start at the user folder - the home directory, marked with the Tilde symbol (<kbd>~</kbd>) . `pwd` shows us where we are. to see the items in the folder, we write `ls`. we change directories with the `cd` command,  

#### Accessing Folders
we can press <kbd>TAB</kbd> to get auto-completion. if we have a space in folder name, we must escape it with **\\** (backslash symbol).

#### Absolute vs Relative Paths
paths can be relative to the current working directory, or absolute (starting from the root directory).the `pwd` command prints the absolute path.

#### Creating & Deleting Files

we create files with `touch` and delete them with `rm`. removing folder is done with `rmdir`, but it requires the folder to be empty.

#### Introducing Flags and Removing Data

flags allow us to add more options to a command, like `ls -s` to see file sizes, or `ls -l` to see the long format. we can combine flags `ls -ls`. we can check the manual of each command by typing `man <command>`.\
if we want to remove a none-empty directory, we can use the `rm -r` flag, which removes the folder and the internal files. (r stand for recursive).

#### Copying & Moving Files & Folders
we can copy files and folder with the `cp` and `mv` commands, we use the `mv` to rename files. the commands take the source path and the destination. the paths can be relative or absolute.

if we want to copy multiple files in a folder, we can use `cp -r source target`, we might need to add a slash `/` if we don't want to create a new folder.

#### Mac Terminal - Core Commands Overview
- `pwd` - print working directory
- `ls` - list files
- `cd` - change directory
  - `cd ..` - go one level up
  - `cd ~` - home directory
  - `cd /` - root directort
  - `cd /Users` - users folder
- `clear` - clear screen
- `touch` - create file
- `mkdir` - create folder
- `rm` - delete file entirely
- `rmdir` - remove empty folder
- `cp` - copy files and folder
- `mv` - move / rename files and folders

</details>

### Windows Command Prompt
<details>
<summary>
Windows Command Prompt.
</summary>

we use `command prompt` as the windows cli tool

- home directory - c:/users/\<username\>
- users directory - c:/users
- root directory - c:/

```sh
dir  #list files
cd .. #move one folder "up"
d: # move to drive D
```

#### Absolute vs Relative Paths

we have relative and abasolute paths. relative path start from the current directory, and we navigate based on the relation from the current location. abasolute path always use the complete path, the abasolute path starts from the root folder (the drive), and the same absolutes path will always bring us to where we want to be,no matter where we are.

```sh
cd / #root
cd ~ #home
```

relative paths are easier to navigate, while absolute paths help us to be consistent.

#### Creating & Deleting Files & Folders

```sh
mkdir folder1 # create new folder
echo our first file > text.txt #create new file
type test.txt #see contents of file
del test.txt #delete file
rmdir folder1 #remove folder
```

#### Copying & Moving Files

```sh
copy test2.txt folder #copy file into folder
move test3.txt folder #move file into folder
```

#### Windows Command Prompt - Core Commands Overview

- `cd` - print current directory
- `cd ..` - move on level up
- `dir` - list files
- `cls` - clear screen
- `echo` - 
- `mkdir` - create fp;der
- `del` - delete file
- `rmdir` - remove directory
- `copy` - copy file/folder
- `move` - move file/folder

</details>


 </details>


## Section 3 - Version Management with Git - The Basics
<details>
<summary>
Understanding Version management and Git.
</summary>

in this module, we will look at the theory behind version control and git. we will also learn to install git and to setup the development environment, and lastly, the basic features of git.

### Git Theory
<details>
<summary>
Git theory - terminology: Commits, Repository, Branches.
</summary>

#### How Git Works
all the files managed by git in a folder are refereed to as "the working directory". we use **commits** to create snapshots of how the working directory looks at a certain time.

when we change the project, we create a new commit, which stores all the differences between the current change and the previous state. each commit is simply the tracking of the changes done over time.

the commits are stored in a **branch**, the initial branch is the "master/main" branch.

#### Working Directory vs Repository

when we start managing a folder with git, a hidden folder, named ".git", which stores the repository data. it has two areas:
- Staging Area - index File
- Commits - Objects folder

the flow to change code is to first move the files into the staging area, and then we create a commit object.

working directory isn't the same as a Repository.

git doesn't store each version of the file, it has the initial file, and then stores the changes done in each commit.

#### Understanding Branches

the working director is our project file, the changes and tracking is done in the git repository folder.

a branch is a way to have parallel development path, like having a copy of the project which we use to create a new feature. a branch is a way to have a copy of the main branch, and we can work on it and add commits to it. and we can then **merge** these changes back to the main branch.


</details>

### Installing Git and Vscode
<details>
<summary>
Getting Git and vscode
</summary>

#### Installing Git on Windows and MacOs

in windows, we download git from the official website, and run the installer. we can select what features we want, we can decide what the default branch name is ("master", "main", etc... ). we can also decide how git is used from the command line, security, style, default behavior of some commands, credentials managements, caching, experimental options, and so on.

in Macos, we install git by using the package management, like homebrew.

we get homebrew (if we don't have it), and run it from the console, and then we use homebrew to install git.

`brew install git`

#### Installing Visual Studio Code

we get [VScode](https://code.visualstudio.com/) as an IDE (integrated development environment). we select th

</details>

### Initializing the Repository & Creating the First Commit (`git init` & `git commit`)

in our empty folder, we add new files, such as "initial-commit.txt".

we can run the `git status` command, and then we'll get an error, because we still didn't initialize the working directory and the the repository. we can fix this by running `git init` and now we check the status. we will see that we have untracked files, so we add them by using `git add`, followed either by the file name or a dot symbol. now the status shows that we have a file in the staging area. to add the file as a commit, we need to run `git commit` and use the `-m` flag to write a message.

in a new project, we might get an error, saying that our user name and email isn't set.

we can set this data globally, which will store this for all repositories on this machine, or omit the `--global` flag to store it locally in this repository.
```sh
git config --global user.email "user.email.com"
git config --global user.name "user name"
```

### Diving Deeper Into Commits with `git log`

we can see the history of commits by running `git log`, we see the author, the date, the comment, and the unique commit id. 
if we add more commits, we can see the earlier versions, so we can retrun to a previous state by running `git checkout <commit-id>`. now we are transported back to the snapshot we chose.

we can return to the main branch by running `git checkout main`.

### Understanding & Creating Branches

we can a have copy of the working directory by creating a branch. we can see the branches by running `git branch`. if we want to create a new branch, we run `git branch` with the new branch name. the branch name cannot contain spaces. to move between branches, we run `git checkout <branch name>`.

the new branch is identical to the main branch, it has the same history.

a different way to create branches is by running `git checkout -b <new branch name>`. if we add commits to the branch, they will only show for this branch, and not for the other branches, this allows us to work in parallel on different features, or have different people work at the same time without effecting one another.

when we want to bring changes from a side branch into the main branch, we use something called **merging**.

### Merging Branches - The Basics

to merge changes from one branch to another, we first move to the target branch (usually main), and we run `git merge <other branch name>` to pull in the chanages.

### Understanding the HEAD

in the log output, we can sometimes see the **HEAD** commit, this means the latest commit at the branch. when we switch between branches, we automatically take the latest commit, the HEAD.

each branch has it's own HEAD, which can be different from another branche HEAD.

### The "detached HEAD"

a **Detached HEAD** is what happens when we checkout a commit from a which isn't the latest commit. it's not the latest Head for any branch. when we are at a detached-head, we don't belong to any branch, when we view the branches, we see that we are currently not in any branch.

### Branches & `git switch` (Git 2.23)

a new command is `git switch`, which works solely for branches, unlike `git checkout` (which work for both commits and branches). 

to move between branches we run `git switch branch-name`, and we create new branches with `git switch -c branch-name`.

### Deleting Data
<details>
<summary>
Deleting data from version management.
</summary>

different types of 'deletion', removing date from version control.

#### Deleting Working Directory Files

if we have a file which we want to remove, we can delete the file, add this staged 'deletion' to the staging area and commit it.

we can also use `git rm` to move the deleted file to the staged area, and then commit it.

#### Undoing Un-staged Changes
if we modify a tracked file, we can resore it to the state it was, if we haven't staged it yet (we didn't `git add` it), we can run `git checkout -- <file>` and restore a file, or run `git checkout -- .` to restore all files.

another option is to use a specific command `git restore <file>` .

if we have new file (untracked), we can run `git clean` to remove files, we can use the `-d`,`-n`,and -`f` flags.

#### Undoing Staged Changes

if we staged a file and we want to undo those changes, we can use the new command `git restore --staged <file>`, but in the past we had to use two commands. simply checking out the file wouldn't work. we first need to run `git reset <file>`  which copies the latest committed changes, into the staging area, and then run `git checkout`

#### Deleting Commits with `git reset`

`git reset` also allows us to reset the heads of our branches, thereby undoing commits. we can run `git reset --soft HEAD~1` to go back one step and delete the commit, but not the data. the default behavior also removes the changes from the stagin area.

using the `--hard` flag removes the commit, from the staging area and from the tracked files.

#### Deleting Branches

we can remove branches with `git branch` and adding either `-d` or `-D` and the branch name, the `-D` option allows us to delete branches even if they weren't merged into the main branch. we can remove multiple branches at once by passing more than one branch name.


</details>

### Committing "detached HEAD" Changes

if we checkout a commit from an earlier stage of the branch, we might want to make stages there.

when we make changes from a detached mode, and we commit it, we have a detached commit, and we can lose it. if we move to another branch, we will see a warning about having a detached head commit floating without branches.

To preserve it, we need to create a branch from this commit. `git branch new-branch-name <commit id>`, we can now merge this branch into the main branch if needed.

### Understanding .gitignore

some files shouldn't be shared and tracked, such as log files, or specific IDE configurations. we can control this by using the ".gitignore" file, which specifies which files shouldn't be tracked by git.

we do need to add this file and track it, of course.

if we want to ignore files we can add the complete path to the .gitignore file, or use wildcard `*` as a pattern, and then override the rule by with `!`. we can ignore folders completely

```
test.log
*.log
!important.log
ignoredFolder/*
```

there is a way to create a global ".gitignore" file.

### Assignment 1 - Practicing the Git Basics

> Git Basics Assignment - Your Instructions
> 
> 1. Create a new folder and initialize the repository
> 2. Paste the "instructions.txt" file into this folder
> 3. Add a .txt file named "file-1" containing any text of your choice to the working directory
> 4. Create a second .txt file named "file-2"
> 5. Add "file-1" and "file-2" to the staging area - don't add "instructions.txt"
> 6. Change the initial text you added to "file-1"
> 7. Now add all working directory files to the staging area
> 8. Create the first commit
> 9. Create a second branch named "feature" (two commands are possible)
> 10. Add a third .txt file ("file-3.txt") to this branch
> 11. Create a new commit
> 12. Add the following text to "file-3": "I will be deleted"
> 13. Add the updated file to the staging area
> 14. Undo the staged change
> 15. Add the following text: "Please add me to the master/main branch"
> 16. Commit this latest change
> 17. Merge the "master" (or "main") branch with "feature"
> 18. Delete the "feature" branch
> 
> And most importantly: Have fun with the assignment :)

my solution:
```sh
#1
mdkir assignment
cd assignment
git init
#2
cp ./instructions.txt ./
#3
echo "text" > file1.txt
#4
echo "other text"> file2.txt
#5
git add file1.txt file2.txt
#6
echo "different text" > file1.txt
#7
git add .
#8
git commit -m "first commit"
#9
git switch -c feature
#10
echo "feature" > file3.txt
#11
git add file3.txt
git commit -m "feature commit"
#12
echo "I'll be deleted" > file3.txt
#13
git add file3.txt
#14
git restore --staged file3.txt
#15
echo "please add me to main branch" > file2.txt
#16
git add file2.txt
git commit -m "from branch"
#17
git switch master
git merge feature
#18
git branch -d feature
```

I forgot to run `git checkout` after restoring file from staged area. i simply got the file out of the staged area, but i didn't reset the contents.

### Useful Resources & Links

</details>


## Section 4 - Diving Deeper Into Git
<details>
<summary>
managing different branches, different kinds of merging branches. stashing and retriving deleted data.
</summary>

Diving Deeper into commit, managing and combining different branches, resolving merge conflicts.

### Understanding the Stash (`git stash`)

the `git stash` command is a way to preserve progress without a commit. the stash is an internal memory that holds uncomitted / upstaged changes.

running `git stash` takes all of our changes and stashes them away. `git stash apply` retrives the changes.

each call to `git stash` creates a stash, which we can see with `git stash list`, we can get specific version by the index `git stash apply 1`.

if we want to better track our stash, we can use the full command `git stash push -m "msg"` to store a message connected to this stash.

`git stash pop` takes the changes off the stash and into the project. it's similar to `git stash apply`, but it also removes the stash from the store.

to remove a stash without applying it, we can use `git stash drop` for one entry, or `git stash clear` for all entries.

### Bringing Lost Data Back with `git reflog`

if we deleted something, like a branch, then we can also get them back with `git reflog`.

if we we use `git reset --hard HEAD~1`, we move the head back, but if we use `git reflog`, we can see a log of our past actions, with this we can grab lost commits and reset back to them.

this can also help us when we delete branches,
we can see the lost commits on the deleted branch. we need to create a new branch for the commit to exists on.

```sh
git checkout <commit>
# we are now on a detached head
git checkout -b new-feature <commit>
```

### Combining Branches - What & Why?

we usually have a main branch (master, development, trunk, etc), and we create feature branches based on it. 

we sometimes need to combine the branches, we might need to get the latest changes from the master branch into the feature, or bring our changes from the feature branch into the master branch by merging.

### Understanding Merge Types

there are two types of merging branches
- **fast-foreward**
- non fast-forward:
  - **recursive** (this is the common one)
  - ours
  - octopus
  - subtree


when we have a main branch and feature branch, if we just worked on the feature branch and the main branch stayed the same, then we can use the *fast-forward* merge. the master HEAD is set to the HEAD of the feature branch, and no new commits are created.

### Applying the Fast-Forward Merge

lets start in a new environment and create a fast-forward merge. as long as the target branch doesn't change, we can use fast-forward merging, without creating new commits.

```sh
git init

#working on main branch
mkdir master
echo "first" > master/m1.txt
git add .
git commit -m "m1"
echo "second" > master/m2.txt
git add .
git commit -m "m2"

# work on feature branch
git switch -c feature
mkdir feature
echo "feature" > feature/f1.txt
git add .
git commit -m "f1"
echo "feature-next" > feature/f2.txt
git add .
git commit -m "f2"

# merge
git branch -v
git switch master

git merge feature
git log

# undo merging
git reset --hard HEAD~2
git log
git switch feature
git log
```
we might want to have single commit of all the changes from the feature branch, rather than carry around individual commits in the main branch. this is done with the `--squash` flag. this means a new commit.

```sh
git switch master
git merge --squash
git status
git add .
git commit -m "squashed"
git log
```

lets go back
```sh
git reset --hard HEAD~1
```

### The Recursive Merge (Non-Fast-Forward)

back in our main branch, we can force git to use a non-fast-forward merge with the `--no-ff` flag. this is the recursive strategy. it creates a new commit about the merge.

```sh
git merge --no-ff feature
git log
```

if we have two branches, and this time, the master branch has also changed. so we can't do a regular fast-forward merge. 

when we reset, we reset to remove the merge commit. we don't care about all the other commits from the feature branch.
```sh
git log
git reset --hard HEAD~1

# change main branch
echo "move main" > ./master/m3.txt
git add .
git commit -m "m3"

# merge
git merge feature

# reset
git reset --hard HEAD~1

# merge squash
git merge --squash feature
git add .
git commit -m "master and feature merged"
```

### Rebasing - Theory

`git rebase` is a way to add the commits at a diffrent location. we make the new HEAD commit become the base commit of the feature commits.

rebase doesn't move commits, it **recreates** them, and it's dangerous to use in a shared project. 

the rebased commits will have different ids.
```sh
#restore
git log
git reset --hard HEAD~1

# in the feature branch
git switch feature
git log # remember the ids
git rebase maser
git log #  ids are different
```

in large projects, rebasing can mess up the history, so it might not be a good idea.

```sh
git switch master
git merge feature #fast forward merge
```

> - New commits in master branch while working in feature branch
> - feature relies on additional commits in master branch. rebase master into feature branch.
> - feature is finished - implementation into master without merge commit. merge master into feature + fast forward merge feature into master.

### Handling Merge Conflicts

sometimes a merge fails, this happens when there are conflicts between branches. sometimes different branches change the same file

```sh
# in master
echo "from master" > feature/f1.txt
git add .
git commit -m "f1 from master"

# in feature
git switch feature
echo "from feature!" > feature/f1.txt
git add .
git commit -m "f1 from master"

# merge
git switch master
git merge feature #merge conflict
```

now the merge fails, and we need to fix it. vscode gives us visual interface to see the differences.

we can see the data in `git status`, use `git log --merge`, `git diff` or abort the merge with `git merge --abort`.

if we fix the conflict, we need to commit the changes.

### Merge vs Rebase vs Cherry Pick

> - merge (no fast-forward) - create marge commit - new commit.
> - rebase - change single commit's parent - new commit IDs.
> - cherry pick - add a specific commit to branch - copies commit with new ID.


sometimes we want just one commit from another branch, without taking (merging) the entire branch.
```sh
# in master
git switch master
echo "typo" > master/m1.txt
git add .
git commit -m "with typo"

#
git switch -c feature2
mkdir feature2
echo "new feature2" > feature2/f-new-1.txt
git add .
git commit -m "f-new-1"
echo "fix typo!" > master/m1.txt
git add .
git commit -m "type fix in m1"
echo "new feature2-2" > feature2/f-new-2.txt
git add .
git commit -m "f-new-2"

# in master
git switch main
git cherry-pick <commit-id--branch>
git log
```

the cherry-picked commit will have a different id.

### Working with Tags (`git tag`)
`git tag` allows us to create tagged commit, a tag is a label, like a milestone of a project.

```sh
git init
git echo "a" >a1.txt
git add .
git commit -m "a1"
git echo "b" >a2.txt
git commit -am "a2"
git echo "c" >a3.txt
git commit -am "a3"

git tag #show tags
git tag tag-name <commit-id> # light weight tag
git tag #show tag
git show tag-name
git checkout tag-name #checkout commit by tag
git checkout master
git tag -d tag-name #remove tag
git tag -a 2.0 -m "latest version" #annotated tag
git show 2.0
```

there are lightweigh tags and annotated tags. an annotated tag is a real object, so it holds data about who created it.

### Useful Resources & Links


</details>

## Section 5 - From Local To Remote - Understanding GitHub
<details>
<summary>
Basic ways to work with github.
</summary>

Leaving the local git environment and moving to the cloud on github. GitHub is a repository hosting service.


### From Local to Remote Repository - Theory

we have an existing git repository on the local machine, and we want to move it to github.

we need to establish a connection betwee the local repository to the remote one.

`git remote add origin <url>` - origin is how we refer to the remote repository, it's an alias to the url. the url is the address of the remote repository.

we then push our local repository onto the remote by calling `git push`, and we get the data from the remote repository with `git pull`.

### Creating a GitHub Account & Introducing GitHub

we go to the github website, set up an account (use the free plan), we can create a new repository or import them from another provider (like gitlab), there all kinds of options.

### Creating a Remote Repository

in the github page, we click <kbd>Create Repository</kbd> or in the repositories page we can click <kbd>New</kbd>.

we can choose the owner of the repository, the name, a description, set the access level (public/ private), and initialize the repositrory with **README** file, a **.gitignore** file and a license file.

once we create the repository, we get some options of how to connect to it.

### Connecting Local & Remote Repositories

since we have an existing repository, we can push it from them the local machine.


```sh
git init
echo "hello world" > m1.txt
git add .
git commit -m "first commit!"

git remote add origin <address>
git branch -m main #rename branch to main
git push -u origin main #push local to remote
```

this doesn't work yet, because we aren't identified as our github user. we get a pop-up to sign in into github (which won't be supported in the future), or use a personal access token.

### Understanding the Personal Access Token

in github web page, click the profile and choose <kbd>settings</kbd>, then <kbd>developer settings</kbd> and we select <kbd>personal access tokens</kbd>, we <kbd>Generate new token</kbd>, assign permissions, give it an experssion time, and copy the created token. and store it somewhere safe.

we also fill it in the the popup, so now we have connected our vscode ide to github.

### Pushing a Second Commit

adding another file, just like the first, `git add`, `git commit` and `git push`. git push doesn't work if we don't set the upstream branch, so for now we still need to use `git push origin/master`.

in github we can see the commit history.

our credentials (the access token) is stored in **windows credentials manager**.

### From Local to Remote - Understanding the Workflow

<details>
<summary>
All kinds of branches, local, remote and tracking.
</summary>

running `git branch -a` shows us new data. it shows us the **remote tracking branch**, a local copy that connects the local and remote branch.

this also is used when we 'pull' from the remote branch, a tracking branch is created, and it is merged into the local branch.

#### Remote Tracking Branches in Practice

we create a new local branch, make changes and push it to github, now we have two remote tracking branches.

but we can also create a branch on github, we don't see it in our branches, but we can still list it.

`git ls-remote`

to get this branch, we can run `git fetch origin` to grab everything into the remote tracking branches. to merge this, we can run `git pull origin master` (nothing changed).

remote tracking branches are read-only.

#### Understanding Local Tracking Branches

- local branch
- local tracking branch - local reference to a remote tracking branch (git pull, push)
- remote tracking branch - local copy of the remote (git fetch)
- remote branch


local tracking branch are editable - the pull/push operations act on this branch. if we have a local tracking branch configured, we don't need to specify the `origin master` each time.


#### Creating Local Tracking Branches

to create a local tracking branch:
`git branch --track feature-remote-local origin/feature-remote`  (this isn't actually what we want).

the names need to match.

#### Remote & Tracking Branches - Command Overview

`git remote` shows the curret remote services, `git remote show origin` gives us more details.

</details>

### Cloning a Remote Repository

to get an existing remote repository, we can run `git clone`. we get the URL from github. we don't need to run `git init` before. we only get the master branch locally, for the rest of the branches, we have remote tracking branches, but not local. if we want to work on one of them, we can create a local tracking branch.

if we don't specify which branch to track, the default behavior is to track the master branch.

### Understanding the Upstream

the `-u` flag for git push creates an upstream, which is a local tracking branch, it's a bit easier to create and manage.

### Deleting Remote Branches & Public Commits

it's easy to delete local branches. to delete remote branches we add the `--remote` flag to the delete command, we just delete a remote tracking branch.


`git push origin --delete branch_name` - to delete a branch from github.

undoing commits is done by reseting the head, `git reset --hard HEAD~1` and then push, which fails initially, but we can add the `--force` flag to force the push.

</details>

## Section 6 - GitHub Deep Dive - Collaboration & Contribution
<details>
<summary>
Using Github in Practice. 
</summary>


Using Github in collaboration with others, exchanging information and code with other developers:
- Account types
- Repository types
- Contributing  to opensource projects

### The 4 GitHub Usecases

core uses of github, some are useful for a single user:
- use github as a cloud storage - durable and available.
- portfolio page - public facing 

for multiple users:
- collaberation on a project, either as a simple user or as part of an organization.
- contributing to other projects, even if it doesn't belong to you.

### Understanding GitHub Account Types

Account Types (pricing):
- Free - Personal user account
- Team - Organization account
- Enterprise - Enterprise account

the personal user account can have public and private repositories, and work with unlimited collaborators for projects.

the organization account is a shared account for groups, they have the same features plus some extras, either the basic set of the team plan or the advanced enterprise plan.

the enterprise account manages multiple github accounts, it's a paid service with the github enterprise cloud and server options.

we can see our account type under <kbd>settings</kbd>, <kbd>Organizations</kbd>.

### Changing the Repository Type from Public to Private

Repositories can be private or public, under <kbd>settings</kbd> we can change the repository visibility and make a repository public or private.

### Pushing Commits to a Public Repository

we can clone any public repository,but we can't push to repositories which we don't own (have an access token).

### How GitHub Manages Account Security

The personal access token provides github access via Git. we can set different permissions levels.

if a different user wants make changes to our code, then he is a collaborator, how this user can interact with our code depends on whether they are inside the same enterprise, if they are part of a team which has the repository, etc...

### Understanding & Adding a Collaborator to a Private User Account

in the repository settings, we can click <kbd>Manage access</kbd> and see who has access. we can invite other github users to become collaborators.

then those collaborators can use their own access tokens, and they'll only have access to what we give them.

### Collaborating in Private Repositories

if we change the visibility of the repository, the collaborators still have the ability to contribute and make changes.

### Comparing Owner & Collaborator Rights

[permissions docs](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-personal-account-settings/permission-levels-for-a-personal-account-repository)

note: in a private user account and a private account, we can't have 'read-only' permissions to another user, so we can't simply invite some to only read the repository.

### Limiting Interactions

in this context, interaction means commenting, opening issues and creating pull requests.

under the <kbd>Settings</kbd> page in our profile, in <kbd>Interaction Limits</kbd> we can restrict what other users can do with our repositories. these limits will win over repository specific limitations. we can restrict actions to a period of time.\
This is a broad-strokes approach, which only makes sense when dealing with public repositories. there are also similar options for individual repositories.

### Organizations

<details>
<summary>
Several Accounts inside an organization.
</summary>

#### Introducing Organizations

member-role access for repositories in a large organization.

#### Creating an Organization

Under <kbd>settings</kbd>, under <kbd>organization</kbd>, we can transform any account into an organization account (as long as it's not part of an organization), or create a new organization from the account. 

if we choose to create an organization, we then choose the plan (free, team, enterprise) and fill in the details.

from the personal user, we can switch into the organization user account

#### Exploring Member Repository Permissions

creating repositories is the same as with the personal account, but in addition to "Direct Access" and "Private Repository" settings, we have "Base Role" options. these settings effect users who are member of the organization, and not outside collaborators.

we can also set some other options, such as creating repositories, forking, creating pages, etc...

#### Adding Outside Collaborators

outside collaborators don't belong to the organization, collaborators have some new permissions - read, triage(managa issues, not code), write, maintain(no destructive actions permitted) and admin.

#### Adding Organization Members

under the <kbd>People</kbd> tab, we can invite members to the organization. the other user needs to accept the invitetation. 

#### Failing to Manage Access for Individual Repositories

member level access is for all repositories, not specific.

</details>

### Introducing Teams

under the <kbd>Teams</kbd> tab, we can create teams, which have more flexible levels of access. teams have a name, ano optional description, sub teams (teams within a team), and visibility level (visible or secret).

in the team page, we can create discussion and write messages. we assign repositores to a team, and then we set the access level for each. 

### Managing Team Repository Access Efficiently

the members will still have permissions based on he "base permission" level, so we need to set it to **no permissions** so that members of the team won't have access to other repositories of the organization.

### Understanding Forks & Pull Requests

contributing to projects without being assigned to them.

cloning vs forking:\
cloning/pushing usually requires the owner to know about the other people who wish to work on the code. a fork is a 'copy' of the remote repository that resides in the external user github account. forking is a github terminology, it isn't native to git.\
the contributing user works on the forked repository and pushes changes to it. when the development is done, a **pull request** is created, and the original owner can decide to accept the PR and merge it into the main repository.

we can also use the forking/pulling process inside the a team, and take advantage of he pull-request features for code review.

### Forking a Repository

we must be signed it to github in order to fork repositories, and we need to have access to it (like it being a public repository). we simply click the <kbd>fork</kbd> button. the forked repository shows the original repo, and we can see the <kbd>contribute</kbd> and <kbd>fetch upstream</kbd> options.

### Pull Requests in Practice

the owner of the repositroy can see who forked the repositories, by looking at the <kbd>insights</kbd> tab.

to start a new pull request, we go to the <kbd>pull requests</kbd> tab and click <kbd>create new pull request</kbd>. we need to choose which is the source and which is the target branch, usually the source is inside the forked repository, while the target is the original.

when we create a pull request, we write a comment to descrive the changes, and we see if there are merge conflicts. if we wish to accept the changes we can click <kbd>Merge pull requests</kbd>. we can also close a pull request if we don't wish to accept it. we could also <kbd>Revert</kbd> a merge, which is another commit.

### Opening & Closing Issues

The <kbd>Issues</kbd> tab, we can visit actual repositroy, like [vueJS](https://github.com/vuejs/vue). the issues tab documents what open issues exists, issues are points for improvements, like bugs, feature requests, etc. a well written issue describes a probelm or a need, and documents the effects, or how to re-produce it. and then other people can work on fixing it.

in our repository, we can create a <kbd>New Issue</kbd>, and fill in the text. we can assign another user to work on the issue, apply labels, attach to a pull requests, etc....


### Working with GitHub Projects

in the <kbd>Projects</kbd> tab, we can <kbd>Create a New Project</kbd>. there are pre-defined templates. a project is a board which tracks the issues and similar items. a project board has cards (or notes) which we can create and move around. it's like a board in Jira, but intergrated into github. when we create an issue, we can assign it to a project, and then it will appear on the board and we can place it in the correct column.

we can click <kbd>Manage Automation</kbd> and have the issues move between the boards based on their state.

### Creating a README File in a Repository

we can create an optional "README.md" file at each folder, and at the repository root. md stand for markdown, a lightweight markup langauge to create formatted text. The file is displayed when we view the repository.\
It contains a description of the repository, basic instructions, and helps managing the contribution rules to the repository.

the file is part of the repository, and we commit it just like any other file.

### Presenting Yourself as Developer on GitHub

we also have a markdown page for our personal account, if we access our profile page, we can create repository that has the same name as the username. this repository can contain a "README.md" that will be displayed when someone visits the user profile page.\
This acts as a github landing-page.

### About GitHub Stars

Github stars are a "reputation" measurement, like facebook "like", instagram "heart" or up-votes.

</details>

## Section 7 - Real Project Example: Git & GitHub Applied
<details>
<summary>
Learning by doing. Step by step example of working on a github project.
</summary>

Applying the knowledge to create a real project on github. we will build a web project on the local git, and push it to github. there will also be a collabration part.

### Preparing the Project
The website is built with React, it's a food shopping website, it's just a UI, with no backend for the code. the project requires nodeJS to run locally.

### Project Setup
we start by downloading the source code and extracting it. to start the website locally, we need to run some commands.

```sh
npm install
npm start
```
the project start at an incomplete state.

### Creating our First Commit Locally
we first initialize the project as a local git repository. we first run `git init`. we might get a warning about the "node_modules" being too much to track. we can delete the folder and re-create it with `npm install`.  we don't need to track those files in git, so we can add the folder to the ".gitignore" file.

now we add the first commit with all the files

```sh
git add .
git commit -m "initial commit"
git status
```

we can rename the master branch into "main"
`git branch -m master main`.


### Changing the ReactJS Project Code
we start this section by changing the app code. this is just react code.

```ts
import {useState} from 'react';


function APP() {

const [cartIsShown, setCartIsShown]=useState(false);

function showCartHandler(){
  setCartIsShown(true);
}
function hideCartHandler(){
  setCartIsShown(false);
}
  return (
    <CartProvider>
      {cartIsShown && <div>Cart...</div>}
      <Header onShowCart={showCartHandler} />
      <main>
        <Meals />
      </main>
    </CartProvider>
  );
}
```

this is a new feature, so we want it as a feature branch.

```sh
git checkout -b feat/cart-logic
git add .
git commit -m "added cart visibility state"
```

### Pushing the Code to the Owner's Remote Repository
we have the commit locally, but it needs to be pushed to github, so other collabrators could work on it as well.

we create the repository from the github page, and we don't need to clone it, we set it as remote repository. We need to have git configured before hand, by running `git config` at some time.

```sh
git remote add origin <url>
git push origin main
```

now we need to login into github, usually with personal access token. from <kbd>Account Settings</kbd>,<kbd>Developer Settings</kbd>, and <kbd>Personal access tokens</kbd>.

if we have done everything correctly, our changes will be on github.

### Cloning the Project & Changing the Code

we are now it the "collabration" part, we want to clone the project from github into our local machine. we simply run `git clone` with the url. if we want the project to be cloned in the current directory, we add a space and a dot to the clone command.

Because this is a node project, we run `npm install`, but this has nothing to do with git or github. we can start the project with `npm start`.

we play with react code, under "/src/component/Cart" we add "Cart.js" and "Cart.module.css" files. in the "app.js" file, we add an import statement

```js
import Cart from './components/Cart/Cart';

function APP(){
  return (
    <CartProvider>
      <Cart />
      <Header />
      <main>
        <Meals />
      </main>
    </CartProvider>
  );
}
```
now we want to push the code back to github.

### Pushing the Commit & Why it Fails
we want our code to be in a seprate branch.

```sh
git checkout -b feat/car
git branch
git add .
git commit -m "cart module"
git push origin feat/cart
```
this fails because we aren't authorized to push into the github repository. we only cloned it, but we still aren't designated as project collabratores.

### Adding a Collaborator

in the github repository, we need to add ourself as a collaborator. this is better than sharing a personal access token.

in the project <kbd>Settings</kbd>, under <kbd>Manage Access</kbd> we click <kbd>Invite a collaborator</kbd>.

### Creating the Personal Access Token & Pushing Successfully

A collaborator can't change the repository settings, it's something that only the owner can do.

we access github from the vscode editor by using a personal access token. under the <kbd>profile</kbd>, <kbd>Settings</kbd>, <kbd>Developer Settings</kbd>, we can create a token with *repo* permissions.\
if everything worked correctly, then we can now create branches.

### Merging the Owner Branches

from the repository owner, we need to `git pull` and then we can run `git branch -a` to see the branches. we want to merge changes.
```sh
git merge feat/cart-logic #merge into current branch
```

### Fixing Merge Conflicts

local branch, tracking branches...

```sh
git checkout origin/feat/cart
git checkout feat/cart # create local tracking branch
git merge feat/cart # merge into main
```
but now we have a conflict. conflicts are marked in the file with "<<<<<< HEAD", "=======" and ">>>>>> Branch Name". we fix the code.

```js
return (
    <CartProvider>
      {cartIsShown && <Cart onClose={hideCartHandler}/>}
      <Header onShowCart={ShowCartHandler}/>
      <main>
        <Meals />
      </main>
    </CartProvider>
  );
```
we commit the code and push it back to github.
```sh
git add .
git commit -m "merge conflicts"
git push origin main
```

### Accessing Code via Forks

a third option to modify the repository is by using pull requests, this doesn't require us to be an owner or designated as a collaborator. we can **fork** the repository, fix an issue and push it back with a **pull request**. forking just requires the repository to be visible, and it part of the github capability, rather than git's. after forking, we clone the branch to the local development machine.

### Contributing to the Project by Improving the Code

we fix the code

```js
const cartItemRemoveHandler = (id) => {
  cartCtx.removeItem(id);
};

const cartItemAddHandler = (item) => {
  cartCtx.addItem({...,item,amount: 1});
};
```
we push the commit to our github repository.

### Creating the Pull Request

in the original repository, under the <kbd>pull requests</kbd> tab, and the <kbd>new pull request</kbd>, the default behavior is to create a pull request from a branch, so we need to move it into "compare accross forks", and then we can see our forked repository. we choose the branches we want to merge between, and then we see the changes and commits. to confirm we click <kbd>Create pull request</kbd> and fill the form with description of our changes.

### Accepting the Pull Request

The code owner sees the pull request to the repository, and can review the changes, comment on them and start a conversation. if everything is fine, we can <kbd>merge pull request</kbd>, if not, we might need to solve any conflicts. accepting a pull request creates a commit. once accepted, a pull request is closed.

</details>

## Takeaways
<details>
<summary>
Stuff worth remembering
</summary>

### Git Commands

command | use |  notes
---|---|---|
`git version` | which version was installed | like `git --version`
`git init` | initialize git in folder | 
`git add` | add files to track | 
`git status` | check status |  see tracked and untracked files
`git log` | see commit logs | exit by pressing `q`
`git config` | change configuration |
`git checkout` | checkout branch or commit | `--` current branch|
`git branch` | see or create branches | 
`git merge <other-branch>` | merge changes |
`git switch` | branch operation| similar to checkout, but solely for branches
`git ls-files` | which files are part of the staging area (which are tracked)
`git restore` | restore files | 
`git clean` | remove untracked files |
`git reset` bring back latest status to the staging area |  `git restore --staged` is a new way of doing this
`git stash` | stash changes without a commit | deafult behavior is `push`
`git reflog` | retrieve deleted data | default behavior is `show`
`git rebase` | recreate commits and change base commit
`git diff` | see differences
`git cherry-pick` | get specific commit
`git tag` | label commits
`git show` | view objects (default HEAD) | show commits, tags, trees, blobs
`git remote` | connect to a remote hosting
`git clone` | clone remote repository

[git version](https://git-scm.com/docs/git-version)
- `--build-options` - more detailed information

[git status](https://git-scm.com/docs/git-status)
- `-s --short` - short form
- `-b --branch` - show branch info when running short form
- `--long` - long form

[git branch](https://git-scm.com/docs/git-branch)
- `-d --delete` - delete branch if merged
- `-D` - delete branch even if wasn't merged, same as `--delete --force`.
- `-v -vv` - verbose
- `-m` - rename
- `-M` - rename force
- `-a --all ` - see all branches (including tracking branches)
- `-r` - remote branches
- `--track` - create tracking branch

[git switch](https://git-scm.com/docs/git-switch)
- `-c --create` - create branch if doesn't exist
- `-C --Create` - create branch, if exists override it

[git log](https://git-scm.com/docs/git-log)
- `-n --max-count` - limit number of log entries
- `--merge`

[git clean](https://git-scm.com/docs/git-clean)
- `-d` - directory recursion when no path given
- `--dry-run -n` - just list files
- `--force -f` -  remove files
- `--interactive -i` - interactive mode

[git reset](https://git-scm.com/docs/git-reset)
- mode:
  - `--soft` -
  - `--mixed` (default) -
  - `--hard` -

[git restore](https://git-scm.com/docs/git-restore)
- `--staged -S`- remove staged changes
- 

[git stash](https://git-scm.com/docs/git-stash)
- `push` - push a dirty snapshot, `-m` flag to add a msg. this is the default behavior.
- `apply` - retrieve from stash, we can use the index to get a specific stash.
- `pop` - retrieve and remove from stash, we can use the index to get a specific stash.
- `list` - see stash list
- `show`
- `drop` - remove a stash by index.
- `clear` - remove all stashes

[git reflog](https://git-scm.com/docs/git-reflog)
- `show` - default behavior, log of user actions (moving between branches, etc)
- `expire`
- `delete`
- `exists`

[git merge](https://git-scm.com/docs/git-merge)
- `squash`
- `--no-ff` - don't do fast-forward merge
- `--abort` - abort conflicted merge

[git rebase](https://git-scm.com/docs/git-rebase)


[git cherry-pick](https://git-scm.com/docs/git-cherry-pick)

[git tag](https://git-scm.com/docs/git-tag)
- `--list -l` - list tags, default
- `-a` - annotated tag
  - `-m` - add message to annotated tag
- `-d` - remove tag

[git show](https://git-scm.com/docs/git-show)

[git remote](https://git-scm.com/docs/git-remote)
- `add origin <url>` - add remote
- `show origin` - show detailed configuration

[git pull](https://git-scm.com/docs/git-pull)

[git push](https://git-scm.com/docs/git-push)
- `-u` - set upstream
- `--delete` - delete remote branches

### .gitignore file

each line is a pattern:

- `*` - as wild card
- `!` - at the start of the line to override ignore rules (force tracking)
- `#` - comments
 
</details>
