# GitHub Action - Expect Label Prefix

Inspired by: [trstringer/require-label-prefix](https://github.com/trstringer/require-label-prefix)

## What it does

Checks whether GitHub issue or pull request have label with specific prefix

This action is inspired by [@trstringer's](https://github.com/trstringer/require-label-prefix) but differs from that one:
This action handles only a single issue or pull request just emitted an event (does NOT scan the entire repo).

## Usage

```yaml
    steps:
      - name: Assert that the issue has expected labels
        uses: Rindrics/expect-label-prefix@v1
        with:
          token: ${{ github.TOKEN }}

          # [label_prefix]
          # The prefix you expect the issue to have.
          # If you expect size labels (e.g. "size/S", "size/L") are enforced,
          # the prefix would be "size".
          label_prefix: size

          # [add_label]
          # Whethe or not to add 'default_label' (explained below) to the issue
          # which does not have labels with expected prefix.
          # Options: "true", "false" (default).
          # add_label: false

          # [default_label]
          # The label to be used if `add_label=true`.
          # default_label: "size/needed"

          # [label_separator]
          # The character which divides label prefix and label body
          # Default value: "/"
          # label_separator: "/"

          # [comment]
          # The comment body to be used if `add_label=false`
          # Default value: "Label with expected prefix not found."
          # comment: ""
```
