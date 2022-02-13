# Business-directory-backend
Manages user information and serves API to the front-end mobile


## API Middlewares
1. Security: Encrypt and Decrypt api request and response with AES
2. Client Validation: Uses HMAC encryption to encrypt request body to validate request. request header has two keys 'CID' for passing client ID, and 'CS' for passing encrypted request body.
3. Authentication: Validates users JWT if expired return invalid session else proceed to next middleware or controller. JWT are encrypted using AES to protect its content. 
4. CORS