# GoGroupDigest
GoGroupDigest automates email digest creation using Go, Gmail API, and OpenAI's GPT. It fetches, summarizes, and circulates essential action items from gmail Groups, streamlining communication and enhancing productivity in corporate environments.

Setting Up OAuth for goGroupDigest
To use goGroupDigest with your own Gmail account, you need to set up OAuth credentials:

# Getting Started with goGroupDigest
Prerequisites
Install Go: Download the latest version of Go from the official Go website. Follow the installation instructions for your operating system.
Setting Up a Workspace
Create a Workspace Directory: Create a directory for your Go workspace (e.g., mkdir ~/goGroupDigestWorkspace).
Workspace Structure: Inside your workspace, create three subdirectories: src, bin, and pkg.
Project Setup
Clone the Repository: Clone goGroupDigest into the src directory of your workspace.
Navigate to Project: cd ~/goGroupDigestWorkspace/src/goGroupDigest.
Installing Dependencies
Initialize Go Modules: In the project directory, run go mod init.
Install Required Libraries: Run the following commands:
bash
Copy code
go get golang.org/x/oauth2
go get google.golang.org/api/gmail/v1
Update go.mod and go.sum: Running go get updates these files with the new dependencies.
Running the Project
Build and Run: Inside the project directory, use go build to compile and ./goGroupDigest to run the project.
Further Steps
Follow the project-specific instructions for setting up OAuth credentials and other configurations.


# Setting Up OAuth for goGroupDigest
To use goGroupDigest with your own Gmail account, you need to set up OAuth credentials:

Google Cloud Project:

Go to the Google Cloud Console.
Create a new project.
Enable the Gmail API in the "Library" section.
Create OAuth Credentials:

In the project dashboard, go to "Credentials".
Click "Create Credentials" and choose "OAuth client ID".
Set up the consent screen as prompted.
Choose the application type (e.g., Desktop app) and name it.
Download the JSON file containing your OAuth client ID and secret.
Configure goGroupDigest:

Place the downloaded JSON file in the root directory of goGroupDigest or a designated configuration folder.
Rename the file to credentials.json (or update the application's configuration to recognize the file's given name).
Follow these steps to authenticate with your Google account securely and start using goGroupDigest.
