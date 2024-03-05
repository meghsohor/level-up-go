# 7. Balanced Brackets

In this challenge, we need to implement a solution which will evaluate whether a mathematical expression given by user is valid or not.

The following brackets will be evaluated:

- Round brackets: `(  )`
- Square brackets: `[  ]`
- Curly brackets: `{  }`

### Examples:

```bash
{1/[2+(1+3)]*5}    // Valid expression.
[1+{2*(3+5)]      // Inalid expression. Missing closing square bracket "}".
```

### Requirements:

- Implement a `Stack` with `push` and `pop` methods. **(Go doesn't have built-in Stack data type)**
