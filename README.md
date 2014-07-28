git-show-pr
===========

shell command for showing Github pull request

## INSTALL

`make install`

## USAGE

0. add this line `fetch = +refs/pull/*/head:refs/pull/origin/*` to `[remote "origin"]` in `.git/config` to setup refspec
0. `git show-pr`
