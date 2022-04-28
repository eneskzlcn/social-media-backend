# social-media-backend


Planning the architecture

###Models
* User
  * Id
  * Profile
  * followers []User
  * followings []User
  * preferences Preferences

* Post
    * content , date, comments Comment[], publisher []User, whoSaw []User, whoLiked []User
* Groups
    * creationDate Date
    * participants []User
    * posts []Post
* Profile
    * name string, 
    * posts []Post
* Preferences
    * any user chosen app properties,
    * user preferences.

* MessageRoom // May need a message service for the message rooms not thought yet.
    
* Messages
  * sender User, 
  * whoSaw []User,
  * publishDate Date, 
  * 

* Comment
  * sender []User, 
  * publishDate Date , 
  * whoLiked []User ,
  * writedFor Post, 

    