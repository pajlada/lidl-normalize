# lidl-normalize
Normalize special unicode characters to latin readable ones. Uses a combination of sources to "unconfuse" characters

## Usage

### Update table.go
If you have updated the confusable files under `files/`, you need to regenerate `table.go`. Do this by running the `generate.py` script:
`$ python3 generate.py`
