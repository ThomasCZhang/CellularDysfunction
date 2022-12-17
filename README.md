# CellularDysfunction
02601-Programming Group Project <br>
Members: Sumitra Lele, Shamieraah Jamal, Aruneshwar Venkat, Thomas Zhang

Link to github project: https://github.com/ThomasCZhang/CellularDysfunction
Code can also be acquired by cloning the repository from the link above.

## Setup And Running The Code:

1) Place the folder CellularDysfunction in go/src.
2) On a terminal, navigate to the CeullarDysfunction folder (where main.go is located) 
3) Compile the code with "go build"
4) Run the file with "./CellularDysfunction.exe" (without the ./ for windows)
5) A local host address "http://localhost:5000" should appear in the console. Copy that and paste it in any browser.

## The Web App:

7 Fields will appear in the web app. These fields are input parameters to simulate cells in the ECM.

Inputs Fields:
1) Number of Generations (int): The number of generations to simulate the ECM for. It is recommended to keep this relatively low (less than 300). Each generation has to be drawn to a gif, so the more generations there are the longer the code takes to run.

2) Time Step (float64): The amount of time that passes between each generation in hours. Recommended to keep this value between 0 and 1. Values too small barely show any cell movement. Values too large cause cells to move very erratically. 

3) Number of Cells (integer): The number of cells to place on the ECM.

4) Number of Fibres (integer): The number of fibres to place on the ECM. Recommended to keep this between 5000 and 15000.

5) Stiffness (float64): The stiffness of the ECM. Must input a value between 0 and 1. 0 means the ECM is not very stiff at all, 1 means the ECM is very stiff. Default value is 0.95.

6) Cell Speed (float64): The speed the cell travels in micrometers/hour. Recommended to keep this value between 10 and 20.

7) Width (float64): The width of the ECM "board". The ECM board is a square so the width is also the length. Recommended to keep this between 500 and 1000.

Once all the fields have been filled in. Click on the "Submit Query" button. This will
begin the simulation. The simulation should finish very quickly, however the time to draw
the gif may take a while. With 200 generations it takes around 2-3 minutes.

The generated gif will also be saved on the local computer in the folder ".\CellularDysfunction\gifs" as "CellMigration.out.gif".

## Video Walkthrough:

https://cmu.zoom.us/rec/share/QckshbcHYS1JBdKmLsVhL_TIXrVTwBqXpsqMxdqNB-9l7JlIATcZVGA_Jmt9LqHa.FVW5W2MJody5AJtt?startTime=1671250392000 <br>
Passcode: nNB8?hL*

There is also a copy of the video tutorial in the subfolder "Documents" with the name "CellularDysfunction_Tutorial.mp4"