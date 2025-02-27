# Trie - prefix tree algorithm

CLI application for prefix tree, using the [Go](https://go.dev/) language and the [Table](https://github.com/aquasecurity/table) library.

This project was made based on the video from the channel [NeuralNine](https://www.youtube.com/@NeuralNine).

## How to run on your machine
1. Clone the repository:
   ```
   git clone https://github.com/alexsandro49/trie.git
   ```
2. Access the project files:
   ```
   cd trie/
   ```
3. Compile the application:
   ```
   go build .
   ```
4. Change the permissions of the binary:
   ```
   chmod +x ./trie
   ```
5. Run the application:
   ```
   ./trie -h
   ```

## Application commands
* `-h` - To list all commands.
* `-add` - Add a new word.
* `-delete` - Delete a word.
* `-search` - Search a word. 
* `-list` - List all registered words.

## References
### Original project video:
   ```
   https://www.youtube.com/watch?v=y3qN18t-AhQ
   ```

## License
- [MIT](https://github.com/alexsandro49/cli_todo_list/blob/main/LICENSE)
