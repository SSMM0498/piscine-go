#  QUEST 01
## Introduction
>   [Here it is](../hello.sh)

The mainly purpose of this tasks is to teach us :
+   How to retrieve our piscine repo on any computer in the cluster
+   How to upload our changes to that repo

When we start a new session on a computer we first need to:
+   Ensure that git will remember our password with the following command:
    
    ```sh
    git config --global credential.helper store
    ```

+   Then we can clone our repo to the desktop with:
    
    ```sh
    git clone https://learn.zone01dakar.sn/git/serignmbaye/piscine-go.git
    ```

+   And start working on the daily task

To send our changes to the server we need to:
+   Add all changed files to the next commit
    
    ```sh
    git add hello.sh
    ```

+   Making the commit with a specifix message

    
    ```sh
    git commit -m "My very first commit"
    ```

+   Push all the stuff

    ```sh
    git push
    ```

## Make it better
>   [Here it is](../done.tar)

I don't know what this one is for. But for beginners will get them inform there permission which define how to access to files...
It can also consolidate concepts learn in the previous task.

⚠️NOTE: I NEED TO ENTIRE STUDY THIS GIANT COMMAND IN DETAILS⚠️

```sh
TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
```

## To git or not to git
>   [Here it is](../to-git-or-not-to-git.sh)

A very cool one, It enables to me to refresh my knowledge about 
`cut` command

ℹ️ I SAW THAT MAJOR PART OF MY TEAMMATES DO NOT USE A VARIABLE TO STORE THE CURL'S RETURN. 
THEY DIRECTLY USE A PIPE AFTER IT

😏 I TRY TO IT [FIRST](https://learn.zone01dakar.sn/git/serignmbaye/piscine-go/commit/61be403663a4d7c4c621a853371f586917db4d7a) WITH 
`grep -Po '8521'` BUT IT'S OBVIOUSLY NOT WORK


## Who are you
>   [Here it is](../who-are-you.sh)

Finally learn something who is entirely new to me: **JQ COMMAND**.
It's use for reading JSON data with some filters or not
The main structure of the command is :

```sh
jq '<filters>' <source>
```

The filters, I learned here was:
+   `.[]` : For retrieving the entire JSON
+   `select(.<attribute>==<value>)` : For retrieving the data with a condition
+   `.<attribute>` for just getting the attributes we need


For this task the answer was :

```sh
jq '.[]  | select(.id==70) | .name' list.json
```

## Console Camp 1
>   [Here it is](../mastertheLS)

Like its name suggest, It is mainly focus for teaching us that we can use command with a series of options.
Unfortunalety, I first use this hard understanding command because I don't look first for all options of `ls` command. But I improve my skills in using `sed`.

```sh
ls -tF | tr '\r\n' ',' | sed 's/\(.*\),/\1/' | sed 's/*//g'
```


## Console Camp 2
>   [Here it is](../r)

No comment

## Console Camp 3
>   [Here it is](../look)

I litterally adapt a command from this stackoverflow [links](https://unix.stackexchange.com/questions/94009/how-search-for-a-file-beginning-with-either-a-or-z-and-ending-with-a-or-z)

## Console Camp 4
>   [Here it is](../myfamily.sh)

I saw from this [link](https://vic.demuzere.be/articles/using-bash-variables-in-jq/) that I need a `--arg` options for getting the environment variable into the `jq` command but all of my teammates get it withou it.

A very confuse task for me, I push 8 commits before getting the **BIM**.



## Console Camp 5
>   [Here it is](../lookagain.sh)

This [link](https://stackoverflow.com/questions/22727107/how-to-find-the-last-field-using-cut) was very helpfull to know how to cut and get the last element of the cut.
It's done with a very tricky means :
+   You revert the result
+   Cut it and take the first one due of the reversion
+   And finally revert it again to reset the first one effect

## Console Camp 6
>   [Here it is](../countfiles.sh)

An easy one the answer is in introduction video

## Console Camp 7
>   [Here it is](../touchspe.sh)

The main issue with this one, is to create a file with special characters. [Here](https://stackoverflow.com/questions/49988312/how-do-i-create-files-with-special-characters-in-linux) is how to do it.
I JUST NEED TO ESCAPE SPECIAL CHARACTERS WITH THE BACKSLASH SYMBOL (\).

## Console Camp 8
>   [Here it is](../skip.sh)

No comment
[link 1](https://stackoverflow.com/questions/604864/print-a-file-skipping-the-first-x-lines-in-bash)
[link 2](https://superuser.com/questions/852404/what-does-n-option-in-sed-do)

## THE COMMAND LINE MURDERS
[Here](resolving.md) is my note about the section

