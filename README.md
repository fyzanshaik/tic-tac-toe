# Tic-Tac-Toe with Hand Gesture Control

This project features two components:

1. **Gesture-Controlled Tic-Tac-Toe**: A Tic-Tac-Toe game where moves are made using hand gestures recognized via a camera.
2. **Terminal-Based Tic-Tac-Toe**: A classic Tic-Tac-Toe game running in the terminal.

## Project Structure

-  `game-gesture/`: Contains the code for the gesture-controlled Tic-Tac-Toe game.
-  `game-go/`: Contains the code for the terminal-based Tic-Tac-Toe game.

## Setup

### Gesture-Controlled Tic-Tac-Toe

1. **Install Go** and **Python**.
2. **Navigate to `gesturegame/` folder**:

   ```sh
   cd game-gesture
   ```

3. **Build and run the Go application**:
   ```sh
   go build -o tic-tac-toe
   ./tic-tac-toe
   ```
4. **Install Python dependencies**:
   ```sh
   pip install opencv-python mediapipe
   ```
5. **Run the Python script for gesture recognition**:
   ```sh
   python scripts/gesture_recognition.py
   ```

### Terminal-Based Tic-Tac-Toe

1. **Navigate to `terminalapp/` folder**:
   ```sh
   cd game-go
   ```
2. **Build and run the Go application**:
   ```sh
   go build -o tic-tac-toe
   ./tic-tac-toe
   ```

## Control Flow

1. **Gesture-Controlled Game**:
   -  Runs the Go game logic with input from a Python script that recognizes hand gestures.
2. **Terminal Game**:
   -  A standard terminal-based Tic-Tac-Toe game with text-based input and output.

### Summary:

-  **Project Structure**: Describes the organization of your folders.
-  **Setup**: Provides setup instructions for both components.
-  **Control Flow**: Explains the basic operation of each component.
-  **Contact**: Provides contact information for further inquiries.

```

```
