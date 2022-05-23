# Preflight

## Usage

### list all checkers

```shell
preflight list 
```

### skip run checkers

run all checkers expect checker type is port and os

```shell
preflight run --skip port,os
```

### ignore checkers result

run all checkers and ignore cpu and mem errors,just warning.

```shell
preflight run --ignore-errors cpu,mem
```

## Examples
