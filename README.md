git-show-pr
===========

shell command for showing Github pull request

## INSTALL

`make install`

## USAGE

0. add this line `fetch = +refs/pull/*/head:refs/pull/origin/*` to `[remote "origin"]` in `.git/config` to setup refspec
0. `git fetch origin`
0. `git show-pr`

## OPTIONS

0. `--author=`, similar to `git log`'s `--author` option

## WHY USE THIS

As soon as you use `git show-pr` to look at the pull requests, you can easily checkout it locally by `git checkout -b pull/123 pull/origin/123`
