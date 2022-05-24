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

### preflight result show

```shell
[root@iZbp10a38f9badoq9dpp1eZ tmp]# preflight run
{
    "passed": [
        {
            "checker_name": "cpu:2",
            "checker_type": "cpu"
        },
        {
            "checker_name": "memory:1700",
            "checker_type": "memory"
        }
    ],
    "failed": [
        {
            "checker_name": "port:6443",
            "checker_type": "port",
            "error_message": "Port 6443 is in use,listen tcp :6443: bind: address already in use",
            "metadata": {
                "description": "Check the the port is available",
                "level": "fatal",
                "explain": "an open port is a network port that accepts incoming packets from remote locations",
                "suggestion": "Maybe you should check your machine of the port is available for use"
            }
        }
    ]
}

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
