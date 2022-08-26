<!--
// cSpell:ignore Schwarzmüller 
 -->

# Learn Git & GitHub and master working with commits, branches, the stash, cherry picking, rebasing, pull requests & more!

udemy course [Learn Git & GitHub and master working with commits, branches, the stash, cherry picking, rebasing, pull requests & more!](https://www.udemy.com/course/git-github-practical-guide/) by *Manuel Lorenz* and *Maximilian Schwarzmüller*. 


1. Introduction
1. Mac Terminal & Windows Prompt Introduction
1. Version Management with Git - The Basics
1. Diving Deeper Into Git
1. From Local To Remote - Understanding GitHub
1. GitHub Deep Dive - Collaboration & Contribution
1. Real Project Example: Git & GitHub Applied

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

### Please Read! Windows & Mac Users

### Mac Terminal - The Basics

### Accessing Folders

### Absolute vs Relative Paths

### Creating & Deleting Files

### Introducing Flags and Removing Data

### Copying & Moving Files & Folders

### Mac Terminal - Core Commands Overview

### Windows Command Prompt - The Basics

### Absolute vs Relative Paths

### Creating & Deleting Files & Folders

### Copying & Moving Files

### Windows Command Prompt - Core Commands Overview

### Useful Resources & Links

### </details>


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
<!-- <details> -->
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

if we staged a file and we want to undo those changes

#### Deleting Commits with `git reset`
#### Deleting Branches

</details>

### Committing "detached HEAD" Changes

### Understanding .gitignore

### Wrap Up & Basic Commands Overview

### Assignment 1 - Practicing the Git Basics

### Useful Resources & Links

</details>


## Section 4 - Diving Deeper Into Git
<details>
<summary>

</summary>
</details>

## Section 5 - From Local To Remote - Understanding GitHub
<details>
<summary>

</summary>
</details>

## Section 6 - GitHub Deep Dive - Collaboration & Contribution
<details>
<summary>

</summary>
</details>

## Section 7 - Real Project Example: Git & GitHub Applied
<details>
<summary>

</summary>
</details>


## Takeaways
<details>
<summary>
Stuff worth remembering
</summary>

git commands

command | use | flags | notes
---|---|---|---
`git version` | which version was installed | `--build-options` for additional data | like `git --version`
`git init` | initialize git in folder | 
`git status` | check status | `-s --short` `-b --branch` | see tracked and untracked files
`git log` | see commit logs || exit by pressing `q`
`git config` | change configuration |`--global`
`git checkout` | checkout branch or commit | `--` current branch|
`git branch` | see or create branches | `-v` for verbose info
`git merge <other-branch>` | merge changes ||
`git switch` | branch operations | `-c` to create new branches | similar to checkout, but solely for branches
`git ls-files` | which files are part of the staging area (which are tracked)
`git restore` | restore files | `--staged`
`git clean` | remove untracked files | `-d`,`-n`,`-f`|


git status
- `-s --short` - short form
- `-b --branch` - show branch info when running short form
- `--long` - long from
</details>
