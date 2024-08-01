import cv2
import mediapipe as mp
import time

def get_input(prompt):
    print(prompt)
    cap = cv2.VideoCapture(0)
    mp_hands = mp.solutions.hands
    hands = mp_hands.Hands()
    mp_draw = mp.solutions.drawing_utils

    while True:
        success, img = cap.read()
        img_rgb = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)
        results = hands.process(img_rgb)

        cv2.putText(img, prompt, (10, 70), cv2.FONT_HERSHEY_SIMPLEX, 1, (255, 0, 0), 2, cv2.LINE_AA)

        if results.multi_hand_landmarks:
            for handLms in results.multi_hand_landmarks:
                mp_draw.draw_landmarks(img, handLms, mp_hands.HAND_CONNECTIONS)

                # Example gesture detection logic
                thumb_tip_y = handLms.landmark[mp_hands.HandLandmark.THUMB_TIP].y
                index_tip_y = handLms.landmark[mp_hands.HandLandmark.INDEX_FINGER_TIP].y

                if thumb_tip_y < 0.5 and index_tip_y < 0.5:
                    cap.release()
                    cv2.destroyAllWindows()
                    time.sleep(2)  # Wait for 2 seconds
                    return 0  # Example: Row 0

                if thumb_tip_y > 0.5 and index_tip_y < 0.5:
                    cap.release()
                    cv2.destroyAllWindows()
                    time.sleep(2)  # Wait for 2 seconds
                    return 1  # Example: Row 1

                # Add more gestures as needed

        cv2.imshow("Hand Gesture Input", img)

        if cv2.waitKey(1) & 0xFF == ord('q'):
            break

    cap.release()
    cv2.destroyAllWindows()
    return -1  # Return invalid index if no gesture detected

if __name__ == "__main__":
    prompt = "Give row input"
    row = get_input(prompt)
    print(f"Row input received: {row}")
