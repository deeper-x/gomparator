# gomparator

[WIP]

Given an input github URL, you get repo basic informations -usually used to evaluate project-, and store it for next comparisons.

Ideally, you should use the app to compare libs you want to pick up on your current project, so you have a window like that:

```bash
_________________________________________________
NAME  | STARS | CONTRIBUTORS | LAST COMMIT       |
_________________________________________________
LIB1  |   200 |   25         | 2020-01-23  12:16 |
_________________________________________________
LIB2  |   20  |   21         | 2019-01-01  11:12 |
_________________________________________________
LIB3  |   120 |   15         | 2020-11-11  11:11 |
_________________________________________________
```

LIB1 looks better because of more stars, more contribs and recent updates 

## USAGE [WIP]

Ideally, you feed the system with a set of repos you'd like to evaluate, one by one

```bash
home: http://127.0.0.1:8000?query={home}
```
