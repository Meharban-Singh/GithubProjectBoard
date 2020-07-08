GithubProjectBoard
==================

A mobile app to interact with github project boards.

## Tools

We are using Expo based React Native for client and Go for the server.  

---

## SET UP
##### First time collaborating? Follow these steps to begin :)

###### Client Setup

1. Install Node.js.
    * https://nodejs.org/en/


2. Open a terminal in your cloned directory. Change directory to the ``client``
folder. Run this command:``npm install``. This should install all the
dependencies required to run the program on your mobile device or on an emulator.


3. Install Expo.
  * Run this anywhere in the terminal once you have node installed:
  ``npm install -g expo-cli``


4. Run ``expo start`` in the ``client`` folder in the cloned directory.
This should start an expo server and you should see a QR code.
Now, you can download and install Expo mobile app on your phone and scan this
QR code to see the mobile app, or use an emulator to run the app on your machine.
For more information refer to this link: https://docs.expo.io

###### Server Setup

1. Install Go.
    * https://golang.org/doc/install


2. For testing, you need access to your Github account. To do that, follow these
steps:
    1. Login to your GitHub account.
    2. Click on ``settings``. Then click on ``Developer settings``.
    3. Select ``Personal Access Tokens``. Click on ``Generate new token ``.
    4. Fill in some note: ``PAT for testing githubprojectboards`` and give it
    all permissions for ``notifications``, ``user`` and ``repos``.
    5. Finally click ``Generate``. Now you need to copy this token. We will need
    it in the next step.


3. Change directory to the ``server`` folder in your terminal and create a file
called ``config.json``. Add these lines in it:
```javascript
{
    "pat": "<YOUR PERSONAL AUTHORIZATION TOKEN FROM GITHUB>"
}
```


4. Change directory into the ``server`` folder and type in the command:
``go run server.go``. This will run the server in your machine.

Now you are all set up to run and test the app :)

---
