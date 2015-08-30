---
title: "Getting Started with Github"
date: "2015-01-20"
menu: "Resources"
---

# Getting Started with Github

For this course you will need to use [github](http://github.com) to
manage your code and submit your assignments.  [Git](http://gitscm.com/)
is a widely used version control system originally developed by Linus
Torvalds for use in managing development of the Linux operating system
kernel.  Github is company that provides free web-hosting for git
repositories.

Follow the instructions below to get started using git and set up the
repository you will use for this class (we'll talk more about git in
class, but this should be enough to get started):

1. Register for a student github account: [https://github.com/edu](https://github.com/edu)

      Note: you need to do this even if you already have a public github
      account, unless you have a paid account to support non-public
      repositories.  With the student account, you will be able to have
      non-public repositories for free.

      You will need to verify your email address (which must be a .edu
      address), and then go back to the
      [https://github.com/edu](https://github.com/edu) page to fill in a
      form to request a student account.  You should then receive a
      "Powerup get!" message from github, and be able to create private
      repositories with your account.

2. Set up git on the machine you will use for course assignments. Follow the instructions here:
      [https://help.github.com/articles/set-up-git](https://help.github.com/articles/set-up-git).
      Follow the instructions to download and install a git client for your OS
      ([Mac](http://mac.github.com/),
      [Windows](http://windows.github.com/)).

# Using git for CCC

For some of the assignments in this course, we will provide starting
code in a git repository and you will need to use separate private
repository for each assignment.  (After the assignment is due, you can
make this repository public, both to share it as you wish, and to free
up another private repository since the number of private repositories
you get with your free education account is limited.)

Typically, you need to fetch the code skeleton from the respective
public repository of the course every time before you start working.

The general process of working on projects is to (we'll go through each
step in detail below):

1. Create a private repository on github.

2. **Clone** the repository, so you have a local working copy.

3. Get the code skeleton from the course repo, merge it into your
working copy, and push to the main repository.

4. Edit code and text files in your local copy.

5. **push** the local changes to the main repository.

As you do the assignment, you should commit and push your changes regularly.

## Setting up the Repository

Let's take [PS1]({{< relref "ps/ps1.md" >}}) as an example. Here's what you should do.

<b>1. Create a new repository named `ps1`.</b>

- Visit [https://github.com/repositories/new](https://github.com/repositories/new)
- Select "Private" for the type of repository.  (If you got a student discount correctly, this should be available for free.  If it asks you for
          a credit card, go back to step 1 of [setting up github](|filename|github.md).)
- Keep the "Initialize this repository with a README" unchecked, since you will fetch it later from a public repository.

<b>2. Clone the empty private repository to your working
environment.</b> Instead of _mygithubname_ below, use your github
username.

    :::text
    git clone https://github.com/mygithubname/ps1.git

If you haven't set up SSH keys for github, you'll need to enter your
github username and password.

<b>3. Fetch the code skeleton from the course repository to your private
repository.</b> Enter the working directory of your empty repository and
add a remote repository named `starting`, merge the code, and push it to
your private repository by executing:

    :::text
    git remote add starting git@github.com:CryptoCurrencyCabal/ps1.git
    git pull starting master
    git push --tags origin master

## Working with Files

Now you can write your own code in editor and save it to files. You can see what files have changed by running:
`git status`.  You should see something like this:

    :::text
    # On branch master
    # Untracked files:
    #   (use "git add <file>..." to include in what will be committed)
    #
    #	break_daves_key.go

In this example, nothing has been added to commit but untracked files
are present.

To add the files to the main repository:

    :::text
    git add break_daves_key.go

You should try running `git status` again to see the files that will be
committed.  Note that add stages the file in its current state!  If you
modify the file after the add, the version of the file when you did the
add is the version that will be committed.

Then,

    :::text
    git commit

commits the staged files to the main repository.  This will launch and
editor for you to enter a commit message.  You can use `-m` to provide a
message at the command line instead:

    :::text
    git commit -m 'I am no longer afraid of commitment!'

Your actual commit messages you use should be **clear** and **useful**.
It is tempting to use [lazy commit message](http://whatthecommit.com/),
but you will regret this as your projects get more complex and code
breaks in mysterious ways.

If you want to skip the two-stage commit, you can do:

    :::text
    git commit -a

to automatically add all new and changed files to the commit.

At this point, the changes are stored in your local repository, but not
yet in the main repository.  Once you have changes you want to push to
the main repository:

    :::text
    git push

After this, all the changes are now pushed to the main repository at
github.  Visit your repository in github to see the result.  

This seems like a lot of steps, but by providing a full local repository
and making the steps of staging and pushing to the remote repository
explicit, git provides developers with a lot of control over
repositories that becomes very useful when several (or many) developers
are working on the same project and trying to combine each others
changes (or recover from mistakes).

There are many resources for learning more about git.  The free book,
[Pro Git](http://git-scm.com/book/en/) by Scott Chacon is recommended.

