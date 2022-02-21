# Business-directory-backend
Manages user information and serves API to the front-end mobile

## Requirements
1. Go version 1.17
2. Mongo database

## Get started
`git clone https://github.com/TechBuilder-360/business-directory-backend.git`

`go install`

`go run main.go`

## API Middlewares
1. **Security**: Encrypt and Decrypt api request and response with AES
2. **Client Validation**: Uses HMAC encryption to encrypt request body to validate request. request header has two keys 'CID' for passing client ID, and 'CS' for passing encrypted request body.
3. **Authentication**: Validates users JWT if expired return invalid session else proceed to next middleware or controller. JWT are encrypted using AES to protect its content. 
4. **CORS**

## Authentication Flow

### User enrollment
1. User registers using email address.
2. Confirmation mail get sent to user to verify/activate account.
3. User enters activation code to proceed to use the app.

### User sign-in
1. User enters email.
2. Login code gets sent to email address.
3. User enters code and gets redirected to dashboard.

## Organization 
Payment is a way to track the proof of life of an organization. 
### Creation 
1. A registered user can create an organization and by default become the owner of that organization.
2. Organization has 6 months free enrollment before they become inactive (not visible on search).
3. An organization running a free trial will be given a bronze star to indicate that it was registered less than 6 months on the platform.
4. Organization running free trial would have access to limited features.
5. An organization can opt to pay monthly or yearly subscription, notification would be sent 1 month before an account would be turned inactive if payment isn't made before then.

### Organization features
1. Has a minimum of one branch (HQ)
2. Location is optional
3. An organization without location cannot create a branch.