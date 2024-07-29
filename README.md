# BeatBoxBox

<p align="center">
  <img src="./frontend/src/assets/images/logo.png" alt="BeatBoxBox Logo">
</p>
BeatBoxBox is a free online music streaming platform for everyone.


## Getting Started

To get the project running locally:

1. Install the dependencies:
   - Docker
   - Docker Compose

2. Clone the repository (you can also download the repository as a zip file and extract it):
`git clone https://github.com/BoxBoxJason/BeatBoxBox.git`

3. Setup your TLS certificates (to use https):
   1. `cd BeatBoxBox`
   2. `mkdir secret`
   3. `openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout secret/key.pem -out secret/cert.pem`

4. Select your environment variables:
   - Edit the `.env` file to set the environment variables for the project.

5. Build and run the project:
   1. `cd BeatBoxBox`
   2. For linux & mac users: `./quickstart.sh`
   3. For windows users (will never be tested so hope it works): `quickstart.ps1`

6. Access the web interface with your web browser at `https://localhost:3000`
