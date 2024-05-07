# Zeebe Changelog

[![Build Status](https://travis-ci.com/zeebe-io/zeebe-changelog.svg?branch=master)](https://travis-ci.com/zeebe-io/zeebe-changelog)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeebe-io/zeebe-changelog?style=flat-square)](https://goreportcard.com/report/github.com/zeebe-io/zeebe-changelog)
[![Release](https://img.shields.io/github/release/zeebe-io/zeebe-changelog.svg?style=flat-square)](https://github.com/zeebe-io/zeebe-changelog/releases/latest)
[![codecov](https://codecov.io/gh/zeebe-io/zeebe-changelog/branch/master/graph/badge.svg)](https://codecov.io/gh/zeebe-io/zeebe-changelog)

Generate changelog for [Zeebe](github.com/zeebe-io/zeebe) project.


## Example usage

```sh
  export ZCL_FROM_REV=PREV_VERSION
  export ZCL_TARGET_REV=TARGET_VERSION

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
```
