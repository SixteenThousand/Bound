# Bound
### Sometimes you don't want to manage, you just want to measure
Bound is a project management tool, or more precisely a project *measurement* 
tool; it helps you to measure how much time is spent on a project so you can 
plan (very rough!) estimates for future projects.

The idea of Bound is to help manage some variant of a pomodoro workflow; that 
is, set a timer => work for that time => set another, shorter timer => 
break, and then back to the first step again, but crucially allows the user 
to choose how long the timers they use will be.

It will also get other clock app features (alarms, stopwatch, etc.) as 
development goes on, but this is not the focus.

---

## Usage
This project is still very early in development, and so does not have usable 
features yet. See the [roadmap](#development-roadmap) to see how bound will 
work.

---

## Development Roadmap
### version 0.0.0
The aim here is make the absolute minimum usable version of Bound possible, 
i.e. a command-line tool that allows the user to set a timer with some metadata 
and then save that data to a database.
- [ ] Ability to set a timer that shows a simple countdown
- [ ] Bound sends a system notification (on windows) when the timer is done
- [ ] Bound sends metadata  about the timer (project name, a set of tags, the 
  start time & date, and the length of the timer) to a database
- [ ] Bound can store metadata about timer & the length of that timer for 
  future use in a list of "Saved Timers"; Bound can also list those timers

### version 1.0.0
Next, the aim will be to make a basic TUI interface for Bound, and add a 
couple of useful features to make it more of a full clock app.
- [ ] Ability to have multiple timers running at the same time
- [ ] Alarm system
- [ ] Stopwatch system

### version 2.0.0
Next, the aim will be to make tools to analyse data produced and to set a 
timetable for future work.
- [ ] Bound can create a calendar showing what work has been done between two 
  given dates in HTML format
- [ ] Bound can parse a csv and/or ods file to make an *actual* timetable to 
  plan future work around, with the ability to export the timetable to HTML (as 
  in the above bullet point), and to compare the timetable to actual work done
- [ ] Bound can calculate some statistics from work data:
    - The total time spent on a given project
    - The total time spent in pomodoros with a given tag
    - The % of time spent on a project in a given period of the day (say, in 
      the morning)

### version 3.0.0
Finally, the aim will be to build a GUI frontend to Bound, from which all of 
the functionality above should be available. This will be done in one (or 
possibly both) of the following ways:
- A browser frontend, where essentially bound will act as a server connected to 
  the browser via localhost
- A GTK frontend

Most functionality should be achieved via simple forms and the like, but in 
particular timetable creation should have a drag-and-drop interface.
