# Preflight

## CLI Usage

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

### ignore checkers result

specify the runner option whether return immediately when an error is reported. if `--not-tolerable`=true, will return
immediately and skip remaining checks.

```shell
preflight run --not-tolerable
```

## PKG Usage

```shell
func main() {
	checkList := []checker.Interface{
		checker.PortCheck{Port: 22},
		checker.NumCPUCheck{NumCPU: 2},
	}

	r, err := runner.NewCheckRunner(checkList)
	if err != nil {
		fmt.Println("failed to init runner")
	}

	resp := result.NewDefaultFormatter(r.Execute()).Format()

	responseJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println("format json failed")
	}

	fmt.Println(string(responseJSON))
}

```

## Examples
