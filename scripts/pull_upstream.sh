#!/bin/bash

if [ $# -lt 1 ]; then
	echo "Pull from specified upstream staged repository. Syncs with upstream master if branch isn't specified."
	echo "usage: $0 <remote> [<branch>]"
	exit 0
fi

source "$(dirname $0)/utils.sh"

upstream_remote_url=$(git remote get-url "$1")
upstream_remote_name="$(cat $repo_list | grep $upstream_remote_url | awk '{print $1;}')"
remote_dir="$staging_dir/$upstream_remote_name"
split_branch="$upstream_remote_name-$(date +%s)"
target_branch=${2:-master}

if [ "$remote_dir" == "$staging_dir/" ] || [ ! -d "$remote_dir" ]; then
	echo "Missing remote from $repo_list"
	exit 1
fi

git fetch -t $upstream_remote_name $target_branch
git subtree split --prefix=$staging_dir/$upstream_remote_name --rejoin -b $split_branch 

git subtree pull --squash -m "Sync upstream $upstream_remote_name $target_branch" --prefix=$staging_dir/$upstream_remote_name $upstream_remote_name $target_branch
git branch -D $split_branch || true

for staged_dep in $(ls $staging_dir | grep -v "^$remote_name$"); do
	staged_mod=$(cd $staging_dir/$staged_dep && go list -m)
	grep "$staged_mod" $remote_dir/go.mod && sh -c "cd $remote_dir && \
							go mod edit -require $staged_mod@v0.0.0-00010101000000-000000000000 && \
							go mod edit -replace $staged_mod=../$staged_dep
							go mod vendor"
done
sh -c "cd $remote_dir && \
	go mod edit -replace $downstream_repo=../../ && \
	go mod vendor && \
	git add go.mod go.sum vendor"

# remove nested OWNERS file for openshift CI
git rm $remote_dir/OWNERS
git commit --amend --no-edit

git checkout ${current_branch}
git merge --squash -s recursive -X theirs -m "Sync upstream $upstream_remote_name $target_branch" ${temp_branch}

printf "\\n** Upstream merge complete! **\\n"
echo "** You can now inspect the branch. **"
echo ""
git diff --dirstat ${current_branch}..${temp_branch}
echo "** Push the changes to remote with **"
echo ""
echo "$ git checkout $temp_branch"
echo "$ git push origin $temp_branch:<BRANCH>"
echo ""
# echo "$ git checkout $temp_branch"
# echo "$ git push origin $temp_branch:master"

