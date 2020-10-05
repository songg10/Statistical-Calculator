# Statistical Calculator

## General Goal
- The initial idea of the project was to be a full fledged calculator with many areas such as Calculus, Statistics, Linear Algebrea and IP Subnet. However due to time constraint, I have only able to implement a simple summary statistics calculator, which would be nice for anyone to quickly gather some statistics about their sample through just the usage of the internet.  

## Implementation
For the implementation, the following languages were used:
- **Python** (REST API Server): Python is used for communicating between the front-end and the back-end (implemented using Flask).
- **GoLang** (Calculation): Go is a system-language which is very effecient with calculations. Leveraging the strength of parallelism in GoLang, the calculation generated for the client would be much faster than single-threaded. Hence, improves the responsiveness.
- **JavaScript** (Front-end): Javascript is a perfect language to implement the logic for front-end as it is intergrated on almost all modern browsers and allows you to incoporate logic into the front-end. ReactJs were used to create a better looking UI and speed up the development process.

## Communication
- As we have a seperated front-end and back-end, communication using **Rest API** would make the most sense as it allows us to send data back-and-fourth between the side. In this context, Javascript and Python communicate with each other through HTTP Request.
- For calculation, the foreign function interface in Python **ctypes** was used to call GoLang function from Python. 

## Instructions
### **Before proceeding, please make sure port 8080 and 3000 are free on your host system.** ###
### **If you are not running on CSIL Lab, please replace the file in `/project/cookbooks/polyglot/recipes/defaul.rb` by the file in `non-CSIL/default.rb`** ###
- Clone or download this repository to your computer.
- `cd` into the project directory where we have the **Vagrantfile**.
- Run the following command:
```
vagrant up
vagrant provision
vagrant ssh
```
- Now the server should be up and running, you could make query to the server by yourself using `curl`. However, the main purpose of it is to provide data for the front-end, it is not really useful to do so.
- `cd/project/calculator` to get into the front-end directoy.
- Run the following command:
`npm install --no-bin-links`
- Wait for it to finish installing dependencies and run the following command:
`npm start`
- Now you should be able to access the front-end through your browser through [http://localhost:3000](http://localhost:3000).
- You can now enter the data into the text box. The text box will take a list of number seperated by commas. Every illegal character will be dropped automatically. Then, you could press calculate to get the result in the following table.

## Features
- As the instruction and general idea sections has outlined, this is a calculator for getting the summary of a sample data. 
- There were more features planned initially. However, the other features are still faulty and need more time to develope.
- If one more day were given, I would be able to implement an interactive calculator and some tools for inferential statistical test such as T-test, Mann-Whitney U-test, e.t.c. 
- In the future, I am planning to implement a feature that allows people to enter their data by uploading a .csv file as it is very common for people to record data to a .csv file.