# Remove Dups - Golang Command Line Utility

The utility helps remove copy of files that have filenames such as
example(1).txt example(2).txt, other_example(1).exe, and so on. Once complete 
the utility prints out the total # of files deleted.

### Usage
1. Clone repository and navigate to the directory
2. Build using ``` go build ```
3. Run ``` remove_dups.exe <path> ```

### Example
```remove_dupes.exe C:\Users\Mh\Downloads```
The utility also supports three different flags:
- ```-all``` 
  * Removes all copies of a file regardless of number in parentheses
- ```-v```
  * Outputs to stdout each file that was deleted
- ```-dup``` 
  * Used to indicate which copy of a file should be removed
  * If ```remove_dups.exe -dup=2 .``` is ran then only files like example(2) are deleted

#### Note: 
- **If no flags are provided the utility removes all copies of a file**
- **If both -all and -dup flags are used, -all overrides the -dup flag**


