# planisphere

**Definition from Wikipedia**

In astronomy, a planisphere is a star chart analog computing instrument in the form of two adjustable disks that rotate on a common pivot. It can be adjusted to display the visible stars for any time and date. It is an instrument to assist in learning how to recognize stars and constellations. The astrolabe, an instrument that has its origins in Hellenistic astronomy, is a predecessor of the modern planisphere. The term planisphere contrasts with armillary sphere, where the celestial sphere is represented by a three-dimensional framework of rings.

## planisphere is DWMY

planisphere is a routine manager application focused on providing "multidimentional" routine planning.

At this point in time, I can be sure that in about a week something in my current morning routine is going to change, even if slightly. So I divised a system called DWMY, standing for Day - Week - Month - Year, which is simply to go through what you'd like your day, week, month, year would look like. Planisphere will organize these in a local SQLite database and, optionally, also output markdown or just plain text files.

The idea is to keep checking in in your current DWMY. It will prompt you to review your routine weekly by default. But this can be changed in the settings. planisphere will use `git` to manage revisions, so you can always restore or merge with a past version for **max flex**ibility.

By default, the file structure used for outputting files will be as follows:

```bash
dwmy
├── 2024
│   ├── day.md
│   ├── month-08.md
│   ├── week-34.md
│   └── year.md
└── template
    ├── day.md
    ├── month-00.md
    ├── week-00.md
    └── year.md
```
