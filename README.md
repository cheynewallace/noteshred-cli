# noteshred-cli
A simple command line based tool for sending encrypted, private information with minimal effort & keystrokes using the NoteShred API.

## Usage
#### Create A Note
```noteshred new "the password for your new machine is 72psjajyhd%#"```

*Result*
```
############################################
Note Successfully Created!
Token:    56cbd17e8e
Password: TSuj#5gMkaFZiegl
URL:      http://shred.io/56cbd17e8e
CLI:      noteshred show 56cbd17e8e TSuj#5gMkaFZiegl
############################################
```

#### Read A Note
```noteshred show <note token> <password>```

*Example using the above note*

```noteshred show 56cbd17e8e TSuj#5gMkaFZiegl```

*Result*
```
############################################
Title:   Command Line Note
From:    Cheyne Wallace
Message: the password for your new machine is 72psjajyhd%#
############################################
```

#### Things To Note
For simplicity and to avoid having to type a bunch of parameters most note options found on the web application are automated.

* Note passwords are randomly generated 16 byte strings
* Note titles are generic and set to "CLI Note"
* All notes are set to "shred after reading"


## OSX Setup
#### Get An API Key
Sign up at www.noteshred.com.
From the dashboard click `Settings` then scroll down and copy the API key
#### Add The API Key To Your .bash_profile
Replace `<your API key>` in the following line with the copied key and run

```echo "export NOTESHRED_API_KEY=<your API key>" >> ~/.bash_profile```
#### Reload Your .bash_profile
Run ```source ~/.bash_profile```
