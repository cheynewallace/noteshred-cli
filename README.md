# noteshred-cli
A simple command line based tool for sending encrypted, private information with minimal effort & keystrokes using the NoteShred API.

## Usage
#### Create A Note
```noteshred new "the password for your new machine is 72psjajyhd%#"```

#### Create A Note And Send The Details To Someone Via Email

```noteshred new "the password for your new machine is 72psjajyhd%#" someone@gmail.com```

*Result*
```
############################################
Note Successfully Created!
Token:    56cbd17e8e
Password: TSuj#5gMkaFZiegl
URL:      shred.io/56cbd17e8e
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
Notes created via the CLI can still be viewed via the web application and each one will still receive its own unique URL resembling `shred.io/<token>`. In the event the recipient does not use the CLI tool, they can just use a browser.

For simplicity and to avoid having to type a bunch of parameters most note options found on the web application are automated.

* Note passwords are randomly generated 16 byte strings
* Note titles are generic and resemble "May 18 Note"
* All notes are set to "shred after reading"


## OSX Setup
#### Get An API Key
Sign up at www.noteshred.com.
From the dashboard click `Settings` then scroll down and copy the API key
#### Add The API Key To Your .bash_profile
Replace `<your API key>` in the following line with the copied key and run this command from inside the cloned directory

```echo "export NOTESHRED_API_KEY=<your API key>" >> ~/.bash_profile; ln -s $(pwd)/bin/noteshred /usr/local/bin/noteshred```
#### Reload Your .bash_profile
Run ```source ~/.bash_profile```
