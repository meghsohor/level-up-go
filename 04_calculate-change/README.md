# 4. Calculate Change

In this challenge, we need to implement a function that calculates the change for a given amount of money and a list of coins. The function should return the minimum number of coins required for the change.

### Available list of coins:

```go
var coins = []coin{
	{name: "1 pound", value: 1},
	{name: "50 pence", value: 0.50},
	{name: "20 pence", value: 0.20},
	{name: "10 pence", value: 0.10},
	{name: "5 pence", value: 0.05},
	{name: "1 penny", value: 0.01},
}
```

### If the user input is `1.50`, the output should be:

```bash
"1 pound": 1
"50 pence": 1
```
