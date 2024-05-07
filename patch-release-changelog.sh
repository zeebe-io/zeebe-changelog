#!/bin/bash
#
# Script to make it easier to generate the change log for Zeebe patch releases
#
# It highly depends on the Zeebe changelong tool https://github.com/zeebe-io/zeebe-changelog
# Make sure you have the latest version of the Zeebe Changelog tool (zcl) installed. 
# 
# Additionally, you'll also need a Github 'Personal Access Token' that can be generated in your Github's settings. 
# The token needs to have public_repo scope. For completeness, we will pass the token into zcl as a parameter but you can set it as an environment variable GITHUB_TOKEN.
#
set -exo pipefail

if [ -z $GITHUB_TOKEN ]
then
  echo "You need to set a GITHUB_TOKEN in your environment in order to use this script."
  exit 1
fi

if [ -z $1 ]
then
  echo "Please provide a release to generate the changelog and update the related issues."
  exit 1
fi

version=$1


if [[ "$version" =~ ^8\.[1-9][0-9]*\.[1-9][0-9]*$ ]]
then
  # PATCH VERSION
  #
  # Before continuing, make sure to fetch the latest changes from the repository, as otherwise you may not get the right change log.
  git fetch --all --tags
  git checkout "release-$version"
  git rebase

  # Next, to add the release labels to release's issues, specifying the previous and current release in place of ZCL_FROM_REV and ZCL_TARGET_REV, respectively.
  #
  # ZCL_TARGET_REV should be replaced with the tag name for the version you are releasing, and 
  # ZCL_FROM_REV should be replaced as the tag name for the previous version, based on the release type:
  #
  # ALPHA: the tag for the latest alpha on the current minor version. e.g. if you're releasing 1.2.0-alpha2, then ZCL_FROM_REV=1.2.0-alpha1. If there was no alpha version yet, then it would be the previous minor version. e.g. if you're releasing 1.2.0-alpha1, then ZCL_FROM_REV=1.1.0
  # PATCH: the tag for the previous patch version on the same minor branch. e.g. if you're releasing 1.2.3, then ZCL_FROM_REV=1.2.2.
  # MINOR: the tag for the previous minor version. e.g. if you're releasing 1.3.0, then ZCL_FROM_REV=1.2.0.
  #   As there may be many patch release on a single version, it's simpler for users to take the change log as the list of changes from version X since version X-1, as otherwise they would have to figure out where the versions branched off.

  # To find the previous patch version we extract the patch version
  # and subsctract by one
  patchVersion=$(echo $version | sed 's/8\.[1-9][0-9]\*\.//')
  majorMinor=$(echo $version | sed 's/\.[1-9][0-9]\*$//')
  export ZCL_FROM_REV="$majorMinor".$(( patchVersion - 1 ))
  export ZCL_TARGET_REV="$version"

  # This will add labels to the issues in GitHub. You can verify this step by looking at closed issues. They should now be tagged with the release.
  zcl add-labels \
    --token=$GITHUB_TOKEN \
    --from=$ZCL_FROM_REV \
    --target=$ZCL_TARGET_REV \
    --label="version:$ZCL_TARGET_REV" \
    --org camunda --repo zeebe

  # This command will print markdown code to the console. You will need to manually insert this output into the release draft.
  zcl generate \
     --token=$GITHUB_TOKEN \
     --label="version:$ZCL_TARGET_REV" \
     --org camunda --repo zeebe

else
  echo "Sorry but we support the generation of changelog and updating issues only for patch releases. Please help yourself"
  exit 1
fi
