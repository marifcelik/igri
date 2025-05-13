# Chat App Back-end

### Notes
- ~~I'm gonna remove the session package and use the JWT instead. Because I couldn't find a session package I was aiming for and writing it is not my priority right now.~~ Solved, back to the session package.

---

## API Documentation

This document outlines the API endpoints available in the Go chat backend application.

### General Configuration

*   **CORS:**
    *   **Allowed Methods:** `DELETE`, `GET`, `OPTIONS`, `PATCH`, `POST`, `PUT`, `UPDATE`
    *   **Exposed Headers:** `Accept`, `Authorization`, "Content-Type", `X-CSRF-Token`, and session-related headers (`config.C.HeaderKey.Session`, `config.C.HeaderKey.Expiry`).
    *   **Allowed Headers:** `Accept`, `Authorization`, "Content-Type", `X-CSRF-Token`, and session-related headers.
    *   **Allowed Origins:** `*` (All origins)
    *   **Allow Credentials:** `true`

### Health Check

*   **Endpoint:** `GET /healthz`
    *   **Description:** Returns the health status of the server.
    *   **Responses:**
        *   `200 OK`: Server is healthy.

### Root Endpoint (Session Test)

*   **Endpoint:** `GET /`
    *   **Middleware:** `Auth` (Requires authentication)
    *   **Description:** A test endpoint that increments a counter in the user's session.
    *   **Request:** None
    *   **Responses:**
        *   `202 Accepted`:
            *   Body: `string` (The new count as a string, e.g., "1\n")
        *   `401 Unauthorized`: If the user is not authenticated.

---

## Authentication Module (`/auth`)

Handles user registration, login, and logout.

### 1. User Login

*   **Endpoint:** `POST /auth/login`
    *   **Middleware:** `LoggedIn` (Likely checks if a user is *not* already logged in, or handles session renewal)
    *   **Description:** Authenticates a user and creates a session.
    *   **Request Body:** `application/json`
        ```json
        {
          "username": "string",
          "password": "string"
        }
        ```
        *   `username`: User's username (string, required).
        *   `password`: User's password (string, required, min length 8).
    *   **Responses:**
        *   `200 OK`: Login successful.
            ```json
            {
              "status": "success",
              "data": {
                "id": "string", // User ID (MongoDB ObjectID Hex)
                "name": "string",
                "username": "string",
                "createdAt": "string" // ISO 8601 datetime string
              }
            }
            ```
        *   `400 Bad Request`: Validation error (e.g., missing fields, password too short) or JSON parsing error.
            ```json
            {
              "status": "error",
              "message": "Validation error",
              "data": { /* Detailed validation errors */ }
            }
            ```
        *   `401 Unauthorized`: Incorrect username or password.
            ```json
            {
              "status": "error",
              "data": "username or password is incorrect"
            }
            ```
        *   `500 Internal Server Error`: Server-side error (e.g., database issue).

### 2. User Logout

*   **Endpoint:** `POST /auth/logout`
    *   **Description:** Logs out the current user by destroying their session.
    *   **Request:** None (Session-based)
    *   **Responses:**
        *   `204 No Content`: Logout successful.
        *   `406 Not Acceptable`: If no active session found or error during session destruction.

### 3. User Registration

*   **Endpoint:** `POST /auth/register`
    *   **Description:** Creates a new user account.
    *   **Request Body:** `application/json`
        ```json
        {
          "name": "string",
          "username": "string",
          "password": "string",
          "passwordConfirm": "string"
        }
        ```
        *   `name`: User's full name (string, required).
        *   `username`: Desired username (string, required).
        *   `password`: Desired password (string, required, min length 8).
        *   `passwordConfirm`: Confirmation of the password (string, required, must match `password`).
    *   **Responses:**
        *   `201 Created`: Registration successful. The user is also logged in.
            *   Headers: `Location: /users/{id}` (where `{id}` is the new user's ID)
            ```json
            {
              "status": "success",
              "data": {
                "id": "string", // User ID (MongoDB ObjectID Hex)
                "name": "string",
                "username": "string",
                "createdAt": "string" // ISO 8601 datetime string
              }
            }
            ```
        *   `400 Bad Request`: Validation error (e.g., missing fields, passwords don't match) or JSON parsing error.
            ```json
            {
              "status": "error",
              "message": "Validation error",
              "data": { /* Detailed validation errors */ }
            }
            ```
        *   `409 Conflict`: Username already exists.
            ```json
            {
              "status": "error",
              "data": "username already exists"
            }
            ```
        *   `500 Internal Server Error`: Server-side error (e.g., database issue, password hashing error).

---

## Message Module (`/messages`)

Handles fetching messages and conversations. All endpoints in this module require authentication (`Auth` middleware).

### 1. Get Conversation Messages

*   **Endpoint:** `GET /messages/{conversationID}`
    *   **Description:** Retrieves messages for a specific conversation.
    *   **Path Parameter:**
        *   `conversationID`: The ID of the conversation (string, MongoDB ObjectID Hex).
    *   **Responses:**
        *   `200 OK`: Successfully retrieved messages.
            ```json
            {
              "status": "success",
              "data": [
                {
                  "id": "string", // Message ID (MongoDB ObjectID Hex)
                  "content": "string",
                  "senderID": "string", // Sender's User ID (MongoDB ObjectID Hex)
                  "createdAt": "string" // Formatted as time.UnixDate (e.g., "Mon Jan _2 15:04:05 MST 2006")
                }
                // ... more messages
              ]
            }
            ```
        *   `400 Bad Request`: Invalid `conversationID` format or validation error.
            ```json
            {
              "status": "error",
              "message": "Validation error", // or "Invalid conversation id"
              "data": { /* Detailed validation errors, if applicable */ }
            }
            ```
        *   `500 Internal Server Error`: Server-side error (e.g., database issue).

### 2. Get User Conversations

*   **Endpoint:** `GET /messages/conversations/{userID}`
    *   **Description:** Retrieves all conversations for a specific user.
    *   **Path Parameter:**
        *   `userID`: The ID of the user (string, MongoDB ObjectID Hex).
    *   **Responses:**
        *   `200 OK`: Successfully retrieved conversations.
            ```json
            {
              "status": "success",
              "data": [
                {
                  "id": "string", // Conversation ID (MongoDB ObjectID Hex)
                  "name": "string,omitempty", // Name of the other participant (for normal conversations) or group name
                  "username": "string,omitempty", // Username of the other participant (for normal conversations)
                  "participants": ["string"], // Array of participant User IDs (MongoDB ObjectID Hex)
                  "lastMessage": {
                    "id": "string",
                    "content": "string",
                    "senderID": "string",
                    "createdAt": "string" // Formatted as time.UnixDate
                  }
                }
                // ... more conversations
              ]
            }
            ```
        *   `400 Bad Request`: Invalid `userID` format or validation error.
            ```json
            {
              "status": "error",
              "message": "Validation error", // or "Invalid user id"
              "data": { /* Detailed validation errors, if applicable */ }
            }
            ```
        *   `500 Internal Server Error`: Server-side error (e.g., database issue).

---

## WebSocket Module (`/_ws`)

Handles real-time messaging via WebSockets.

*   **Endpoint:** `GET /_ws`
    *   **Middleware:** `UpgradeChecher`, `WsHeader`, `Auth` (Requires authentication)
    *   **Description:** Upgrades the HTTP connection to a WebSocket connection for real-time communication.
    *   **Connection:**
        *   Once connected, the server stores the connection associated with the authenticated user ID.
    *   **Messages (Client to Server):**
        *   All messages from client to server should be JSON objects.
        *   **Primary Message Structure (`ws.MessageDTO`):**
            ```json
            {
              "type": "NORMAL" | "GROUP", // enums.ConversationType: Type of conversation
              "senderID": "string", // User ID of the sender (MongoDB ObjectID Hex)
              "recipientUsername": "string,omitempty", // Username of the recipient (required for new NORMAL conversations)
              "conversationID": "string,omitempty", // Existing conversation ID (MongoDB ObjectID Hex, required if not a new NORMAL conversation)
              "content": "string", // Message content
              "createdAt": "string" // Client-side timestamp (e.g., ISO 8601 or Unix timestamp string)
            }
            ```
        *   **Special "ping" message:** For some browsers (e.g., Chrome), a raw string `"ping"` (4 bytes) might be sent as a keep-alive, to which the server will respond with a pong.
    *   **Messages (Server to Client):**
        *   All messages from server to client are JSON objects with the following envelope (`ws.WSMessage`):
            ```json
            {
              "type": "RESULT" | "CONVERSATION", // enums.WSMessageType
              "data": {} // Payload, structure depends on "type"
            }
            ```
        *   **If `type` is `RESULT` (`ws.ResultDTO`):**
            *   Sent to the sender to confirm message status.
            ```json
            {
              "status": "SUCCESS" | "FAILED", // enums.ResultStatus
              "conversationID": "string", // Conversation ID (MongoDB ObjectID Hex)
              "messageID": "string", // Message ID (MongoDB ObjectID Hex)
              "message": "string,omitempty" // Error message if status is FAILED
            }
            ```
        *   **If `type` is `CONVERSATION` (payload is `ws.MessageDTO`):**
            *   Sent to the recipient of a message. The structure is the same as the `ws.MessageDTO` sent by the client.
            ```json
            {
              "type": "NORMAL" | "GROUP",
              "senderID": "string",
              // "recipientUsername" might be omitted or present based on server logic
              "conversationID": "string",
              "content": "string",
              "createdAt": "string" // This is the createdAt timestamp from the original sender's message
            }
            ```
    *   **Error Handling (WebSocket):**
        *   If validation fails (e.g., invalid `senderID`, `recipientUsername` not found), a string message describing the error might be sent directly over the WebSocket before closing or as part of a `FAILED` `RESULT` message.
        *   Examples: `"invalid sender id"`, `"user in \"senderID\" field is not found"`, `"username not found in session"`.
