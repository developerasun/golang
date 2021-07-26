import pyautogui
import schedule
import time

# Set a classtime in a dictionary form. 
classTime = dict({
    1 : "09:05",
    2 : "10:05",
    3 : "11:05",
    4 : "12:05",
    5 : "14:05",
    6 : "15:05",
    7 : "16:05",
    8 : "17:05",
})

# Screenshot automator ver 1.0
# Create a screenShot function to take a screenshot.
def screenShot(i): 
    myFileName = str(i) + "교시" + ".png"
    myShot = pyautogui.screenshot()
    myShot.save(r'./' + myFileName)

for i in range(1,9):    
    schedule.every().day.at(classTime[i]).do(screenShot,i)

while True: 
    schedule.run_pending()
    time.sleep(1)

