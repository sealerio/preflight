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

list build-in checkers

```shell
[root@iZbp10a38f9badoq9dpp1eZ tmp]# preflight list
+---------------------+--------------+-------+--------------------------------+
|    CHECKER NAME     | CHECKER TYPE | LEVEL |          DESCRIPTION           |
+---------------------+--------------+-------+--------------------------------+
| memory:${arg}       | memory       | fatal | Check the number of megabytes  |
|                     |              |       | of memory required             |
| cpu:${arg}          | cpu          | fatal | Check number of CPUs required  |
| fileexisting:${arg} | fileexisting | warn  | Check the given file does is   |
|                     |              |       | already exist                  |
| port:${arg}         | port         | fatal | Check the the port is          |
|                     |              |       | available                      |
| os:${arg}           | os           | panic | Check host operating system    |
|                     |              |       | info                           |
+---------------------+--------------+-------+--------------------------------+

```

execute build-in runner

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
