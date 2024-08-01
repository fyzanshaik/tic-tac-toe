# Tic-Tac-Toe with Hand Gesture Control

This project integrates a classic Tic-Tac-Toe game with hand gesture recognition. The game allows users to make moves by using hand gestures detected through a camera. The backend is written in Go,
and the hand gesture recognition is handled by a Python script using OpenCV and MediaPipe.

## Project Structure

-  `game/`: Contains Go code for the Tic-Tac-Toe game logic.
-  `scripts/`: Contains Python scripts for hand gesture recognition.
-  `README.md`: This file.

## Dependencies

### Go

-  Go 1.18 or later

### Python

-  Python 3.7 or later
-  OpenCV
-  MediaPipe

## Setup

### Go Backend

1. **Install Go**: Follow the [Go installation guide](https://golang.org/doc/install) for your operating system.

2. **Clone the repository**:

   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

3. **Build the Go application**:

   ```sh
   go build -o tic-tac-toe
   ```

4. **Run the Go application**:
   ```sh
   ./tic-tac-toe
   ```

### Python Script for Hand Gesture Recognition

1. **Install Python**: Follow the [Python installation guide](https://www.python.org/downloads/) for your operating system.

2. **Install required Python packages**:

   ```sh
   pip install opencv-python mediapipe
   ```

3. **Run the Python script**:
   ```sh
   python scripts/gesture_recognition.py
   ```

## Control Flow

1. **Game Initialization**: The Go application initializes the Tic-Tac-Toe game board and prints it to the console.

2. **Hand Gesture Input**: The Go application calls the Python script (`gesture_recognition.py`) to get row and column inputs based on hand gestures. The script captures video from the camera,
   processes it using MediaPipe to detect hand landmarks, and returns mock values.

3. **Move Validation**: The Go application validates the input, checks if the chosen cell is available, and updates the game board accordingly.

4. **Computer Move**: After the user's move, the Go application calculates a random move for the computer and updates the board.

5. **Game Status**: The application checks if the game is over (win or draw) and displays the result.

## Why MediaPipe and OpenCV?

-  **MediaPipe**: Provides advanced hand gesture recognition with high accuracy and efficiency. It is suitable for real-time applications and offers a comprehensive set of hand landmarks.
-  **OpenCV**: Facilitates image processing and real-time video capture, making it an ideal choice for building gesture recognition systems.

## File Descriptions

-  `game/main.go`: Contains the main game logic, including functions for starting the game, handling moves, and checking game status.
-  `scripts/gesture_recognition.py`: Python script for hand gesture recognition. It captures video, processes hand landmarks, and provides mock input values based on detected gestures.

## Contact

For any questions or issues, please contact [fyzan.shaik@gmail.com](mailto:fyzan.shaik@gmail.com).

### Summary of Sections:

1. **Project Structure**: Describes the organization of the project.
2. **Dependencies**: Lists required software and libraries.
3. **Setup**: Provides instructions for setting up the Go backend and Python script.
4. **Control Flow**: Explains the sequence of operations in the game.
5. **Why MediaPipe and OpenCV?**: Justifies the choice of libraries used for hand gesture recognition.
6. **File Descriptions**: Briefly describes the purpose of key files in the project.
7. **Contact**: Contact information for further inquiries.
