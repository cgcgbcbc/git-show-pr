git-show-pr
===========

shell command for showing Github pull request

## INSTALL

`cp -f git-show-pr /usr/local/bin/git-show-pr`

## USAGE

0. add this line `fetch = +refs/pull/*/head:refs/pull/origin/*` to `[remote "origin"]` in `.git/config` to setup refspec
0. `git show-pr`
