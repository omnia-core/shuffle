# Shuffle: Random Team Generator 🔀

`shuffle` is a command-line tool to shuffle and randomly distribute team members into teams. You can specify names directly or load them from a file. It’s perfect for team assignments, random grouping, or just having fun with lists of names.

---

## Installation

To install the latest version of `shuffle`, run the following command:

```bash
go install github.com/omnia-core/shuffle/cmd/shuffle@latest
```

---

## Usage

### About

Shuffle a list of team members into random groups.

### Usage

`shuffle [flags]`

### Flags

- **`-n, --names`**  
  **Type**: `string`  
  Description: List of names separated by commas.  
  Example: `"Alice,Bob,Charlie,David,Eve"`

- **`-f, --file`**  
  **Type**: `string`  
  Description: File containing a list of names, one name per line.

- **`-t, --teams`**  
  **Type**: `int`  
  Description: Number of teams to create.

- **`-s, --size`**  
  **Type**: `int`  
  Description: Size of each team.

### Note

Only **one** of the flags `-n` or `-f` should be provided at a time.

---

## Example

Shuffling names and creating 2 teams of size 2:

```bash
shuffle -n "Alice,Bob,Charlie,David,Eve" -t 2 -s 2
```

### Output:
```
Team 1: [Eve Alice]
Team 2: [Bob Charlie]
```

Shuffling names from a file and creating 5 teams of size 4:
```bash
shuffle -f names.txt -t 2 -s 2
```

### Output:
```
Team 1: [Rachel Quinn Ivy Mia]
Team 2: [Jack Alice Olivia Tina]
Team 3: [Eve Liam Steve Hank]
Team 4: [Diana Frank Paul Karen]
Team 5: [Charlie Grace Bob Noah]
```

---

## License

This project is licensed under the MIT License.

---

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue for feature requests or bugs.

---

## Author

Created by [omnia-core](https://github.com/omnia-core).
