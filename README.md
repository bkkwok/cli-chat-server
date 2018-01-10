# go-chat
simple chat system made with golang

Simple as in, no rooms do not persist if empty.
User's names do not persist either.

## Users
  ### v1
  * can create a room
  * can join a room
  * can leave a room
  * msgs broadcast to all users in room
  * destroy a room that they own
  ### v2
  Following is only possible if user data persists beyond connection close
  * change name
  * change password
  * friends list
  * blacklist
  
## Rooms
  ### v1
  * anyone can join a room
  * any user can create a room
  * only owner can destroy a room
  * all users are booted if room is destroyed
  ### v2
  Following is only possible if user data persists beyond connection close or pointless without peristence
  * rooms can be given a name, and persist even if empty
  * owner can kick users
  * owner can block users from entering room (ip blocking)
  * owner can give permissions to other users
  Possible without persistence
  * mute user cmd
