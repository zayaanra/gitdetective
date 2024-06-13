# gitdetective

<div align="center">
  <img src="https://github.com/zayaanra/gitdetective/blob/main/assets/gitdetective-high-resolution-logo.png" height="250">
</div>

### Example Usage
Idea is to run the command as such:
- `gd`
In order to properly `gitdetective`, you first need to be in the working directory of the git repository that you would like to analyze. The `gd` command will pretty-print basic statistics for the repo. This includes the repo name, when it was created, the number of files, total lines of code, total number of commits, and the number of authors.

- `gd -c -t`
The `-t` flag produces commit statistics for only today. Specifically, it will show a table consisting of each hour of the day and the number of commits per hour.

- `gd -c -w`
The `-w` flag produces commit statistics for the past week.

- `gd -c -m`
The `-m` flag producs commit statistics for the past month.

- `gd -c -y`
The `-y` flag producs commit statistics for the past year.

In addition to the previous 3 commands, you can provide a duration (e.g. `gd -c -w 3`) which shows the the commit statistics from the past 3 weeks.

- `gd -c save -path=<filepath> -filename=<filename>`
Saves all information recorded in the output of `gd commits` into a CSV file where `<path>` is.

- `gd -a`
This provides information about each author. We'll show a pie chart with each author's contribution as a percentage, a table with each author and the amount of commits they made, as well as a list of each files along with the number of commits each file has.

- `gd -a save -path=<filepath> -filename=<filename>`
Saves all information recorded in the output of `gd -a` into a CSV file where `<path>` is.

- `gd -a -t`
Does the same as `gd -a` but only for the past 24 hours.

- `gd -a -w`
Does the same `gd -a` but for the past week. Num of weeks can be specified.

- `gd -a -m`
Does the same `gd -a` but for the past month. Num of months can be specified.

- `gd -a -y`
Does the same `gd -a` but for the past year. Num of years can be specified.

### Additional Features
Some features that would be interesting:
- Serve stats on a webpage
- Use AI/ML to predict future `git` usage
- Use LLms to generate nicer and cleaner visualizations