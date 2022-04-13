# fnm-nushell

fnm -> this -> load-env = use fnm in nushell

This is just a quick, simple, non-hacky-script solution for https://github.com/Schniz/fnm/issues/463

## usage

Install:

```
go install github.com/Southclaws/fnm-nushell@latest
```

Load into shell:

```
fnm env --shell powershell | fnm-nushell | from json | load-env
```

Breakdown:

1. we use powershell because it's the easiest to parse without surprises

```
fnm env --shell powershell
```

2. fnm-nushell just turns this into JSON

```
fnm env --shell powershell | fnm-nushell
```

3. turn that json into a nushell table

```
fnm env --shell powershell | fnm-nushell | from json
```

4. load that data table as env vars

```
fnm env --shell powershell | fnm-nushell | from json | load-env
```

