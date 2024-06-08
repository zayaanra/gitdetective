# gitdetective

### Example Usage
Idea is to run the command as such:
- `gd`
In order to properly `gitdetective`, you first need to be in the working directory of the git repository that you would like to analyze. The `gd` command will pretty-print basic statistics for the repo. This includes the repo name, when it was created, the number of files, total lines of code, total number of commits, and the number of authors.

- `gd commits -t`
The `-t` flag produces repository statistics for only today. Specifically, it will show a table consisting of each hour of the day and the number of commits per hour.

- `gd commits -w`
The `-w` flag produces repository statistics for the past week.

- `gd commits -m`
The `-m` flag producs repository statistics for the past month.

- `gd commits -y`
The `-y` flag producs repository statistics for the past year.

In addition to the previous 3 commands, you can provide a duration (e.g. `gd commits -w 3`) which shows the the repository statistics from the past 3 weeks.

- `gd authors`
This provides information about each author.

- `gd authors -p`
This will provide author-related statistics in the form of a pie chart. Each slice of the pie is how much they have each contributed as a percentage.








